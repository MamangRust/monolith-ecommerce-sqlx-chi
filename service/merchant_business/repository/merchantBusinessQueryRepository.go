package repository

import (
	"context"
	"database/sql"
	"errors"

	db "github.com/MamangRust/monolith-ecommerce-pkg/database/schema"
	"github.com/MamangRust/monolith-ecommerce-shared/domain/requests"
	merchantbusiness_errors "github.com/MamangRust/monolith-ecommerce-shared/errors/merchant_business"
	"github.com/jmoiron/sqlx"
)

const getMerchantsBusinessInformation = `SELECT
    mbi.merchant_business_info_id,
    mbi.merchant_id,
    mbi.business_type,
    mbi.tax_id,
    mbi.established_year,
    mbi.number_of_employees,
    mbi.website_url,
    mbi.created_at,
    mbi.updated_at,
    m.name AS merchant_name,
    COUNT(*) OVER () AS total_count
FROM
    merchant_business_information mbi
    JOIN merchants m ON mbi.merchant_id = m.merchant_id
WHERE
    mbi.deleted_at IS NULL
    AND LOWER(m.name) LIKE LOWER(CONCAT('%', $1::text, '%'))
LIMIT $2
OFFSET
    $3`

const getMerchantsBusinessInformationActive = `SELECT
    mbi.merchant_business_info_id,
    mbi.merchant_id,
    mbi.business_type,
    mbi.tax_id,
    mbi.established_year,
    mbi.number_of_employees,
    mbi.website_url,
    mbi.created_at,
    mbi.updated_at,
    mbi.deleted_at,
    m.name AS merchant_name,
    COUNT(*) OVER () AS total_count
FROM
    merchant_business_information mbi
    JOIN merchants m ON mbi.merchant_id = m.merchant_id
WHERE
    mbi.deleted_at IS NULL
    AND m.deleted_at IS NULL
    AND LOWER(m.name) LIKE LOWER(CONCAT('%', $1::text, '%'))
LIMIT $2
OFFSET
    $3`

const getMerchantsBusinessInformationTrashed = `SELECT
    mbi.merchant_business_info_id,
    mbi.merchant_id,
    mbi.business_type,
    mbi.tax_id,
    mbi.established_year,
    mbi.number_of_employees,
    mbi.website_url,
    mbi.created_at,
    mbi.updated_at,
    mbi.deleted_at,
    m.name AS merchant_name,
    COUNT(*) OVER () AS total_count
FROM
    merchant_business_information mbi
    JOIN merchants m ON mbi.merchant_id = m.merchant_id
WHERE
    mbi.deleted_at IS NOT NULL
    AND LOWER(m.name) LIKE LOWER(CONCAT('%', $1::text, '%'))
LIMIT $2
OFFSET
    $3`

const getMerchantBusinessInformation = `SELECT mbi.merchant_business_info_id, mbi.merchant_id, mbi.business_type, mbi.tax_id, mbi.established_year, mbi.number_of_employees, mbi.website_url, mbi.created_at, mbi.updated_at
FROM
    merchant_business_information mbi
WHERE
    merchant_business_info_id = $1
    AND deleted_at IS NULL`

type merchantBusinessQueryRepository struct {
	db *sqlx.DB
}

func NewMerchantBusinessQueryRepository(
	db *sqlx.DB,
) *merchantBusinessQueryRepository {
	return &merchantBusinessQueryRepository{
		db: db,
	}
}

func (r *merchantBusinessQueryRepository) FindAll(ctx context.Context, req *requests.FindAllMerchant) ([]*db.GetMerchantsBusinessInformationRow, error) {
	offset := (req.Page - 1) * req.PageSize

	var businesses []*db.GetMerchantsBusinessInformationRow
	err := r.db.SelectContext(ctx, &businesses, getMerchantsBusinessInformation,
		req.Search,
		int32(req.PageSize),
		int32(offset),
	)

	if err != nil {
		return nil, merchantbusiness_errors.ErrFindAllMerchantBusinesses.WithInternal(err)
	}

	return businesses, nil
}

func (r *merchantBusinessQueryRepository) FindActive(ctx context.Context, req *requests.FindAllMerchant) ([]*db.GetMerchantsBusinessInformationActiveRow, error) {
	offset := (req.Page - 1) * req.PageSize

	var businesses []*db.GetMerchantsBusinessInformationActiveRow
	err := r.db.SelectContext(ctx, &businesses, getMerchantsBusinessInformationActive,
		req.Search,
		int32(req.PageSize),
		int32(offset),
	)

	if err != nil {
		return nil, merchantbusiness_errors.ErrFindActiveMerchantBusinesses.WithInternal(err)
	}

	return businesses, nil
}

func (r *merchantBusinessQueryRepository) FindTrashed(ctx context.Context, req *requests.FindAllMerchant) ([]*db.GetMerchantsBusinessInformationTrashedRow, error) {
	offset := (req.Page - 1) * req.PageSize

	var businesses []*db.GetMerchantsBusinessInformationTrashedRow
	err := r.db.SelectContext(ctx, &businesses, getMerchantsBusinessInformationTrashed,
		req.Search,
		int32(req.PageSize),
		int32(offset),
	)

	if err != nil {
		return nil, merchantbusiness_errors.ErrFindTrashedMerchantBusinesses.WithInternal(err)
	}

	return businesses, nil
}

func (r *merchantBusinessQueryRepository) FindByID(ctx context.Context, user_id int) (*db.GetMerchantBusinessInformationRow, error) {
	var business db.GetMerchantBusinessInformationRow
	err := r.db.GetContext(ctx, &business, getMerchantBusinessInformation, int32(user_id))

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, merchantbusiness_errors.ErrMerchantBusinessNotFound
		}
		return nil, merchantbusiness_errors.ErrMerchantBusinessInternal.WithInternal(err)
	}

	return &business, nil
}
