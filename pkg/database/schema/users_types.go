// Code migrated from sqlc-generated sources; domain types kept for repository layer.
// source: users.sql

package db

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type CreateUserParams struct {
	Firstname        string `db:"firstname" json:"firstname"`
	Lastname         string `db:"lastname" json:"lastname"`
	Email            string `db:"email" json:"email"`
	Password         string `db:"password" json:"password"`
	VerificationCode string `db:"verification_code" json:"verification_code"`
	IsVerified       *bool  `db:"is_verified" json:"is_verified"`
}

type CreateUserRow struct {
	UserID    int32            `db:"user_id" json:"user_id"`
	Firstname string           `db:"firstname" json:"firstname"`
	Lastname  string           `db:"lastname" json:"lastname"`
	Email     string           `db:"email" json:"email"`
	Password  string           `db:"password" json:"password"`
	CreatedAt pgtype.Timestamp `db:"created_at" json:"created_at"`
	UpdatedAt pgtype.Timestamp `db:"updated_at" json:"updated_at"`
}

type GetUserByEmailAndVerifyRow struct {
	UserID    int32            `db:"user_id" json:"user_id"`
	Firstname string           `db:"firstname" json:"firstname"`
	Lastname  string           `db:"lastname" json:"lastname"`
	Email     string           `db:"email" json:"email"`
	Password  string           `db:"password" json:"password"`
	CreatedAt pgtype.Timestamp `db:"created_at" json:"created_at"`
	UpdatedAt pgtype.Timestamp `db:"updated_at" json:"updated_at"`
}

type GetUserByEmailWithPasswordRow struct {
	UserID   int32  `db:"user_id" json:"user_id"`
	Email    string `db:"email" json:"email"`
	Password string `db:"password" json:"password"`
}

type GetUserByIDRow struct {
	UserID    int32            `db:"user_id" json:"user_id"`
	Firstname string           `db:"firstname" json:"firstname"`
	Lastname  string           `db:"lastname" json:"lastname"`
	Email     string           `db:"email" json:"email"`
	Password  string           `db:"password" json:"password"`
	CreatedAt pgtype.Timestamp `db:"created_at" json:"created_at"`
	UpdatedAt pgtype.Timestamp `db:"updated_at" json:"updated_at"`
}

type GetUserByIdAndVerifyRow struct {
	UserID int32  `db:"user_id" json:"user_id"`
	Email  string `db:"email" json:"email"`
}

type GetUserByVerificationCodeRow struct {
	UserID    int32            `db:"user_id" json:"user_id"`
	Firstname string           `db:"firstname" json:"firstname"`
	Lastname  string           `db:"lastname" json:"lastname"`
	Email     string           `db:"email" json:"email"`
	Password  string           `db:"password" json:"password"`
	CreatedAt pgtype.Timestamp `db:"created_at" json:"created_at"`
	UpdatedAt pgtype.Timestamp `db:"updated_at" json:"updated_at"`
}

type GetUserTrashedParams struct {
	Column1 string `db:"column1" json:"column_1"`
	Limit   int32  `db:"limit" json:"limit"`
	Offset  int32  `db:"offset" json:"offset"`
}

type GetUserTrashedRow struct {
	UserID     int32            `db:"user_id" json:"user_id"`
	Firstname  string           `db:"firstname" json:"firstname"`
	Lastname   string           `db:"lastname" json:"lastname"`
	Email      string           `db:"email" json:"email"`
	Password   string           `db:"password" json:"password"`
	CreatedAt  pgtype.Timestamp `db:"created_at" json:"created_at"`
	UpdatedAt  pgtype.Timestamp `db:"updated_at" json:"updated_at"`
	DeletedAt  pgtype.Timestamp `db:"deleted_at" json:"deleted_at"`
	TotalCount int64            `db:"total_count" json:"total_count"`
}

type GetUsersParams struct {
	Column1 string `db:"column1" json:"column_1"`
	Limit   int32  `db:"limit" json:"limit"`
	Offset  int32  `db:"offset" json:"offset"`
}

type GetUsersRow struct {
	UserID     int32            `db:"user_id" json:"user_id"`
	Firstname  string           `db:"firstname" json:"firstname"`
	Lastname   string           `db:"lastname" json:"lastname"`
	Email      string           `db:"email" json:"email"`
	Password   string           `db:"password" json:"password"`
	CreatedAt  pgtype.Timestamp `db:"created_at" json:"created_at"`
	UpdatedAt  pgtype.Timestamp `db:"updated_at" json:"updated_at"`
	TotalCount int64            `db:"total_count" json:"total_count"`
}

type GetUsersActiveParams struct {
	Column1 string `db:"column1" json:"column_1"`
	Limit   int32  `db:"limit" json:"limit"`
	Offset  int32  `db:"offset" json:"offset"`
}

type GetUsersActiveRow struct {
	UserID     int32            `db:"user_id" json:"user_id"`
	Firstname  string           `db:"firstname" json:"firstname"`
	Lastname   string           `db:"lastname" json:"lastname"`
	Email      string           `db:"email" json:"email"`
	Password   string           `db:"password" json:"password"`
	CreatedAt  pgtype.Timestamp `db:"created_at" json:"created_at"`
	UpdatedAt  pgtype.Timestamp `db:"updated_at" json:"updated_at"`
	DeletedAt  pgtype.Timestamp `db:"deleted_at" json:"deleted_at"`
	TotalCount int64            `db:"total_count" json:"total_count"`
}

type RestoreUserRow struct {
	UserID    int32            `db:"user_id" json:"user_id"`
	Firstname string           `db:"firstname" json:"firstname"`
	Lastname  string           `db:"lastname" json:"lastname"`
	Email     string           `db:"email" json:"email"`
	CreatedAt pgtype.Timestamp `db:"created_at" json:"created_at"`
	UpdatedAt pgtype.Timestamp `db:"updated_at" json:"updated_at"`
	DeletedAt pgtype.Timestamp `db:"deleted_at" json:"deleted_at"`
}

type TrashUserRow struct {
	UserID    int32            `db:"user_id" json:"user_id"`
	Firstname string           `db:"firstname" json:"firstname"`
	Lastname  string           `db:"lastname" json:"lastname"`
	Email     string           `db:"email" json:"email"`
	CreatedAt pgtype.Timestamp `db:"created_at" json:"created_at"`
	UpdatedAt pgtype.Timestamp `db:"updated_at" json:"updated_at"`
	DeletedAt pgtype.Timestamp `db:"deleted_at" json:"deleted_at"`
}

type UpdateUserParams struct {
	UserID    int32  `db:"user_id" json:"user_id"`
	Firstname string `db:"firstname" json:"firstname"`
	Lastname  string `db:"lastname" json:"lastname"`
	Email     string `db:"email" json:"email"`
	Password  string `db:"password" json:"password"`
}

type UpdateUserIsVerifiedParams struct {
	UserID     int32 `db:"user_id" json:"user_id"`
	IsVerified *bool `db:"is_verified" json:"is_verified"`
}

type UpdateUserIsVerifiedRow struct {
	UserID    int32            `db:"user_id" json:"user_id"`
	Firstname string           `db:"firstname" json:"firstname"`
	Lastname  string           `db:"lastname" json:"lastname"`
	Email     string           `db:"email" json:"email"`
	Password  string           `db:"password" json:"password"`
	CreatedAt pgtype.Timestamp `db:"created_at" json:"created_at"`
	UpdatedAt pgtype.Timestamp `db:"updated_at" json:"updated_at"`
}

type UpdateUserPasswordParams struct {
	UserID   int32  `db:"user_id" json:"user_id"`
	Password string `db:"password" json:"password"`
}

type UpdateUserPasswordRow struct {
	UserID    int32            `db:"user_id" json:"user_id"`
	Firstname string           `db:"firstname" json:"firstname"`
	Lastname  string           `db:"lastname" json:"lastname"`
	Email     string           `db:"email" json:"email"`
	Password  string           `db:"password" json:"password"`
	CreatedAt pgtype.Timestamp `db:"created_at" json:"created_at"`
	UpdatedAt pgtype.Timestamp `db:"updated_at" json:"updated_at"`
}
