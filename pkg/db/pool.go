package db

import (
	"context"
	"log/slog"

	"github.com/jackc/pgx/v5/pgxpool"
)

func CreatePool(ctx context.Context, cfg Config, logger *slog.Logger) (*pgxpool.Pool, error) {
	pgxConf, err := pgxpool.ParseConfig(cfg.String())
	if err != nil {
		return nil, err
	}

	pgxConf.ConnConfig.Tracer = NewQueryTracer(logger)

	dbPool, err := pgxpool.NewWithConfig(ctx, pgxConf)
	if err != nil {
		return nil, err
	}

	return dbPool, nil
}
