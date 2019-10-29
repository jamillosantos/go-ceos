package models

import (
	"errors"

	"github.com/jamillosantos/go-ceous/generator/helpers"
	myasthurts "github.com/lab259/go-my-ast-hurts"
)

type (
	ModelFieldModifier func() string

	ModelField struct {
		Name      string
		FieldName string
		Type      myasthurts.RefType
		Modifiers []ModelFieldModifier
	}

	Model struct {
		_s        *myasthurts.Struct
		PK        *ModelField
		Name      string
		TableName string
		Fields    []*ModelField
	}
)

func fieldPK() string {
	return "ceous.FieldPK"
}

func fieldAutoInc() string {
	return "ceous.FieldAutoIncrement"
}

var SkipField = errors.New("field skipped")

func NewModelField(m *Model, field *myasthurts.Field) (*ModelField, error) {
	t := field.Tag.TagParamByName("ceous")
	if t == nil {
		return nil, SkipField
	}
	if t.Value == "-" {
		return nil, SkipField
	}
	f := &ModelField{
		Name:      field.Name,
		FieldName: field.Name, // TODO(jota): Let the developer to choose its default naming convention...
		Type:      field.RefType,
		Modifiers: make([]ModelFieldModifier, 0), // TODO(jota): To check this..
	}

	switch field.RefType.Type().(type) {
	case *myasthurts.Struct:
		// field.
	}

	if t.Value != "" {
		f.FieldName = t.Value
	}
	// TODO(jota): To support multiple options...
	for _, o := range t.Options {
		switch o {
		case "pk":
			if m.PK != nil {
				panic(errors.New("PK already defined")) // TODO(jota): To formalize this error.
			}
			f.Modifiers = append(f.Modifiers, fieldPK)
			m.PK = f
		case "autoincr":
			f.Modifiers = append(f.Modifiers, fieldAutoInc)
			m.PK = f
		}
	}

	return f, nil
}

func isModel(r myasthurts.RefType) bool {
	return r.Pkg().Name == "ceous" && r.Name() == "Model"
}

func NewModel(s *myasthurts.Struct) (*Model, error) {
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
		mField, err := NewModelField(m, field)
		if err == SkipField {
			continue
		} else if err != nil {
			return nil, err
		}
		m.Fields = append(m.Fields, mField)
	}
	return m, nil
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
