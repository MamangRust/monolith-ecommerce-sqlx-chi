package repository

import (
	"context"
	"database/sql"
	"errors"

	db "github.com/MamangRust/monolith-ecommerce-pkg/database/schema"
	"github.com/MamangRust/monolith-ecommerce-shared/domain/requests"
	merchant_policy_errors "github.com/MamangRust/monolith-ecommerce-shared/errors/merchant_policy_errors"
	"github.com/jmoiron/sqlx"
)

const getMerchantPolicies = `SELECT
    mp.merchant_policy_id,
    mp.merchant_id,
    mp.policy_type,
    mp.title,
    mp.description,
    mp.created_at,
    mp.updated_at,
    m.name AS merchant_name,
    COUNT(*) OVER () AS total_count
FROM
    merchant_policies mp
    JOIN merchants m ON mp.merchant_id = m.merchant_id
WHERE
    mp.deleted_at IS NULL
    AND LOWER(m.name) LIKE LOWER(CONCAT('%', $1::text, '%'))
LIMIT $2
OFFSET
    $3`

const getMerchantPoliciesActive = `SELECT
    mp.merchant_policy_id,
    mp.merchant_id,
    mp.policy_type,
    mp.title,
    mp.description,
    mp.created_at,
    mp.updated_at,
    mp.deleted_at,
    m.name AS merchant_name,
    COUNT(*) OVER () AS total_count
FROM
    merchant_policies mp
    JOIN merchants m ON mp.merchant_id = m.merchant_id
WHERE
    mp.deleted_at IS NULL
    AND m.deleted_at IS NULL
    AND LOWER(m.name) LIKE LOWER(CONCAT('%', $1::text, '%'))
LIMIT $2
OFFSET
    $3`

const getMerchantPoliciesTrashed = `SELECT
    mp.merchant_policy_id,
    mp.merchant_id,
    mp.policy_type,
    mp.title,
    mp.description,
    mp.created_at,
    mp.updated_at,
    mp.deleted_at,
    m.name AS merchant_name,
    COUNT(*) OVER () AS total_count
FROM
    merchant_policies mp
    JOIN merchants m ON mp.merchant_id = m.merchant_id
WHERE
    mp.deleted_at IS NOT NULL
    AND LOWER(m.name) LIKE LOWER(CONCAT('%', $1::text, '%'))
LIMIT $2
OFFSET
    $3`

const getMerchantPolicy = `SELECT mp.merchant_policy_id, mp.merchant_id, mp.policy_type, mp.title, mp.description, mp.created_at, mp.updated_at
FROM merchant_policies mp
WHERE
    merchant_policy_id = $1
    AND deleted_at IS NULL`

type merchantPolicyQueryRepository struct {
	db *sqlx.DB
}

func NewMerchantPolicyQueryRepository(db *sqlx.DB) *merchantPolicyQueryRepository {
	return &merchantPolicyQueryRepository{
		db: db,
	}
}

func (r *merchantPolicyQueryRepository) FindAll(ctx context.Context, req *requests.FindAllMerchant) ([]*db.GetMerchantPoliciesRow, error) {
	offset := (req.Page - 1) * req.PageSize

	var policies []*db.GetMerchantPoliciesRow
	err := r.db.SelectContext(ctx, &policies, getMerchantPolicies,
		req.Search,
		int32(req.PageSize),
		int32(offset),
	)

	if err != nil {
		return nil, merchant_policy_errors.ErrFindAllMerchantPolicies.WithInternal(err)
	}

	return policies, nil
}

func (r *merchantPolicyQueryRepository) FindActive(ctx context.Context, req *requests.FindAllMerchant) ([]*db.GetMerchantPoliciesActiveRow, error) {
	offset := (req.Page - 1) * req.PageSize

	var policies []*db.GetMerchantPoliciesActiveRow
	err := r.db.SelectContext(ctx, &policies, getMerchantPoliciesActive,
		req.Search,
		int32(req.PageSize),
		int32(offset),
	)

	if err != nil {
		return nil, merchant_policy_errors.ErrFindActiveMerchantPolicies.WithInternal(err)
	}

	return policies, nil
}

func (r *merchantPolicyQueryRepository) FindTrashed(ctx context.Context, req *requests.FindAllMerchant) ([]*db.GetMerchantPoliciesTrashedRow, error) {
	offset := (req.Page - 1) * req.PageSize

	var policies []*db.GetMerchantPoliciesTrashedRow
	err := r.db.SelectContext(ctx, &policies, getMerchantPoliciesTrashed,
		req.Search,
		int32(req.PageSize),
		int32(offset),
	)

	if err != nil {
		return nil, merchant_policy_errors.ErrFindTrashedMerchantPolicies.WithInternal(err)
	}

	return policies, nil
}

func (r *merchantPolicyQueryRepository) FindByID(ctx context.Context, user_id int) (*db.GetMerchantPolicyRow, error) {
	var policy db.GetMerchantPolicyRow
	err := r.db.GetContext(ctx, &policy, getMerchantPolicy, int32(user_id))

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, merchant_policy_errors.ErrMerchantPolicyNotFound
		}
		return nil, merchant_policy_errors.ErrFindMerchantPolicyByID.WithInternal(err)
	}

	return &policy, nil
}
