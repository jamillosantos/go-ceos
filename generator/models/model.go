package models

import (
	"strings"

	"github.com/jamillosantos/go-ceous/generator/helpers"
	"github.com/jamillosantos/go-ceous/generator/reporters"
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
		ctx     *GenContext
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

func NewModelType(ctx *GenContext, refType myasthurts.RefType) *ModelType {
	return &ModelType{
		ctx:     ctx,
		RefType: refType,
	}
}

func NewModelCondition(ctx *GenContext, field string, refType myasthurts.RefType) *ModelCondition {
	nfM := strings.Replace(field, ".", "", 0) // Prepare name for method by removing possible dots from property members.
	return &ModelCondition{
		Field:         field,
		NameForMethod: nfM,
		Type:          NewModelType(ctx, refType),
		Conditions:    make([]*ModelConditionCondition, 0),
	}
}

func isStructEmbedded(s *myasthurts.Struct) bool {
	for _, f := range s.Fields {
		if isEmbedded(f.RefType) {
			return true
		}
	}
	return false
}

// isModel returns whether or not a given `r` is represents a model.
func isModel(r myasthurts.RefType) bool {
	return r.Pkg().Name == "ceous" && r.Name() == "Model"
}

func isEmbedded(r myasthurts.RefType) bool {
	return r.Pkg().Name == "ceous" && r.Name() == "Embedded"
}

func NewModel(s *myasthurts.Struct) *Model {
	return &Model{
		Name:         s.Name(),
		TableName:    s.Name(), // TODO(jota): Set the table name convention.
		Connection:   "Default",
		SchemaFields: make([]*ModelField, 0),
		Columns:      make([]*ModelColumn, 0),
		Fields:       make([]*ModelField, 0),
		ColumnsMap:   make(map[string]int, 0),
		Relations:    make([]*ModelRelation, 0),
	}
}

type FieldContext struct {
	Gen      *GenContext
	Model    *Model
	Schema   *Schema
	Reporter reporters.Reporter
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
