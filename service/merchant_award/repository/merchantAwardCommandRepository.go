package repository

import (
	"context"
	"database/sql"
	"errors"
	"time"

	db "github.com/MamangRust/monolith-ecommerce-pkg/database/schema"
	"github.com/MamangRust/monolith-ecommerce-shared/convert"
	"github.com/MamangRust/monolith-ecommerce-shared/domain/requests"
	merchantaward_errors "github.com/MamangRust/monolith-ecommerce-shared/errors/merchant_award"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jmoiron/sqlx"
)

const createMerchantCertificationOrAward = `INSERT INTO
    merchant_certifications_and_awards (
        merchant_id,
        title,
        description,
        issued_by,
        issue_date,
        expiry_date,
        certificate_url
    )
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING
    merchant_certification_id,
    merchant_id,
    title,
    description,
    issued_by,
    issue_date,
    expiry_date,
    certificate_url,
    created_at,
    updated_at`

const updateMerchantCertificationOrAward = `UPDATE merchant_certifications_and_awards
SET
    title = $2,
    description = $3,
    issued_by = $4,
    issue_date = $5,
    expiry_date = $6,
    certificate_url = $7,
    updated_at = CURRENT_TIMESTAMP
WHERE
    merchant_certification_id = $1
    AND deleted_at IS NULL
RETURNING
    merchant_certification_id,
    merchant_id,
    title,
    description,
    issued_by,
    issue_date,
    expiry_date,
    certificate_url,
    created_at,
    updated_at`

const trashMerchantCertificationOrAward = `UPDATE merchant_certifications_and_awards
SET
    deleted_at = CURRENT_TIMESTAMP
WHERE
    merchant_certification_id = $1
    AND deleted_at IS NULL
RETURNING
    merchant_certification_id,
    merchant_id,
    title,
    description,
    issued_by,
    issue_date,
    expiry_date,
    certificate_url,
    created_at,
    updated_at,
    deleted_at`

const restoreMerchantCertificationOrAward = `UPDATE merchant_certifications_and_awards
SET
    deleted_at = NULL
WHERE
    merchant_certification_id = $1
    AND deleted_at IS NOT NULL
RETURNING
    merchant_certification_id,
    merchant_id,
    title,
    description,
    issued_by,
    issue_date,
    expiry_date,
    certificate_url,
    created_at,
    updated_at,
    deleted_at`

const deleteMerchantCertificationOrAwardPermanently = `DELETE FROM merchant_certifications_and_awards
WHERE
    merchant_certification_id = $1
    AND deleted_at IS NOT NULL`

const restoreAllMerchantCertificationsAndAwards = `UPDATE merchant_certifications_and_awards
SET
    deleted_at = NULL
WHERE
    deleted_at IS NOT NULL`

const deleteAllPermanentMerchantCertificationsAndAwards = `DELETE FROM merchant_certifications_and_awards
WHERE
    deleted_at IS NOT NULL`

type merchantAwardCommandRepository struct {
	db *sqlx.DB
}

func NewMerchantAwardCommandRepository(db *sqlx.DB) *merchantAwardCommandRepository {
	return &merchantAwardCommandRepository{
		db: db,
	}
}

func (r *merchantAwardCommandRepository) Create(
	ctx context.Context,
	request *requests.CreateMerchantCertificationOrAwardRequest,
) (*db.CreateMerchantCertificationOrAwardRow, error) {

	var award db.CreateMerchantCertificationOrAwardRow
	err := r.db.GetContext(ctx, &award, createMerchantCertificationOrAward,
		int32(request.MerchantID),
		request.Title,
		convert.NullableString(request.Description),
		convert.NullableString(request.IssuedBy),
		parseDateToPgDate(request.IssueDate),
		parseDateToPgDate(request.ExpiryDate),
		convert.NullableString(request.CertificateUrl),
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, merchantaward_errors.ErrMerchantAwardNotFound
		}
		return nil, merchantaward_errors.ErrCreateMerchantAward.WithInternal(err)
	}

	return &award, nil
}

func (r *merchantAwardCommandRepository) Update(ctx context.Context, request *requests.UpdateMerchantCertificationOrAwardRequest) (*db.UpdateMerchantCertificationOrAwardRow, error) {
	var award db.UpdateMerchantCertificationOrAwardRow
	err := r.db.GetContext(ctx, &award, updateMerchantCertificationOrAward,
		int32(*request.MerchantCertificationID),
		request.Title,
		convert.NullableString(request.Description),
		convert.NullableString(request.IssuedBy),
		parseDateToPgDate(request.IssueDate),
		parseDateToPgDate(request.ExpiryDate),
		convert.NullableString(request.CertificateUrl),
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, merchantaward_errors.ErrMerchantAwardNotFound
		}
		return nil, merchantaward_errors.ErrUpdateMerchantAward.WithInternal(err)
	}

	return &award, nil
}

func (r *merchantAwardCommandRepository) Trash(ctx context.Context, award_id int) (*db.MerchantCertificationsAndAward, error) {
	var award db.MerchantCertificationsAndAward
	err := r.db.GetContext(ctx, &award, trashMerchantCertificationOrAward, int32(award_id))

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, merchantaward_errors.ErrMerchantAwardNotFound
		}
		return nil, merchantaward_errors.ErrTrashedMerchantAward.WithInternal(err)
	}

	return &award, nil
}

func (r *merchantAwardCommandRepository) Restore(ctx context.Context, award_id int) (*db.MerchantCertificationsAndAward, error) {
	var award db.MerchantCertificationsAndAward
	err := r.db.GetContext(ctx, &award, restoreMerchantCertificationOrAward, int32(award_id))

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, merchantaward_errors.ErrMerchantAwardNotFound
		}
		return nil, merchantaward_errors.ErrRestoreMerchantAward.WithInternal(err)
	}

	return &award, nil
}

func (r *merchantAwardCommandRepository) DeletePermanent(ctx context.Context, award_id int) (bool, error) {
	if _, err := r.db.ExecContext(ctx, deleteMerchantCertificationOrAwardPermanently, int32(award_id)); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, merchantaward_errors.ErrMerchantAwardNotFound
		}
		return false, merchantaward_errors.ErrDeleteMerchantAwardPermanent.WithInternal(err)
	}

	return true, nil
}

func (r *merchantAwardCommandRepository) RestoreAll(ctx context.Context) (bool, error) {
	if _, err := r.db.ExecContext(ctx, restoreAllMerchantCertificationsAndAwards); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, merchantaward_errors.ErrMerchantAwardNotFound
		}
		return false, merchantaward_errors.ErrRestoreAllMerchantAwards.WithInternal(err)
	}

	return true, nil
}

func (r *merchantAwardCommandRepository) DeleteAll(ctx context.Context) (bool, error) {
	if _, err := r.db.ExecContext(ctx, deleteAllPermanentMerchantCertificationsAndAwards); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, merchantaward_errors.ErrMerchantAwardNotFound
		}
		return false, merchantaward_errors.ErrDeleteAllMerchantAwardsPermanent.WithInternal(err)
	}

	return true, nil
}

func parseDateToPgDate(dateStr string) pgtype.Date {
	if dateStr == "" {
		return pgtype.Date{Valid: false}
	}

	t, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		return pgtype.Date{Valid: false}
	}

	return pgtype.Date{
		Time:  t,
		Valid: true,
	}
}
