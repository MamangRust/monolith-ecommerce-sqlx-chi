package repository

import (
	"context"
	"database/sql"
	"errors"

	db "github.com/MamangRust/monolith-ecommerce-pkg/database/schema"
	"github.com/MamangRust/monolith-ecommerce-shared/convert"
	"github.com/MamangRust/monolith-ecommerce-shared/domain/requests"
	merchant_errors "github.com/MamangRust/monolith-ecommerce-shared/errors/merchant"
	"github.com/jmoiron/sqlx"
)

type merchantDocumentCommandRepository struct {
	db *sqlx.DB
}

func NewMerchantDocumentCommandRepository(db *sqlx.DB) *merchantDocumentCommandRepository {
	return &merchantDocumentCommandRepository{
		db: db,
	}
}

const createMerchantDocument = `INSERT INTO
    merchant_documents (
        merchant_id,
        document_type,
        document_url,
        status,
        note,
        uploaded_at,
        updated_at
    )
VALUES (
        $1,
        $2,
        $3,
        $4,
        $5,
        current_timestamp,
        current_timestamp
    )
RETURNING
    document_id,
    merchant_id,
    document_type,
    document_url,
    status,
    note,
    uploaded_at,
    created_at,
    updated_at`

const updateMerchantDocument = `UPDATE merchant_documents
SET
    document_type = $2,
    document_url = $3,
    status = $4,
    note = $5,
    updated_at = current_timestamp
WHERE
    document_id = $1
    AND deleted_at IS NULL
RETURNING
    document_id,
    merchant_id,
    document_type,
    document_url,
    status,
    note,
    uploaded_at,
    created_at,
    updated_at`

const updateMerchantDocumentStatus = `UPDATE merchant_documents
SET
    status = $2,
    note = $3,
    updated_at = current_timestamp
WHERE
    document_id = $1
    AND deleted_at IS NULL
RETURNING
    document_id,
    merchant_id,
    document_type,
    document_url,
    status,
    note,
    uploaded_at,
    created_at,
    updated_at`

const trashMerchantDocument = `UPDATE merchant_documents
SET
    deleted_at = current_timestamp
WHERE
    document_id = $1
    AND deleted_at IS NULL
RETURNING
    document_id,
    merchant_id,
    document_type,
    document_url,
    status,
    note,
    uploaded_at,
    created_at,
    updated_at,
    deleted_at`

const restoreMerchantDocument = `UPDATE merchant_documents
SET
    deleted_at = NULL
WHERE
    document_id = $1
    AND deleted_at IS NOT NULL
RETURNING
    document_id,
    merchant_id,
    document_type,
    document_url,
    status,
    note,
    uploaded_at,
    created_at,
    updated_at,
    deleted_at`

const deleteMerchantDocumentPermanently = `DELETE FROM merchant_documents
WHERE
    document_id = $1
    AND deleted_at IS NOT NULL`

const restoreAllMerchantDocuments = `UPDATE merchant_documents
SET
    deleted_at = NULL
WHERE
    deleted_at IS NOT NULL`

const deleteAllPermanentMerchantDocuments = `DELETE FROM merchant_documents WHERE deleted_at IS NOT NULL`

func (r *merchantDocumentCommandRepository) Create(ctx context.Context, request *requests.CreateMerchantDocumentRequest) (*db.CreateMerchantDocumentRow, error) {
	return r.createMerchantDocument(ctx, r.db, request)
}

// CreateInTx persists the document inside the given database transaction so the
// caller can commit the business write and its outbox event atomically (Phase 6
// — transactional outbox).
func (r *merchantDocumentCommandRepository) CreateInTx(ctx context.Context, tx *sqlx.Tx, request *requests.CreateMerchantDocumentRequest) (*db.CreateMerchantDocumentRow, error) {
	return r.createMerchantDocument(ctx, tx, request)
}

func (r *merchantDocumentCommandRepository) createMerchantDocument(ctx context.Context, q sqlx.QueryerContext, request *requests.CreateMerchantDocumentRequest) (*db.CreateMerchantDocumentRow, error) {
	var doc db.CreateMerchantDocumentRow
	if err := sqlx.GetContext(ctx, q, &doc, createMerchantDocument,
		int32(request.MerchantID),
		request.DocumentType,
		request.DocumentUrl,
		"pending",
		convert.NullableString(""),
	); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, merchant_errors.ErrMerchantNotFound
		}
		return nil, merchant_errors.ErrMerchantInternal.WithInternal(err)
	}

	return &doc, nil
}

func (r *merchantDocumentCommandRepository) Update(ctx context.Context, request *requests.UpdateMerchantDocumentRequest) (*db.UpdateMerchantDocumentRow, error) {
	var doc db.UpdateMerchantDocumentRow
	if err := r.db.GetContext(ctx, &doc, updateMerchantDocument,
		int32(*request.DocumentID),
		request.DocumentType,
		request.DocumentUrl,
		request.Status,
		convert.NullableString(request.Note),
	); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, merchant_errors.ErrMerchantNotFound
		}
		return nil, merchant_errors.ErrMerchantInternal.WithInternal(err)
	}

	return &doc, nil
}

func (r *merchantDocumentCommandRepository) UpdateStatus(ctx context.Context, request *requests.UpdateMerchantDocumentStatusRequest) (*db.UpdateMerchantDocumentStatusRow, error) {
	return r.updateMerchantDocumentStatus(ctx, r.db, request)
}

// UpdateStatusInTx updates the document status inside the given database
// transaction so the caller can commit the business write and its outbox event
// atomically (Phase 6 — transactional outbox).
func (r *merchantDocumentCommandRepository) UpdateStatusInTx(ctx context.Context, tx *sqlx.Tx, request *requests.UpdateMerchantDocumentStatusRequest) (*db.UpdateMerchantDocumentStatusRow, error) {
	return r.updateMerchantDocumentStatus(ctx, tx, request)
}

func (r *merchantDocumentCommandRepository) updateMerchantDocumentStatus(ctx context.Context, q sqlx.QueryerContext, request *requests.UpdateMerchantDocumentStatusRequest) (*db.UpdateMerchantDocumentStatusRow, error) {
	var doc db.UpdateMerchantDocumentStatusRow
	if err := sqlx.GetContext(ctx, q, &doc, updateMerchantDocumentStatus,
		int32(*request.DocumentID),
		request.Status,
		convert.NullableString(request.Note),
	); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, merchant_errors.ErrMerchantNotFound
		}
		return nil, merchant_errors.ErrMerchantInternal.WithInternal(err)
	}

	return &doc, nil
}

func (r *merchantDocumentCommandRepository) Trash(ctx context.Context, documentID int) (*db.MerchantDocument, error) {
	var doc db.MerchantDocument
	if err := r.db.GetContext(ctx, &doc, trashMerchantDocument, int32(documentID)); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, merchant_errors.ErrMerchantNotFound
		}
		return nil, merchant_errors.ErrMerchantInternal.WithInternal(err)
	}

	return &doc, nil
}

func (r *merchantDocumentCommandRepository) Restore(ctx context.Context, documentID int) (*db.MerchantDocument, error) {
	var doc db.MerchantDocument
	if err := r.db.GetContext(ctx, &doc, restoreMerchantDocument, int32(documentID)); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, merchant_errors.ErrMerchantNotFound
		}
		return nil, merchant_errors.ErrMerchantInternal.WithInternal(err)
	}

	return &doc, nil
}

func (r *merchantDocumentCommandRepository) DeletePermanent(ctx context.Context, documentID int) (bool, error) {
	if _, err := r.db.ExecContext(ctx, deleteMerchantDocumentPermanently, int32(documentID)); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, merchant_errors.ErrMerchantNotFound
		}
		return false, merchant_errors.ErrMerchantInternal.WithInternal(err)
	}

	return true, nil
}

func (r *merchantDocumentCommandRepository) RestoreAll(ctx context.Context) (bool, error) {
	if _, err := r.db.ExecContext(ctx, restoreAllMerchantDocuments); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, merchant_errors.ErrMerchantNotFound
		}
		return false, merchant_errors.ErrMerchantInternal.WithInternal(err)
	}

	return true, nil
}

func (r *merchantDocumentCommandRepository) DeleteAll(ctx context.Context) (bool, error) {
	if _, err := r.db.ExecContext(ctx, deleteAllPermanentMerchantDocuments); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, merchant_errors.ErrMerchantNotFound
		}
		return false, merchant_errors.ErrMerchantInternal.WithInternal(err)
	}

	return true, nil
}
