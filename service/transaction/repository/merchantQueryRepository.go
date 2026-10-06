package repository

import (
	"context"

	merchantadapter "github.com/MamangRust/monolith-ecommerce-pkg/adapter/merchant"
	db "github.com/MamangRust/monolith-ecommerce-pkg/database/schema"
)

// merchantQueryRepository delegates merchant lookups to the shared gRPC adapter.
// Failures keep propagating the dependency's raw gRPC status, as they did
// before the adapter migration.
type merchantQueryRepository struct {
	adapter merchantadapter.QueryRepository
}

func NewMerchantQueryRepository(adapter merchantadapter.QueryRepository) *merchantQueryRepository {
	return &merchantQueryRepository{
		adapter: adapter,
	}
}

func (r *merchantQueryRepository) FindByID(ctx context.Context, merchant_id int) (*db.Merchant, error) {
	return r.adapter.FindByID(ctx, merchant_id)
}
