package ceous

import (
	"context"
	"database/sql"
	"fmt"

	sq "github.com/elgris/sqrl"
	"github.com/pkg/errors"
)

var (
	ErrConditionTypeNotSupported = errors.New("condition type not supported")
)

type Query interface {
	RawQuery() (*sql.Rows, error)
	RawQueryContext(context.Context) (*sql.Rows, error)
	RawQueryRow() sq.RowScanner
	RawQueryRowContext(context.Context) sq.RowScanner
}

type BaseQuery struct {
	_modified *sq.SelectBuilder
	Schema    Schema
	db        *sql.DB
	runner    sq.DBProxy

	where          []interface{}
	selectedFields []SchemaField
	fieldsExcluded map[SchemaField]bool

	builder *sq.SelectBuilder

	limit  uint64
	offset uint64

	disableCache              bool
	IsDefaultScenarioDisabled bool
}

type CeousOption func(q interface{})

func NewBaseQuery(options ...CeousOption) *BaseQuery {
	q := &BaseQuery{}
	// Apply all options to the recently created query.
	for _, option := range options {
		option(q)
	}
	if q.disableCache {
		// q.runner = q.db // TODO(jota): implement this
		panic("not implemented")
	} else {
		q.runner = sq.NewStmtCacheProxy(q.db)
	}
	return q
}

// DisableDefaultScenario sets the flag IsDefaultScenarioDisabled to true.
//
// TODO(jota): See more DefaultScenario for queries.
func DisableDefaultScenario(q *BaseQuery) {
	q.IsDefaultScenarioDisabled = true
}

// WithDB returns a query option for creating .
func WithDB(db *sql.DB) CeousOption {
	return func(obj interface{}) {
		switch q := obj.(type) {
		case *BaseQuery:
			q.db = db
		case *BaseStore:
			q.db = db
		default:
			panic(errors.New(fmt.Sprintf("invalid obj: %T", obj)))
		}
	}
}

// WithCache returns a query option that will enable or disable the cache.
func WithCache(useCache bool) CeousOption {
	return func(obj interface{}) {
		switch q := obj.(type) {
		case *BaseQuery:
			q.disableCache = useCache
		case *BaseStore:
			q.disableCache = useCache
		default:
			panic(errors.New(fmt.Sprintf("invalid obj: %T", obj)))
		}
	}
}

// WithSchema returns a query option that will set the schema of a Query. Useful
// for using aliases.
func WithSchema(schema Schema) CeousOption {
	return func(obj interface{}) {
		switch q := obj.(type) {
		case *BaseQuery:
			q.Schema = schema
		case *BaseStore:
			q.schema = schema
		default:
			panic(errors.New(fmt.Sprintf("invalid obj: %T", obj)))
		}
	}
}

func (q *BaseQuery) Select(fields ...SchemaField) {
	if len(fields) == 0 {
		return
	}

	// If we don't have any selected fields
	if q.selectedFields == nil {
		q.selectedFields = fields
	} else {
		q.selectedFields = append(q.selectedFields, fields...)
	}
}

func (q *BaseQuery) ExcludeFields(fields ...SchemaField) {
	if len(fields) == 0 {
		return
	}

	if q.fieldsExcluded == nil {
		// initializes
		q.fieldsExcluded = make(map[SchemaField]bool, len(fields))
	}

	// Defines the excluded fields...
	for _, field := range fields {
		q.fieldsExcluded[field] = true
	}
}

func (q *BaseQuery) Where(pred interface{}, args ...interface{}) {
	switch c := pred.(type) {
	case Condition:
		q.where = append(q.where, c)
	case string:
		q.where = append(q.where, &sqlCondition{c, args})
	case *string:
		q.where = append(q.where, &sqlCondition{*c, args})
	default:
		q.where = append(q.where, c)
	}
}

// Builder will prepare a *sq.SelectBuilder and return it with all fields,
// conditions and limits.
func (q *BaseQuery) Builder() (*sq.SelectBuilder, error) {
	if q._modified != nil {
		// If it was not modified since last `compile`, just
		// return the cached.
		return q._modified, nil
	}
	var (
		fields         []string
		selectedFields []SchemaField
	)
	if len(q.selectedFields) == 0 {
		selectedFields = q.Schema.Columns()
	} else {
		selectedFields = q.selectedFields
	}
	fields = make([]string, 0, len(selectedFields)-len(q.fieldsExcluded)) // All fields of the schema without the excluded fields.

	// Generate the fields list for selection
	for _, schemaField := range selectedFields {
		if _, ok := q.fieldsExcluded[schemaField]; ok {
			continue // Skip field...
		}
		if aliasField, ok := schemaField.(AliasedSchemaField); ok {
			fields = append(fields, aliasField.Reference())
			continue
		}
		fields = append(fields, schemaField.QualifiedName(q.Schema))
	}

	// Format the table
	tableName := q.Schema.Table()
	if q.Schema.Alias() != "" {
		tableName += " " + q.Schema.Alias()
	}

	// Creates the initial select
	sqQuery := sq.Select(fields...).From(tableName)

	// If we have conditions to be added ...
	if len(q.where) > 0 {
		for _, cond := range q.where {
			switch c := cond.(type) {
			case Condition:
				sql, args, err := c(q.Schema).ToSql()
				if err != nil {
					return nil, err
				}
				sqQuery.Where(sql, args...)
			case Sqlizer:
				sql, args, err := c.ToSql()
				if err != nil {
					return nil, err
				}
				sqQuery.Where(sql, args...)
			default:
				return nil, errors.Wrap(ErrConditionTypeNotSupported, fmt.Sprintf("%T not recognized", c))
			}
		}
	}

	// Applies the limit
	if q.limit > 0 {
		sqQuery.Limit(q.limit)
	}

	// Applies the offset
	if q.offset > 0 {
		sqQuery.Offset(q.offset)
	}

	q._modified = sqQuery

	return sqQuery, nil
}

// Limit will update the limit directive of this query.
//
// Warning: If you use the method `Builder` directly, be aware that this will
// affect the builder returned.
func (q *BaseQuery) Limit(limit uint64) *BaseQuery {
	q.limit = limit
	if q._modified != nil {
		q._modified.Limit(limit)
	}
	return q
}

// Offset will update the offset directive of this query.
//
// Warning: If you use the method `Builder` directly, be aware that this will
// affect the builder returned.
func (q *BaseQuery) Offset(offset uint64) *BaseQuery {
	q.offset = offset
	if q._modified != nil {
		q._modified.Offset(offset)
	}
	return q
}

func (q *BaseQuery) RawQuery() (*sql.Rows, error) {
	return q._modified.RunWith(q.runner).Query()
}

func (q *BaseQuery) RawQueryContext(ctx context.Context) (*sql.Rows, error) {
	return q._modified.RunWith(q.runner).QueryContext(ctx)
}

func (q *BaseQuery) RawQueryRow() sq.RowScanner {
	return q._modified.RunWith(q.runner).QueryRow()
}

func (q *BaseQuery) RawQueryRowContext(ctx context.Context) sq.RowScanner {
	return q._modified.RunWith(q.runner).QueryRowContext(ctx)
}
