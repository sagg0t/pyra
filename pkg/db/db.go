// Package db - database abstraction
package db

import (
	"context"
	"database/sql"

	_ "github.com/jackc/pgx/v5/stdlib"

	"pyra/pkg/log"
)

type DB struct {
	conn *sql.DB
	log  *log.Logger
}

func New(ctx context.Context, cfg Config, logger *log.Logger) (DBTX, error) {
	pool, err := sql.Open(cfg.Adapter, cfg.String())
	if err != nil {
		return nil, err
	}

	if err := pool.PingContext(ctx); err != nil {
		return nil, err
	}

	return &DB{
		conn: pool,
		log:  logger,
	}, nil
}

func (db *DB) Begin() (DBTX, error) {
	db.log.Trace("BEGIN transaction")
	tx, err := db.conn.Begin()

	return &TX{tx: tx, log: db.log}, err
}

func (db *DB) BeginTx(ctx context.Context, opts *sql.TxOptions) (DBTX, error) {
	db.log.TraceContext(ctx, "BEGIN transaction")
	tx, err := db.conn.BeginTx(ctx, opts)

	return &TX{tx: tx, log: db.log}, err
}

func (db *DB) Commit() error {
	panic("DB doesn't implement Commit(). You probably meant using a transaction.")
}

func (db *DB) Rollback() error {
	panic("DB doesn't implement Rollback(). You probably meant using a transaction.")
}

func (db *DB) Prepare(query string) (*sql.Stmt, error) {
	db.log.Trace("preparing query: " + query)

	return db.conn.Prepare(query)
}

func (db *DB) PrepareContext(ctx context.Context, query string) (*sql.Stmt, error) {
	db.log.TraceContext(ctx, "preparing query: " + query)
	return db.conn.PrepareContext(ctx, query)
}

func (db *DB) Exec(query string, args ...any) (sql.Result, error) {
	db.log.Trace(query, "args", args)
	return db.conn.Exec(query, args...)
}

func (db *DB) ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error) {
	db.log.TraceContext(ctx, query, "args", args)
	return db.conn.ExecContext(ctx, query, args...)
}

func (db *DB) Query(query string, args ...any) (*sql.Rows, error) {
	db.log.Trace(query, "args", args)
	return db.conn.Query(query, args...)
}

func (db *DB) QueryContext(ctx context.Context, query string, args ...any) (*sql.Rows, error) {
	db.log.TraceContext(ctx, query, "args", args)
	return db.conn.QueryContext(ctx, query, args...)
}

func (db *DB) QueryRow(query string, args ...any) *sql.Row {
	db.log.Trace(query, "args", args)
	return db.conn.QueryRow(query, args...)
}

func (db *DB) QueryRowContext(ctx context.Context, query string, args ...any) *sql.Row {
	db.log.TraceContext(ctx, query, "args", args)
	return db.conn.QueryRowContext(ctx, query, args...)
}

func (db *DB) Close() error {
	db.log.Trace("closing DB connection")
	return db.conn.Close()
}
