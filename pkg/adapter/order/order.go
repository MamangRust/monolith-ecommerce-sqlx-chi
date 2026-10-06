// Package order adapts the Order service gRPC query API into the shared
// pkg/database sqlc row types.
package order

import (
	"context"

	pborder "github.com/MamangRust/monolith-ecommerce-pb/order"
	"github.com/MamangRust/monolith-ecommerce-pkg/adapter"
	"github.com/MamangRust/monolith-ecommerce-pkg/adapter/convert"
	db "github.com/MamangRust/monolith-ecommerce-pkg/database/schema"
	"github.com/MamangRust/monolith-ecommerce-pkg/resilience"
	order_errors "github.com/MamangRust/monolith-ecommerce-shared/errors/order_errors"
)

// QueryRepository is the contract consumers depend on for order reads.
type QueryRepository interface {
	FindByID(ctx context.Context, orderID int) (*db.Order, error)
}

type repository struct {
	client pborder.OrderQueryServiceClient
	guard  *resilience.DependencyGuard
}

// New builds an order query adapter. Passing zero options leaves the guard nil,
// which makes DependencyGuard.Call a plain passthrough.
func New(client pborder.OrderQueryServiceClient, opts ...adapter.GuardOption) *repository {
	r := &repository{client: client}
	for _, opt := range opts {
		opt(r)
	}
	return r
}

// NewQueryAdapter returns the adapter typed as the read-only contract.
func NewQueryAdapter(client pborder.OrderQueryServiceClient, opts ...adapter.GuardOption) QueryRepository {
	return New(client, opts...)
}

func (r *repository) SetGuard(g *resilience.DependencyGuard) {
	r.guard = g
}

var _ QueryRepository = (*repository)(nil)

func (r *repository) FindByID(ctx context.Context, orderID int) (*db.Order, error) {
	resp, err := adapter.Call(r.guard, ctx, func(callCtx context.Context) (*pborder.ApiResponseOrder, error) {
		return r.client.FindById(callCtx, &pborder.FindByIdOrderRequest{Id: int32(orderID)})
	})
	if err != nil {
		return nil, order_errors.ErrFindById.WithInternal(err)
	}
	if resp == nil || resp.Data == nil {
		return nil, order_errors.ErrFindById
	}

	return mapOrder(resp.Data), nil
}

// mapOrder converts a proto order into the sqlc row type.
func mapOrder(o *pborder.OrderResponse) *db.Order {
	if o == nil {
		return nil
	}
	return &db.Order{
		OrderID:    o.Id,
		UserID:     o.UserId,
		MerchantID: o.MerchantId,
		TotalPrice: o.TotalPrice,
		CreatedAt:  convert.Timestamp(o.CreatedAt),
		UpdatedAt:  convert.Timestamp(o.UpdatedAt),
	}
}
