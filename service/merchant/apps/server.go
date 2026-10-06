package apps

import (
	"context"
	"fmt"
	"time"

	"github.com/MamangRust/monolith-ecommerce-merchant/cache"
	"github.com/MamangRust/monolith-ecommerce-merchant/handler"
	merchantKafka "github.com/MamangRust/monolith-ecommerce-merchant/kafka"
	"github.com/MamangRust/monolith-ecommerce-merchant/repository"
	"github.com/MamangRust/monolith-ecommerce-merchant/service"
	pbmerchant "github.com/MamangRust/monolith-ecommerce-pb/merchant"
	pbmerchant_document "github.com/MamangRust/monolith-ecommerce-pb/merchant_document"
	pbuser "github.com/MamangRust/monolith-ecommerce-pb/user"
	"github.com/MamangRust/monolith-ecommerce-pkg/adapter"
	"github.com/MamangRust/monolith-ecommerce-pkg/kafka"
	"github.com/MamangRust/monolith-ecommerce-pkg/outbox"
	"github.com/MamangRust/monolith-ecommerce-pkg/resilience"
	"github.com/MamangRust/monolith-ecommerce-pkg/server"
	"github.com/MamangRust/monolith-ecommerce-shared/observability"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// kafkaOutboxPublisher adapts the ecommerce *kafka.Kafka (whose SendMessage
// takes no context) to the outbox.OutboxPublisher contract.
type kafkaOutboxPublisher struct {
	k *kafka.Kafka
}

func (p kafkaOutboxPublisher) SendMessage(_ context.Context, topic, key string, value []byte) error {
	return p.k.SendMessage(topic, key, value)
}

func NewServer(cfg *server.Config) (*server.GRPCServer, error) {
	srv, err := server.New(cfg)
	if err != nil {
		return nil, err
	}

	userAddr := viper.GetString("GRPC_USER_ADDR")

	userConn, err := grpc.NewClient(
		userAddr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to user service: %w", err)
	}

	userQueryClient := pbuser.NewUserQueryServiceClient(userConn)

	guardUser := resilience.NewDependencyGuard("user", 5, 30, 100, 3*time.Second, srv.Logger)

	repos := repository.NewRepositories(srv.DB, userQueryClient,
		repository.GuardOptions{
			User: []adapter.GuardOption{
				adapter.WithDependencyGuard(guardUser),
			},
		},
	)
	myKafka := kafka.NewKafka(srv.Logger, []string{viper.GetString("KAFKA_BROKERS")})
	mencache := cache.NewMencache(srv.CacheStore)
	obs, _ := observability.NewObservability(viper.GetString("merchant-server"), srv.Logger)

	outboxService := outbox.NewOutboxService(srv.DB, kafkaOutboxPublisher{k: myKafka}, srv.Logger)

	svc := service.NewService(&service.Deps{
		Kafka:         myKafka,
		Repositories:  repos,
		Mencache:      mencache,
		Pool:          srv.DB,
		Outbox:        outboxService,
		Logger:        srv.Logger,
		Observability: obs,
	})

	// Start the outbox relay so events committed with the business writes are
	// published to Kafka with durable retry and dead-letter semantics.
	go outboxService.Start(srv.Ctx, outbox.OutboxRelayInterval, outbox.OutboxRelayBatchSize)

	h := handler.NewHandler(&handler.Deps{
		Service: svc,
		Logger:  srv.Logger,
	})

	if err := myKafka.StartConsumersWithContext(srv.Ctx, []string{"merchant-service-topic-transaction-event"}, "merchant-service-group", merchantKafka.NewTransactionConsumer(srv.Ctx, mencache.MerchantCommandCache, srv.Logger)); err != nil {
		return nil, fmt.Errorf("failed to start merchant transaction consumer: %w", err)
	}

	srv.RegisterServices = func(gs *grpc.Server) {
		pbmerchant.RegisterMerchantQueryServiceServer(gs, h.MerchantQuery)
		pbmerchant.RegisterMerchantCommandServiceServer(gs, h.MerchantCommandHandler)
		pbmerchant_document.RegisterMerchantDocumentQueryServiceServer(gs, h.MerchantDocumentQuery)
		pbmerchant_document.RegisterMerchantDocumentCommandServiceServer(gs, h.MerchantDocumentCommand)
	}

	return srv, nil
}
