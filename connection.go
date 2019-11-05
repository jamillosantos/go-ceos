package ceous

import (
	"context"
	"database/sql"
	"io"

	sq "github.com/elgris/sqrl"
)

type (
	Pinger interface {
		Ping() error
	}

	PingerContext interface {
		PingContext(ctx context.Context) error
	}

	Transactioner interface {
		Begin() (*sql.Tx, error)
		BeginTx(ctx context.Context, opts *sql.TxOptions) (*sql.Tx, error)
	}

	Statisticer interface {
		Stats() sql.DBStats
	}

	DBProxy interface {
		sq.Execer
		sq.ExecerContext
		sq.Queryer
		sq.QueryerContext
		QueryRow(string, ...interface{}) *sql.Row
		QueryRowContext(context.Context, string, ...interface{}) *sql.Row
		sq.Preparer
		io.Closer
		Statisticer
		Transactioner
	}

	Connection interface {
		DB() DBProxy
	}

	BaseConnection struct {
		_db DBProxy
	}
)

func NewConnection(db DBProxy) *BaseConnection {
	return &BaseConnection{
		_db: db,
	}
}

// DB returns the real connection object for the database connection.
func (conn *BaseConnection) DB() DBProxy {
	return conn._db
}
