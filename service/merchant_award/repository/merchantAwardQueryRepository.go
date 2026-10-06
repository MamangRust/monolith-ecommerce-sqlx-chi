package repository

import (
	"context"
	"database/sql"
	"errors"

	db "github.com/MamangRust/monolith-ecommerce-pkg/database/schema"
	"github.com/MamangRust/monolith-ecommerce-shared/domain/requests"
	merchantaward_errors "github.com/MamangRust/monolith-ecommerce-shared/errors/merchant_award"
	"github.com/jmoiron/sqlx"
)

const getMerchantCertificationsAndAwards = `SELECT
    mca.merchant_certification_id,
    mca.merchant_id,
    mca.title,
    mca.description,
    mca.issued_by,
    mca.issue_date,
    mca.expiry_date,
    mca.certificate_url,
    mca.created_at,
    mca.updated_at,
    m.name AS merchant_name,
    COUNT(*) OVER () AS total_count
FROM
    merchant_certifications_and_awards mca
    JOIN merchants m ON mca.merchant_id = m.merchant_id
WHERE
    mca.deleted_at IS NULL
    AND LOWER(m.name) LIKE LOWER(CONCAT('%', $1::text, '%'))
LIMIT $2
OFFSET
    $3`

const getMerchantCertificationsAndAwardsActive = `SELECT
    mca.merchant_certification_id,
    mca.merchant_id,
    mca.title,
    mca.description,
    mca.issued_by,
    mca.issue_date,
    mca.expiry_date,
    mca.certificate_url,
    mca.created_at,
    mca.updated_at,
    mca.deleted_at,
    m.name AS merchant_name,
    COUNT(*) OVER () AS total_count
FROM
    merchant_certifications_and_awards mca
    JOIN merchants m ON mca.merchant_id = m.merchant_id
WHERE
    mca.deleted_at IS NULL
    AND m.deleted_at IS NULL
    AND LOWER(m.name) LIKE LOWER(CONCAT('%', $1::text, '%'))
LIMIT $2
OFFSET
    $3`

const getMerchantCertificationsAndAwardsTrashed = `SELECT
    mca.merchant_certification_id,
    mca.merchant_id,
    mca.title,
    mca.description,
    mca.issued_by,
    mca.issue_date,
    mca.expiry_date,
    mca.certificate_url,
    mca.created_at,
    mca.updated_at,
    mca.deleted_at,
    m.name AS merchant_name,
    COUNT(*) OVER () AS total_count
FROM
    merchant_certifications_and_awards mca
    JOIN merchants m ON mca.merchant_id = m.merchant_id
WHERE
    mca.deleted_at IS NOT NULL
    AND LOWER(m.name) LIKE LOWER(CONCAT('%', $1::text, '%'))
LIMIT $2
OFFSET
    $3`

const getMerchantCertificationOrAward = `SELECT mca.merchant_certification_id, mca.merchant_id, mca.title, mca.description, mca.issued_by, mca.issue_date, mca.expiry_date, mca.certificate_url, mca.created_at, mca.updated_at
FROM
    merchant_certifications_and_awards mca
WHERE
    merchant_certification_id = $1
    AND deleted_at IS NULL`

type merchantAwardQueryRepository struct {
	db *sqlx.DB
}

func NewMerchantAwardQueryRepository(db *sqlx.DB) *merchantAwardQueryRepository {
	return &merchantAwardQueryRepository{
		db: db,
	}
}

func (r *merchantAwardQueryRepository) FindAll(ctx context.Context, req *requests.FindAllMerchant) ([]*db.GetMerchantCertificationsAndAwardsRow, error) {
	offset := (req.Page - 1) * req.PageSize

	var awards []*db.GetMerchantCertificationsAndAwardsRow
	err := r.db.SelectContext(ctx, &awards, getMerchantCertificationsAndAwards,
		req.Search,
		int32(req.PageSize),
		int32(offset),
	)

	if err != nil {
		return nil, merchantaward_errors.ErrFindAllMerchantAwards.WithInternal(err)
	}

	return awards, nil
}

func (r *merchantAwardQueryRepository) FindActive(ctx context.Context, req *requests.FindAllMerchant) ([]*db.GetMerchantCertificationsAndAwardsActiveRow, error) {
	offset := (req.Page - 1) * req.PageSize

	var awards []*db.GetMerchantCertificationsAndAwardsActiveRow
	err := r.db.SelectContext(ctx, &awards, getMerchantCertificationsAndAwardsActive,
		req.Search,
		int32(req.PageSize),
		int32(offset),
	)

	if err != nil {
		return nil, merchantaward_errors.ErrFindByActiveMerchantAwards.WithInternal(err)
	}

	return awards, nil
}

func (r *merchantAwardQueryRepository) FindTrashed(ctx context.Context, req *requests.FindAllMerchant) ([]*db.GetMerchantCertificationsAndAwardsTrashedRow, error) {
	offset := (req.Page - 1) * req.PageSize

	var awards []*db.GetMerchantCertificationsAndAwardsTrashedRow
	err := r.db.SelectContext(ctx, &awards, getMerchantCertificationsAndAwardsTrashed,
		req.Search,
		int32(req.PageSize),
		int32(offset),
	)

	if err != nil {
		return nil, merchantaward_errors.ErrFindByTrashedMerchantAwards.WithInternal(err)
	}

	return awards, nil
}

func (r *merchantAwardQueryRepository) FindByID(ctx context.Context, user_id int) (*db.GetMerchantCertificationOrAwardRow, error) {
	var award db.GetMerchantCertificationOrAwardRow
	err := r.db.GetContext(ctx, &award, getMerchantCertificationOrAward, int32(user_id))

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, merchantaward_errors.ErrMerchantAwardNotFound.WithInternal(err)
		}
		return nil, merchantaward_errors.ErrFindByIdMerchantAward.WithInternal(err)
	}

	return &award, nil
}
