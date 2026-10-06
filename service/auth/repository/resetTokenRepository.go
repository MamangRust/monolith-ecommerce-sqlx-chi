package repository

import (
	"context"
	"database/sql"
	"errors"
	"time"

	db "github.com/MamangRust/monolith-ecommerce-pkg/database/schema"
	"github.com/MamangRust/monolith-ecommerce-shared/domain/requests"
	sharedErrors "github.com/MamangRust/monolith-ecommerce-shared/errors"
	resettoken_errors "github.com/MamangRust/monolith-ecommerce-shared/errors/reset_token_errors"
	"github.com/jmoiron/sqlx"
)

// resetTokenRepository is a struct that implements the ResetTokenRepository interface
type resetTokenRepository struct {
	db *sqlx.DB
}

// NewResetTokenRepository creates a new instance of resetTokenRepository.
//
// Args:
// db: a pointer to the sqlx.DB connection, providing database query capabilities.
// mapper: a ResetTokenRecordMapping object to map database records to domain records.
//
// Returns:
// A pointer to a newly initialized resetTokenRepository struct.
func NewResetTokenRepository(db *sqlx.DB) *resetTokenRepository {
	return &resetTokenRepository{
		db: db,
	}
}

// FindByToken retrieves a reset token record by token string.
//
// Parameters:
//   - ctx: the context for the database operation
//   - token: the reset token to search for
//
// Returns:
//   - A ResetTokenRecord if found, or an error if not found or operation fails.
func (r *resetTokenRepository) FindByToken(ctx context.Context, code string) (*db.ResetToken, error) {
	var token db.ResetToken
	err := r.db.GetContext(ctx, &token, getResetToken, code)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, resettoken_errors.ErrTokenNotFound.WithInternal(err)
		}
		return nil, sharedErrors.ErrInternal.WithInternal(err)
	}

	return &token, nil
}

// CreateResetToken inserts a new reset token into the database.
//
// Parameters:
//   - ctx: the context for the database operation
//   - req: the request payload containing user ID and token info
//
// Returns:
//   - The created ResetTokenRecord, or an error if the operation fails.
//
// CreateResetTokenInTx persists the reset token inside the given database
// transaction so the caller can commit the token write and its outbox event
// atomically (Phase 6 — transactional outbox).
func (r *resetTokenRepository) CreateResetTokenInTx(ctx context.Context, tx *sqlx.Tx, req *requests.CreateResetTokenRequest) (*db.ResetToken, error) {
	return r.createResetToken(ctx, tx, req)
}

func (r *resetTokenRepository) CreateResetToken(ctx context.Context, req *requests.CreateResetTokenRequest) (*db.ResetToken, error) {
	return r.createResetToken(ctx, r.db, req)
}

func (r *resetTokenRepository) createResetToken(ctx context.Context, q sqlx.QueryerContext, req *requests.CreateResetTokenRequest) (*db.ResetToken, error) {
	expiryDate, err := time.Parse("2006-01-02 15:04:05", req.ExpiredAt)
	if err != nil {
		return nil, err
	}

	var token db.ResetToken
	if err := sqlx.GetContext(ctx, q, &token, createResetToken,
		int64(req.UserID),
		req.ResetToken,
		expiryDate,
	); err != nil {
		return nil, resettoken_errors.ErrCreateResetToken.WithInternal(err)
	}

	return &token, nil
}

// DeleteResetToken removes the reset token associated with the given user ID.
//
// Parameters:
//   - ctx: the context for the database operation
//   - userID: the user ID whose token should be deleted
//
// Returns:
//   - An error if the deletion fails.
func (r *resetTokenRepository) DeleteResetToken(ctx context.Context, user_id int) error {
	if _, err := r.db.ExecContext(ctx, deleteResetToken, int64(user_id)); err != nil {
		return resettoken_errors.ErrDeleteByUserID.WithInternal(err)
	}

	return nil
}

const createResetToken = `INSERT INTO
    reset_tokens (user_id, token, expiry_date)
VALUES ($1, $2, $3)
RETURNING
    id, user_id, token, expiry_date`

const deleteResetToken = `DELETE FROM reset_tokens WHERE user_id = $1`

const getResetToken = `SELECT id, user_id, token, expiry_date FROM reset_tokens WHERE token = $1`
