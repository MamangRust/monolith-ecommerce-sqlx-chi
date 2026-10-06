package repository

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"errors"
	"testing"

	shared_errors "github.com/MamangRust/monolith-ecommerce-shared/errors"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jmoiron/sqlx"
)

type userPermanentDeleteDriver struct{}

func (userPermanentDeleteDriver) Open(string) (driver.Conn, error) {
	return nil, errors.New("use userPermanentDeleteConnector")
}

type userPermanentDeleteConnector struct{}

func (userPermanentDeleteConnector) Connect(context.Context) (driver.Conn, error) {
	return &userPermanentDeleteDriverConn{}, nil
}

func (userPermanentDeleteConnector) Driver() driver.Driver {
	return userPermanentDeleteDriver{}
}

type userPermanentDeleteDriverConn struct{}

func (c *userPermanentDeleteDriverConn) Prepare(string) (driver.Stmt, error) {
	return nil, errors.New("not implemented")
}

func (c *userPermanentDeleteDriverConn) Close() error { return nil }

func (c *userPermanentDeleteDriverConn) Begin() (driver.Tx, error) {
	return nil, errors.New("not implemented")
}

func (c *userPermanentDeleteDriverConn) ExecContext(context.Context, string, []driver.NamedValue) (driver.Result, error) {
	return nil, &pgconn.PgError{Code: "23503", ConstraintName: "merchants_user_id_fkey"}
}

func newUserPermanentDeleteRepository() UserCommandRepository {
	handle := sql.OpenDB(&userPermanentDeleteConnector{})
	return NewUserCommandRepository(sqlx.NewDb(handle, "pgx"))
}

func TestUserCommandRepositoryDeletePermanentReturnsConflictOnForeignKey(t *testing.T) {
	repo := newUserPermanentDeleteRepository()

	deleted, err := repo.DeletePermanent(context.Background(), 7)
	assertUserPermanentDeleteConflict(t, deleted, err)
}

func TestUserCommandRepositoryDeleteAllReturnsConflictOnForeignKey(t *testing.T) {
	repo := newUserPermanentDeleteRepository()

	deleted, err := repo.DeleteAll(context.Background())
	assertUserPermanentDeleteConflict(t, deleted, err)
}

func assertUserPermanentDeleteConflict(t *testing.T, deleted bool, err error) {
	t.Helper()

	if deleted {
		t.Fatal("delete reported success after a foreign-key violation")
	}

	var appErr *shared_errors.AppError
	if !errors.As(err, &appErr) {
		t.Fatalf("error = %T %v, want AppError", err, err)
	}
	if appErr.Type != shared_errors.ErrorTypeConflict || appErr.Code != 409 {
		t.Fatalf("error type/code = %s/%d, want CONFLICT/409", appErr.Type, appErr.Code)
	}
}
