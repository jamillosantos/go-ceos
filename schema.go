package ceous

import (
	"fmt"
)

type SchemaField interface {
	// String returns the string representation of the field. That is, its name.
	String() string
	// QualifiedString returns the name of the field qualified by the alias of
	// the given schema.
	QualifiedName(Schema) string
}

type BaseSchemaField struct {
	name string
}

func NewSchemaField(name string) *BaseSchemaField {
	return &BaseSchemaField{
		name: name,
	}
}

func (field *BaseSchemaField) String() string {
	return field.name
}

func (field *BaseSchemaField) QualifiedName(schema Schema) string {
	return fmt.Sprintf("%s.%s", schema.Alias(), field.name)
}

type Schema interface {
	Alias() string
	Table() string
	Columns() []SchemaField
}

type BaseSchema struct {
	tableName string
	alias     string
	columns   []SchemaField
}

func NewBaseSchema(tableName, alias string, columns ...SchemaField) *BaseSchema {
	baseSchema := &BaseSchema{
		tableName: tableName,
		alias:     alias,
		columns:   columns,
	}
	return baseSchema
}

func (schema *BaseSchema) Alias() string {
	return schema.alias
}

func (schema *BaseSchema) Table() string {
	return schema.tableName
}

func (schema *BaseSchema) Columns() []SchemaField {
	return schema.columns
}
