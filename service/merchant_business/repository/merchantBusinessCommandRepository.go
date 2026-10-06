package repository

import (
	"context"
	"database/sql"
	"errors"

	db "github.com/MamangRust/monolith-ecommerce-pkg/database/schema"
	"github.com/MamangRust/monolith-ecommerce-shared/convert"
	"github.com/MamangRust/monolith-ecommerce-shared/domain/requests"
	merchantbusiness_errors "github.com/MamangRust/monolith-ecommerce-shared/errors/merchant_business"
	"github.com/jmoiron/sqlx"
)

const createMerchantBusinessInformation = `INSERT INTO
    merchant_business_information (
        merchant_id,
        business_type,
        tax_id,
        established_year,
        number_of_employees,
        website_url
    )
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING
    merchant_business_info_id,
    merchant_id,
    business_type,
    tax_id,
    established_year,
    number_of_employees,
    website_url,
    created_at,
    updated_at`

const updateMerchantBusinessInformation = `UPDATE merchant_business_information
SET
    business_type = $2,
    tax_id = $3,
    established_year = $4,
    number_of_employees = $5,
    website_url = $6,
    updated_at = CURRENT_TIMESTAMP
WHERE
    merchant_business_info_id = $1
    AND deleted_at IS NULL
RETURNING
    merchant_business_info_id,
    merchant_id,
    business_type,
    tax_id,
    established_year,
    number_of_employees,
    website_url,
    created_at,
    updated_at`

const trashMerchantBusinessInformation = `UPDATE merchant_business_information
SET
    deleted_at = CURRENT_TIMESTAMP
WHERE
    merchant_business_info_id = $1
    AND deleted_at IS NULL
RETURNING
    merchant_business_info_id,
    merchant_id,
    business_type,
    tax_id,
    established_year,
    number_of_employees,
    website_url,
    created_at,
    updated_at,
    deleted_at`

const restoreMerchantBusinessInformation = `UPDATE merchant_business_information
SET
    deleted_at = NULL
WHERE
    merchant_business_info_id = $1
    AND deleted_at IS NOT NULL
RETURNING
    merchant_business_info_id,
    merchant_id,
    business_type,
    tax_id,
    established_year,
    number_of_employees,
    website_url,
    created_at,
    updated_at,
    deleted_at`

const deleteMerchantBusinessInformationPermanently = `DELETE FROM merchant_business_information
WHERE
    merchant_business_info_id = $1
    AND deleted_at IS NOT NULL`

const restoreAllMerchantBusinessInformation = `UPDATE merchant_business_information
SET
    deleted_at = NULL
WHERE
    deleted_at IS NOT NULL`

const deleteAllPermanentMerchantBusinessInformation = `DELETE FROM merchant_business_information
WHERE
    deleted_at IS NOT NULL`

type merchantBusinessCommandRepository struct {
	db *sqlx.DB
}

func NewMerchantBusinessCommandRepository(
	db *sqlx.DB,
) *merchantBusinessCommandRepository {
	return &merchantBusinessCommandRepository{
		db: db,
	}
}

func (r *merchantBusinessCommandRepository) Create(
	ctx context.Context,
	request *requests.CreateMerchantBusinessInformationRequest,
) (*db.CreateMerchantBusinessInformationRow, error) {

	var business db.CreateMerchantBusinessInformationRow
	err := r.db.GetContext(ctx, &business, createMerchantBusinessInformation,
		int32(request.MerchantID),
		convert.NullableString(request.BusinessType),
		convert.NullableString(request.TaxID),
		int32Ptr(request.EstablishedYear),
		int32Ptr(request.NumberOfEmployees),
		convert.NullableString(request.WebsiteUrl),
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, merchantbusiness_errors.ErrMerchantBusinessNotFound
		}
		return nil, merchantbusiness_errors.ErrCreateMerchantBusiness.WithInternal(err)
	}

	return &business, nil
}

func (r *merchantBusinessCommandRepository) Update(ctx context.Context, request *requests.UpdateMerchantBusinessInformationRequest) (*db.UpdateMerchantBusinessInformationRow, error) {
	var business db.UpdateMerchantBusinessInformationRow
	err := r.db.GetContext(ctx, &business, updateMerchantBusinessInformation,
		int32(*request.MerchantBusinessInfoID),
		convert.NullableString(request.BusinessType),
		convert.NullableString(request.TaxID),
		int32Ptr(request.EstablishedYear),
		int32Ptr(request.NumberOfEmployees),
		convert.NullableString(request.WebsiteUrl),
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, merchantbusiness_errors.ErrMerchantBusinessNotFound
		}
		return nil, merchantbusiness_errors.ErrUpdateMerchantBusiness.WithInternal(err)
	}

	return &business, nil
}

func (r *merchantBusinessCommandRepository) Trash(ctx context.Context, merchant_business_info_id int) (*db.MerchantBusinessInformation, error) {
	var business db.MerchantBusinessInformation
	err := r.db.GetContext(ctx, &business, trashMerchantBusinessInformation, int32(merchant_business_info_id))

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, merchantbusiness_errors.ErrMerchantBusinessNotFound
		}
		return nil, merchantbusiness_errors.ErrTrashMerchantBusiness.WithInternal(err)
	}

	return &business, nil
}

func (r *merchantBusinessCommandRepository) Restore(ctx context.Context, merchant_business_info_id int) (*db.MerchantBusinessInformation, error) {
	var business db.MerchantBusinessInformation
	err := r.db.GetContext(ctx, &business, restoreMerchantBusinessInformation, int32(merchant_business_info_id))

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, merchantbusiness_errors.ErrMerchantBusinessNotFound
		}
		return nil, merchantbusiness_errors.ErrRestoreMerchantBusiness.WithInternal(err)
	}

	return &business, nil
}

func (r *merchantBusinessCommandRepository) DeletePermanent(ctx context.Context, merchant_business_info_id int) (bool, error) {
	if _, err := r.db.ExecContext(ctx, deleteMerchantBusinessInformationPermanently, int32(merchant_business_info_id)); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, merchantbusiness_errors.ErrMerchantBusinessNotFound
		}
		return false, merchantbusiness_errors.ErrDeletePermanentMerchantBusiness.WithInternal(err)
	}

	return true, nil
}

func (r *merchantBusinessCommandRepository) RestoreAll(ctx context.Context) (bool, error) {
	if _, err := r.db.ExecContext(ctx, restoreAllMerchantBusinessInformation); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, merchantbusiness_errors.ErrMerchantBusinessNotFound
		}
		return false, merchantbusiness_errors.ErrRestoreAllMerchantBusinesses.WithInternal(err)
	}

	return true, nil
}

func (r *merchantBusinessCommandRepository) DeleteAll(ctx context.Context) (bool, error) {
	if _, err := r.db.ExecContext(ctx, deleteAllPermanentMerchantBusinessInformation); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, merchantbusiness_errors.ErrMerchantBusinessNotFound
		}
		return false, merchantbusiness_errors.ErrDeleteAllPermanentMerchantBusinesses.WithInternal(err)
	}

	return true, nil
}

func int32Ptr(v int) *int32 {
	if v == 0 {
		return nil
	}
	val := int32(v)
	return &val
}
