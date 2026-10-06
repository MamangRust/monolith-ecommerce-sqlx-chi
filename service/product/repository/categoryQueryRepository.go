package repository

import (
	"context"

	categoryadapter "github.com/MamangRust/monolith-ecommerce-pkg/adapter/category"
	db "github.com/MamangRust/monolith-ecommerce-pkg/database/schema"
	category_errors "github.com/MamangRust/monolith-ecommerce-shared/errors/category_errors"
)

// categoryQueryRepository delegates category lookups to the shared gRPC adapter
// and re-wraps failures in this service's own error so callers keep the status
// codes they had before the adapter migration.
type categoryQueryRepository struct {
	adapter categoryadapter.QueryRepository
}

func NewCategoryQueryRepository(adapter categoryadapter.QueryRepository) *categoryQueryRepository {
	return &categoryQueryRepository{
		adapter: adapter,
	}
}

func (r *categoryQueryRepository) FindByID(ctx context.Context, category_id int) (*db.Category, error) {
	res, err := r.adapter.FindByID(ctx, category_id)
	if err != nil {
		return nil, category_errors.ErrFindCategoryById.WithInternal(err)
	}

	return res, nil
}
