// Code migrated from sqlc-generated sources; domain types kept for repository layer.
// source: merchant_detail.sql

package db

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type CreateMerchantDetailParams struct {
	MerchantID       int32   `db:"merchant_id" json:"merchant_id"`
	DisplayName      *string `db:"display_name" json:"display_name"`
	CoverImageUrl    *string `db:"cover_image_url" json:"cover_image_url"`
	LogoUrl          *string `db:"logo_url" json:"logo_url"`
	ShortDescription *string `db:"short_description" json:"short_description"`
	WebsiteUrl       *string `db:"website_url" json:"website_url"`
}

type CreateMerchantDetailRow struct {
	MerchantDetailID int32            `db:"merchant_detail_id" json:"merchant_detail_id"`
	MerchantID       int32            `db:"merchant_id" json:"merchant_id"`
	DisplayName      *string          `db:"display_name" json:"display_name"`
	CoverImageUrl    *string          `db:"cover_image_url" json:"cover_image_url"`
	LogoUrl          *string          `db:"logo_url" json:"logo_url"`
	ShortDescription *string          `db:"short_description" json:"short_description"`
	WebsiteUrl       *string          `db:"website_url" json:"website_url"`
	CreatedAt        pgtype.Timestamp `db:"created_at" json:"created_at"`
	UpdatedAt        pgtype.Timestamp `db:"updated_at" json:"updated_at"`
}

type GetMerchantDetailRow struct {
	MerchantDetailID int32            `db:"merchant_detail_id" json:"merchant_detail_id"`
	MerchantID       int32            `db:"merchant_id" json:"merchant_id"`
	DisplayName      *string          `db:"display_name" json:"display_name"`
	CoverImageUrl    *string          `db:"cover_image_url" json:"cover_image_url"`
	LogoUrl          *string          `db:"logo_url" json:"logo_url"`
	ShortDescription *string          `db:"short_description" json:"short_description"`
	WebsiteUrl       *string          `db:"website_url" json:"website_url"`
	CreatedAt        pgtype.Timestamp `db:"created_at" json:"created_at"`
	UpdatedAt        pgtype.Timestamp `db:"updated_at" json:"updated_at"`
	MerchantName     string           `db:"merchant_name" json:"merchant_name"`
	SocialMediaLinks []byte           `db:"social_media_links" json:"social_media_links"`
}

type GetMerchantDetailTrashedRow struct {
	MerchantDetailID int32            `db:"merchant_detail_id" json:"merchant_detail_id"`
	MerchantID       int32            `db:"merchant_id" json:"merchant_id"`
	DisplayName      *string          `db:"display_name" json:"display_name"`
	CoverImageUrl    *string          `db:"cover_image_url" json:"cover_image_url"`
	LogoUrl          *string          `db:"logo_url" json:"logo_url"`
	ShortDescription *string          `db:"short_description" json:"short_description"`
	WebsiteUrl       *string          `db:"website_url" json:"website_url"`
	CreatedAt        pgtype.Timestamp `db:"created_at" json:"created_at"`
	UpdatedAt        pgtype.Timestamp `db:"updated_at" json:"updated_at"`
}

type GetMerchantDetailsParams struct {
	Column1 string `db:"column1" json:"column_1"`
	Limit   int32  `db:"limit" json:"limit"`
	Offset  int32  `db:"offset" json:"offset"`
}

type GetMerchantDetailsRow struct {
	MerchantDetailID int32            `db:"merchant_detail_id" json:"merchant_detail_id"`
	MerchantID       int32            `db:"merchant_id" json:"merchant_id"`
	DisplayName      *string          `db:"display_name" json:"display_name"`
	CoverImageUrl    *string          `db:"cover_image_url" json:"cover_image_url"`
	LogoUrl          *string          `db:"logo_url" json:"logo_url"`
	ShortDescription *string          `db:"short_description" json:"short_description"`
	WebsiteUrl       *string          `db:"website_url" json:"website_url"`
	CreatedAt        pgtype.Timestamp `db:"created_at" json:"created_at"`
	UpdatedAt        pgtype.Timestamp `db:"updated_at" json:"updated_at"`
	MerchantName     string           `db:"merchant_name" json:"merchant_name"`
	TotalCount       int64            `db:"total_count" json:"total_count"`
	SocialMediaLinks []byte           `db:"social_media_links" json:"social_media_links"`
}

type GetMerchantDetailsActiveParams struct {
	Column1 string `db:"column1" json:"column_1"`
	Limit   int32  `db:"limit" json:"limit"`
	Offset  int32  `db:"offset" json:"offset"`
}

type GetMerchantDetailsActiveRow struct {
	MerchantDetailID int32            `db:"merchant_detail_id" json:"merchant_detail_id"`
	MerchantID       int32            `db:"merchant_id" json:"merchant_id"`
	DisplayName      *string          `db:"display_name" json:"display_name"`
	CoverImageUrl    *string          `db:"cover_image_url" json:"cover_image_url"`
	LogoUrl          *string          `db:"logo_url" json:"logo_url"`
	ShortDescription *string          `db:"short_description" json:"short_description"`
	WebsiteUrl       *string          `db:"website_url" json:"website_url"`
	CreatedAt        pgtype.Timestamp `db:"created_at" json:"created_at"`
	UpdatedAt        pgtype.Timestamp `db:"updated_at" json:"updated_at"`
	DeletedAt        pgtype.Timestamp `db:"deleted_at" json:"deleted_at"`
	MerchantName     string           `db:"merchant_name" json:"merchant_name"`
	TotalCount       int64            `db:"total_count" json:"total_count"`
	SocialMediaLinks []byte           `db:"social_media_links" json:"social_media_links"`
}

type GetMerchantDetailsTrashedParams struct {
	Column1 string `db:"column1" json:"column_1"`
	Limit   int32  `db:"limit" json:"limit"`
	Offset  int32  `db:"offset" json:"offset"`
}

type GetMerchantDetailsTrashedRow struct {
	MerchantDetailID int32            `db:"merchant_detail_id" json:"merchant_detail_id"`
	MerchantID       int32            `db:"merchant_id" json:"merchant_id"`
	DisplayName      *string          `db:"display_name" json:"display_name"`
	CoverImageUrl    *string          `db:"cover_image_url" json:"cover_image_url"`
	LogoUrl          *string          `db:"logo_url" json:"logo_url"`
	ShortDescription *string          `db:"short_description" json:"short_description"`
	WebsiteUrl       *string          `db:"website_url" json:"website_url"`
	CreatedAt        pgtype.Timestamp `db:"created_at" json:"created_at"`
	UpdatedAt        pgtype.Timestamp `db:"updated_at" json:"updated_at"`
	DeletedAt        pgtype.Timestamp `db:"deleted_at" json:"deleted_at"`
	MerchantName     string           `db:"merchant_name" json:"merchant_name"`
	TotalCount       int64            `db:"total_count" json:"total_count"`
	SocialMediaLinks []byte           `db:"social_media_links" json:"social_media_links"`
}

type UpdateMerchantDetailParams struct {
	MerchantDetailID int32   `db:"merchant_detail_id" json:"merchant_detail_id"`
	DisplayName      *string `db:"display_name" json:"display_name"`
	CoverImageUrl    *string `db:"cover_image_url" json:"cover_image_url"`
	LogoUrl          *string `db:"logo_url" json:"logo_url"`
	ShortDescription *string `db:"short_description" json:"short_description"`
	WebsiteUrl       *string `db:"website_url" json:"website_url"`
}

type UpdateMerchantDetailRow struct {
	MerchantDetailID int32            `db:"merchant_detail_id" json:"merchant_detail_id"`
	MerchantID       int32            `db:"merchant_id" json:"merchant_id"`
	DisplayName      *string          `db:"display_name" json:"display_name"`
	CoverImageUrl    *string          `db:"cover_image_url" json:"cover_image_url"`
	LogoUrl          *string          `db:"logo_url" json:"logo_url"`
	ShortDescription *string          `db:"short_description" json:"short_description"`
	WebsiteUrl       *string          `db:"website_url" json:"website_url"`
	CreatedAt        pgtype.Timestamp `db:"created_at" json:"created_at"`
	UpdatedAt        pgtype.Timestamp `db:"updated_at" json:"updated_at"`
}
