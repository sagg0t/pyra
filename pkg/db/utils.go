package db

import (
	"context"

	"pyra/pkg/log"
)

func RollbackGuard(ctx context.Context, tx DBTX, err *error) {
	if *err == nil {
		return
	}

	if rollbackErr := tx.Rollback(); rollbackErr != nil {
		log.FromContext(ctx).ErrorContext(ctx, "failed to tollback transaction",
			"error", rollbackErr)
	}
}
