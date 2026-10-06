// Code migrated from sqlc-generated sources; domain types kept for repository layer.
// source: merchant_social_links.sql

package db

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type CreateMerchantSocialMediaLinkParams struct {
	MerchantDetailID int32  `db:"merchant_detail_id" json:"merchant_detail_id"`
	Platform         string `db:"platform" json:"platform"`
	Url              string `db:"url" json:"url"`
}

type CreateMerchantSocialMediaLinkRow struct {
	MerchantSocialID int32            `db:"merchant_social_id" json:"merchant_social_id"`
	MerchantDetailID int32            `db:"merchant_detail_id" json:"merchant_detail_id"`
	Platform         string           `db:"platform" json:"platform"`
	Url              string           `db:"url" json:"url"`
	CreatedAt        pgtype.Timestamp `db:"created_at" json:"created_at"`
	UpdatedAt        pgtype.Timestamp `db:"updated_at" json:"updated_at"`
}

type UpdateMerchantSocialMediaLinkParams struct {
	MerchantSocialID int32  `db:"merchant_social_id" json:"merchant_social_id"`
	Platform         string `db:"platform" json:"platform"`
	Url              string `db:"url" json:"url"`
}

type UpdateMerchantSocialMediaLinkRow struct {
	MerchantSocialID int32            `db:"merchant_social_id" json:"merchant_social_id"`
	MerchantDetailID int32            `db:"merchant_detail_id" json:"merchant_detail_id"`
	Platform         string           `db:"platform" json:"platform"`
	Url              string           `db:"url" json:"url"`
	CreatedAt        pgtype.Timestamp `db:"created_at" json:"created_at"`
	UpdatedAt        pgtype.Timestamp `db:"updated_at" json:"updated_at"`
}
