package ceous

type (
	SchemaField interface {
		// String returns the string representation of the field. That is, its name.
		String() string

		// QualifiedString returns the name of the field qualified by the alias of
		// the given schema.
		QualifiedName(Schema) string

		IsPK() bool
		IsAutoInc() bool
	}

	AliasedSchemaField interface {
		SchemaField
		Reference() string
	}

	BaseSchemaField struct {
		name    string
		pk      bool
		autoInc bool
	}

	aliasSchemaField struct {
		SchemaField
		schema Schema
	}

	Schema interface {
		PrimaryKey() SchemaField
		Alias() string
		Table() string
		Columns() []SchemaField
		As(string) Schema
	}

	BaseSchema struct {
		id         SchemaField
		tableName  string
		alias      string
		ColumnsArr []SchemaField
	}

	aliasSchema struct {
		Schema
		alias string
	}
)

type FieldOption func(*BaseSchemaField)

func NewSchemaField(name string, options ...FieldOption) *BaseSchemaField {
	field := &BaseSchemaField{
		name: name,
	}
	for _, option := range options {
		option(field)
	}
	return field
}

func FieldPK(field *BaseSchemaField) {
	field.pk = true
}

func FieldAutoIncrement(field *BaseSchemaField) {
	field.autoInc = true
}

func (field *BaseSchemaField) String() string {
	return field.name
}

func (field *BaseSchemaField) QualifiedName(schema Schema) string {
	alias := schema.Alias()
	if alias != "" {
		return alias + "." + field.name
	}
	return field.name
}

func (field *BaseSchemaField) IsAutoInc() bool {
	return field.autoInc
}

func (field *BaseSchemaField) IsPK() bool {
	return field.pk
}

func NewAliasSchemaField(schema Schema, field SchemaField) AliasedSchemaField {
	return &aliasSchemaField{field, schema}
}

func (field *aliasSchemaField) Reference() string {
	return field.QualifiedName(field.schema)
}

func NewBaseSchema(tableName, alias string, columns ...SchemaField) *BaseSchema {
	baseSchema := &BaseSchema{
		tableName:  tableName,
		alias:      alias,
		ColumnsArr: columns,
	}
	return baseSchema
}

// FieldAlias creates a function that will create a SchemaField that will be
// bound to the schema passed.
//
// Example:
func FieldAlias(schema Schema) func(SchemaField) AliasedSchemaField {
	return func(field SchemaField) AliasedSchemaField {
		return NewAliasSchemaField(schema, field)
	}
}

func (schema *BaseSchema) PrimaryKey() SchemaField {
	return schema.id
}

func (schema *BaseSchema) Alias() string {
	return schema.alias
}

func (schema *BaseSchema) Table() string {
	return schema.tableName
}

func (schema *BaseSchema) Columns() []SchemaField {
	return schema.ColumnsArr
}

func (schema *BaseSchema) As(alias string) Schema {
	return &aliasSchema{schema, alias}
}

func (schema *aliasSchema) Alias() string {
	if schema.Schema.Alias() == "" {
		return schema.alias
	}
	return schema.Schema.Alias() + "_" + schema.alias
}
