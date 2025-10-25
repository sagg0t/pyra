package test

import (
	"context"
	"fmt"
	"strings"
	"sync"
	"testing"
	"time"

	pyradb "pyra/pkg/db"

	"github.com/docker/go-connections/nat"
	tc "github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

func Must[T any](v T, err error) T {
	if err != nil {
		panic(err)
	}

	return v
}

type TestDB struct {
	pyradb.DBTX
}

type DBTruncateFunc func()

var (
	dbConn                 *TestDB
	truncateAllTablesQuery string
	mx                     sync.Mutex
)

func DB(t testing.TB) *TestDB {
	if dbConn == nil {
		mx.Lock()
		defer mx.Unlock()
		conn := SetupTestDB(t)
		dbConn = &TestDB{conn}
	}

	dbConn.Truncate()

	return dbConn
}

// Truncate - truncates all tables.
func (conn *TestDB) Truncate() {
	if truncateAllTablesQuery == "" {
		panic("truncate function not configured, call test.SetupTestDB first")
	}

	ctx := context.Background()
	if _, err := conn.ExecContext(ctx, truncateAllTablesQuery); err != nil {
		fmt.Printf("[pyra] error while truncating DB: %v\n", err)
	}
}

func SetupTestDB(t testing.TB) pyradb.DBTX {
	fmt.Printf("[pyra] Setting up test DB: %s\n", time.Now().String())

	host := "127.0.0.1"
	var port uint = 5433

	if false {
		var mappedPort nat.Port
		host, mappedPort = StartDBContainer(t.Context())

		port = uint(mappedPort.Int())
	}

	dbConfig := pyradb.NewConfig("pgx")
	dbConfig.Host = host
	dbConfig.Port = port
	dbConfig.DBName = "pyra_test"

	logger := NewLogger(t.Output())
	dbPool := Must(pyradb.New(t.Context(), dbConfig, logger))

	rows := Must(dbPool.QueryContext(t.Context(), `
SELECT table_name
FROM information_schema.tables
WHERE table_schema = 'public'
	AND table_name != 'schema_migrations';`))

	var tableNames []string
	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			panic(err)
		}

		tableNames = append(tableNames, name)
	}

	names := strings.Join(tableNames, ", ")
	truncateAllTablesQuery = fmt.Sprintf("TRUNCATE %s;", names)

	return dbPool
}

func StartDBContainer(ctx context.Context) (host string, port nat.Port) {
	fmt.Println("[pyra] Spinning test DB container")

	// TODO: need to apply migrations as well.
	container, err := tc.Run(ctx, "postgres:18-alpine",
		tc.WithExposedPorts("5432/tcp"),
		tc.WithEnv(map[string]string{
			"POSTGRES_USER":     "pyra",
			"POSTGRES_PASSWORD": "pyra",
			"POSTGRES_DB":       "pyra_test",
		}),
		tc.WithWaitStrategy(wait.ForListeningPort("5432/tcp")))
	if err != nil {
		panic(fmt.Sprintf("could not start database container: %v", err))
	}

	host, err = container.Host(ctx)
	if err != nil {
		panic(err)
	}

	port, err = container.MappedPort(ctx, "5432/tcp")
	if err != nil {
		panic(err)
	}

	fmt.Printf("--> PORT: %s\n", port.Port())

	return host, port
}
