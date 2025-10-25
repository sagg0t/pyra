package db

import (
	"context"
	"database/sql"
	"pyra/pkg/log"
)

type TX struct {
	tx *sql.Tx
	log *log.Logger
}

func (tx *TX) Begin() (DBTX, error) {
	panic("TX doesn't implement Begin(). You probably meant to use a DB connection.")
}

func (tx *TX) BeginTx(ctx context.Context, opts *sql.TxOptions) (DBTX, error) {
	panic("TX doesn't implement BeginTx(). You probably meant to use a DB connection.")
}

func (tx *TX) Commit() error {
	tx.log.Trace("COMMIT transaction")
	return tx.tx.Commit()
}

func (tx *TX) Rollback() error {
	tx.log.Trace("ROLLBACK transaction")
	return tx.tx.Rollback()
}

func (tx *TX) Prepare(query string) (*sql.Stmt, error) {
	tx.log.Trace("preparing query: " + query)

	return tx.tx.Prepare(query)
}

func (tx *TX) PrepareContext(ctx context.Context, query string) (*sql.Stmt, error) {
	tx.log.TraceContext(ctx, "preparing query: " + query)
	return tx.tx.PrepareContext(ctx, query)
}

func (tx *TX) Exec(query string, args ...any) (sql.Result, error) {
	tx.log.Trace(query, "args", args)
	return tx.tx.Exec(query, args...)
}

func (tx *TX) ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error) {
	tx.log.TraceContext(ctx, query, "args", args)
	return tx.tx.ExecContext(ctx, query, args...)
}

func (tx *TX) Query(query string, args ...any) (*sql.Rows, error) {
	tx.log.Trace(query, "args", args)
	return tx.tx.Query(query, args...)
}

func (tx *TX) QueryContext(ctx context.Context, query string, args ...any) (*sql.Rows, error) {
	tx.log.TraceContext(ctx, query, "args", args)
	return tx.tx.QueryContext(ctx, query, args...)
}

func (tx *TX) QueryRow(query string, args ...any) *sql.Row {
	tx.log.Trace(query, "args", args)
	return tx.tx.QueryRow(query, args...)
}

func (tx *TX) QueryRowContext(ctx context.Context, query string, args ...any) *sql.Row {
	tx.log.TraceContext(ctx, query, "args", args)
	return tx.tx.QueryRowContext(ctx, query, args...)
}

func (tx *TX) Close() error {
	panic("TX doesn't implement Close(). You probably meant to close a DB connection.")
}
