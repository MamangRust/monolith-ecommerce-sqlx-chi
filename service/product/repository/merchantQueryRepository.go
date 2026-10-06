package repository

import (
	"context"

	merchantadapter "github.com/MamangRust/monolith-ecommerce-pkg/adapter/merchant"
	db "github.com/MamangRust/monolith-ecommerce-pkg/database/schema"
	merchant_errors "github.com/MamangRust/monolith-ecommerce-shared/errors/merchant"
)

// merchantQueryRepository delegates merchant lookups to the shared gRPC adapter
// and re-wraps failures in this service's own error so callers keep the status
// codes they had before the adapter migration.
type merchantQueryRepository struct {
	adapter merchantadapter.QueryRepository
}

func NewMerchantQueryRepository(adapter merchantadapter.QueryRepository) *merchantQueryRepository {
	return &merchantQueryRepository{
		adapter: adapter,
	}
}

func (r *merchantQueryRepository) FindByID(ctx context.Context, merchant_id int) (*db.Merchant, error) {
	res, err := r.adapter.FindByID(ctx, merchant_id)
	if err != nil {
		return nil, merchant_errors.ErrMerchantInternal.WithInternal(err)
	}

	return res, nil
}
