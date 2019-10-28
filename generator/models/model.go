package models

import (
	"fmt"

	"github.com/jamillosantos/go-ceous/generator/helpers"
	myasthurts "github.com/lab259/go-my-ast-hurts"
)

type (
	ModelFieldModifier interface{} // TODO(jota): To specify this...

	ModelField struct {
		Name      string
		FieldName string
		Type      myasthurts.RefType
		Modifiers []ModelFieldModifier
	}

	Model struct {
		_s        *myasthurts.Struct
		Name      string
		TableName string
		Fields    []*ModelField
	}
)

func NewModelField(field *myasthurts.Field) *ModelField {
	f := &ModelField{
		Name:      field.Name,
		FieldName: field.Name, // TODO(jota): Let the developer to choose its default naming convention...
		Type:      field.RefType,
		Modifiers: make([]ModelFieldModifier, 0), // TODO(jota): To check this..
	}

	for _, t := range field.Tag.Params {
		switch t.Name {
		case "ceous":
			if t.Value != "" {
				f.FieldName = t.Value
			}
			// TODO(jota): To support multiple options...
			for _, o := range t.Options {
				fmt.Println(" *", o)
			}
		}
	}

	return f
}

func isModel(r myasthurts.RefType) bool {
	return r.Pkg().Name == "ceous" && r.Name() == "Model"
}

func NewModel(s *myasthurts.Struct) *Model {
	m := &Model{
		Name:      s.Name(),
		TableName: s.Name(), // TODO(jota): Set the table name convention.
		Fields:    make([]*ModelField, 0, len(s.Fields)),
	}
	for _, field := range s.Fields {
		if isModel(field.RefType) {
			tableNameTag := field.Tag.TagParamByName("tableName")
			if tableNameTag != nil {
				m.TableName = tableNameTag.Value
			}
		}
		if field.Name == "" {
			continue
		}
		m.Fields = append(m.Fields, NewModelField(field))
	}
	return m
}

func (m *Model) SchemaName() string {
	return "schema" + m.Name
}

func (m *Model) BaseSchemaName() string {
	return "baseSchema" + m.Name
}

func (m *Model) QueryName() string {
	return helpers.CamelCase(m.Name) + "Query"
}

func (m *Model) StoreName() string {
	return helpers.CamelCase(m.Name) + "Store"
}
