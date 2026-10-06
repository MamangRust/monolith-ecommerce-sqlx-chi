package repository

import (
	"context"
	"database/sql"
	"errors"

	db "github.com/MamangRust/monolith-ecommerce-pkg/database/schema"
	"github.com/MamangRust/monolith-ecommerce-shared/domain/requests"
	sharedErrors "github.com/MamangRust/monolith-ecommerce-shared/errors"
	"github.com/MamangRust/monolith-ecommerce-shared/errors/role_errors"
	"github.com/jmoiron/sqlx"
)

const getRoles = `SELECT
    role_id,
    role_name,
    created_at,
    updated_at,
    deleted_at,
    COUNT(*) OVER() AS total_count
FROM
    roles
WHERE
    $1::TEXT IS NULL OR role_name ILIKE '%' || $1 || '%'
ORDER BY
    created_at ASC
LIMIT $2 OFFSET $3`

const getRole = `SELECT
    role_id,
    role_name,
    created_at,
    updated_at,
    deleted_at
FROM
    roles
WHERE
    role_id = $1`

const getRoleByName = `SELECT
    role_id,
    role_name,
    created_at,
    updated_at,
    deleted_at
FROM
    roles
WHERE
    role_name = $1`

const getUserRoles = `SELECT
    r.role_id,
    r.role_name,
    r.created_at,
    r.updated_at,
    r.deleted_at
FROM
    roles r
JOIN
    user_roles ur ON ur.role_id = r.role_id
WHERE
    ur.user_id = $1
ORDER BY
    r.created_at ASC`

const getActiveRoles = `SELECT
    role_id,
    role_name,
    created_at,
    updated_at,
    deleted_at,
    COUNT(*) OVER() AS total_count
FROM
    roles
WHERE
    deleted_at IS NULL
    AND ($1::TEXT IS NULL OR role_name ILIKE '%' || $1 || '%')
ORDER BY
    created_at ASC
LIMIT $2 OFFSET $3`

const getTrashedRoles = `SELECT
    role_id,
    role_name,
    created_at,
    updated_at,
    deleted_at,
    COUNT(*) OVER() AS total_count
FROM
    roles
WHERE
    deleted_at IS NOT NULL
    AND ($1::TEXT IS NULL OR role_name ILIKE '%' || $1 || '%')
ORDER BY
    deleted_at DESC
LIMIT $2 OFFSET $3`

type roleQueryRepository struct {
	db *sqlx.DB
}

func NewRoleQueryRepository(db *sqlx.DB) *roleQueryRepository {
	return &roleQueryRepository{
		db: db,
	}
}

func (r *roleQueryRepository) FindAll(ctx context.Context, req *requests.FindAllRole) ([]*db.GetRolesRow, error) {
	offset := (req.Page - 1) * req.PageSize

	var res []*db.GetRolesRow
	err := r.db.SelectContext(ctx, &res, getRoles,
		req.Search,
		int32(req.PageSize),
		int32(offset),
	)

	if err != nil {
		return nil, role_errors.ErrFindAllRoles.WithInternal(err)
	}

	return res, nil
}

func (r *roleQueryRepository) FindByID(ctx context.Context, id int) (*db.Role, error) {
	var role db.Role
	err := r.db.GetContext(ctx, &role, getRole, int32(id))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, role_errors.ErrRoleNotFound.WithInternal(err)
		}
		return nil, sharedErrors.ErrInternal.WithInternal(err)
	}

	return &role, nil
}

func (r *roleQueryRepository) FindByName(ctx context.Context, name string) (*db.Role, error) {
	var role db.Role
	err := r.db.GetContext(ctx, &role, getRoleByName, name)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, role_errors.ErrRoleNotFound.WithInternal(err)
		}

		return nil, sharedErrors.ErrInternal.WithInternal(err)
	}

	return &role, nil
}

func (r *roleQueryRepository) FindByUserId(ctx context.Context, user_id int) ([]*db.Role, error) {
	var res []*db.Role
	err := r.db.SelectContext(ctx, &res, getUserRoles, int32(user_id))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, role_errors.ErrRoleNotFound.WithInternal(err)
		}

		return nil, sharedErrors.ErrInternal.WithInternal(err)
	}

	return res, nil
}

func (r *roleQueryRepository) FindActive(ctx context.Context, req *requests.FindAllRole) ([]*db.GetActiveRolesRow, error) {
	offset := (req.Page - 1) * req.PageSize

	var res []*db.GetActiveRolesRow
	err := r.db.SelectContext(ctx, &res, getActiveRoles,
		req.Search,
		int32(req.PageSize),
		int32(offset),
	)

	if err != nil {
		return nil, role_errors.ErrFindActiveRoles.WithInternal(err)
	}

	return res, nil
}

func (r *roleQueryRepository) FindTrashed(ctx context.Context, req *requests.FindAllRole) ([]*db.GetTrashedRolesRow, error) {
	offset := (req.Page - 1) * req.PageSize

	var res []*db.GetTrashedRolesRow
	err := r.db.SelectContext(ctx, &res, getTrashedRoles,
		req.Search,
		int32(req.PageSize),
		int32(offset),
	)

	if err != nil {
		return nil, role_errors.ErrFindTrashedRoles.WithInternal(err)
	}

	return res, nil
}
