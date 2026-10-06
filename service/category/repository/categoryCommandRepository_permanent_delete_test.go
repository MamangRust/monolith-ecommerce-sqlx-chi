package repository

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"errors"
	"testing"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jmoiron/sqlx"

	shared_errors "github.com/MamangRust/monolith-ecommerce-shared/errors"
)

// categoryPermanentDeleteFakeDB is a minimal database/sql driver whose
// ExecContext returns a canned error, allowing the foreign-key conflict
// mapping to be tested without a real database.
type categoryPermanentDeleteFakeDB struct {
	err error
}

func (f *categoryPermanentDeleteFakeDB) Connect(context.Context) (driver.Conn, error) {
	return f, nil
}

func (f *categoryPermanentDeleteFakeDB) Driver() driver.Driver { return f }

func (f *categoryPermanentDeleteFakeDB) Open(string) (driver.Conn, error) { return f, nil }

func (f *categoryPermanentDeleteFakeDB) Prepare(string) (driver.Stmt, error) {
	return nil, errors.New("categoryPermanentDeleteFakeDB: Prepare not supported")
}

func (f *categoryPermanentDeleteFakeDB) Close() error { return nil }

func (f *categoryPermanentDeleteFakeDB) Begin() (driver.Tx, error) {
	return nil, errors.New("categoryPermanentDeleteFakeDB: Begin not supported")
}

func (f *categoryPermanentDeleteFakeDB) ExecContext(context.Context, string, []driver.NamedValue) (driver.Result, error) {
	return driver.RowsAffected(0), f.err
}

func (f *categoryPermanentDeleteFakeDB) QueryContext(context.Context, string, []driver.NamedValue) (driver.Rows, error) {
	return nil, errors.New("categoryPermanentDeleteFakeDB: QueryContext not supported")
}

func newFakeSQLxDBForCategory(t *testing.T, execErr error) *sqlx.DB {
	t.Helper()
	return sqlx.NewDb(sql.OpenDB(&categoryPermanentDeleteFakeDB{err: execErr}), "pgx")
}

func TestCategoryCommandRepositoryDeletePermanentReturnsConflictOnForeignKey(t *testing.T) {
	repo := NewCategoryCommandRepository(newFakeSQLxDBForCategory(t,
		&pgconn.PgError{Code: "23503", ConstraintName: "products_category_id_fkey"},
	))

	deleted, err := repo.DeletePermanent(context.Background(), 7)
	assertCategoryPermanentDeleteConflict(t, deleted, err)
}

func TestCategoryCommandRepositoryDeleteAllReturnsConflictOnForeignKey(t *testing.T) {
	repo := NewCategoryCommandRepository(newFakeSQLxDBForCategory(t,
		&pgconn.PgError{Code: "23503", ConstraintName: "products_category_id_fkey"},
	))

	deleted, err := repo.DeleteAll(context.Background())
	assertCategoryPermanentDeleteConflict(t, deleted, err)
}

func assertCategoryPermanentDeleteConflict(t *testing.T, deleted bool, err error) {
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
