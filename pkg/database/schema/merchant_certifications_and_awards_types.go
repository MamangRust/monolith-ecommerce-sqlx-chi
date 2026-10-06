// Code migrated from sqlc-generated sources; domain types kept for repository layer.
// source: merchant_certifications_and_awards.sql

package db

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type CreateMerchantCertificationOrAwardParams struct {
	MerchantID     int32       `db:"merchant_id" json:"merchant_id"`
	Title          string      `db:"title" json:"title"`
	Description    *string     `db:"description" json:"description"`
	IssuedBy       *string     `db:"issued_by" json:"issued_by"`
	IssueDate      pgtype.Date `db:"issue_date" json:"issue_date"`
	ExpiryDate     pgtype.Date `db:"expiry_date" json:"expiry_date"`
	CertificateUrl *string     `db:"certificate_url" json:"certificate_url"`
}

type CreateMerchantCertificationOrAwardRow struct {
	MerchantCertificationID int32            `db:"merchant_certification_id" json:"merchant_certification_id"`
	MerchantID              int32            `db:"merchant_id" json:"merchant_id"`
	Title                   string           `db:"title" json:"title"`
	Description             *string          `db:"description" json:"description"`
	IssuedBy                *string          `db:"issued_by" json:"issued_by"`
	IssueDate               pgtype.Date      `db:"issue_date" json:"issue_date"`
	ExpiryDate              pgtype.Date      `db:"expiry_date" json:"expiry_date"`
	CertificateUrl          *string          `db:"certificate_url" json:"certificate_url"`
	CreatedAt               pgtype.Timestamp `db:"created_at" json:"created_at"`
	UpdatedAt               pgtype.Timestamp `db:"updated_at" json:"updated_at"`
}

type GetMerchantCertificationOrAwardRow struct {
	MerchantCertificationID int32            `db:"merchant_certification_id" json:"merchant_certification_id"`
	MerchantID              int32            `db:"merchant_id" json:"merchant_id"`
	Title                   string           `db:"title" json:"title"`
	Description             *string          `db:"description" json:"description"`
	IssuedBy                *string          `db:"issued_by" json:"issued_by"`
	IssueDate               pgtype.Date      `db:"issue_date" json:"issue_date"`
	ExpiryDate              pgtype.Date      `db:"expiry_date" json:"expiry_date"`
	CertificateUrl          *string          `db:"certificate_url" json:"certificate_url"`
	CreatedAt               pgtype.Timestamp `db:"created_at" json:"created_at"`
	UpdatedAt               pgtype.Timestamp `db:"updated_at" json:"updated_at"`
}

type GetMerchantCertificationsAndAwardsParams struct {
	Column1 string `db:"column1" json:"column_1"`
	Limit   int32  `db:"limit" json:"limit"`
	Offset  int32  `db:"offset" json:"offset"`
}

type GetMerchantCertificationsAndAwardsRow struct {
	MerchantCertificationID int32            `db:"merchant_certification_id" json:"merchant_certification_id"`
	MerchantID              int32            `db:"merchant_id" json:"merchant_id"`
	Title                   string           `db:"title" json:"title"`
	Description             *string          `db:"description" json:"description"`
	IssuedBy                *string          `db:"issued_by" json:"issued_by"`
	IssueDate               pgtype.Date      `db:"issue_date" json:"issue_date"`
	ExpiryDate              pgtype.Date      `db:"expiry_date" json:"expiry_date"`
	CertificateUrl          *string          `db:"certificate_url" json:"certificate_url"`
	CreatedAt               pgtype.Timestamp `db:"created_at" json:"created_at"`
	UpdatedAt               pgtype.Timestamp `db:"updated_at" json:"updated_at"`
	MerchantName            string           `db:"merchant_name" json:"merchant_name"`
	TotalCount              int64            `db:"total_count" json:"total_count"`
}

type GetMerchantCertificationsAndAwardsActiveParams struct {
	Column1 string `db:"column1" json:"column_1"`
	Limit   int32  `db:"limit" json:"limit"`
	Offset  int32  `db:"offset" json:"offset"`
}

type GetMerchantCertificationsAndAwardsActiveRow struct {
	MerchantCertificationID int32            `db:"merchant_certification_id" json:"merchant_certification_id"`
	MerchantID              int32            `db:"merchant_id" json:"merchant_id"`
	Title                   string           `db:"title" json:"title"`
	Description             *string          `db:"description" json:"description"`
	IssuedBy                *string          `db:"issued_by" json:"issued_by"`
	IssueDate               pgtype.Date      `db:"issue_date" json:"issue_date"`
	ExpiryDate              pgtype.Date      `db:"expiry_date" json:"expiry_date"`
	CertificateUrl          *string          `db:"certificate_url" json:"certificate_url"`
	CreatedAt               pgtype.Timestamp `db:"created_at" json:"created_at"`
	UpdatedAt               pgtype.Timestamp `db:"updated_at" json:"updated_at"`
	DeletedAt               pgtype.Timestamp `db:"deleted_at" json:"deleted_at"`
	MerchantName            string           `db:"merchant_name" json:"merchant_name"`
	TotalCount              int64            `db:"total_count" json:"total_count"`
}

type GetMerchantCertificationsAndAwardsTrashedParams struct {
	Column1 string `db:"column1" json:"column_1"`
	Limit   int32  `db:"limit" json:"limit"`
	Offset  int32  `db:"offset" json:"offset"`
}

type GetMerchantCertificationsAndAwardsTrashedRow struct {
	MerchantCertificationID int32            `db:"merchant_certification_id" json:"merchant_certification_id"`
	MerchantID              int32            `db:"merchant_id" json:"merchant_id"`
	Title                   string           `db:"title" json:"title"`
	Description             *string          `db:"description" json:"description"`
	IssuedBy                *string          `db:"issued_by" json:"issued_by"`
	IssueDate               pgtype.Date      `db:"issue_date" json:"issue_date"`
	ExpiryDate              pgtype.Date      `db:"expiry_date" json:"expiry_date"`
	CertificateUrl          *string          `db:"certificate_url" json:"certificate_url"`
	CreatedAt               pgtype.Timestamp `db:"created_at" json:"created_at"`
	UpdatedAt               pgtype.Timestamp `db:"updated_at" json:"updated_at"`
	DeletedAt               pgtype.Timestamp `db:"deleted_at" json:"deleted_at"`
	MerchantName            string           `db:"merchant_name" json:"merchant_name"`
	TotalCount              int64            `db:"total_count" json:"total_count"`
}

type UpdateMerchantCertificationOrAwardParams struct {
	MerchantCertificationID int32       `db:"merchant_certification_id" json:"merchant_certification_id"`
	Title                   string      `db:"title" json:"title"`
	Description             *string     `db:"description" json:"description"`
	IssuedBy                *string     `db:"issued_by" json:"issued_by"`
	IssueDate               pgtype.Date `db:"issue_date" json:"issue_date"`
	ExpiryDate              pgtype.Date `db:"expiry_date" json:"expiry_date"`
	CertificateUrl          *string     `db:"certificate_url" json:"certificate_url"`
}

type UpdateMerchantCertificationOrAwardRow struct {
	MerchantCertificationID int32            `db:"merchant_certification_id" json:"merchant_certification_id"`
	MerchantID              int32            `db:"merchant_id" json:"merchant_id"`
	Title                   string           `db:"title" json:"title"`
	Description             *string          `db:"description" json:"description"`
	IssuedBy                *string          `db:"issued_by" json:"issued_by"`
	IssueDate               pgtype.Date      `db:"issue_date" json:"issue_date"`
	ExpiryDate              pgtype.Date      `db:"expiry_date" json:"expiry_date"`
	CertificateUrl          *string          `db:"certificate_url" json:"certificate_url"`
	CreatedAt               pgtype.Timestamp `db:"created_at" json:"created_at"`
	UpdatedAt               pgtype.Timestamp `db:"updated_at" json:"updated_at"`
}
