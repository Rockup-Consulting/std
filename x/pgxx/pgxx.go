package pgxx

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
)

type Args struct {
	Host              string
	Port              string
	Database          string
	UserName          string
	Password          string
	ConnectionTimeout time.Duration
}

var (
	DevArgs = Args{
		Host:              "localhost",
		Port:              "5432",
		Database:          "celbux",
		UserName:          "postgres",
		Password:          "",
		ConnectionTimeout: 1,
	}
)

func parseString(a Args) (string, error) {
	var str string = "postgres://"

	if a.UserName == "" {
		return "", errors.New("parsing connection string: missing UserName")
	}
	str += a.UserName

	if a.Password != "" {
		str = fmt.Sprintf("%s:%s", str, a.Password)
	}

	if a.Host == "" {
		return "", errors.New("parsing connection string: missing Host")
	}

	str = fmt.Sprintf("%s@%s", str, a.Host)

	if a.Database == "" {
		return "", errors.New("parsing connection string: missing Database")
	}

	str = fmt.Sprintf("%s/%s", str, a.Database)

	if a.ConnectionTimeout != 0 {
		str = fmt.Sprintf("%s?connect_timeout=%d", str, int(a.ConnectionTimeout.Seconds()))
	}

	return str, nil
}

// NewDriver opens a Postgres connection using the jackc/pgx driver, it then returns a sql.Conn
//
// Docs: https://github.com/jackc/pgx/wiki/Getting-started-with-pgx-through-database-sql
func NewDriver(ctx context.Context, args Args) (*sql.DB, error) {
	dbURI, err := parseString(args)
	if err != nil {
		return nil, err
	}

	fmt.Println(dbURI)

	db, err := sql.Open("pgx", dbURI)
	if err != nil {
		return nil, err
	}

	if err := db.PingContext(ctx); err != nil {
		return nil, err
	}

	return db, nil
}
