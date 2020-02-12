package ceous

import (
	"context"
	"database/sql"

	sq "github.com/elgris/sqrl"
)

type (
	TxProxy interface {
		DBRunner
		Commit() error
		Rollback() error
	}

	BaseTxRunner struct {
		tx *sql.Tx
	}
)

// NewTx returns a new transaction abstraction.
func NewTx(tx *sql.Tx) *BaseTxRunner {
	return &BaseTxRunner{
		tx: tx,
	}
}

func (tx *BaseTxRunner) Exec(query string, args ...interface{}) (sql.Result, error) {
	return tx.tx.Exec(query, args...)
}

func (tx *BaseTxRunner) ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	return tx.tx.ExecContext(ctx, query, args...)
}

func (tx *BaseTxRunner) Query(query string, args ...interface{}) (*sql.Rows, error) {
	return tx.tx.Query(query, args...)
}

func (tx *BaseTxRunner) QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error) {
	return tx.tx.QueryContext(ctx, query, args...)
}

func (tx *BaseTxRunner) QueryRow(query string, args ...interface{}) *sql.Row {
	return tx.tx.QueryRow(query, args)
}

func (tx *BaseTxRunner) QueryRowContext(ctx context.Context, query string, args ...interface{}) sq.RowScanner {
	return tx.tx.QueryRowContext(ctx, query, args...)
}

func (tx *BaseTxRunner) Prepare(query string) (*sql.Stmt, error) {
	return tx.tx.Prepare(query)
}

func (tx *BaseTxRunner) PrepareContext(ctx context.Context, query string) (*sql.Stmt, error) {
	return tx.tx.PrepareContext(ctx, query)
}

func (tx *BaseTxRunner) Commit() error {
	return tx.tx.Commit()
}

func (tx *BaseTxRunner) Rollback() error {
	return tx.tx.Rollback()
}
