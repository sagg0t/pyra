package main

import (
	"log/slog"
	"os"
	"strconv"

	"github.com/olehvolynets/pyra/internal/db"
	"github.com/olehvolynets/pyra/pkg/migrate"
)

const (
	dir                    = "database/migrations"
	migrationEngineInitErr = "pyra: failed to initialize migration engine"
)

var migrEngine *migrate.Engine

func main() {
	config := db.NewConfig("postgres")
	config.Attrs.Add("sslmode", fetchEnv("DB_SSLMODE", "disable"))

	migrEngine = migrate.NewBuilder().SetDir(dir).SetDb(nil).Done()
}

func fetchEnv(name, fallback string) string {
	val, ok := os.LookupEnv(name)
	if !ok {
		return fallback
	}

	return val
}

func parseCount(sCount string) uint64 {
	count, err := strconv.ParseUint(sCount, 10, strconv.IntSize)
	if err != nil {
		slog.Error("failed to read migrations count", "error", err)
	}

	return count
}
