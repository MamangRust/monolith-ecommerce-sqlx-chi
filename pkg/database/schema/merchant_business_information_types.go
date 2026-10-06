// Code migrated from sqlc-generated sources; domain types kept for repository layer.
// source: merchant_business_information.sql

package db

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type CreateMerchantBusinessInformationParams struct {
	MerchantID        int32   `db:"merchant_id" json:"merchant_id"`
	BusinessType      *string `db:"business_type" json:"business_type"`
	TaxID             *string `db:"tax_id" json:"tax_id"`
	EstablishedYear   *int32  `db:"established_year" json:"established_year"`
	NumberOfEmployees *int32  `db:"number_of_employees" json:"number_of_employees"`
	WebsiteUrl        *string `db:"website_url" json:"website_url"`
}

type CreateMerchantBusinessInformationRow struct {
	MerchantBusinessInfoID int32            `db:"merchant_business_info_id" json:"merchant_business_info_id"`
	MerchantID             int32            `db:"merchant_id" json:"merchant_id"`
	BusinessType           *string          `db:"business_type" json:"business_type"`
	TaxID                  *string          `db:"tax_id" json:"tax_id"`
	EstablishedYear        *int32           `db:"established_year" json:"established_year"`
	NumberOfEmployees      *int32           `db:"number_of_employees" json:"number_of_employees"`
	WebsiteUrl             *string          `db:"website_url" json:"website_url"`
	CreatedAt              pgtype.Timestamp `db:"created_at" json:"created_at"`
	UpdatedAt              pgtype.Timestamp `db:"updated_at" json:"updated_at"`
}

type GetMerchantBusinessInformationRow struct {
	MerchantBusinessInfoID int32            `db:"merchant_business_info_id" json:"merchant_business_info_id"`
	MerchantID             int32            `db:"merchant_id" json:"merchant_id"`
	BusinessType           *string          `db:"business_type" json:"business_type"`
	TaxID                  *string          `db:"tax_id" json:"tax_id"`
	EstablishedYear        *int32           `db:"established_year" json:"established_year"`
	NumberOfEmployees      *int32           `db:"number_of_employees" json:"number_of_employees"`
	WebsiteUrl             *string          `db:"website_url" json:"website_url"`
	CreatedAt              pgtype.Timestamp `db:"created_at" json:"created_at"`
	UpdatedAt              pgtype.Timestamp `db:"updated_at" json:"updated_at"`
}

type GetMerchantsBusinessInformationParams struct {
	Column1 string `db:"column1" json:"column_1"`
	Limit   int32  `db:"limit" json:"limit"`
	Offset  int32  `db:"offset" json:"offset"`
}

type GetMerchantsBusinessInformationRow struct {
	MerchantBusinessInfoID int32            `db:"merchant_business_info_id" json:"merchant_business_info_id"`
	MerchantID             int32            `db:"merchant_id" json:"merchant_id"`
	BusinessType           *string          `db:"business_type" json:"business_type"`
	TaxID                  *string          `db:"tax_id" json:"tax_id"`
	EstablishedYear        *int32           `db:"established_year" json:"established_year"`
	NumberOfEmployees      *int32           `db:"number_of_employees" json:"number_of_employees"`
	WebsiteUrl             *string          `db:"website_url" json:"website_url"`
	CreatedAt              pgtype.Timestamp `db:"created_at" json:"created_at"`
	UpdatedAt              pgtype.Timestamp `db:"updated_at" json:"updated_at"`
	MerchantName           string           `db:"merchant_name" json:"merchant_name"`
	TotalCount             int64            `db:"total_count" json:"total_count"`
}

type GetMerchantsBusinessInformationActiveParams struct {
	Column1 string `db:"column1" json:"column_1"`
	Limit   int32  `db:"limit" json:"limit"`
	Offset  int32  `db:"offset" json:"offset"`
}

type GetMerchantsBusinessInformationActiveRow struct {
	MerchantBusinessInfoID int32            `db:"merchant_business_info_id" json:"merchant_business_info_id"`
	MerchantID             int32            `db:"merchant_id" json:"merchant_id"`
	BusinessType           *string          `db:"business_type" json:"business_type"`
	TaxID                  *string          `db:"tax_id" json:"tax_id"`
	EstablishedYear        *int32           `db:"established_year" json:"established_year"`
	NumberOfEmployees      *int32           `db:"number_of_employees" json:"number_of_employees"`
	WebsiteUrl             *string          `db:"website_url" json:"website_url"`
	CreatedAt              pgtype.Timestamp `db:"created_at" json:"created_at"`
	UpdatedAt              pgtype.Timestamp `db:"updated_at" json:"updated_at"`
	DeletedAt              pgtype.Timestamp `db:"deleted_at" json:"deleted_at"`
	MerchantName           string           `db:"merchant_name" json:"merchant_name"`
	TotalCount             int64            `db:"total_count" json:"total_count"`
}

type GetMerchantsBusinessInformationTrashedParams struct {
	Column1 string `db:"column1" json:"column_1"`
	Limit   int32  `db:"limit" json:"limit"`
	Offset  int32  `db:"offset" json:"offset"`
}

type GetMerchantsBusinessInformationTrashedRow struct {
	MerchantBusinessInfoID int32            `db:"merchant_business_info_id" json:"merchant_business_info_id"`
	MerchantID             int32            `db:"merchant_id" json:"merchant_id"`
	BusinessType           *string          `db:"business_type" json:"business_type"`
	TaxID                  *string          `db:"tax_id" json:"tax_id"`
	EstablishedYear        *int32           `db:"established_year" json:"established_year"`
	NumberOfEmployees      *int32           `db:"number_of_employees" json:"number_of_employees"`
	WebsiteUrl             *string          `db:"website_url" json:"website_url"`
	CreatedAt              pgtype.Timestamp `db:"created_at" json:"created_at"`
	UpdatedAt              pgtype.Timestamp `db:"updated_at" json:"updated_at"`
	DeletedAt              pgtype.Timestamp `db:"deleted_at" json:"deleted_at"`
	MerchantName           string           `db:"merchant_name" json:"merchant_name"`
	TotalCount             int64            `db:"total_count" json:"total_count"`
}

type UpdateMerchantBusinessInformationParams struct {
	MerchantBusinessInfoID int32   `db:"merchant_business_info_id" json:"merchant_business_info_id"`
	BusinessType           *string `db:"business_type" json:"business_type"`
	TaxID                  *string `db:"tax_id" json:"tax_id"`
	EstablishedYear        *int32  `db:"established_year" json:"established_year"`
	NumberOfEmployees      *int32  `db:"number_of_employees" json:"number_of_employees"`
	WebsiteUrl             *string `db:"website_url" json:"website_url"`
}

type UpdateMerchantBusinessInformationRow struct {
	MerchantBusinessInfoID int32            `db:"merchant_business_info_id" json:"merchant_business_info_id"`
	MerchantID             int32            `db:"merchant_id" json:"merchant_id"`
	BusinessType           *string          `db:"business_type" json:"business_type"`
	TaxID                  *string          `db:"tax_id" json:"tax_id"`
	EstablishedYear        *int32           `db:"established_year" json:"established_year"`
	NumberOfEmployees      *int32           `db:"number_of_employees" json:"number_of_employees"`
	WebsiteUrl             *string          `db:"website_url" json:"website_url"`
	CreatedAt              pgtype.Timestamp `db:"created_at" json:"created_at"`
	UpdatedAt              pgtype.Timestamp `db:"updated_at" json:"updated_at"`
}
