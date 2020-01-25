package models

type (
	BaseSchemaField struct {
		Name       string
		ColumnName string
		Type       string
		IsAutoIncr bool
		IsPK       bool
	}

	BaseSchema struct {
		Name         string
		FullName     string
		TableName    string
		Fields       []*BaseSchemaField
		FieldsIdxMap map[string]int
		FieldsMap    map[string]*BaseSchemaField
	}

	SchemaField struct {
		Name       string
		Type       string
		SchemaName string
		ColumnName string
	}

	Schema struct {
		IsModel    bool
		Name       string
		FullName   string
		BaseSchema *BaseSchema
		Fields     []*SchemaField
	}
)

func NewBaseSchema(name, tableName string) *BaseSchema {
	return &BaseSchema{
		Name:         name,
		FullName:     "baseSchema" + name,
		TableName:    tableName,
		Fields:       make([]*BaseSchemaField, 0),
		FieldsIdxMap: make(map[string]int),
		FieldsMap:    make(map[string]*BaseSchemaField),
	}
}

func NewSchema(name string, baseSchema *BaseSchema) *Schema {
	return &Schema{
		Name:       name,
		FullName:   "schema" + name,
		BaseSchema: baseSchema,
		Fields:     make([]*SchemaField, 0),
	}
}

// AddField adds a new instance of BaseSchemaField to the schema fields and
// returns it.
func (schema *BaseSchema) AddField(name, columnName string, isPK, isAutoIncr bool) *BaseSchemaField {
	field := &BaseSchemaField{
		Name:       name,
		ColumnName: columnName,
		IsPK:       isPK,
		IsAutoIncr: isAutoIncr,
	}
	schema.Fields = append(schema.Fields, field)
	schema.FieldsIdxMap[field.ColumnName] = len(schema.Fields) - 1
	schema.FieldsMap[field.ColumnName] = field
	return field
}

// AddField adds a new instance of SchemaField to the schema fields and returns
// it.
func (schema *Schema) AddField(name, t, schemaName, columnName string) *SchemaField {
	field := &SchemaField{
		Name:       name,
		Type:       t,
		SchemaName: schemaName,
		ColumnName: columnName,
	}
	schema.Fields = append(schema.Fields, field)
	return field
}
