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
	MerchantPoliciesQuery   MerchantPoliciesQueryRepository
	MerchantPoliciesCommand MerchantPoliciesCommandRepository
	MerchantQuery           MerchantQueryRepository
}

func NewRepositories(db *sqlx.DB, merchantQueryClient pbmerchant.MerchantQueryServiceClient, guards ...GuardOptions) *Repositories {
	var g GuardOptions
	if len(guards) > 0 {
		g = guards[0]
	}

	return &Repositories{
		MerchantPoliciesQuery:   NewMerchantPolicyQueryRepository(db),
		MerchantPoliciesCommand: NewMerchantPolicyCommandRepository(db),
		MerchantQuery:           merchantadapter.New(merchantQueryClient, g.Merchant...),
	}
}
