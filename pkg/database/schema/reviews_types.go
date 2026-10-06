// Code migrated from sqlc-generated sources; domain types kept for repository layer.
// source: reviews.sql

package db

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type CreateReviewParams struct {
	UserID    int32  `db:"user_id" json:"user_id"`
	ProductID int32  `db:"product_id" json:"product_id"`
	Name      string `db:"name" json:"name"`
	Comment   string `db:"comment" json:"comment"`
	Rating    int32  `db:"rating" json:"rating"`
}

type CreateReviewRow struct {
	ReviewID  int32            `db:"review_id" json:"review_id"`
	UserID    int32            `db:"user_id" json:"user_id"`
	ProductID int32            `db:"product_id" json:"product_id"`
	Name      string           `db:"name" json:"name"`
	Comment   string           `db:"comment" json:"comment"`
	Rating    int32            `db:"rating" json:"rating"`
	CreatedAt pgtype.Timestamp `db:"created_at" json:"created_at"`
	UpdatedAt pgtype.Timestamp `db:"updated_at" json:"updated_at"`
}

type GetReviewByIDRow struct {
	ReviewID  int32            `db:"review_id" json:"review_id"`
	UserID    int32            `db:"user_id" json:"user_id"`
	ProductID int32            `db:"product_id" json:"product_id"`
	Name      string           `db:"name" json:"name"`
	Comment   string           `db:"comment" json:"comment"`
	Rating    int32            `db:"rating" json:"rating"`
	CreatedAt pgtype.Timestamp `db:"created_at" json:"created_at"`
	UpdatedAt pgtype.Timestamp `db:"updated_at" json:"updated_at"`
}

type GetReviewByMerchantIdParams struct {
	MerchantID int32 `db:"merchant_id" json:"merchant_id"`
	Column2    int32 `db:"column2" json:"column_2"`
	Limit      int32 `db:"limit" json:"limit"`
	Offset     int32 `db:"offset" json:"offset"`
}

type GetReviewByMerchantIdRow struct {
	ReviewID      int32            `db:"review_id" json:"review_id"`
	UserID        int32            `db:"user_id" json:"user_id"`
	ProductID     int32            `db:"product_id" json:"product_id"`
	Name          string           `db:"name" json:"name"`
	Comment       string           `db:"comment" json:"comment"`
	Rating        int32            `db:"rating" json:"rating"`
	CreatedAt     pgtype.Timestamp `db:"created_at" json:"created_at"`
	UpdatedAt     pgtype.Timestamp `db:"updated_at" json:"updated_at"`
	DeletedAt     pgtype.Timestamp `db:"deleted_at" json:"deleted_at"`
	TotalCount    int64            `db:"total_count" json:"total_count"`
	ReviewDetails interface{}      `db:"review_details" json:"review_details"`
}

type GetReviewByProductIdParams struct {
	ProductID int32 `db:"product_id" json:"product_id"`
	Column2   int32 `db:"column2" json:"column_2"`
	Limit     int32 `db:"limit" json:"limit"`
	Offset    int32 `db:"offset" json:"offset"`
}

type GetReviewByProductIdRow struct {
	ReviewID      int32            `db:"review_id" json:"review_id"`
	UserID        int32            `db:"user_id" json:"user_id"`
	ProductID     int32            `db:"product_id" json:"product_id"`
	Name          string           `db:"name" json:"name"`
	Comment       string           `db:"comment" json:"comment"`
	Rating        int32            `db:"rating" json:"rating"`
	CreatedAt     pgtype.Timestamp `db:"created_at" json:"created_at"`
	UpdatedAt     pgtype.Timestamp `db:"updated_at" json:"updated_at"`
	DeletedAt     pgtype.Timestamp `db:"deleted_at" json:"deleted_at"`
	TotalCount    int64            `db:"total_count" json:"total_count"`
	ReviewDetails interface{}      `db:"review_details" json:"review_details"`
}

type GetReviewsParams struct {
	Column1 string `db:"column1" json:"column_1"`
	Limit   int32  `db:"limit" json:"limit"`
	Offset  int32  `db:"offset" json:"offset"`
}

type GetReviewsRow struct {
	ReviewID   int32            `db:"review_id" json:"review_id"`
	UserID     int32            `db:"user_id" json:"user_id"`
	ProductID  int32            `db:"product_id" json:"product_id"`
	Name       string           `db:"name" json:"name"`
	Comment    string           `db:"comment" json:"comment"`
	Rating     int32            `db:"rating" json:"rating"`
	CreatedAt  pgtype.Timestamp `db:"created_at" json:"created_at"`
	UpdatedAt  pgtype.Timestamp `db:"updated_at" json:"updated_at"`
	TotalCount int64            `db:"total_count" json:"total_count"`
}

type GetReviewsActiveParams struct {
	Column1 string `db:"column1" json:"column_1"`
	Limit   int32  `db:"limit" json:"limit"`
	Offset  int32  `db:"offset" json:"offset"`
}

type GetReviewsActiveRow struct {
	ReviewID   int32            `db:"review_id" json:"review_id"`
	UserID     int32            `db:"user_id" json:"user_id"`
	ProductID  int32            `db:"product_id" json:"product_id"`
	Name       string           `db:"name" json:"name"`
	Comment    string           `db:"comment" json:"comment"`
	Rating     int32            `db:"rating" json:"rating"`
	CreatedAt  pgtype.Timestamp `db:"created_at" json:"created_at"`
	UpdatedAt  pgtype.Timestamp `db:"updated_at" json:"updated_at"`
	DeletedAt  pgtype.Timestamp `db:"deleted_at" json:"deleted_at"`
	TotalCount int64            `db:"total_count" json:"total_count"`
}

type GetReviewsTrashedParams struct {
	Column1 string `db:"column1" json:"column_1"`
	Limit   int32  `db:"limit" json:"limit"`
	Offset  int32  `db:"offset" json:"offset"`
}

type GetReviewsTrashedRow struct {
	ReviewID   int32            `db:"review_id" json:"review_id"`
	UserID     int32            `db:"user_id" json:"user_id"`
	ProductID  int32            `db:"product_id" json:"product_id"`
	Name       string           `db:"name" json:"name"`
	Comment    string           `db:"comment" json:"comment"`
	Rating     int32            `db:"rating" json:"rating"`
	CreatedAt  pgtype.Timestamp `db:"created_at" json:"created_at"`
	UpdatedAt  pgtype.Timestamp `db:"updated_at" json:"updated_at"`
	DeletedAt  pgtype.Timestamp `db:"deleted_at" json:"deleted_at"`
	TotalCount int64            `db:"total_count" json:"total_count"`
}

type UpdateReviewParams struct {
	ReviewID int32  `db:"review_id" json:"review_id"`
	Name     string `db:"name" json:"name"`
	Comment  string `db:"comment" json:"comment"`
	Rating   int32  `db:"rating" json:"rating"`
}

type UpdateReviewRow struct {
	ReviewID  int32            `db:"review_id" json:"review_id"`
	UserID    int32            `db:"user_id" json:"user_id"`
	ProductID int32            `db:"product_id" json:"product_id"`
	Name      string           `db:"name" json:"name"`
	Comment   string           `db:"comment" json:"comment"`
	Rating    int32            `db:"rating" json:"rating"`
	CreatedAt pgtype.Timestamp `db:"created_at" json:"created_at"`
	UpdatedAt pgtype.Timestamp `db:"updated_at" json:"updated_at"`
}
