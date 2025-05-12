package database

import (
	"context"
	"database/sql"
	"github.com/AlekseyPorandaykin/go-template/pkg/metrics"
	"github.com/jmoiron/sqlx"
	"time"
)

type Database interface {
	NamedExecContext(ctx context.Context, query string, arg interface{}) (sql.Result, error)
	GetContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error
	SelectContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error
	ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error)
	QueryxContext(ctx context.Context, query string, args ...interface{}) (*sqlx.Rows, error)
	NamedQueryContext(ctx context.Context, query string, arg interface{}) (*sqlx.Rows, error)
	QueryRowxContext(ctx context.Context, query string, args ...interface{}) *sqlx.Row
}

type MetricDB struct {
	db Database
}

func NewMetricDB(db Database) *MetricDB {
	return &MetricDB{db: db}
}

func (d *MetricDB) NamedExecContext(ctx context.Context, database, name, query string, arg interface{}) (sql.Result, error) {
	defer func(start time.Time) {
		metrics.DurationExecuteQueryDB(database, name, time.Since(start))
	}(time.Now())
	metrics.IncCountQueryDB(database, name)
	res, err := d.db.NamedExecContext(ctx, query, arg)
	if err != nil {
		metrics.IncErrorQueryDB(database, name)
	}
	return res, err
}
