package ceous

import (
	"bytes"
	"database/sql"
	"fmt"

	sq "github.com/elgris/sqrl"
	"github.com/pkg/errors"
)

var (
	ErrNonNewDocument = errors.New("ceous: cannot insert a non new document")
)

type (
	// BaseStore is the generic implementation of the Store.
	//
	// TODO(jota): To document it...
	BaseStore struct {
		schema       Schema
		db           *sql.DB
		runner       sq.DBProxy
		disableCache bool
	}
)

// NewStore returns a new base implementaiton of a Store. That does not aims to
// be used by itself, but to be used as composition on real stores.
//
// TODO(jota): Improve documentation.
func NewStore(schema Schema, options ...CeousOption) *BaseStore {
	store := &BaseStore{
		schema: schema,
	}
	for _, option := range options {
		option(store)
	}
	if store.disableCache {
		panic("not implemented")
	} else {
		store.runner = sq.NewStmtCacheProxy(store.db)
	}
	return store
}

// Insert will insert a record.
//
// `fields` define what fields are going to be used on the insert.
func (store *BaseStore) Insert(record Record, fields ...SchemaField) error {
	if record.IsPersisted() {
		return ErrNonNewDocument
	}

	if len(fields) == 0 {
		for _, field := range store.schema.Columns() {
			if field.IsAutoInc() {
				continue
			}
			fields = append(fields, field)
		}
	}

	var (
		autoIncPks   = make([]string, 0, 1)
		columnNames  = make([]string, 0, len(fields))
		columnValues = make([]interface{}, 0, len(fields))
		fieldName    string
		fieldValue   interface{}
		err          error
	)
	for _, col := range fields {
		fieldName = col.String()
		if col.IsAutoInc() {
			autoIncPks = append(autoIncPks, fieldName)
			continue
		}
		fieldValue, err = record.Value(fieldName)
		if err != nil {
			return err
		}
		columnNames = append(columnNames, fieldName)
		columnValues = append(columnValues, fieldValue)
	}

	// TODO(jota): To add support for virtual columns.
	// cols = append(cols, virtualCols...)
	// values = append(values, virtualColValues...)

	var colBuf bytes.Buffer
	var valBuf bytes.Buffer

	for i, col := range columnNames {
		if i != 0 {
			colBuf.WriteRune(',')
			valBuf.WriteRune(',')
		}
		colBuf.WriteString(col)
		valBuf.WriteString(fmt.Sprintf("$%d", i+1))
	}

	var query bytes.Buffer
	query.WriteString("INSERT INTO ")
	query.WriteString(store.schema.Table())
	query.WriteString(" (")
	query.Write(colBuf.Bytes())
	query.WriteString(") VALUES (")
	query.Write(valBuf.Bytes())
	query.WriteString(")")

	/*
		TODO(jota): To uncomment this.
			if true {
				var pk interface{}
				pk, err = record.ColumnAddress(store.schema.ID().String())
				if err != nil {
					return err
				}

				query.WriteString(fmt.Sprintf(" RETURNING %s", store.schema.ID().String()))
				//err = s.runner.QueryRow(query.String(), values...).Scan(pk)
				rows, err := store.runner.Query(query.String(), columnValues...)
				if err != nil {
					return err
				}
				if rows.Next() {
					err = rows.Scan(pk)
					rows.Close()
					if err != nil {
						return err
					}
				}
			} else {
	*/
	_, err = store.runner.Exec(query.String(), columnValues...)
	/*
		}
	*/

	if err != nil {
		return err
	}

	// TODO(jota): Uncomment this.
	// record.setWritable(true)
	record.setPersisted()
	return nil
}
