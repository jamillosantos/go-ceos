package models

type (
	BaseSchemaField struct {
		Name       string
		ColumnName string
		Type       string
	}

	BaseSchema struct {
		Name      string
		TableName string
		Fields    []*BaseSchemaField
		FieldsMap map[string]int
	}

	SchemaField struct {
		Name string
		// Type       string
		ColumnName string
	}

	Schema struct {
		IsModel    bool
		Name       string
		BaseSchema *BaseSchema
		Fields     []*SchemaField
	}
)

func NewBaseSchema(name, tableName string) *BaseSchema {
	return &BaseSchema{
		Name:      name,
		TableName: tableName,
		Fields:    make([]*BaseSchemaField, 0),
		FieldsMap: make(map[string]int),
	}
}

func NewSchema(name string, baseSchema *BaseSchema) *Schema {
	return &Schema{
		Name:       name,
		BaseSchema: baseSchema,
		Fields:     make([]*SchemaField, 0),
	}
}

// AddField adds a new instance of BaseSchemaField to the schema fields and
// returns it.
func (schema *BaseSchema) AddField(name, columnName string) *BaseSchemaField {
	field := &BaseSchemaField{
		Name:       name,
		ColumnName: columnName,
	}
	schema.Fields = append(schema.Fields, field)
	schema.FieldsMap[field.ColumnName] = len(schema.Fields) - 1
	return field
}

// AddField adds a new instance of SchemaField to the schema fields and returns
// it.
func (schema *Schema) AddField(name, columnName string) *SchemaField {
	field := &SchemaField{
		Name: name,
		// Type:       t,
		ColumnName: columnName,
	}
	schema.Fields = append(schema.Fields, field)
	return field
}
