package migrate

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"time"
)

var (
	directions   = [2]string{"up", "down"}
	formatRegex  = regexp.MustCompile(`(\d{16})_(.+)\.(\w+)\.sql`)
	initTemplate = `
	CREATE TABLE IF NOT EXISTS "%s" (
		version VARCHAR(24) PRIMARY KEY,
		created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
	);
	`
)

type Engine struct {
	db   *sql.DB
	conf Config
}

func NewEngine(conf Config, dbConf DBConfig) (*Engine, error) {
	dbConn, err := sql.Open(dbConf.Adapter, dbConf.String())
	if err != nil {
		return nil, err
	}

	eng := &Engine{
		conf: conf,
		db:   dbConn,
	}

	err = eng.init()
	if err != nil {
		return nil, err
	}

	return eng, nil
}

func (eng *Engine) init() error {
	ctx := context.Background()
	err := eng.db.PingContext(ctx)
	if err != nil {
		return fmt.Errorf("DB ping failed - %w", err)
	}
	slog.Debug("DB ping successful")

	_, err = eng.db.ExecContext(ctx, fmt.Sprintf(initTemplate, eng.conf.TableName))
	if err != nil {
		return fmt.Errorf("failed to create the control table:\n\t %w", err)
	}
	slog.Debug("control table initialisation success")

	return nil
}

// Creates a migration file.
// Returns a list of created file names
func (eng *Engine) CreateMigration(name string) ([]string, error) {
	version := strconv.FormatInt(time.Now().UnixMicro(), 10)
	versionGlob := filepath.Join(eng.conf.Dir, version+"_*.sql")
	matches, err := filepath.Glob(versionGlob)
	if err != nil {
		return nil, err
	}

	if len(matches) > 0 {
		return nil, fmt.Errorf("duplicate migration version: %s, %w", version, err)
	}

	if err = os.MkdirAll(eng.conf.Dir, os.ModePerm); err != nil {
		return nil, err
	}

	createdFiles := []string{}
	for _, direction := range directions {
		baseName := fmt.Sprintf("%s_%s.%s.sql", version, name, direction)
		fileName := filepath.Join(eng.conf.Dir, baseName)

		f, err := os.Create(fileName)
		if err != nil {
			if len(createdFiles) != 0 {
				for _, f := range createdFiles {
					os.Remove(f)
				}
			}

			return nil, err
		}
		f.Close()

		createdFiles = append(createdFiles, fileName)
	}

	return createdFiles, nil
}

func (eng *Engine) CurrentVersion() (Migration, error) {
	ctx := context.Background()
	row := eng.db.QueryRowContext(ctx, "SELECT * FROM schema_migrations ORDER BY version DESC LIMIT 1;")

	m := Migration{}

	err := row.Scan(&m.Version, &m.AppliedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return Migration{Version: "0", AppliedAt: time.Time{}}, nil
		}

		return m, err
	}

	return m, nil
}

// Returns slice of migrations combining all local migrations as well as the migrations, that have been
// previously applied, but which files are missing.
func (eng *Engine) Status() ([]Migration, error) {
	localMigrations, err := eng.ScanDir()
	if err != nil {
		return nil, err
	}

	migrationHash := make(map[string]*Migration, len(localMigrations))
	for _, mig := range localMigrations {
		m := new(Migration)
		*m = mig
		migrationHash[mig.Version] = m
	}

	remoteMigrations, err := eng.remoteVersions()
	if err != nil {
		return nil, err
	}

	for _, remoteMig := range remoteMigrations {
		pMig, ok := migrationHash[remoteMig.Version]

		if ok {
			pMig.AppliedAt = remoteMig.AppliedAt
		} else {
			m := new(Migration)
			*m = remoteMig

			migrationHash[m.Version] = m
		}
	}

	migs := make([]Migration, 0)
	for _, mig := range migrationHash {
		migs = append(migs, *mig)
	}

	sort.Sort(migrations(migs))

	return migs, nil
}

func (eng *Engine) Apply(count uint64) (appliedMigrations []Migration, outErr error) {
	pendingMigrations, err := eng.PendingMigrations()
	if err != nil {
		return nil, err
	}

	if count > 0 {
		pendingMigrations = pendingMigrations[:count]
	}

	appliedMigrations = make([]Migration, 0)
	for _, mig := range pendingMigrations {
		query, err := eng.readMigrationFile(mig.UpFile)
		if err != nil {
			outErr = err
			break
		}

		err = eng.applyMigration(fmt.Sprint(mig.Version), string(query), eng.onApplyMigration)
		if err != nil {
			outErr = err
			break
		}

		appliedMigrations = append(appliedMigrations, mig)
	}

	if len(appliedMigrations) != 0 && len(appliedMigrations) != len(pendingMigrations) {
		outErr = fmt.Errorf("some DB migrations failed to apply - %w", outErr)
	}

	return appliedMigrations, outErr
}

func (eng *Engine) Rollback(count uint64) (rolledBack []Migration, outErr error) {
	appliedVersions, err := eng.AppliedMigrations()
	if err != nil {
		return nil, err
	}

	if len(appliedVersions) == 0 {
		return nil, fmt.Errorf("nothing to rollback")
	}

	if count == 0 {
		count = uint64(len(appliedVersions))
	}

	sort.Sort(sort.Reverse(migrations(appliedVersions)))

	rolledBack = make([]Migration, 0, count)
	for _, mig := range appliedVersions {
		if count == 0 {
			break
		}

		query, err := eng.readMigrationFile(mig.DownFile)
		if err != nil {
			outErr = err
			break
		}

		err = eng.applyMigration(fmt.Sprint(mig.Version), string(query), eng.onRollBackMigration)
		if err != nil {
			outErr = err
			break
		}

		count--
		rolledBack = append(rolledBack, mig)
	}

	if count != 0 {
		outErr = fmt.Errorf("some DB migrations failed to apply - %w", outErr)
	}

	return rolledBack, outErr
}

// Intended for internal usage, but who knows...
// Returns a slice of Migration's for each migration version present in the corresponding directory.
func (eng *Engine) ScanDir() ([]Migration, error) {
	files, err := os.ReadDir(eng.conf.Dir)
	if err != nil {
		return nil, err
	}

	migrationHash := make(map[string]*Migration)
	for _, file := range files {
		scan := formatRegex.FindStringSubmatch(file.Name())

		version := scan[1]

		m, ok := migrationHash[version]
		if !ok {
			migrationHash[version] = &Migration{Version: version}
			m = migrationHash[version]
		}

		m.Name = scan[2]

		switch scan[3] {
		case upDirection:
			m.UpFile = file.Name()
		case downDirection:
			m.DownFile = file.Name()
		default:
			return nil, fmt.Errorf("invalid migration file format - direction is invalid - %s", file.Name())
		}
	}

	migrations := make([]Migration, 0)
	for _, m := range migrationHash {
		migrations = append(migrations, *m)
	}

	return migrations, nil
}

func (eng *Engine) readMigrationFile(filename string) ([]byte, error) {
	filepath := filepath.Join(eng.conf.Dir, filename)
	query, err := os.ReadFile(filepath)
	if err != nil {
		return nil, err
	}

	return query, nil
}
