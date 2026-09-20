package internal_test

import (
	"database/sql"
	"database/sql/driver"
	"errors"
	"io"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vidarandrebo/nutrition-tracker/api/internal"
)

type mockHealthDriver struct{}

type mockHealthConn struct {
	fail bool
}

type mockHealthRows struct {
	closed bool
}

func (d *mockHealthDriver) Open(name string) (driver.Conn, error) {
	if name == "fail" {
		return &mockHealthConn{fail: true}, nil
	}
	return &mockHealthConn{fail: false}, nil
}

func (c *mockHealthConn) Prepare(query string) (driver.Stmt, error) {
	return nil, errors.New("not implemented")
}

func (c *mockHealthConn) Close() error {
	return nil
}

func (c *mockHealthConn) Begin() (driver.Tx, error) {
	return nil, errors.New("not implemented")
}

func (c *mockHealthConn) Query(query string, args []driver.Value) (driver.Rows, error) {
	if c.fail {
		return nil, errors.New("database error")
	}
	return &mockHealthRows{}, nil
}

func (r *mockHealthRows) Columns() []string {
	return []string{"?column?"}
}

func (r *mockHealthRows) Close() error {
	r.closed = true
	return nil
}

func (r *mockHealthRows) Next(dest []driver.Value) error {
	if r.closed {
		return io.EOF
	}
	dest[0] = int64(1)
	r.closed = true
	return nil
}

func init() {
	sql.Register("mock_health_driver", &mockHealthDriver{})
}

func TestHealthEndpoint_ServeHTTP_Success(t *testing.T) {
	db, err := sql.Open("mock_health_driver", "ok")
	assert.NoError(t, err)
	defer db.Close()

	endpoint := internal.NewHealthEndpoint(db, slog.Default())

	req := httptest.NewRequest(http.MethodGet, "/healthz", nil)
	rec := httptest.NewRecorder()

	endpoint.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
}

func TestHealthEndpoint_ServeHTTP_DBError(t *testing.T) {
	db, err := sql.Open("mock_health_driver", "fail")
	assert.NoError(t, err)
	defer db.Close()

	endpoint := internal.NewHealthEndpoint(db, slog.Default())

	req := httptest.NewRequest(http.MethodGet, "/healthz", nil)
	rec := httptest.NewRecorder()

	endpoint.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusInternalServerError, rec.Code)
}

func TestHealthEndpoint_ServeHTTP_NilDB(t *testing.T) {
	endpoint := internal.NewHealthEndpoint(nil, slog.Default())

	req := httptest.NewRequest(http.MethodGet, "/healthz", nil)
	rec := httptest.NewRecorder()

	endpoint.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusInternalServerError, rec.Code)
}
