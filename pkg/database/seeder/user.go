package seeder

import (
	"context"

	db "github.com/MamangRust/monolith-ecommerce-pkg/database/schema"
	"github.com/MamangRust/monolith-ecommerce-pkg/hash"
	"github.com/MamangRust/monolith-ecommerce-pkg/logger"
	"github.com/jmoiron/sqlx"

	"fmt"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

const seedCreateUser = `INSERT INTO
    users (
        firstname,
        lastname,
        email,
        password,
        verification_code,
        is_verified,
        created_at,
        updated_at
    )
VALUES (
        $1,
        $2,
        $3,
        $4,
        $5,
        $6,
        current_timestamp,
        current_timestamp
    )
RETURNING
    user_id,
    firstname,
    lastname,
    email,
    password,
    created_at,
    updated_at`

const seedTrashUser = `UPDATE users
SET
    deleted_at = current_timestamp
WHERE
    user_id = $1
    AND deleted_at IS NULL
RETURNING
    user_id,
    firstname,
    lastname,
    email,
    created_at,
    updated_at,
    deleted_at`

type userSeeder struct {
	db     *sqlx.DB
	hash   hash.HashPassword
	ctx    context.Context
	logger logger.LoggerInterface
}

func NewUserSeeder(db *sqlx.DB, hash hash.HashPassword, ctx context.Context, logger logger.LoggerInterface) *userSeeder {
	return &userSeeder{
		db:     db,
		hash:   hash,
		ctx:    ctx,
		logger: logger,
	}
}

func (r *userSeeder) Seed() error {
	for i := 1; i <= 10; i++ {
		email := fmt.Sprintf("user_%s@example.com", uuid.NewString())
		rawPassword := fmt.Sprintf("password%d", i)

		hashedPassword, err := r.hash.HashPassword(rawPassword)
		if err != nil {
			r.logger.Error("failed to hash password", zap.Int("user", i), zap.Error(err))
			return fmt.Errorf("failed to hash password for user %d: %w", i, err)
		}

		var createdUser db.CreateUserRow
		err = r.db.GetContext(r.ctx, &createdUser, seedCreateUser,
			fmt.Sprintf("User%d", i),
			fmt.Sprintf("Last%d", i),
			email,
			hashedPassword,
			"",
			nil,
		)
		if err != nil {
			r.logger.Error("failed to seed user", zap.Int("user", i), zap.Error(err))
			return fmt.Errorf("failed to seed user %d: %w", i, err)
		}

		if i > 5 {
			var trashedUser db.TrashUserRow
			err = r.db.GetContext(r.ctx, &trashedUser, seedTrashUser, createdUser.UserID)
			if err != nil {
				r.logger.Error("failed to trash user", zap.Int("user", i), zap.Error(err))
				return fmt.Errorf("failed to trash user %d: %w", i, err)
			}
		}
	}

	r.logger.Info("User seeding completed successfully")
	return nil
}
