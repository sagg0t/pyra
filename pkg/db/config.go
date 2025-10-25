package db

import (
	"context"
	"fmt"
	"net/url"

	"github.com/sethvargo/go-envconfig"
)

// Config - database connection config.
// Available params https://www.postgresql.org/docs/15/libpq-connect.html#LIBPQ-PARAMKEYWORDS
//
// Defaults:
//
//	user: OS user name
//	host: localhost
//	port: 5432 ?
//	dbname: same as user name
//	sslmode: prefered
type Config struct {
	Adapter string

	Scheme   string `env:"DB_SCHEME,default=postgresql"`
	User     string `env:"DB_USER,default=pyra"`
	Password string `env:"DB_PASSWORD,default=pyra"`
	DBName   string `env:"DB_NAME,default=pyra_dev"`
	Host     string `env:"DB_HOST,default=0.0.0.0"`
	Port     uint   `env:"DB_PORT,default=5432"`

	Attrs url.Values
}

func (c Config) String() string {
	u := url.URL{
		Scheme:   c.Scheme,
		User:     url.UserPassword(c.User, c.Password),
		Host:     fmt.Sprintf("%s:%d", c.Host, c.Port),
		Path:     c.DBName,
		RawQuery: c.Attrs.Encode(),
	}

	out := u.String()
	if c.DBName == "" {
		out += "/"
	}

	return out
}

func NewConfig(adapter string) Config {
	c := Config{Adapter: adapter, Attrs: url.Values{}}

	err := envconfig.Process(context.Background(), &c)
	if err != nil {
		panic(err)
	}

	return c
}
