package mocks

import (
	"context"
	"database/sql"

	"pyra/pkg/db"
)

type Tx struct {
	BeginTxFn  func(ctx context.Context, opts *sql.TxOptions) (db.DBTX, error)
	CommitFn   func() error
	RollbackFn func() error

	CloseFn func() error

	PrepareContextFn  func(ctx context.Context, query string) (*sql.Stmt, error)
	ExecContextFn     func(ctx context.Context, query string, args ...any) (sql.Result, error)
	QueryContextFn    func(ctx context.Context, query string, args ...any) (*sql.Rows, error)
	QueryRowContextFn func(ctx context.Context, query string, args ...any) *sql.Row
}

var _ db.DBTX = &Tx{}

func (mock *Tx) BeginTx(ctx context.Context, opts *sql.TxOptions) (db.DBTX, error) {
	if mock.BeginTxFn == nil {
		notImplemented()
	}

	return mock.BeginTxFn(ctx, opts)
}

func (mock *Tx) Commit() error {
	if mock.CommitFn == nil {
		notImplemented()
	}

	return mock.CommitFn()
}

func (mock *Tx) Rollback() error {
	if mock.RollbackFn == nil {
		notImplemented()
	}

	return mock.RollbackFn()
}

func (mock *Tx) Close() error {
	if mock.CloseFn == nil {
		notImplemented()
	}

	return mock.CloseFn()
}

func (mock *Tx) PrepareContext(ctx context.Context, query string) (*sql.Stmt, error) {
	if mock.PrepareContextFn == nil {
		notImplemented()
	}

	return mock.PrepareContextFn(ctx, query)
}

func (mock *Tx) ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error) {
	if mock.ExecContextFn == nil {
		notImplemented()
	}

	return mock.ExecContextFn(ctx, query, args...)
}

func (mock *Tx) QueryContext(ctx context.Context, query string, args ...any) (*sql.Rows, error) {
	if mock.QueryContextFn == nil {
		notImplemented()
	}

	return mock.QueryContextFn(ctx, query, args...)
}

func (mock *Tx) QueryRowContext(ctx context.Context, query string, args ...any) *sql.Row {
	if mock.QueryRowContextFn == nil {
		notImplemented()
	}

	return mock.QueryRowContextFn(ctx, query, args...)
}
