package main

import (
	"context"

	"pyra/pkg/db"
	"pyra/pkg/log"
)

var l = log.NewLogger()

func main() {
	dbConf := db.NewConfig("pgx")
	// dbConf.Attrs.Add("sslmode", fetchEnv("DB_SSLMODE", "disable"))
	dbConn, err := db.New(context.Background(), dbConf, l)
	if err != nil {
		panic(err)
	}

	s := seeder{conn: dbConn}

	s.seed(seedProducts)
	s.seed(seedDishes)
}

type seeder struct {
	conn db.DBTX
}

type seedFn func(db.DBTX) error

func (s *seeder) seed(f seedFn) {
	if err := f(s.conn); err != nil {
		l.Error("failed to seed", "fn", f)
		panic(err)
	}
}
