package repository

import (
	"context"

	useradapter "github.com/MamangRust/monolith-ecommerce-pkg/adapter/user"
	db "github.com/MamangRust/monolith-ecommerce-pkg/database/schema"
	"github.com/MamangRust/monolith-ecommerce-shared/domain/requests"
	"github.com/MamangRust/monolith-ecommerce-shared/errors/user_errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// userRepository delegates user access to the shared gRPC adapter and re-wraps
// failures in this service's own errors so callers keep the status codes they
// had before the adapter migration.
type userRepository struct {
	adapter *useradapter.Repository
}

// NewUserRepository returns a new instance of userRepository backed by the
// shared gRPC user adapter.
func NewUserRepository(adapter *useradapter.Repository) *userRepository {
	return &userRepository{
		adapter: adapter,
	}
}

// FindByID retrieves a user by their unique ID.
func (r *userRepository) FindByID(ctx context.Context, user_id int) (*db.User, error) {
	res, err := r.adapter.FindByID(ctx, user_id)
	if err != nil {
		return nil, user_errors.ErrUserNotFound.WithInternal(err)
	}

	return res, nil
}

// FindByEmail retrieves a user by their email address. A missing user is
// reported as (nil, nil) — the login/register flows rely on that to tell
// "unknown email" apart from a failing dependency.
func (r *userRepository) FindByEmail(ctx context.Context, email string) (*db.User, error) {
	res, err := r.adapter.FindByEmail(ctx, email)
	if err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, nil
		}
		return nil, user_errors.ErrUserNotFound.WithInternal(err)
	}

	return res, nil
}

// FindByEmailAndVerify retrieves a user by their email address for the login
// path. Unlike FindByEmail it never short-circuits on codes.NotFound.
func (r *userRepository) FindByEmailAndVerify(ctx context.Context, email string) (*db.User, error) {
	res, err := r.adapter.FindByEmailAndVerify(ctx, email)
	if err != nil {
		return nil, user_errors.ErrUserNotFound.WithInternal(err)
	}

	return res, nil
}

// FindByVerificationCode retrieves a user by their verification code.
func (r *userRepository) FindByVerificationCode(ctx context.Context, verification_code string) (*db.User, error) {
	res, err := r.adapter.FindByVerificationCode(ctx, verification_code)
	if err != nil {
		return nil, user_errors.ErrUserNotFound.WithInternal(err)
	}

	return res, nil
}

// CreateUser inserts a new user.
func (r *userRepository) CreateUser(ctx context.Context, request *requests.RegisterRequest) (*db.User, error) {
	res, err := r.adapter.CreateUser(ctx, request)
	if err != nil {
		return nil, user_errors.ErrCreateUser.WithInternal(err)
	}

	return res, nil
}

// UpdateUserIsVerified updates the verification status of a user.
func (r *userRepository) UpdateUserIsVerified(ctx context.Context, user_id int, is_verified bool) (*db.User, error) {
	res, err := r.adapter.UpdateUserIsVerified(ctx, user_id, is_verified)
	if err != nil {
		return nil, user_errors.ErrUpdateUserVerificationCode.WithInternal(err)
	}

	return res, nil
}

// UpdateUserPassword updates a user's password.
func (r *userRepository) UpdateUserPassword(ctx context.Context, user_id int, password string) (*db.User, error) {
	res, err := r.adapter.UpdateUserPassword(ctx, user_id, password)
	if err != nil {
		return nil, user_errors.ErrUpdateUserPassword.WithInternal(err)
	}

	return res, nil
}
