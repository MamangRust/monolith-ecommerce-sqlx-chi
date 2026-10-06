package repository

import (
	"context"
	"database/sql"
	"errors"

	db "github.com/MamangRust/monolith-ecommerce-pkg/database/schema"
	"github.com/MamangRust/monolith-ecommerce-shared/domain/requests"
	review_errors "github.com/MamangRust/monolith-ecommerce-shared/errors/review"
	"github.com/jmoiron/sqlx"
)

const getReviews = `SELECT
    review_id,
    user_id,
    product_id,
    name,
    comment,
    rating,
    created_at,
    updated_at,
    COUNT(*) OVER () AS total_count
FROM reviews
WHERE
    deleted_at IS NULL
    AND (
        $1::TEXT IS NULL
        OR review_id::TEXT ILIKE '%' || $1 || '%'
        OR name ILIKE '%' || $1 || '%'
    )
ORDER BY created_at DESC
LIMIT $2
OFFSET
    $3`

const getReviewByProductId = `SELECT
    r.review_id,
    r.user_id,
    r.product_id,
    r.name,
    r.comment,
    r.rating,
    r.created_at,
    r.updated_at,
    r.deleted_at,
    COUNT(*) OVER () AS total_count,
    COALESCE(
        (
            SELECT json_agg(
                    jsonb_build_object(
                        'detail_id', rd.review_detail_id, 'type', rd.type, 'url', rd.url, 'caption', rd.caption, 'created_at', rd.created_at
                    )
                )
            FROM review_details rd
            WHERE
                rd.review_id = r.review_id
        ),
        '[]'
    ) AS review_details
FROM reviews r
WHERE
    r.deleted_at IS NULL
    AND r.product_id = $1
    AND (
        $2::INT IS NULL
        OR r.rating = $2
    )
ORDER BY r.created_at DESC
LIMIT $3
OFFSET
    $4`

const getReviewByMerchantId = `SELECT
    r.review_id,
    r.user_id,
    r.product_id,
    r.name,
    r.comment,
    r.rating,
    r.created_at,
    r.updated_at,
    r.deleted_at,
    COUNT(*) OVER () AS total_count,
    COALESCE(
        (
            SELECT json_agg(
                    jsonb_build_object(
                        'detail_id', rd.review_detail_id, 'type', rd.type, 'url', rd.url, 'caption', rd.caption, 'created_at', rd.created_at
                    )
                )
            FROM review_details rd
            WHERE
                rd.review_id = r.review_id
        ),
        '[]'
    ) AS review_details
FROM reviews r
    JOIN products p ON r.product_id = p.product_id
WHERE
    r.deleted_at IS NULL
    AND p.merchant_id = $1
    AND (
        $2::INT IS NULL
        OR r.rating = $2
    )
ORDER BY r.created_at DESC
LIMIT $3
OFFSET
    $4`

const getReviewsActive = `SELECT
    review_id,
    user_id,
    product_id,
    name,
    comment,
    rating,
    created_at,
    updated_at,
    deleted_at,
    COUNT(*) OVER () AS total_count
FROM reviews
WHERE
    deleted_at IS NULL
    AND (
        $1::TEXT IS NULL
        OR review_id::TEXT ILIKE '%' || $1 || '%'
        OR name ILIKE '%' || $1 || '%'
    )
ORDER BY created_at DESC
LIMIT $2
OFFSET
    $3`

const getReviewsTrashed = `SELECT
    review_id,
    user_id,
    product_id,
    name,
    comment,
    rating,
    created_at,
    updated_at,
    deleted_at,
    COUNT(*) OVER () AS total_count
FROM reviews
WHERE
    deleted_at IS NOT NULL
    AND (
        $1::TEXT IS NULL
        OR review_id::TEXT ILIKE '%' || $1 || '%'
        OR name ILIKE '%' || $1 || '%'
    )
ORDER BY created_at DESC
LIMIT $2
OFFSET
    $3`

const getReviewByID = `SELECT
    review_id,
    user_id,
    product_id,
    name,
    comment,
    rating,
    created_at,
    updated_at
FROM reviews
WHERE
    review_id = $1
    AND deleted_at IS NULL`

type reviewQueryRepository struct {
	db *sqlx.DB
}

func NewReviewQueryRepository(db *sqlx.DB) *reviewQueryRepository {
	return &reviewQueryRepository{
		db: db,
	}
}

func (r *reviewQueryRepository) FindAll(ctx context.Context, req *requests.FindAllReview) ([]*db.GetReviewsRow, error) {
	offset := (req.Page - 1) * req.PageSize

	var res []*db.GetReviewsRow
	err := r.db.SelectContext(ctx, &res, getReviews,
		req.Search,
		int32(req.PageSize),
		int32(offset),
	)

	if err != nil {
		return nil, review_errors.ErrFindAllReviews.WithInternal(err)
	}

	return res, nil
}

func (r *reviewQueryRepository) FindByProduct(ctx context.Context, req *requests.FindAllReviewByProduct) ([]*db.GetReviewByProductIdRow, error) {
	offset := (req.Page - 1) * req.PageSize

	var res []*db.GetReviewByProductIdRow
	err := r.db.SelectContext(ctx, &res, getReviewByProductId,
		int32(req.ProductID),
		int32(req.Rating),
		int32(req.PageSize),
		int32(offset),
	)

	if err != nil {
		return nil, review_errors.ErrFindReviewsByProduct.WithInternal(err)
	}

	return res, nil
}

func (r *reviewQueryRepository) FindByMerchant(ctx context.Context, req *requests.FindAllReviewByMerchant) ([]*db.GetReviewByMerchantIdRow, error) {
	offset := (req.Page - 1) * req.PageSize

	var res []*db.GetReviewByMerchantIdRow
	err := r.db.SelectContext(ctx, &res, getReviewByMerchantId,
		int32(req.MerchantID),
		int32(req.Rating),
		int32(req.PageSize),
		int32(offset),
	)

	if err != nil {
		return nil, review_errors.ErrFindReviewsByMerchant.WithInternal(err)
	}

	return res, nil
}

func (r *reviewQueryRepository) FindActive(ctx context.Context, req *requests.FindAllReview) ([]*db.GetReviewsActiveRow, error) {
	offset := (req.Page - 1) * req.PageSize

	var res []*db.GetReviewsActiveRow
	err := r.db.SelectContext(ctx, &res, getReviewsActive,
		req.Search,
		int32(req.PageSize),
		int32(offset),
	)

	if err != nil {
		return nil, review_errors.ErrFindActiveReviews.WithInternal(err)
	}

	return res, nil
}

func (r *reviewQueryRepository) FindTrashed(ctx context.Context, req *requests.FindAllReview) ([]*db.GetReviewsTrashedRow, error) {
	offset := (req.Page - 1) * req.PageSize

	var res []*db.GetReviewsTrashedRow
	err := r.db.SelectContext(ctx, &res, getReviewsTrashed,
		req.Search,
		int32(req.PageSize),
		int32(offset),
	)

	if err != nil {
		return nil, review_errors.ErrFindTrashedReviews.WithInternal(err)
	}

	return res, nil
}

func (r *reviewQueryRepository) FindByID(ctx context.Context, id int) (*db.GetReviewByIDRow, error) {
	var res db.GetReviewByIDRow
	err := r.db.GetContext(ctx, &res, getReviewByID, int32(id))

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, review_errors.ErrReviewNotFound
		}
		return nil, review_errors.ErrFindReviewByID.WithInternal(err)
	}

	return &res, nil
}
