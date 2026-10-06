package repository

import (
	"context"
	"database/sql"
	"errors"

	db "github.com/MamangRust/monolith-ecommerce-pkg/database/schema"
	"github.com/MamangRust/monolith-ecommerce-shared/domain/requests"
	"github.com/MamangRust/monolith-ecommerce-shared/errors/role_errors"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jmoiron/sqlx"
)

const createRole = `INSERT INTO roles (
    role_name,
    created_at,
    updated_at
) VALUES (
    $1,
    current_timestamp,
    current_timestamp
) RETURNING
    role_id,
    role_name,
    created_at,
    updated_at,
    deleted_at`

const updateRole = `UPDATE roles
SET
    role_name = $2,
    updated_at = current_timestamp
WHERE
    role_id = $1
RETURNING
    role_id,
    role_name,
    created_at,
    updated_at,
    deleted_at`

const trashRole = `UPDATE roles
SET
    deleted_at = current_timestamp
WHERE
    role_id = $1
RETURNING role_id, role_name, created_at, updated_at, deleted_at`

const restoreRole = `UPDATE roles
SET
    deleted_at = NULL
WHERE
    role_id = $1
RETURNING role_id, role_name, created_at, updated_at, deleted_at`

const deletePermanentRole = `DELETE FROM roles
WHERE
    role_id = $1 AND deleted_at IS NOT NULL`

const restoreAllRoles = `UPDATE roles
SET
    deleted_at = NULL
WHERE
    deleted_at IS NOT NULL`

const deleteAllPermanentRoles = `DELETE FROM roles
WHERE
    deleted_at IS NOT NULL`

type roleCommandRepository struct {
	db *sqlx.DB
}

func NewRoleCommandRepository(db *sqlx.DB) *roleCommandRepository {
	return &roleCommandRepository{
		db: db,
	}
}

func (r *roleCommandRepository) Create(ctx context.Context, req *requests.CreateRoleRequest) (*db.Role, error) {
	var role db.Role
	err := r.db.GetContext(ctx, &role, createRole, req.Name)

	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return nil, role_errors.ErrRoleConflict
		}
		return nil, role_errors.ErrCreateRole.WithInternal(err)
	}

	return &role, nil
}

func (r *roleCommandRepository) Update(ctx context.Context, req *requests.UpdateRoleRequest) (*db.Role, error) {
	var role db.Role
	err := r.db.GetContext(ctx, &role, updateRole,
		int32(*req.ID),
		req.Name,
	)

	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return nil, role_errors.ErrRoleConflict
		}
		return nil, role_errors.ErrUpdateRole.WithInternal(err)
	}

	return &role, nil
}

func (r *roleCommandRepository) Trash(ctx context.Context, id int) (*db.Role, error) {
	var role db.Role
	err := r.db.GetContext(ctx, &role, trashRole, int32(id))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, role_errors.ErrRoleNotFound
		}
		return nil, role_errors.ErrTrashedRole.WithInternal(err)
	}
	return &role, nil
}

func (r *roleCommandRepository) Restore(ctx context.Context, id int) (*db.Role, error) {
	var role db.Role
	err := r.db.GetContext(ctx, &role, restoreRole, int32(id))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, role_errors.ErrRoleNotFound
		}
		return nil, role_errors.ErrRestoreRole.WithInternal(err)
	}
	return &role, nil
}

func (r *roleCommandRepository) DeletePermanent(ctx context.Context, role_id int) (bool, error) {
	if _, err := r.db.ExecContext(ctx, deletePermanentRole, int32(role_id)); err != nil {
		return false, role_errors.ErrDeleteRolePermanent.WithInternal(err)
	}
	return true, nil
}

func (r *roleCommandRepository) RestoreAll(ctx context.Context) (bool, error) {
	if _, err := r.db.ExecContext(ctx, restoreAllRoles); err != nil {
		return false, role_errors.ErrRestoreAllRoles.WithInternal(err)
	}

	return true, nil
}

func (r *roleCommandRepository) DeleteAll(ctx context.Context) (bool, error) {
	if _, err := r.db.ExecContext(ctx, deleteAllPermanentRoles); err != nil {
		return false, role_errors.ErrDeleteAllRoles.WithInternal(err)
	}

	return true, nil
}
