package ceous

import (
	"context"
	"database/sql"

	sq "github.com/elgris/sqrl"
)

type Query interface {
	RawQuery() (*sql.Rows, error)
	RawQueryContext(context.Context) (*sql.Rows, error)
	RawQueryRow() sq.RowScanner
	RawQueryRowContext(context.Context) sq.RowScanner
}

type QueryOption func(q Query)

type BaseQuery struct {
	Builder  sq.SelectBuilder
	schema   Schema
	runner   sq.DBProxy
	useCache bool
}

func NewBaseQuery(schema Schema, options ...QueryOption) *BaseQuery {
	q := &BaseQuery{
		schema: schema,
	}
	// Apply all options to the recently created query.
	for _, option := range options {
		option(q)
	}
	return q
}

func WithDB(db sq.DBProxy) QueryOption {
	return func(q Query) {
		if bq, ok := q.(*BaseQuery); ok {
			bq.runner = db
		}
	}
}

func WithCache(useCache bool) QueryOption {
	return func(q Query) {
		if bq, ok := q.(*BaseQuery); ok {
			bq.useCache = useCache
		}
	}
}

func (q *BaseQuery) Where(cond Condition) *BaseQuery {
	q.Builder.Where(cond(q.schema))
	return q
}

func (q *BaseQuery) RawQuery() (*sql.Rows, error) {
	return q.Builder.RunWith(q.runner).Query()
}

func (q *BaseQuery) RawQueryContext(ctx context.Context) (*sql.Rows, error) {
	return q.Builder.RunWith(q.runner).QueryContext(ctx)
}

func (q *BaseQuery) RawQueryRow() sq.RowScanner {
	return q.Builder.RunWith(q.runner).QueryRow()
}

func (q *BaseQuery) RawQueryRowContext(ctx context.Context) sq.RowScanner {
	return q.Builder.RunWith(q.runner).QueryRowContext(ctx)
}
