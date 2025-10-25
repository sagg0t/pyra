package test

import (
	"pyra/pkg/session"

	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load("/Users/sagg0t/devel/pyra/.env.test"); err != nil {
		panic(err)
	}

	session.SetupSessionStore()
}
