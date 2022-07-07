package connects

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	conf "main/configs/autoload"
	"os"
	"strconv"
)

func PostgresConnect() *pgx.Conn {
	host := conf.Postgres.Host
	port, _ := strconv.Atoi(conf.Postgres.Port)
	user := conf.Postgres.Username
	password := conf.Postgres.Password
	database := conf.Postgres.Database

	dbUrl := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, database)
	conn, err := pgx.Connect(context.Background(), dbUrl)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	return conn
}
