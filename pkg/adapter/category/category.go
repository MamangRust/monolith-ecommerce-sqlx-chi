// Package category adapts the Category service gRPC API into the shared
// pkg/database sqlc row types.
package category

import (
	"context"

	pbcategory "github.com/MamangRust/monolith-ecommerce-pb/category"
	"github.com/MamangRust/monolith-ecommerce-pkg/adapter"
	"github.com/MamangRust/monolith-ecommerce-pkg/adapter/convert"
	db "github.com/MamangRust/monolith-ecommerce-pkg/database/schema"
	"github.com/MamangRust/monolith-ecommerce-pkg/resilience"
	category_errors "github.com/MamangRust/monolith-ecommerce-shared/errors/category_errors"
)

// QueryRepository is the contract consumers depend on for category reads.
type QueryRepository interface {
	FindByID(ctx context.Context, categoryID int) (*db.Category, error)
}

type repository struct {
	client pbcategory.CategoryQueryServiceClient
	guard  *resilience.DependencyGuard
}

// New builds a category adapter. Passing zero options leaves the guard nil,
// which makes DependencyGuard.Call a plain passthrough.
func New(client pbcategory.CategoryQueryServiceClient, opts ...adapter.GuardOption) *repository {
	r := &repository{client: client}
	for _, opt := range opts {
		opt(r)
	}
	return r
}

// NewQueryAdapter returns the adapter typed as the read-only contract.
func NewQueryAdapter(client pbcategory.CategoryQueryServiceClient, opts ...adapter.GuardOption) QueryRepository {
	return New(client, opts...)
}

func (r *repository) SetGuard(g *resilience.DependencyGuard) {
	r.guard = g
}

var _ QueryRepository = (*repository)(nil)

func (r *repository) FindByID(ctx context.Context, categoryID int) (*db.Category, error) {
	resp, err := adapter.Call(r.guard, ctx, func(callCtx context.Context) (*pbcategory.ApiResponseCategory, error) {
		return r.client.FindById(callCtx, &pbcategory.FindByIdCategoryRequest{Id: int32(categoryID)})
	})
	if err != nil {
		return nil, category_errors.ErrFindCategoryById.WithInternal(err)
	}
	if resp == nil || resp.Data == nil {
		return nil, category_errors.ErrFindCategoryById
	}

	return mapCategory(resp.Data), nil
}

// mapCategory converts a proto category into the sqlc row type.
func mapCategory(c *pbcategory.CategoryResponse) *db.Category {
	if c == nil {
		return nil
	}
	return &db.Category{
		CategoryID:    c.Id,
		Name:          c.Name,
		Description:   convert.NullableString(c.Description),
		SlugCategory:  convert.NullableString(c.SlugCategory),
		ImageCategory: convert.NullableString(c.ImageCategory),
		CreatedAt:     convert.Timestamp(c.CreatedAt),
		UpdatedAt:     convert.Timestamp(c.UpdatedAt),
	}
}
