package models

type (
	SchemaField struct {
		Name       string
		ColumnName string
		Type       string
	}

	Schema struct {
		Name    string
		IsModel bool
		Fields  []*SchemaField
	}
)

func NewSchema(name string) *Schema {
	return &Schema{
		Name:   name,
		Fields: make([]*SchemaField, 0),
	}
}

// AddField adds a new instance of SchemaField to the schema fields and returns
// it.
func (schema *Schema) AddField(name, columnName string) *SchemaField {
	field := &SchemaField{
		Name:       name,
		ColumnName: columnName,
	}
	schema.Fields = append(schema.Fields, field)
	return field
}
