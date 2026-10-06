// Package merchant adapts the Merchant service gRPC API into the shared
// pkg/database sqlc row types used by the merchant-domain consumers.
package merchant

import (
	"context"

	pbmerchant "github.com/MamangRust/monolith-ecommerce-pb/merchant"
	"github.com/MamangRust/monolith-ecommerce-pkg/adapter"
	"github.com/MamangRust/monolith-ecommerce-pkg/adapter/convert"
	db "github.com/MamangRust/monolith-ecommerce-pkg/database/schema"
	"github.com/MamangRust/monolith-ecommerce-pkg/resilience"
	merchant_errors "github.com/MamangRust/monolith-ecommerce-shared/errors/merchant"
)

// QueryRepository is the contract consumers depend on for merchant reads.
type QueryRepository interface {
	FindByID(ctx context.Context, merchantID int) (*db.Merchant, error)
}

type repository struct {
	client pbmerchant.MerchantQueryServiceClient
	guard  *resilience.DependencyGuard
}

// New builds a merchant adapter. Passing zero options leaves the guard nil,
// which makes DependencyGuard.Call a plain passthrough.
func New(client pbmerchant.MerchantQueryServiceClient, opts ...adapter.GuardOption) *repository {
	r := &repository{client: client}
	for _, opt := range opts {
		opt(r)
	}
	return r
}

// NewQueryAdapter returns the adapter typed as the read-only contract.
func NewQueryAdapter(client pbmerchant.MerchantQueryServiceClient, opts ...adapter.GuardOption) QueryRepository {
	return New(client, opts...)
}

func (r *repository) SetGuard(g *resilience.DependencyGuard) {
	r.guard = g
}

var _ QueryRepository = (*repository)(nil)

func (r *repository) FindByID(ctx context.Context, merchantID int) (*db.Merchant, error) {
	resp, err := adapter.Call(r.guard, ctx, func(callCtx context.Context) (*pbmerchant.ApiResponseMerchant, error) {
		return r.client.FindById(callCtx, &pbmerchant.FindByIdMerchantRequest{Id: int32(merchantID)})
	})
	if err != nil {
		return nil, merchant_errors.ErrMerchantInternal.WithInternal(err)
	}
	if resp == nil || resp.Data == nil {
		return nil, merchant_errors.ErrMerchantNotFound
	}

	return mapMerchant(resp.Data), nil
}

// mapMerchant converts a proto merchant into the sqlc row type.
func mapMerchant(m *pbmerchant.MerchantResponse) *db.Merchant {
	if m == nil {
		return nil
	}
	return &db.Merchant{
		MerchantID:   m.Id,
		UserID:       m.UserId,
		Name:         m.Name,
		Description:  convert.NullableString(m.Description),
		Address:      convert.NullableString(m.Address),
		ContactEmail: convert.NullableString(m.ContactEmail),
		ContactPhone: convert.NullableString(m.ContactPhone),
		Status:       m.Status,
		CreatedAt:    convert.Timestamp(m.CreatedAt),
		UpdatedAt:    convert.Timestamp(m.UpdatedAt),
	}
}
