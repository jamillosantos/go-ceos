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

type (
	SelectForType        = sq.SelectForType
	SelectForLockingType = sq.SelectForLockingType
)

const (
	// ForUpdate marks a use a FOR UPDATE.
	ForUpdate = sq.ForUpdate
	// ForNoKeyUpdate marks a use a FOR NO KEY UPDATE.
	ForNoKeyUpdate = sq.ForNoKeyUpdate
	// ForShare marks a use a FOR SHARE.
	ForShare = sq.ForShare
	// ForKeyShare marks a use a FOR KEY SHARE.
	ForKeyShare = sq.ForKeyShare
)

const (
	// ForUpdateTypeNone will generate the FOR [OPTION] with no modifier.
	ForUpdateTypeNone = sq.ForUpdateTypeNone
	// SkipLocked will generate the FOR [OPTION] SKIP LOCKED modifier.
	SkipLocked = sq.SkipLocked
	// NoWait will generate the FOR [OPTION] NOWAIT modifier.
	NoWait = sq.NoWait
)

type (
	Query interface {
		RawQuery() (*sql.Rows, error)
		RawQueryContext(context.Context) (*sql.Rows, error)
		RawQueryRow() sq.RowScanner
		RawQueryRowContext(context.Context) sq.RowScanner
	}

	BaseQuery struct {
		_modified *sq.SelectBuilder
		Schema    Schema
		Runner    DBRunner
		runner    DBRunner

		where          []interface{}
		selectedFields []SchemaField
		fieldsExcluded map[SchemaField]bool
		Relations      []Relation

		builder *sq.SelectBuilder

		limit  uint64
		offset uint64
		order  []string

		RecordScanner RecordScanner

		forType        *sq.SelectForType
		forLockingType sq.SelectForLockingType

		IsDefaultScenarioDisabled bool
	}

	CeousOption func(q interface{})
)

func NewBaseQuery(options ...CeousOption) *BaseQuery {
	q := &BaseQuery{}
	// Apply all options to the recently created query.
	for _, option := range options {
		option(q)
	}
	q.runner = q.Runner
	return q
}

// DisableDefaultScenario sets the flag IsDefaultScenarioDisabled to true.
//
// TODO(jota): See more DefaultScenario for queries.
func DisableDefaultScenario(q *BaseQuery) {
	q.IsDefaultScenarioDisabled = true
}

// WithConn returns a query option for creating .
func WithConn(conn Connection) CeousOption {
	return func(obj interface{}) {
		switch q := obj.(type) {
		case *BaseQuery:
			q.Runner = conn.DB()
		case *BaseStore:
			q._runner = conn.DB()
		default:
			panic(errors.New(fmt.Sprintf("invalid obj: %T", obj))) // TODO(jota): To formalize this error
		}
	}
}

// WithRunner returns a query option for setting the runner for a transaction.
func WithRunner(runner DBRunner) CeousOption {
	return func(obj interface{}) {
		switch q := obj.(type) {
		case *BaseQuery:
			q.Runner = runner
		case *BaseStore:
			q._runner = runner
		default:
			panic(errors.New(fmt.Sprintf("invalid obj: %T", obj))) // TODO(jota): To formalize this error
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

func (q *BaseQuery) With(d interface{}) {
	switch modifier := d.(type) {
	case RecordScannerColumns:
		q.Select(modifier.SelectColumns()...)
	case RecordScanner:
		q.RecordScanner = modifier
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

func (q *BaseQuery) applyConditions(sqQuery *sq.SelectBuilder) error {
	for _, cond := range q.where {
		switch c := cond.(type) {
		case Condition:
			sql, args, err := c(q.Schema).ToSql()
			if err != nil {
				return err
			}
			sqQuery.Where(sql, args...)
		case Sqlizer:
			sql, args, err := c.ToSql()
			if err != nil {
				return err
			}
			sqQuery.Where(sql, args...)
		default:
			return errors.Wrap(ErrConditionTypeNotSupported, fmt.Sprintf("%T not recognized", c))
		}
	}
	return nil
}

func (q *BaseQuery) OrderBy(fields ...interface{}) {
	if q.order == nil {
		q.order = make([]string, 0, len(fields))
	}
	for _, field := range fields {
		switch f := field.(type) {
		case string:
			q.order = append(q.order, f)
		case *string:
			q.order = append(q.order, *f)
		case fmt.Stringer:
			q.order = append(q.order, f.String())
		default:
			q.order = append(q.order, fmt.Sprint(f))
		}
	}
}

// Builder will prepare a *sq.SelectBuilder and return it with all fields,
// conditions and limits.
func (q *BaseQuery) Builder() (*sq.SelectBuilder, error) {
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

	sqQuery.PlaceholderFormat(sq.Dollar) // TODO(jota): To parametrize it.

	// If we have conditions to be added ...
	if len(q.where) > 0 {
		q.applyConditions(sqQuery)
	}

	// Applies the limit
	if q.limit > 0 {
		sqQuery.Limit(q.limit)
	}

	// Applies the offset
	if q.offset > 0 {
		sqQuery.Offset(q.offset)
	}

	if len(q.order) > 0 {
		sqQuery.OrderBy(q.order...)
	}

	if q.forType != nil {
		sqQuery.For(*q.forType, q.forLockingType)
	}
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

func (q *BaseQuery) Count() (int64, error) {
	tableName := q.Schema.Table()
	if q.Schema.Alias() != "" {
		tableName += " " + q.Schema.Alias()
	}

	sqQuery := sq.Select("COUNT(*)").From(tableName)
	// If we have conditions to be added ...
	if len(q.where) > 0 {
		q.applyConditions(sqQuery)
	}

	var n int64
	err := sqQuery.PlaceholderFormat(sq.Dollar).RunWith(q.runner).QueryRow().Scan(&n)
	return n, err
}

func (q *BaseQuery) RawQuery() (*sql.Rows, error) {
	builder, err := q.Builder()
	if err != nil {
		return nil, err
	}
	builder.PlaceholderFormat(sq.Dollar) // TODO(jota): Make this placeholder configurable and optimize this.
	sql, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}
	rows, err := q.runner.Query(sql, args...)
	return rows, err
}

func (q *BaseQuery) RawQueryContext(ctx context.Context) (*sql.Rows, error) {
	builder, err := q.Builder()
	if err != nil {
		return nil, err
	}
	builder.PlaceholderFormat(sq.Dollar) // TODO(jota): Make this placeholder configurable and optimize this.
	return builder.RunWith(q.runner).QueryContext(ctx)
}

func (q *BaseQuery) RawQueryRow() sq.RowScanner {
	q._modified.PlaceholderFormat(sq.Dollar) // TODO(jota): Make this placeholder configurable and optimize this.
	return q._modified.RunWith(q.runner).QueryRow()
}

func (q *BaseQuery) RawQueryRowContext(ctx context.Context) sq.RowScanner {
	q._modified.PlaceholderFormat(sq.Dollar) // TODO(jota): Make this placeholder configurable and optimize this.
	return q._modified.RunWith(q.runner).QueryRowContext(ctx)
}

func (q *BaseQuery) For(t SelectForType, lockingType ...SelectForLockingType) {
	var lt sq.SelectForLockingType
	if len(lockingType) > 0 {
		lt = lockingType[0]
		q.forLockingType = lt
	}
	q.forType = &t
	if q._modified != nil {
		q._modified.For(*q.forType, q.forLockingType)
	}
}
