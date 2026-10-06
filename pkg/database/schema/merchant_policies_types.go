// Code migrated from sqlc-generated sources; domain types kept for repository layer.
// source: merchant_policies.sql

package db

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type CreateMerchantPolicyParams struct {
	MerchantID  int32  `db:"merchant_id" json:"merchant_id"`
	PolicyType  string `db:"policy_type" json:"policy_type"`
	Title       string `db:"title" json:"title"`
	Description string `db:"description" json:"description"`
}

type CreateMerchantPolicyRow struct {
	MerchantPolicyID int32            `db:"merchant_policy_id" json:"merchant_policy_id"`
	MerchantID       int32            `db:"merchant_id" json:"merchant_id"`
	PolicyType       string           `db:"policy_type" json:"policy_type"`
	Title            string           `db:"title" json:"title"`
	Description      string           `db:"description" json:"description"`
	CreatedAt        pgtype.Timestamp `db:"created_at" json:"created_at"`
	UpdatedAt        pgtype.Timestamp `db:"updated_at" json:"updated_at"`
}

type GetMerchantPoliciesParams struct {
	Column1 string `db:"column1" json:"column_1"`
	Limit   int32  `db:"limit" json:"limit"`
	Offset  int32  `db:"offset" json:"offset"`
}

type GetMerchantPoliciesRow struct {
	MerchantPolicyID int32            `db:"merchant_policy_id" json:"merchant_policy_id"`
	MerchantID       int32            `db:"merchant_id" json:"merchant_id"`
	PolicyType       string           `db:"policy_type" json:"policy_type"`
	Title            string           `db:"title" json:"title"`
	Description      string           `db:"description" json:"description"`
	CreatedAt        pgtype.Timestamp `db:"created_at" json:"created_at"`
	UpdatedAt        pgtype.Timestamp `db:"updated_at" json:"updated_at"`
	MerchantName     string           `db:"merchant_name" json:"merchant_name"`
	TotalCount       int64            `db:"total_count" json:"total_count"`
}

type GetMerchantPoliciesActiveParams struct {
	Column1 string `db:"column1" json:"column_1"`
	Limit   int32  `db:"limit" json:"limit"`
	Offset  int32  `db:"offset" json:"offset"`
}

type GetMerchantPoliciesActiveRow struct {
	MerchantPolicyID int32            `db:"merchant_policy_id" json:"merchant_policy_id"`
	MerchantID       int32            `db:"merchant_id" json:"merchant_id"`
	PolicyType       string           `db:"policy_type" json:"policy_type"`
	Title            string           `db:"title" json:"title"`
	Description      string           `db:"description" json:"description"`
	CreatedAt        pgtype.Timestamp `db:"created_at" json:"created_at"`
	UpdatedAt        pgtype.Timestamp `db:"updated_at" json:"updated_at"`
	DeletedAt        pgtype.Timestamp `db:"deleted_at" json:"deleted_at"`
	MerchantName     string           `db:"merchant_name" json:"merchant_name"`
	TotalCount       int64            `db:"total_count" json:"total_count"`
}

type GetMerchantPoliciesTrashedParams struct {
	Column1 string `db:"column1" json:"column_1"`
	Limit   int32  `db:"limit" json:"limit"`
	Offset  int32  `db:"offset" json:"offset"`
}

type GetMerchantPoliciesTrashedRow struct {
	MerchantPolicyID int32            `db:"merchant_policy_id" json:"merchant_policy_id"`
	MerchantID       int32            `db:"merchant_id" json:"merchant_id"`
	PolicyType       string           `db:"policy_type" json:"policy_type"`
	Title            string           `db:"title" json:"title"`
	Description      string           `db:"description" json:"description"`
	CreatedAt        pgtype.Timestamp `db:"created_at" json:"created_at"`
	UpdatedAt        pgtype.Timestamp `db:"updated_at" json:"updated_at"`
	DeletedAt        pgtype.Timestamp `db:"deleted_at" json:"deleted_at"`
	MerchantName     string           `db:"merchant_name" json:"merchant_name"`
	TotalCount       int64            `db:"total_count" json:"total_count"`
}

type GetMerchantPolicyRow struct {
	MerchantPolicyID int32            `db:"merchant_policy_id" json:"merchant_policy_id"`
	MerchantID       int32            `db:"merchant_id" json:"merchant_id"`
	PolicyType       string           `db:"policy_type" json:"policy_type"`
	Title            string           `db:"title" json:"title"`
	Description      string           `db:"description" json:"description"`
	CreatedAt        pgtype.Timestamp `db:"created_at" json:"created_at"`
	UpdatedAt        pgtype.Timestamp `db:"updated_at" json:"updated_at"`
}

type UpdateMerchantPolicyParams struct {
	MerchantPolicyID int32  `db:"merchant_policy_id" json:"merchant_policy_id"`
	PolicyType       string `db:"policy_type" json:"policy_type"`
	Title            string `db:"title" json:"title"`
	Description      string `db:"description" json:"description"`
}

type UpdateMerchantPolicyRow struct {
	MerchantPolicyID int32            `db:"merchant_policy_id" json:"merchant_policy_id"`
	MerchantID       int32            `db:"merchant_id" json:"merchant_id"`
	PolicyType       string           `db:"policy_type" json:"policy_type"`
	Title            string           `db:"title" json:"title"`
	Description      string           `db:"description" json:"description"`
	CreatedAt        pgtype.Timestamp `db:"created_at" json:"created_at"`
	UpdatedAt        pgtype.Timestamp `db:"updated_at" json:"updated_at"`
}
