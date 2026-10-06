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

// productPermanentDeleteFakeDB is a minimal database/sql driver whose
// ExecContext returns a canned error, allowing the foreign-key conflict
// mapping to be tested without a real database.
type productPermanentDeleteFakeDB struct {
	err error
}

func (f *productPermanentDeleteFakeDB) Connect(context.Context) (driver.Conn, error) {
	return f, nil
}

func (f *productPermanentDeleteFakeDB) Driver() driver.Driver { return f }

func (f *productPermanentDeleteFakeDB) Open(string) (driver.Conn, error) { return f, nil }

func (f *productPermanentDeleteFakeDB) Prepare(string) (driver.Stmt, error) {
	return nil, errors.New("productPermanentDeleteFakeDB: Prepare not supported")
}

func (f *productPermanentDeleteFakeDB) Close() error { return nil }

func (f *productPermanentDeleteFakeDB) Begin() (driver.Tx, error) {
	return nil, errors.New("productPermanentDeleteFakeDB: Begin not supported")
}

func (f *productPermanentDeleteFakeDB) ExecContext(context.Context, string, []driver.NamedValue) (driver.Result, error) {
	return driver.RowsAffected(0), f.err
}

func (f *productPermanentDeleteFakeDB) QueryContext(context.Context, string, []driver.NamedValue) (driver.Rows, error) {
	return nil, errors.New("productPermanentDeleteFakeDB: QueryContext not supported")
}

func newFakeSQLxDBForProduct(t *testing.T, execErr error) *sqlx.DB {
	t.Helper()
	return sqlx.NewDb(sql.OpenDB(&productPermanentDeleteFakeDB{err: execErr}), "pgx")
}

func TestProductCommandRepositoryDeletePermanentReturnsConflictOnForeignKey(t *testing.T) {
	repo := NewProductCommandRepository(newFakeSQLxDBForProduct(t,
		&pgconn.PgError{Code: "23503", ConstraintName: "reviews_product_id_fkey"},
	))

	deleted, err := repo.DeletePermanent(context.Background(), 7)
	assertProductPermanentDeleteConflict(t, deleted, err)
}

func TestProductCommandRepositoryDeleteAllReturnsConflictOnForeignKey(t *testing.T) {
	repo := NewProductCommandRepository(newFakeSQLxDBForProduct(t,
		&pgconn.PgError{Code: "23503", ConstraintName: "reviews_product_id_fkey"},
	))

	deleted, err := repo.DeleteAll(context.Background())
	assertProductPermanentDeleteConflict(t, deleted, err)
}

func assertProductPermanentDeleteConflict(t *testing.T, deleted bool, err error) {
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
