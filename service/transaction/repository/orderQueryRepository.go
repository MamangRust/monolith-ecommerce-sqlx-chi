package repository

import (
	"context"

	orderadapter "github.com/MamangRust/monolith-ecommerce-pkg/adapter/order"
	db "github.com/MamangRust/monolith-ecommerce-pkg/database/schema"
)

// orderQueryRepository delegates order lookups to the shared gRPC adapter.
// Failures keep propagating the dependency's raw gRPC status, as they did
// before the adapter migration.
type orderQueryRepository struct {
	adapter orderadapter.QueryRepository
}

func NewOrderQueryRepository(adapter orderadapter.QueryRepository) *orderQueryRepository {
	return &orderQueryRepository{
		adapter: adapter,
	}
}

func (r *orderQueryRepository) FindByID(ctx context.Context, order_id int) (*db.Order, error) {
	return r.adapter.FindByID(ctx, order_id)
}
