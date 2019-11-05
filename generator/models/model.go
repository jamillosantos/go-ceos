package models

import (
	"errors"
	"strings"

	"github.com/jamillosantos/go-ceous/generator/helpers"
	myasthurts "github.com/lab259/go-my-ast-hurts"
)

type (
	ModelFieldModifier func() string

	ModelSchemaField struct {
		Name string
		Type *ModelType
	}

	ModelField struct {
		Name       string
		FieldName  string
		Type       *ModelType
		SchemaType string
		Modifiers  []ModelFieldModifier
	}

	ModelColumn struct {
		Column    string
		FullField string
		Modifiers []ModelFieldModifier
	}

	ModelConditionCondition struct {
		SchemaField string
		FullField   string
		Field       string
	}

	ModelCondition struct {
		Field         string
		NameForMethod string
		Type          *ModelType
		Conditions    []*ModelConditionCondition
	}

	ModelType struct {
		ctx     *Ctx
		RefType myasthurts.RefType
	}

	Model struct {
		_s           *myasthurts.Struct
		PK           *ModelField
		Name         string
		TableName    string
		SchemaFields []*ModelField
		Columns      []*ModelColumn
		Fields       []*ModelField
		ColumnsMap   map[string]int
		Conditions   []*ModelCondition
	}
)

func NewModelType(ctx *Ctx, refType myasthurts.RefType) *ModelType {
	return &ModelType{
		ctx:     ctx,
		RefType: refType,
	}
}

func NewModelCondition(ctx *Ctx, field string, refType myasthurts.RefType) *ModelCondition {
	nfM := strings.Replace(field, ".", "", 0) // Prepare name for method by removing possible dots from property members.
	return &ModelCondition{
		Field:         field,
		NameForMethod: nfM,
		Type:          NewModelType(ctx, refType),
		Conditions:    make([]*ModelConditionCondition, 0),
	}
}

func fieldPK() string {
	return "ceous.FieldPK"
}

func fieldAutoInc() string {
	return "ceous.FieldAutoIncrement"
}

var SkipField = errors.New("field skipped")

func isStructEmbedded(s *myasthurts.Struct) bool {
	for _, f := range s.Fields {
		if isEmbedded(f.RefType) {
			return true
		}
	}
	return false
}

func (m *Model) ParseField(ctx *Ctx, field *myasthurts.Field) error {
	t := field.Tag.TagParamByName("ceous")
	if t == nil {
		return SkipField
	}
	if t.Value == "-" {
		return SkipField
	}

	isStructE := false

	f := &ModelField{
		Name:      field.Name,
		FieldName: field.Name, // TODO(jota): Let the developer to choose its default naming convention...
		Type:      NewModelType(ctx, field.RefType),
		Modifiers: make([]ModelFieldModifier, 0), // TODO(jota): To check this..
	}

	ctx.AddRefType(field.RefType)

	if t.Value != "" {
		f.FieldName = t.Value
	}
	// To support multiple options...
	for _, o := range t.Options {
		switch o {
		case "pk":
			if m.PK != nil {
				return errors.New("PK already defined") // TODO(jota): To formalize this error.
			}
			f.Modifiers = append(f.Modifiers, fieldPK)
			m.PK = f
		case "autoincr":
			f.Modifiers = append(f.Modifiers, fieldAutoInc)
		}
	}

	if s, ok := field.RefType.Type().(*myasthurts.Struct); ok && isStructEmbedded(s) {
		isStructE = true
		cond := NewModelCondition(ctx, field.Name, field.RefType)
		for _, embeddedField := range s.Fields {
			ceousTag := embeddedField.Tag.TagParamByName("ceous")
			if ceousTag == nil || ceousTag.Value == "" {
				continue
			}
			columnName := ceousTag.Value
			if columnName == "" {
				columnName = embeddedField.Name // TODO(jota): Apply the default naming convention here.
			}
			column := &ModelColumn{
				Column:    ceousTag.Value,
				FullField: field.Name + "." + embeddedField.Name,
				Modifiers: f.Modifiers,
			}
			m.Columns = append(m.Columns, column)
			m.ColumnsMap[column.Column] = len(m.Columns) - 1

			cond.Conditions = append(cond.Conditions, &ModelConditionCondition{
				FullField:   field.Name + "." + embeddedField.Name,
				Field:       embeddedField.Name,
				SchemaField: m.Name + "." + field.Name + "." + embeddedField.Name,
			})
		}
		m.Conditions = append(m.Conditions, cond)
		f.SchemaType = field.RefType.Name()
	}

	m.Fields = append(m.Fields, f)
	m.SchemaFields = append(m.SchemaFields, f)

	if isStructE { /// If is a embedded struct, it does not need to have the field appended. Instead, only the m.PK will be set.
		return nil
	}
	cond := NewModelCondition(ctx, f.Name, field.RefType)
	cond.Conditions = append(cond.Conditions, &ModelConditionCondition{
		SchemaField: m.Name + "." + field.Name,
		FullField:   field.Name,
	})
	m.Conditions = append(m.Conditions, cond)

	column := &ModelColumn{
		Column:    f.FieldName,
		FullField: field.Name,
		Modifiers: f.Modifiers,
	}
	m.Columns = append(m.Columns, column)
	m.ColumnsMap[f.FieldName] = len(m.Columns) - 1

	return nil
}

func isModel(r myasthurts.RefType) bool {
	return r.Pkg().Name == "ceous" && r.Name() == "Model"
}

func isEmbedded(r myasthurts.RefType) bool {
	return r.Pkg().Name == "ceous" && r.Name() == "Embedded"
}

func NewModel(ctx *Ctx, s *myasthurts.Struct) (*Model, error) {
	m := &Model{
		_s:           s,
		Name:         s.Name(),
		TableName:    s.Name(), // TODO(jota): Set the table name convention.
		SchemaFields: make([]*ModelField, 0),
		Columns:      make([]*ModelColumn, 0),
		Fields:       make([]*ModelField, 0),
		ColumnsMap:   make(map[string]int, 0),
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
		err := m.ParseField(ctx, field)
		if err == SkipField {
			continue
		} else if err != nil {
			return nil, err
		}
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

func (t *ModelType) String() string {
	ctxPkg := t.ctx.AddRefType(t.RefType)
	var r string
	if ctxPkg.Alias != "." && ctxPkg.Alias != "-" {
		r = ctxPkg.Alias + "."
	}
	return r + t.RefType.Name()
}

func (c *ModelConditionCondition) StringField() string {
	if c.Field == "" {
		return ""
	}
	return "." + c.Field
}
