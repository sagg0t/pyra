package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/olehvolynets/pyra/pkg/log"
)

func CreatePool(ctx context.Context, cfg Config, logger *log.Logger) (*pgxpool.Pool, error) {
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
