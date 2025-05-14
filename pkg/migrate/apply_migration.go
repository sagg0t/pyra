package migrate

import "context"

type applyMigrationCb func(version string) error

func (eng *Engine) applyMigration(version, query string, cb applyMigrationCb) error {
	ctx := context.TODO()
	tx, err := eng.db.Begin()
	if err != nil {
		return err
	}

	_, err = tx.ExecContext(ctx, query)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = cb(version)
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()

	return nil
}

func (eng *Engine) onApplyMigration(version string) error {
	ctx := context.TODO()
	_, err := eng.db.ExecContext(ctx, "INSERT INTO schema_migrations (version) VALUES ($1)", version)

	return err
}

func (eng *Engine) onRollBackMigration(version string) error {
	ctx := context.TODO()
	_, err := eng.db.ExecContext(ctx, "DELETE FROM schema_migrations WHERE version = $1", version)

	return err
}
