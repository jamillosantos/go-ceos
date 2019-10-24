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
	ErrNotWritable    = errors.New("ceous: record is not writable")
	ErrNewDocument    = errors.New("ceous: cannot updated a new document")
	ErrNoRowUpdate    = errors.New("ceous: update affected no rows")
	ErrNotFound       = errors.New("ceous: entity not found")
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

// NewStore returns a new base implementation of a Store. That does not aims to
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
		fields = store.schema.Columns()
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

	if len(autoIncPks) > 0 {
		pkRefs := make([]interface{}, len(autoIncPks))
		query.WriteString(" RETURNING ")
		for i, pkField := range autoIncPks {
			if i > 0 {
				query.WriteString(",")
			}
			query.WriteString(pkField)
			pkFieldRef, err := record.ColumnAddress(pkField)
			if err != nil {
				return errors.Wrap(err, fmt.Sprintf("could not find PK %s", pkField))
			}
			pkRefs[i] = pkFieldRef
		}

		rows, err := store.runner.Query(query.String(), columnValues...)
		if err != nil {
			return errors.Wrap(err, "query error")
		}
		defer rows.Close()

		if rows.Next() {
			err = rows.Scan(pkRefs...)
			if err != nil {
				return errors.Wrap(err, "scan error")
			}
		}
	} else {
		_, err = store.runner.Exec(query.String(), columnValues...)
		if err != nil {
			return errors.Wrap(err, " error")
		}
	}

	// TODO(jota): Uncomment this.
	// record.setWritable(true)
	record.setPersisted()
	return nil
}

func (store *BaseStore) Update(record Record, fields ...SchemaField) (int64, error) {
	if !record.IsWritable() {
		return 0, ErrNotWritable
	}

	if !record.IsPersisted() {
		return 0, ErrNewDocument
	}

	// TODO(jota): To take in consideration the primary key.
	/*
		if record.GetID().IsEmpty() {
			return 0, ErrEmptyID
		}
	*/

	if len(fields) == 0 {
		fields = store.schema.Columns()
	}

	var (
		pkNames    = make([]string, 0, 1)
		pkValues   = make([]interface{}, 0, 1)
		fieldName  string
		fieldValue interface{}
		err        error
	)

	for _, col := range store.schema.Columns() {
		if !col.IsPK() {
			continue
		}
		fieldName = col.String()
		fieldValue, err = record.Value(fieldName)
		if err != nil {
			return 0, err
		}
		pkNames = append(pkNames, fieldName)
		pkValues = append(pkValues, fieldValue)
	}

	var (
		columnNames  = make([]string, 0, len(fields))
		columnValues = make([]interface{}, 0, len(fields))
	)
	// remove the ID from there
	for _, col := range fields {
		fieldName = col.String()
		fieldValue, err = record.Value(fieldName)
		if err != nil {
			return 0, err
		}
		columnNames = append(columnNames, fieldName)
		columnValues = append(columnValues, fieldValue)
	}

	/*
		// TODO(jota): Add support for virtual columns.

			virtualCols, virtualColValues := virtualColumns(record, columnNames)
			columnNames = append(columnNames, virtualCols...)
			values = append(values, virtualColValues...)
	*/

	var query bytes.Buffer
	query.WriteString("UPDATE ")
	query.WriteString(store.schema.Table())
	query.WriteString(" SET ")
	for i, col := range columnNames {
		if i > 0 {
			query.WriteRune(',')
		}
		query.WriteString(col)
		query.WriteRune('=')
		query.WriteString(fmt.Sprintf("$%d", i+1))
	}
	query.WriteString(" WHERE ")
	for i, field := range pkNames {
		if i > 0 {
			query.WriteString(" AND ")
		}
		query.WriteString(field)
		query.WriteRune('=')
		query.WriteString(fmt.Sprintf("$%d", len(columnNames)+1+i)) // TODO(jota): Use a placeholder configuration to ensure multi database support.
	}

	result, err := store.runner.Exec(query.String(), append(columnValues, pkValues...)...)
	if err != nil {
		return 0, err
	}

	cnt, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}

	if cnt == 0 {
		return 0, ErrNoRowUpdate
	}

	return cnt, nil
}

func (store *BaseStore) Delete(record Record) error {
	if !record.IsWritable() {
		return ErrNotWritable
	}

	if !record.IsPersisted() {
		return ErrNewDocument
	}

	var (
		pkNames    = make([]string, 0, 1)
		pkValues   = make([]interface{}, 0, 1)
		fieldName  string
		fieldValue interface{}
		err        error
	)

	for _, col := range store.schema.Columns() {
		if !col.IsPK() {
			continue
		}
		fieldName = col.String()
		fieldValue, err = record.Value(fieldName)
		if err != nil {
			return err
		}
		pkNames = append(pkNames, fieldName)
		pkValues = append(pkValues, fieldValue)
	}

	var query bytes.Buffer
	query.WriteString("DELETE FROM ")
	query.WriteString(store.schema.Table())
	query.WriteString(" WHERE ")
	for i, field := range pkNames {
		if i > 0 {
			query.WriteString(" AND ")
		}
		query.WriteString(field)
		query.WriteRune('=')
		query.WriteString(fmt.Sprintf("$%d", 1+i)) // TODO(jota): Use a placeholder configuration to ensure multi database support.
	}

	_, err = store.runner.Exec(query.String(), pkValues...)
	return err
}
