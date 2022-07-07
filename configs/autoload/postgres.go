package autoload

import (
	"os"
)

type PostgresConfig struct {
	Host     string
	Username string
	Password string
	Port     string
	Database string
}

var Postgres PostgresConfig

func init() {
	Postgres = PostgresConfig{
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	}
}
