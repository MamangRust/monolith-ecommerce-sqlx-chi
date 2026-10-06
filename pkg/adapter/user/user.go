// Package user adapts the User service gRPC API into the shared pkg/database
// sqlc row types used by the consumer repositories.
package user

import (
	"context"

	pbuser "github.com/MamangRust/monolith-ecommerce-pb/user"
	"github.com/MamangRust/monolith-ecommerce-pkg/adapter"
	"github.com/MamangRust/monolith-ecommerce-pkg/adapter/convert"
	db "github.com/MamangRust/monolith-ecommerce-pkg/database/schema"
	"github.com/MamangRust/monolith-ecommerce-pkg/resilience"
	"github.com/MamangRust/monolith-ecommerce-shared/domain/requests"
	user_errors "github.com/MamangRust/monolith-ecommerce-shared/errors/user_errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// QueryRepository is the contract consumers depend on for user reads.
type QueryRepository interface {
	FindById(ctx context.Context, userID int) (*db.User, error)
	FindByID(ctx context.Context, userID int) (*db.User, error)
	FindByEmail(ctx context.Context, email string) (*db.User, error)
	FindByEmailAndVerify(ctx context.Context, email string) (*db.User, error)
	FindByVerificationCode(ctx context.Context, code string) (*db.User, error)
}

// CommandRepository is the contract consumers depend on for user writes.
type CommandRepository interface {
	CreateUser(ctx context.Context, request *requests.RegisterRequest) (*db.User, error)
	UpdateUserIsVerified(ctx context.Context, userID int, isVerified bool) (*db.User, error)
	UpdateUserPassword(ctx context.Context, userID int, password string) (*db.User, error)
}

// Repository implements QueryRepository and CommandRepository over the generated
// user query/command clients, guarding every call.
type Repository struct {
	query   pbuser.UserQueryServiceClient
	command pbuser.UserCommandServiceClient
	guard   *resilience.DependencyGuard
}

// New builds a user adapter. Passing zero options leaves the guard nil, which
// makes DependencyGuard.Call a plain passthrough.
func New(query pbuser.UserQueryServiceClient, command pbuser.UserCommandServiceClient, opts ...adapter.GuardOption) *Repository {
	r := &Repository{query: query, command: command}
	for _, opt := range opts {
		opt(r)
	}
	return r
}

// NewQueryAdapter builds a read-only adapter for consumers that never write.
func NewQueryAdapter(query pbuser.UserQueryServiceClient, opts ...adapter.GuardOption) QueryRepository {
	return New(query, nil, opts...)
}

// NewCommandAdapter builds a write-only adapter for consumers that never read.
func NewCommandAdapter(command pbuser.UserCommandServiceClient, opts ...adapter.GuardOption) CommandRepository {
	return New(nil, command, opts...)
}

func (r *Repository) SetGuard(g *resilience.DependencyGuard) {
	r.guard = g
}

var (
	_ QueryRepository   = (*Repository)(nil)
	_ CommandRepository = (*Repository)(nil)
)

func (r *Repository) FindById(ctx context.Context, userID int) (*db.User, error) {
	return r.findByID(ctx, userID)
}

func (r *Repository) FindByID(ctx context.Context, userID int) (*db.User, error) {
	return r.findByID(ctx, userID)
}

func (r *Repository) findByID(ctx context.Context, userID int) (*db.User, error) {
	resp, err := adapter.Call(r.guard, ctx, func(callCtx context.Context) (*pbuser.ApiResponseUser, error) {
		return r.query.FindById(callCtx, &pbuser.FindByIdUserRequest{Id: int32(userID)})
	})
	if err != nil {
		return nil, user_errors.ErrUserNotFound.WithInternal(err)
	}
	if resp == nil || resp.Data == nil {
		return nil, user_errors.ErrUserNotFound
	}

	return mapUser(resp.Data), nil
}

// FindByEmail returns the account with its password hash so login can verify
// credentials. A missing account is reported as (nil, nil) so callers can tell
// "unknown email" apart from a failing dependency.
func (r *Repository) FindByEmail(ctx context.Context, email string) (*db.User, error) {
	resp, err := adapter.Call(r.guard, ctx, func(callCtx context.Context) (*pbuser.ApiResponseUserWithPassword, error) {
		return r.query.FindByEmail(callCtx, &pbuser.FindByEmailRequest{Email: email})
	})
	if err != nil {
		if st, ok := status.FromError(err); ok && st.Code() == codes.NotFound {
			return nil, nil
		}
		return nil, user_errors.ErrUserNotFound.WithInternal(err)
	}
	if resp == nil || resp.Data == nil {
		return nil, nil
	}

	return mapUserWithPassword(resp.Data), nil
}

// FindByEmailAndVerify is the login-time lookup; the User service exposes a
// single email lookup, so this reuses FindByEmail.
func (r *Repository) FindByEmailAndVerify(ctx context.Context, email string) (*db.User, error) {
	resp, err := adapter.Call(r.guard, ctx, func(callCtx context.Context) (*pbuser.ApiResponseUserWithPassword, error) {
		return r.query.FindByEmail(callCtx, &pbuser.FindByEmailRequest{Email: email})
	})
	if err != nil {
		return nil, user_errors.ErrUserNotFound.WithInternal(err)
	}
	if resp == nil || resp.Data == nil {
		return nil, user_errors.ErrUserNotFound
	}

	return mapUserWithPassword(resp.Data), nil
}

// FindByVerificationCode exchanges the emailed code for the account it belongs
// to, backing the VerifyCode flow.
func (r *Repository) FindByVerificationCode(ctx context.Context, code string) (*db.User, error) {
	resp, err := adapter.Call(r.guard, ctx, func(callCtx context.Context) (*pbuser.ApiResponseUser, error) {
		return r.query.FindByVerificationCode(callCtx, &pbuser.FindByVerificationCodeRequest{VerificationCode: code})
	})
	if err != nil {
		return nil, user_errors.ErrUserNotFound.WithInternal(err)
	}
	if resp == nil || resp.Data == nil {
		return nil, user_errors.ErrUserNotFound
	}

	return mapUser(resp.Data), nil
}

func (r *Repository) CreateUser(ctx context.Context, request *requests.RegisterRequest) (*db.User, error) {
	resp, err := adapter.Call(r.guard, ctx, func(callCtx context.Context) (*pbuser.ApiResponseUser, error) {
		return r.command.Create(callCtx, &pbuser.CreateUserRequest{
			Firstname:       request.FirstName,
			Lastname:        request.LastName,
			Email:           request.Email,
			Password:        request.Password,
			ConfirmPassword: request.ConfirmPassword,
		})
	})
	if err != nil {
		return nil, user_errors.ErrCreateUser.WithInternal(err)
	}
	if resp == nil || resp.Data == nil {
		return nil, user_errors.ErrCreateUser
	}

	return mapUser(resp.Data), nil
}

func (r *Repository) UpdateUserIsVerified(ctx context.Context, userID int, isVerified bool) (*db.User, error) {
	resp, err := adapter.Call(r.guard, ctx, func(callCtx context.Context) (*pbuser.ApiResponseUser, error) {
		return r.command.UpdateIsVerified(callCtx, &pbuser.UpdateUserIsVerifiedRequest{
			Id:         int32(userID),
			IsVerified: isVerified,
		})
	})
	if err != nil {
		return nil, user_errors.ErrUpdateUserVerificationCode.WithInternal(err)
	}
	if resp == nil || resp.Data == nil {
		return nil, user_errors.ErrUpdateUserVerificationCode
	}

	return mapUser(resp.Data), nil
}

func (r *Repository) UpdateUserPassword(ctx context.Context, userID int, password string) (*db.User, error) {
	resp, err := adapter.Call(r.guard, ctx, func(callCtx context.Context) (*pbuser.ApiResponseUser, error) {
		return r.command.UpdatePassword(callCtx, &pbuser.UpdateUserPasswordRequest{
			Id:       int32(userID),
			Password: password,
		})
	})
	if err != nil {
		return nil, user_errors.ErrUpdateUserPassword.WithInternal(err)
	}
	if resp == nil || resp.Data == nil {
		return nil, user_errors.ErrUpdateUserPassword
	}

	return mapUser(resp.Data), nil
}

// mapUser converts a proto user into the sqlc row type.
func mapUser(u *pbuser.UserResponse) *db.User {
	if u == nil {
		return nil
	}
	return &db.User{
		UserID:    u.Id,
		Firstname: u.Firstname,
		Lastname:  u.Lastname,
		Email:     u.Email,
		CreatedAt: convert.Timestamp(u.CreatedAt),
		UpdatedAt: convert.Timestamp(u.UpdatedAt),
	}
}

// mapUserWithPassword converts a proto user-with-password into the sqlc row type.
func mapUserWithPassword(u *pbuser.UserResponseWithPassword) *db.User {
	if u == nil {
		return nil
	}
	return &db.User{
		UserID:    u.Id,
		Firstname: u.Firstname,
		Lastname:  u.Lastname,
		Email:     u.Email,
		Password:  u.Password,
		CreatedAt: convert.Timestamp(u.CreatedAt),
		UpdatedAt: convert.Timestamp(u.UpdatedAt),
	}
}
