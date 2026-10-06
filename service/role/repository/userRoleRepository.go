package repository

import (
	"context"

	db "github.com/MamangRust/monolith-ecommerce-pkg/database/schema"
	"github.com/MamangRust/monolith-ecommerce-shared/domain/requests"
	"github.com/MamangRust/monolith-ecommerce-shared/errors/role_errors"
	"github.com/jmoiron/sqlx"
)

const assignRoleToUser = `INSERT INTO user_roles (
    user_id,
    role_id,
    created_at,
    updated_at
) VALUES (
    $1,
    $2,
    current_timestamp,
    current_timestamp
) RETURNING
    user_role_id,
    user_id,
    role_id,
    created_at,
    updated_at,
    deleted_at`

const removeRoleFromUser = `DELETE FROM user_roles
WHERE
    user_id = $1
    AND role_id = $2`

type userRoleRepository struct {
	db *sqlx.DB
}

func NewUserRoleRepository(db *sqlx.DB) UserRoleRepository {
	return &userRoleRepository{
		db: db,
	}
}

func (r *userRoleRepository) AssignRoleToUser(ctx context.Context, req *requests.CreateUserRoleRequest) (*db.UserRole, error) {
	var userRole db.UserRole
	err := r.db.GetContext(ctx, &userRole, assignRoleToUser, int32(req.UserId), int32(req.RoleId))
	if err != nil {
		return nil, role_errors.ErrAssignRole.WithInternal(err)
	}

	return &userRole, nil
}

func (r *userRoleRepository) RemoveRoleFromUser(ctx context.Context, req *requests.RemoveUserRoleRequest) error {
	if _, err := r.db.ExecContext(ctx, removeRoleFromUser, int32(req.UserId), int32(req.RoleId)); err != nil {
		return role_errors.ErrRemoveRole.WithInternal(err)
	}

	return nil
}
