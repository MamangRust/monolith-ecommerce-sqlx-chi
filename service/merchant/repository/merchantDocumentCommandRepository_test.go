package repository

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"errors"
	"io"
	"strings"
	"testing"

	"github.com/MamangRust/monolith-ecommerce-shared/domain/requests"
	"github.com/jmoiron/sqlx"
)

// fakeDocumentUpdateDB is a minimal database/sql driver that records the
// executed SQL and its positional args and returns a canned row for queries.
// It allows testing the repository param building without a real database.
type fakeDocumentUpdateDB struct {
	lastSQL  string
	lastArgs []driver.Value

	columns []string
	row     []driver.Value
	rowErr  error
}

func (f *fakeDocumentUpdateDB) Connect(ctx context.Context) (driver.Conn, error) { return f, nil }

func (f *fakeDocumentUpdateDB) Driver() driver.Driver { return f }

func (f *fakeDocumentUpdateDB) Open(name string) (driver.Conn, error) { return f, nil }

func (f *fakeDocumentUpdateDB) Prepare(query string) (driver.Stmt, error) {
	return nil, errors.New("fakeDocumentUpdateDB: Prepare not supported")
}

func (f *fakeDocumentUpdateDB) Close() error { return nil }

func (f *fakeDocumentUpdateDB) Begin() (driver.Tx, error) {
	return nil, errors.New("fakeDocumentUpdateDB: Begin not supported")
}

func (f *fakeDocumentUpdateDB) ExecContext(ctx context.Context, query string, args []driver.NamedValue) (driver.Result, error) {
	f.lastSQL = query
	f.lastArgs = namedValues(args)
	return driver.RowsAffected(0), nil
}

func (f *fakeDocumentUpdateDB) QueryContext(ctx context.Context, query string, args []driver.NamedValue) (driver.Rows, error) {
	f.lastSQL = query
	f.lastArgs = namedValues(args)
	if f.rowErr != nil {
		return nil, f.rowErr
	}
	return &fakeDocumentUpdateRows{columns: f.columns, vals: f.row}, nil
}

func namedValues(args []driver.NamedValue) []driver.Value {
	vals := make([]driver.Value, len(args))
	for i, a := range args {
		vals[i] = a.Value
	}
	return vals
}

// fakeDocumentUpdateRows returns a single canned row.
type fakeDocumentUpdateRows struct {
	columns []string
	vals    []driver.Value
	done    bool
}

func (r *fakeDocumentUpdateRows) Columns() []string { return r.columns }

func (r *fakeDocumentUpdateRows) Close() error { return nil }

func (r *fakeDocumentUpdateRows) Next(dest []driver.Value) error {
	if r.done {
		return io.EOF
	}
	r.done = true
	copy(dest, r.vals)
	return nil
}

// newFakeDocumentUpdateSQLxDB wraps fakeDocumentUpdateDB in a real *sqlx.DB so
// the repository exercises the exact sqlx code path used in production.
func newFakeDocumentUpdateSQLxDB(t *testing.T, row []driver.Value, rowErr error) (*sqlx.DB, *fakeDocumentUpdateDB) {
	t.Helper()
	f := &fakeDocumentUpdateDB{
		columns: []string{
			"document_id", "merchant_id", "document_type", "document_url",
			"status", "note", "uploaded_at", "created_at", "updated_at",
		},
		row:    row,
		rowErr: rowErr,
	}
	return sqlx.NewDb(sql.OpenDB(f), "postgres"), f
}

// merchantDocumentUpdateRowFixture mirrors the column order of the
// merchant_documents RETURNING scan.
func merchantDocumentUpdateRowFixture(docID, merchantID int) []driver.Value {
	return []driver.Value{
		int64(docID),
		int64(merchantID),
		"business_license",
		"https://example.com/license.pdf",
		"pending",
		nil,
		nil,
		nil,
		nil,
	}
}

func TestMerchantDocumentCommandRepositoryUpdateTargetsDocumentID(t *testing.T) {
	const (
		documentID = 42
		merchantID = 7
	)

	fakeConn, fake := newFakeDocumentUpdateSQLxDB(t, merchantDocumentUpdateRowFixture(documentID, merchantID), nil)
	repo := NewMerchantDocumentCommandRepository(fakeConn)
	note := "replace before review"

	_, err := repo.Update(context.Background(), &requests.UpdateMerchantDocumentRequest{
		DocumentID:   intPtr(documentID),
		MerchantID:   merchantID,
		DocumentType: "business_license",
		DocumentUrl:  "https://example.com/license.pdf",
		Status:       "pending",
		Note:         note,
	})
	if err != nil {
		t.Fatalf("Update() error = %v", err)
	}

	assertMerchantDocumentUpdateTarget(t, fake, documentID, merchantID)
}

func TestMerchantDocumentCommandRepositoryUpdateStatusTargetsDocumentID(t *testing.T) {
	const (
		documentID = 84
		merchantID = 13
	)

	fakeConn, fake := newFakeDocumentUpdateSQLxDB(t, merchantDocumentUpdateRowFixture(documentID, merchantID), nil)
	repo := NewMerchantDocumentCommandRepository(fakeConn)
	note := "verified"

	_, err := repo.UpdateStatus(context.Background(), &requests.UpdateMerchantDocumentStatusRequest{
		DocumentID: intPtr(documentID),
		MerchantID: merchantID,
		Status:     "approved",
		Note:       note,
	})
	if err != nil {
		t.Fatalf("UpdateStatus() error = %v", err)
	}

	assertMerchantDocumentUpdateTarget(t, fake, documentID, merchantID)
}

func assertMerchantDocumentUpdateTarget(t *testing.T, fake *fakeDocumentUpdateDB, documentID, merchantID int) {
	t.Helper()

	if !strings.Contains(fake.lastSQL, "document_id = $1") {
		t.Fatalf("update query target = %q, want document_id", fake.lastSQL)
	}
	if len(fake.lastArgs) == 0 {
		t.Fatal("update query received no arguments")
	}
	gotDocumentID, ok := fake.lastArgs[0].(int64)
	if !ok {
		t.Fatalf("first update argument type = %T, want int64", fake.lastArgs[0])
	}
	if gotDocumentID != int64(documentID) {
		t.Fatalf("first update argument = %d, want DocumentID %d (MerchantID is %d)", gotDocumentID, documentID, merchantID)
	}
	if gotDocumentID == int64(merchantID) {
		t.Fatalf("first update argument incorrectly used MerchantID %d", merchantID)
	}
}

func intPtr(value int) *int {
	return &value
}
