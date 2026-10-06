// Package product adapts the Product service gRPC API into the shared
// pkg/database sqlc row types. Query and command are exposed as separate
// interfaces so consumers can depend on the narrower one.
package product

import (
	"context"

	pbproduct "github.com/MamangRust/monolith-ecommerce-pb/product"
	"github.com/MamangRust/monolith-ecommerce-pkg/adapter"
	"github.com/MamangRust/monolith-ecommerce-pkg/adapter/convert"
	db "github.com/MamangRust/monolith-ecommerce-pkg/database/schema"
	"github.com/MamangRust/monolith-ecommerce-pkg/resilience"
	product_errors "github.com/MamangRust/monolith-ecommerce-shared/errors/product_errors"
)

// QueryRepository reads products from the Product service over gRPC.
type QueryRepository interface {
	FindById(ctx context.Context, productID int) (*db.Product, error)
	FindByID(ctx context.Context, productID int) (*db.Product, error)
}

// CommandRepository mutates product stock via the Product service over gRPC.
type CommandRepository interface {
	UpdateProductCountStock(ctx context.Context, productID int, stock int) (*db.Product, error)
	AdjustProductStock(ctx context.Context, productID int, delta int, operationID string) (*db.Product, error)
}

// Repository implements both QueryRepository and CommandRepository over the
// generated product query/command clients.
type Repository struct {
	query   pbproduct.ProductQueryServiceClient
	command pbproduct.ProductCommandServiceClient
	guard   *resilience.DependencyGuard
}

// New builds a product adapter. Passing zero options leaves the guard nil, which
// makes DependencyGuard.Call a plain passthrough.
func New(query pbproduct.ProductQueryServiceClient, command pbproduct.ProductCommandServiceClient, opts ...adapter.GuardOption) *Repository {
	r := &Repository{query: query, command: command}
	for _, opt := range opts {
		opt(r)
	}
	return r
}

// NewQueryAdapter builds a read-only adapter for consumers that never write.
func NewQueryAdapter(query pbproduct.ProductQueryServiceClient, opts ...adapter.GuardOption) QueryRepository {
	return New(query, nil, opts...)
}

// NewCommandAdapter builds a write-only adapter for consumers that never read.
func NewCommandAdapter(command pbproduct.ProductCommandServiceClient, opts ...adapter.GuardOption) CommandRepository {
	return New(nil, command, opts...)
}

func (r *Repository) SetGuard(g *resilience.DependencyGuard) {
	r.guard = g
}

var (
	_ QueryRepository   = (*Repository)(nil)
	_ CommandRepository = (*Repository)(nil)
)

func (r *Repository) FindById(ctx context.Context, productID int) (*db.Product, error) {
	return r.findByID(ctx, productID)
}

func (r *Repository) FindByID(ctx context.Context, productID int) (*db.Product, error) {
	return r.findByID(ctx, productID)
}

func (r *Repository) findByID(ctx context.Context, productID int) (*db.Product, error) {
	resp, err := adapter.Call(r.guard, ctx, func(callCtx context.Context) (*pbproduct.ApiResponseProduct, error) {
		return r.query.FindById(callCtx, &pbproduct.FindByIdProductRequest{Id: int32(productID)})
	})
	if err != nil {
		return nil, product_errors.ErrProductNotFound.WithInternal(err)
	}
	if resp == nil || resp.Data == nil {
		return nil, product_errors.ErrProductNotFound
	}

	return mapProduct(resp.Data), nil
}

func (r *Repository) UpdateProductCountStock(ctx context.Context, productID int, stock int) (*db.Product, error) {
	resp, err := adapter.Call(r.guard, ctx, func(callCtx context.Context) (*pbproduct.ApiResponseProduct, error) {
		return r.command.UpdateProductCountStock(callCtx, &pbproduct.UpdateProductCountStockRequest{
			ProductId: int32(productID),
			Stock:     int32(stock),
		})
	})
	if err != nil {
		return nil, product_errors.ErrUpdateProductCountStock.WithInternal(err)
	}
	if resp == nil || resp.Data == nil {
		return nil, product_errors.ErrUpdateProductCountStock
	}

	return mapProduct(resp.Data), nil
}

// AdjustProductStock applies a signed delta to a product's stock through the
// owning service, which enforces the atomic `count_in_stock + delta >= 0` guard.
func (r *Repository) AdjustProductStock(ctx context.Context, productID int, delta int, operationID string) (*db.Product, error) {
	resp, err := adapter.Call(r.guard, ctx, func(callCtx context.Context) (*pbproduct.ApiResponseProduct, error) {
		return r.command.AdjustProductStock(callCtx, &pbproduct.AdjustProductStockRequest{
			ProductId:   int32(productID),
			Delta:       int32(delta),
			OperationId: operationID,
		})
	})
	if err != nil {
		return nil, product_errors.ErrUpdateProductCountStock.WithInternal(err)
	}
	if resp == nil || resp.Data == nil {
		return nil, product_errors.ErrUpdateProductCountStock
	}

	return mapProduct(resp.Data), nil
}

// mapProduct converts a proto product into the sqlc row type.
func mapProduct(p *pbproduct.ProductResponse) *db.Product {
	if p == nil {
		return nil
	}
	rating := float64(p.Rating)
	return &db.Product{
		ProductID:    p.Id,
		MerchantID:   p.MerchantId,
		CategoryID:   p.CategoryId,
		Name:         p.Name,
		Description:  convert.NullableString(p.Description),
		Price:        p.Price,
		CountInStock: p.CountInStock,
		Brand:        convert.NullableString(p.Brand),
		Weight:       convert.NullableInt32(p.Weight),
		Rating:       &rating,
		SlugProduct:  convert.NullableString(p.SlugProduct),
		ImageProduct: convert.NullableString(p.ImageProduct),
		CreatedAt:    convert.Timestamp(p.CreatedAt),
		UpdatedAt:    convert.Timestamp(p.UpdatedAt),
	}
}
