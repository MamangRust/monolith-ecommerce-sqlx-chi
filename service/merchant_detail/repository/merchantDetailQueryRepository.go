package repository

import (
	"context"
	"database/sql"
	"errors"

	db "github.com/MamangRust/monolith-ecommerce-pkg/database/schema"
	"github.com/MamangRust/monolith-ecommerce-shared/domain/requests"
	merchantdetail_errors "github.com/MamangRust/monolith-ecommerce-shared/errors/merchant_detail"
	"github.com/jmoiron/sqlx"
)

type merchantDetailQueryRepository struct {
	db *sqlx.DB
}

func NewMerchantDetailQueryRepository(db *sqlx.DB) *merchantDetailQueryRepository {
	return &merchantDetailQueryRepository{
		db: db,
	}
}

const getMerchantDetails = `SELECT
    md.merchant_detail_id,
    md.merchant_id,
    md.display_name,
    md.cover_image_url,
    md.logo_url,
    md.short_description,
    md.website_url,
    md.created_at,
    md.updated_at,
    m.name AS merchant_name,
    COUNT(*) OVER () AS total_count,
    json_agg(
        json_build_object(
            'id',
            sml.merchant_social_id,
            'platform',
            sml.platform,
            'url',
            sml.url
        )
    ) AS social_media_links
FROM
    merchant_details md
    JOIN merchants m ON md.merchant_id = m.merchant_id
    LEFT JOIN merchant_social_media_links sml ON sml.merchant_detail_id = md.merchant_detail_id
WHERE
    md.deleted_at IS NULL
    AND LOWER(m.name) LIKE LOWER(CONCAT('%', $1::text, '%'))
GROUP BY
    md.merchant_detail_id,
    m.merchant_id
LIMIT $2
OFFSET
    $3`

const getMerchantDetailsActive = `SELECT
    md.merchant_detail_id,
    md.merchant_id,
    md.display_name,
    md.cover_image_url,
    md.logo_url,
    md.short_description,
    md.website_url,
    md.created_at,
    md.updated_at,
    md.deleted_at,
    m.name AS merchant_name,
    COUNT(*) OVER () AS total_count,
    json_agg(
        json_build_object(
            'id',
            sml.merchant_social_id,
            'platform',
            sml.platform,
            'url',
            sml.url
        )
    ) AS social_media_links
FROM
    merchant_details md
    JOIN merchants m ON md.merchant_id = m.merchant_id
    LEFT JOIN merchant_social_media_links sml ON sml.merchant_detail_id = md.merchant_detail_id
WHERE
    md.deleted_at IS NULL
    AND m.deleted_at IS NULL
    AND LOWER(m.name) LIKE LOWER(CONCAT('%', $1::text, '%'))
GROUP BY
    md.merchant_detail_id,
    m.merchant_id
LIMIT $2
OFFSET
    $3`

const getMerchantDetailsTrashed = `SELECT
    md.merchant_detail_id,
    md.merchant_id,
    md.display_name,
    md.cover_image_url,
    md.logo_url,
    md.short_description,
    md.website_url,
    md.created_at,
    md.updated_at,
    md.deleted_at,
    m.name AS merchant_name,
    COUNT(*) OVER () AS total_count,
    json_agg(
        json_build_object(
            'id',
            sml.merchant_social_id,
            'platform',
            sml.platform,
            'url',
            sml.url
        )
    ) AS social_media_links
FROM
    merchant_details md
    JOIN merchants m ON md.merchant_id = m.merchant_id
    LEFT JOIN merchant_social_media_links sml ON sml.merchant_detail_id = md.merchant_detail_id
WHERE
    md.deleted_at IS NOT NULL
    AND LOWER(m.name) LIKE LOWER(CONCAT('%', $1::text, '%'))
GROUP BY
    md.merchant_detail_id,
    m.merchant_id
LIMIT $2
OFFSET
    $3`

const getMerchantDetail = `SELECT
    md.merchant_detail_id,
    md.merchant_id,
    md.display_name,
    md.cover_image_url,
    md.logo_url,
    md.short_description,
    md.website_url,
    md.created_at,
    md.updated_at,
    m.name AS merchant_name,
    json_agg(
        json_build_object(
            'id',
            sml.merchant_social_id,
            'platform',
            sml.platform,
            'url',
            sml.url
        )
    ) AS social_media_links
FROM
    merchant_details md
    JOIN merchants m ON md.merchant_id = m.merchant_id
    LEFT JOIN merchant_social_media_links sml ON sml.merchant_detail_id = md.merchant_detail_id
WHERE
    md.merchant_detail_id = $1
    AND md.deleted_at IS NULL
GROUP BY
    md.merchant_detail_id,
    m.merchant_id`

const getMerchantDetailTrashed = `SELECT md.merchant_detail_id, md.merchant_id, md.display_name, md.cover_image_url, md.logo_url, md.short_description, md.website_url, md.created_at, md.updated_at
FROM merchant_details md
WHERE
    merchant_detail_id = $1
    AND deleted_at IS NOT NULL`

func (r *merchantDetailQueryRepository) FindAll(ctx context.Context, req *requests.FindAllMerchant) ([]*db.GetMerchantDetailsRow, error) {
	offset := (req.Page - 1) * req.PageSize

	var details []*db.GetMerchantDetailsRow
	err := r.db.SelectContext(ctx, &details, getMerchantDetails,
		req.Search,
		int32(req.PageSize),
		int32(offset),
	)
	if err != nil {
		return nil, merchantdetail_errors.ErrFindAllMerchantDetails.WithInternal(err)
	}

	return details, nil
}

func (r *merchantDetailQueryRepository) FindActive(ctx context.Context, req *requests.FindAllMerchant) ([]*db.GetMerchantDetailsActiveRow, error) {
	offset := (req.Page - 1) * req.PageSize

	var details []*db.GetMerchantDetailsActiveRow
	err := r.db.SelectContext(ctx, &details, getMerchantDetailsActive,
		req.Search,
		int32(req.PageSize),
		int32(offset),
	)
	if err != nil {
		return nil, merchantdetail_errors.ErrFindActiveMerchantDetails.WithInternal(err)
	}

	return details, nil
}

func (r *merchantDetailQueryRepository) FindTrashed(ctx context.Context, req *requests.FindAllMerchant) ([]*db.GetMerchantDetailsTrashedRow, error) {
	offset := (req.Page - 1) * req.PageSize

	var details []*db.GetMerchantDetailsTrashedRow
	err := r.db.SelectContext(ctx, &details, getMerchantDetailsTrashed,
		req.Search,
		int32(req.PageSize),
		int32(offset),
	)
	if err != nil {
		return nil, merchantdetail_errors.ErrFindTrashedMerchantDetails.WithInternal(err)
	}

	return details, nil
}

func (r *merchantDetailQueryRepository) FindByID(ctx context.Context, user_id int) (*db.GetMerchantDetailRow, error) {
	var detail db.GetMerchantDetailRow
	if err := r.db.GetContext(ctx, &detail, getMerchantDetail, int32(user_id)); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, merchantdetail_errors.ErrMerchantDetailNotFound.WithInternal(err)
		}
		return nil, merchantdetail_errors.ErrMerchantDetailInternal.WithInternal(err)
	}

	return &detail, nil
}

func (r *merchantDetailQueryRepository) FindByIDTrashed(ctx context.Context, user_id int) (*db.GetMerchantDetailTrashedRow, error) {
	var detail db.GetMerchantDetailTrashedRow
	if err := r.db.GetContext(ctx, &detail, getMerchantDetailTrashed, int32(user_id)); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, merchantdetail_errors.ErrMerchantDetailNotFound.WithInternal(err)
		}
		return nil, merchantdetail_errors.ErrFindByIdTrashedMerchantDetail.WithInternal(err)
	}

	return &detail, nil
}
