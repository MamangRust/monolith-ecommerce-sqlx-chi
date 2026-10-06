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

// fakePermanentDeleteDB is a minimal database/sql driver whose ExecContext
// returns a canned error. It allows testing the repository conflict mapping
// without a real database.
type fakePermanentDeleteDB struct {
	execErr error
}

func (f *fakePermanentDeleteDB) Connect(ctx context.Context) (driver.Conn, error) { return f, nil }

func (f *fakePermanentDeleteDB) Driver() driver.Driver { return f }

func (f *fakePermanentDeleteDB) Open(name string) (driver.Conn, error) { return f, nil }

func (f *fakePermanentDeleteDB) Prepare(query string) (driver.Stmt, error) {
	return nil, errors.New("fakePermanentDeleteDB: Prepare not supported")
}

func (f *fakePermanentDeleteDB) Close() error { return nil }

func (f *fakePermanentDeleteDB) Begin() (driver.Tx, error) {
	return nil, errors.New("fakePermanentDeleteDB: Begin not supported")
}

func (f *fakePermanentDeleteDB) ExecContext(ctx context.Context, query string, args []driver.NamedValue) (driver.Result, error) {
	return driver.RowsAffected(0), f.execErr
}

func newFakePermanentDeleteSQLxDB(t *testing.T, execErr error) *sqlx.DB {
	t.Helper()
	return sqlx.NewDb(sql.OpenDB(&fakePermanentDeleteDB{execErr: execErr}), "postgres")
}

func TestMerchantCommandRepositoryDeletePermanentReturnsConflictOnForeignKey(t *testing.T) {
	repo := NewMerchantCommandRepository(newFakePermanentDeleteSQLxDB(t,
		&pgconn.PgError{Code: "23503", ConstraintName: "products_merchant_id_fkey"},
	))

	deleted, err := repo.DeletePermanent(context.Background(), 7)
	assertMerchantPermanentDeleteConflict(t, deleted, err)
}

func TestMerchantCommandRepositoryDeleteAllReturnsConflictOnForeignKey(t *testing.T) {
	repo := NewMerchantCommandRepository(newFakePermanentDeleteSQLxDB(t,
		&pgconn.PgError{Code: "23503", ConstraintName: "products_merchant_id_fkey"},
	))

	deleted, err := repo.DeleteAll(context.Background())
	assertMerchantPermanentDeleteConflict(t, deleted, err)
}

func assertMerchantPermanentDeleteConflict(t *testing.T, deleted bool, err error) {
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
