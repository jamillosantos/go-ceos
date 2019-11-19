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
		Type      *ModelType
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

	ModelRelation struct {
		FromModel      *Model
		FromField      string
		FromColumnType *ModelType
		ToModel        *Model
		ToColumn       string
	}

	Model struct {
		_s           *myasthurts.Struct
		_explored    bool
		Connection   string
		PK           *ModelField
		Name         string
		TableName    string
		SchemaFields []*ModelField
		Columns      []*ModelColumn
		Fields       []*ModelField
		ColumnsMap   map[string]int
		Conditions   []*ModelCondition
		Relations    []*ModelRelation
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

func (m *Model) parseField(ctx *Ctx, t *myasthurts.TagParam, field *myasthurts.Field) error {
	if t.Value == "-" {
		return SkipField
	}

	isStructE := false

	fieldColumnName := ""

	f := &ModelField{
		Name:      field.Name,
		FieldName: field.Name, // TODO(jota): Let the developer to choose its default naming convention...
		Type:      NewModelType(ctx, field.RefType),
		Modifiers: make([]ModelFieldModifier, 0), // TODO(jota): To check this..
	}

	ctx.Imports.AddRefType(field.RefType)

	if t.Value != "" {
		f.FieldName = t.Value
		fieldColumnName = t.Value
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
			ctx.Reporter.Linef(" * PK: %s", f.Name)
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
			if fieldColumnName != "" {
				columnName = fieldColumnName + "_" + columnName
			}
			column := &ModelColumn{
				Column:    columnName,
				Type:      NewModelType(ctx, embeddedField.RefType),
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

	ctx.Reporter.Linef(" + %s: %s", field.Name, field.RefType.Name())
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

func (m *Model) parseFK(ctx *Ctx, t *myasthurts.TagParam, field *myasthurts.Field) error {
	model := ctx.EnsureModel(field.RefType.Name())
	relation := &ModelRelation{
		FromModel:      m,
		FromField:      field.Name,
		FromColumnType: NewModelType(ctx, field.RefType),
		ToModel:        model,
		ToColumn:       t.Value,
	}
	m.Relations = append(m.Relations, relation)
	ctx.ModelsImports.AddRefType(field.RefType)
	ctx.Reporter.Linef("   FK: %s(%s): %s", field.Name, relation.ToColumn, relation.FromColumnType.String())
	return nil
}

func (m *Model) ParseField(ctx *Ctx, field *myasthurts.Field) error {
	for _, t := range field.Tag.Params {
		switch t.Name {
		case "ceous":
			return m.parseField(ctx, &t, field)
		case "fk":
			return m.parseFK(ctx, &t, field)
		default:
			// If there is nothing is detected. No problem, just continue.
		}
	}
	// If nothing is found, just skip the field.
	return SkipField
}

func isModel(r myasthurts.RefType) bool {
	return r.Pkg().Name == "ceous" && r.Name() == "Model"
}

func isEmbedded(r myasthurts.RefType) bool {
	return r.Pkg().Name == "ceous" && r.Name() == "Embedded"
}

func NewModel(name string) *Model {
	return &Model{
		_explored:    true,
		Name:         name,
		TableName:    name, // TODO(jota): Set the table name convention.
		Connection:   "Default",
		SchemaFields: make([]*ModelField, 0),
		Columns:      make([]*ModelColumn, 0),
		Fields:       make([]*ModelField, 0),
		ColumnsMap:   make(map[string]int, 0),
		Relations:    make([]*ModelRelation, 0),
	}
}

func ParseModel(ctx *Ctx, s *myasthurts.Struct) (*Model, error) {
	m, ok := ctx.Models[s.Name()]
	if ok && m._explored {
		return m, nil
	}

	m = NewModel(s.Name())
	m._s = s
	for _, field := range s.Fields {
		if isModel(field.RefType) {
			tableNameTag := field.Tag.TagParamByName("tableName")
			if tableNameTag != nil {
				m.TableName = tableNameTag.Value
			}
			ctx.Reporter.Line("   TableName:", m.TableName)

			// Finds the tag name
			connectionTag := field.Tag.TagParamByName("conn")
			if connectionTag != nil {
				m.Connection = connectionTag.Value
			}
			ctx.Reporter.Line("   Connection:", m.Connection)
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
	ctx.Models[m.Name] = m
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
	ctxPkg := t.ctx.Imports.AddRefType(t.RefType)
	var r string
	if interfaceRef, ok := t.RefType.Type().(*myasthurts.Interface); ok && interfaceRef.Name() == "" {
		return interfaceRef.String()
	}
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

func (r *ModelRelation) RelationName() string {
	return r.FromModel.Name + "Model" + r.FromField + "Relation"
}

func (r *ModelRelation) PkType() *ModelType {
	for _, c := range r.FromModel.Columns {
		if c.Column == r.ToColumn {
			return c.Type
		}
	}
	return nil
}

func (r *ModelRelation) Column() *ModelColumn {
	for _, c := range r.FromModel.Columns {
		if c.Column == r.ToColumn {
			return c
		}
	}
	return nil
}
