package repository

import (
	"context"

	shippingaddressadapter "github.com/MamangRust/monolith-ecommerce-pkg/adapter/shipping_address"
	db "github.com/MamangRust/monolith-ecommerce-pkg/database/schema"
)

// shippingAddressQueryRepository delegates shipping-address lookups to the
// shared gRPC adapter. Failures keep propagating the dependency's raw gRPC
// status, as they did before the adapter migration.
//
// FindByOrder is the operation this service actually needs: it resolves the
// shipping address belonging to an order. The old method was misnamed FindByID
// while calling the FindByOrder RPC underneath.
type shippingAddressQueryRepository struct {
	adapter shippingaddressadapter.QueryRepository
}

func NewShippingAddressQueryRepository(adapter shippingaddressadapter.QueryRepository) *shippingAddressQueryRepository {
	return &shippingAddressQueryRepository{
		adapter: adapter,
	}
}

func (r *shippingAddressQueryRepository) FindByOrder(ctx context.Context, order_id int) (*db.ShippingAddress, error) {
	return r.adapter.FindByOrder(ctx, order_id)
}
