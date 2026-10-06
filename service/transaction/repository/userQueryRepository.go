package repository

import (
	"context"

	useradapter "github.com/MamangRust/monolith-ecommerce-pkg/adapter/user"
	db "github.com/MamangRust/monolith-ecommerce-pkg/database/schema"
)

// userQueryRepository delegates user lookups to the shared gRPC adapter.
// Failures keep propagating the dependency's raw gRPC status, as they did
// before the adapter migration.
type userQueryRepository struct {
	adapter useradapter.QueryRepository
}

func NewUserQueryRepository(adapter useradapter.QueryRepository) *userQueryRepository {
	return &userQueryRepository{
		adapter: adapter,
	}
}

func (r *userQueryRepository) FindByID(ctx context.Context, user_id int) (*db.User, error) {
	return r.adapter.FindByID(ctx, user_id)
}
