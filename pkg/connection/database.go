package connection

import (
	"fmt"
	"github.com/jmoiron/sqlx"

	//drivers
	_ "github.com/jackc/pgx/v5"
	_ "github.com/jackc/pgx/v5/stdlib"
	_ "github.com/mattn/go-sqlite3"
)

type DBConfig struct {
	Driver             string
	Username           string
	Password           string
	Host               string
	Port               string
	Database           string
	MaxOpenConnections int
	MaxIdleConnections int

	PathToDB string //for sqlite
}

func CreateDBConnection(conf DBConfig) (*sqlx.DB, error) {
	conn, err := dbConnection(conf)
	if err != nil {
		return nil, err
	}
	if err := conn.Ping(); err != nil {
		_ = conn.Close()
		return nil, err
	}
	return conn, nil
}

func dbConnection(conf DBConfig) (*sqlx.DB, error) {
	switch conf.Driver {
	case "postgres":
		return postgresConnection(conf)
	case "sqlite":
		return sqliteConnection(conf)
	default:
		return nil, fmt.Errorf("not found dbConnection for driver: %s", conf.Driver)
	}
}

func postgresConnection(conf DBConfig) (*sqlx.DB, error) {
	conn, err := sqlx.Connect(
		"pgx",
		fmt.Sprintf(
			"%s://%s:%s@%s:%s/%s",
			conf.Driver,
			conf.Username,
			conf.Password,
			conf.Host,
			conf.Port,
			conf.Database,
		),
	)
	if err != nil {
		return nil, err
	}
	if conf.MaxOpenConnections > 0 {
		conn.SetMaxOpenConns(conf.MaxOpenConnections)
	}
	if conf.MaxIdleConnections > 0 {
		conn.SetMaxIdleConns(conf.MaxIdleConnections)
	}
	return conn, nil
}

func sqliteConnection(conf DBConfig) (*sqlx.DB, error) {
	conn, err := sqlx.Open("sqlite3", conf.PathToDB)
	if err != nil {
		return nil, err
	}
	if conf.MaxOpenConnections > 0 {
		conn.SetMaxOpenConns(conf.MaxOpenConnections)
	}
	if conf.MaxIdleConnections > 0 {
		conn.SetMaxIdleConns(conf.MaxIdleConnections)
	}
	return conn, nil
}
