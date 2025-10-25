package migrate

import (
	"context"
	"fmt"
	"net/url"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/sethvargo/go-envconfig"
)

type Config struct {
	Dir       string `env:"MIGRATE_DIR,default=database/migrations"`
	TableName string `env:"MIGRATE_TABLE_NAME,default=schema_migrations"`
}

func NewConfig() Config {
	cfg := Config{}

	err := envconfig.Process(context.Background(), &cfg)
	if err != nil {
		panic(err)
	}

	return cfg
}

// DBConfig - database configuration
// Available params https://www.postgresql.org/docs/15/libpq-connect.html#LIBPQ-PARAMKEYWORDS
//
// Defaults:
//
//	user: OS user name
//	host: localhost
//	port: 5432 ?
//	dbname: same as user name
//	sslmode: prefered
type DBConfig struct {
	Adapter string

	Scheme   string `env:"DB_SCHEME,default=postgresql"`
	User     string `env:"DB_USER,default=pyra"`
	Password string `env:"DB_PASSWORD,default=pyra"`
	DBName   string `env:"DB_NAME,default=pyra_dev"`
	Host     string `env:"DB_HOST,default=0.0.0.0"`
	Port     uint   `env:"DB_PORT,default=5432"`

	Attrs url.Values
}

func (c DBConfig) String() string {
	u := url.URL{
		Scheme:   c.Scheme,
		User:     url.UserPassword(c.User, c.Password),
		Host:     fmt.Sprintf("%s:%d", c.Host, c.Port),
		Path:     c.DBName,
		RawQuery: c.Attrs.Encode(),
	}

	return u.String()
}

func NewDBConfig(adapter string) DBConfig {
	c := DBConfig{Adapter: adapter, Attrs: url.Values{}}

	err := envconfig.Process(context.Background(), &c)
	if err != nil {
		panic(err)
	}

	return c
}
