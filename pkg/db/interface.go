package db

import (
	"context"
	"database/sql"
)

type DBTX interface {
	// Begin() (*sql.Tx, error)
	BeginTx(ctx context.Context, opts *sql.TxOptions) (DBTX, error)
	Commit() error
	Rollback() error

	Close() error

	// Prepare(query string) (*sql.Stmt, error)
	PrepareContext(ctx context.Context, query string) (*sql.Stmt, error)

	// Exec(query string, args ...interface{}) (sql.Result, error)
	ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error)

	// Query(query string, args ...interface{}) (*sql.Rows, error)
	QueryContext(ctx context.Context, query string, args ...any) (*sql.Rows, error)

	// QueryRow(query string, args ...interface{}) *sql.Row
	QueryRowContext(ctx context.Context, query string, args ...any) *sql.Row
}
