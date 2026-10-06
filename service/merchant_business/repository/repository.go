package repository

import (
	pbmerchant "github.com/MamangRust/monolith-ecommerce-pb/merchant"
	"github.com/MamangRust/monolith-ecommerce-pkg/adapter"
	merchantadapter "github.com/MamangRust/monolith-ecommerce-pkg/adapter/merchant"
	"github.com/jmoiron/sqlx"
)

type GuardOptions struct {
	Merchant []adapter.GuardOption
}

type Repositories struct {
	MerchantBusinessQuery   MerchantBusinessQueryRepository
	MerchantBusinessCommand MerchantBusinessCommandRepository
	MerchantQuery           MerchantQueryRepository
}

func NewRepositories(db *sqlx.DB, merchantQueryClient pbmerchant.MerchantQueryServiceClient, guards ...GuardOptions) *Repositories {
	var g GuardOptions
	if len(guards) > 0 {
		g = guards[0]
	}

	return &Repositories{
		MerchantBusinessQuery:   NewMerchantBusinessQueryRepository(db),
		MerchantBusinessCommand: NewMerchantBusinessCommandRepository(db),
		MerchantQuery:           merchantadapter.New(merchantQueryClient, g.Merchant...),
	}
}
