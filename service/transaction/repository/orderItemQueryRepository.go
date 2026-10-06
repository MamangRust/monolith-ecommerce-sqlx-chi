package repository

import (
	"context"

	orderitemadapter "github.com/MamangRust/monolith-ecommerce-pkg/adapter/order_item"
	db "github.com/MamangRust/monolith-ecommerce-pkg/database/schema"
)

// orderItemRepository delegates order-item lookups to the shared gRPC adapter.
// Failures keep propagating the dependency's raw gRPC status, as they did
// before the adapter migration.
type orderItemRepository struct {
	adapter orderitemadapter.QueryRepository
}

func NewOrderItemRepository(adapter orderitemadapter.QueryRepository) *orderItemRepository {
	return &orderItemRepository{
		adapter: adapter,
	}
}

func (r *orderItemRepository) FindOrderItemByOrder(ctx context.Context, order_id int) ([]*db.OrderItem, error) {
	return r.adapter.FindOrderItemByOrder(ctx, order_id)
}
