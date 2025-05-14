package db

import (
	"context"
	"database/sql"

	_ "github.com/jackc/pgx/v5/stdlib"

	"pyra/pkg/log"
)

type DBTX interface {
	// Begin() (*sql.Tx, error)
	BeginTx(ctx context.Context, opts *sql.TxOptions) (*sql.Tx, error)
	Close() error

	// Prepare(query string) (*sql.Stmt, error)
	PrepareContext(ctx context.Context, query string) (*sql.Stmt, error)

	// Exec(query string, args ...interface{}) (sql.Result, error)
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)

	// Query(query string, args ...interface{}) (*sql.Rows, error)
	QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)

	// QueryRow(query string, args ...interface{}) *sql.Row
	QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row
}

type DB struct {
	conn *sql.DB
	log  *log.Logger
}

func (db *DB) Begin() (*sql.Tx, error) {
	db.log.Trace("BEGIN transaction")

	return db.conn.Begin()
}

func (db *DB) BeginTx(ctx context.Context, opts *sql.TxOptions) (*sql.Tx, error) {
	db.log.TraceContext(ctx, "BEGIN transaction")

	return db.conn.BeginTx(ctx, opts)
}

func (db *DB) Prepare(query string) (*sql.Stmt, error) {
	return db.conn.Prepare(query)
}

func (db *DB) PrepareContext(ctx context.Context, query string) (*sql.Stmt, error) {
	return db.conn.PrepareContext(ctx, query)
}

func (db *DB) Exec(query string, args ...interface{}) (sql.Result, error) {
	// db.log.Trace(query, args ...any)
	return db.conn.Exec(query, args...)
}

func (db *DB) ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	return db.conn.ExecContext(ctx, query, args...)
}

func (db *DB) Query(query string, args ...interface{}) (*sql.Rows, error) {
	return db.conn.Query(query, args...)
}

func (db *DB) QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error) {
	return db.conn.QueryContext(ctx, query, args...)
}

func (db *DB) QueryRow(query string, args ...interface{}) *sql.Row {
	return db.conn.QueryRow(query, args...)
}

func (db *DB) QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row {
	return db.conn.QueryRowContext(ctx, query, args...)
}

func (db *DB) Close() error {
	db.log.Trace("closing DB connection")

	return db.conn.Close()
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
