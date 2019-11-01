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
		sq.QueryRower
		sq.QueryRowerContext
		sq.Preparer
		io.Closer
		Statisticer
		Transactioner
	}

	Connection interface {
		DB() DBProxy
	}

	BaseConnection struct {
		name string
		_db  DBProxy
	}
)

func NewConnection(name string, db DBProxy) Connection {
	return &BaseConnection{
		_db: db,
	}
}

// DB returns the real connection object for the database connection.
func (conn *BaseConnection) DB() DBProxy {
	return conn._db
}
