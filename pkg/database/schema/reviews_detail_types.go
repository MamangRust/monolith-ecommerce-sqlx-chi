// Code migrated from sqlc-generated sources; domain types kept for repository layer.
// source: reviews_detail.sql

package db

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type CreateReviewDetailParams struct {
	ReviewID int32   `db:"review_id" json:"review_id"`
	Type     string  `db:"type" json:"type"`
	Url      string  `db:"url" json:"url"`
	Caption  *string `db:"caption" json:"caption"`
}

type CreateReviewDetailRow struct {
	ReviewDetailID int32            `db:"review_detail_id" json:"review_detail_id"`
	ReviewID       int32            `db:"review_id" json:"review_id"`
	Type           string           `db:"type" json:"type"`
	Url            string           `db:"url" json:"url"`
	Caption        *string          `db:"caption" json:"caption"`
	CreatedAt      pgtype.Timestamp `db:"created_at" json:"created_at"`
	UpdatedAt      pgtype.Timestamp `db:"updated_at" json:"updated_at"`
}

type GetReviewDetailRow struct {
	ReviewDetailID int32            `db:"review_detail_id" json:"review_detail_id"`
	ReviewID       int32            `db:"review_id" json:"review_id"`
	Type           string           `db:"type" json:"type"`
	Url            string           `db:"url" json:"url"`
	Caption        *string          `db:"caption" json:"caption"`
	CreatedAt      pgtype.Timestamp `db:"created_at" json:"created_at"`
	UpdatedAt      pgtype.Timestamp `db:"updated_at" json:"updated_at"`
}

type GetReviewDetailsParams struct {
	Column1 string `db:"column1" json:"column_1"`
	Limit   int32  `db:"limit" json:"limit"`
	Offset  int32  `db:"offset" json:"offset"`
}

type GetReviewDetailsRow struct {
	ReviewDetailID int32            `db:"review_detail_id" json:"review_detail_id"`
	ReviewID       int32            `db:"review_id" json:"review_id"`
	Type           string           `db:"type" json:"type"`
	Url            string           `db:"url" json:"url"`
	Caption        *string          `db:"caption" json:"caption"`
	CreatedAt      pgtype.Timestamp `db:"created_at" json:"created_at"`
	UpdatedAt      pgtype.Timestamp `db:"updated_at" json:"updated_at"`
	TotalCount     int64            `db:"total_count" json:"total_count"`
}

type GetReviewDetailsActiveParams struct {
	Column1 string `db:"column1" json:"column_1"`
	Limit   int32  `db:"limit" json:"limit"`
	Offset  int32  `db:"offset" json:"offset"`
}

type GetReviewDetailsActiveRow struct {
	ReviewDetailID int32            `db:"review_detail_id" json:"review_detail_id"`
	ReviewID       int32            `db:"review_id" json:"review_id"`
	Type           string           `db:"type" json:"type"`
	Url            string           `db:"url" json:"url"`
	Caption        *string          `db:"caption" json:"caption"`
	CreatedAt      pgtype.Timestamp `db:"created_at" json:"created_at"`
	UpdatedAt      pgtype.Timestamp `db:"updated_at" json:"updated_at"`
	DeletedAt      pgtype.Timestamp `db:"deleted_at" json:"deleted_at"`
	TotalCount     int64            `db:"total_count" json:"total_count"`
}

type GetReviewDetailsTrashedParams struct {
	Column1 string `db:"column1" json:"column_1"`
	Limit   int32  `db:"limit" json:"limit"`
	Offset  int32  `db:"offset" json:"offset"`
}

type GetReviewDetailsTrashedRow struct {
	ReviewDetailID int32            `db:"review_detail_id" json:"review_detail_id"`
	ReviewID       int32            `db:"review_id" json:"review_id"`
	Type           string           `db:"type" json:"type"`
	Url            string           `db:"url" json:"url"`
	Caption        *string          `db:"caption" json:"caption"`
	CreatedAt      pgtype.Timestamp `db:"created_at" json:"created_at"`
	UpdatedAt      pgtype.Timestamp `db:"updated_at" json:"updated_at"`
	DeletedAt      pgtype.Timestamp `db:"deleted_at" json:"deleted_at"`
	TotalCount     int64            `db:"total_count" json:"total_count"`
}

type UpdateReviewDetailParams struct {
	Type           string  `db:"type" json:"type"`
	Url            string  `db:"url" json:"url"`
	Caption        *string `db:"caption" json:"caption"`
	ReviewDetailID int32   `db:"review_detail_id" json:"review_detail_id"`
}

type UpdateReviewDetailRow struct {
	ReviewDetailID int32            `db:"review_detail_id" json:"review_detail_id"`
	ReviewID       int32            `db:"review_id" json:"review_id"`
	Type           string           `db:"type" json:"type"`
	Url            string           `db:"url" json:"url"`
	Caption        *string          `db:"caption" json:"caption"`
	CreatedAt      pgtype.Timestamp `db:"created_at" json:"created_at"`
	UpdatedAt      pgtype.Timestamp `db:"updated_at" json:"updated_at"`
}
