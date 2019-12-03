package models

import (
	"github.com/jamillosantos/go-ceous/generator/reporters"
	myasthurts "github.com/lab259/go-my-ast-hurts"
)

type (
	Field struct {
		Name             string
		FieldPath        string
		Column           string
		IsAutoIncrement  bool
		IsPrimaryKey     bool
		ForeignKeyColumn string
		Type             string
		RefType          myasthurts.RefType
		Fieldable        *Fieldable
	}

	Fieldable struct {
		Name              string
		TableName         string
		Connection        string
		IsModel           bool
		IsEmbedded        bool
		Fields            []*Field
		FieldsByColumnMap map[string]*Field
		PK                *Field
	}

	FieldableContext struct {
		InputPkg      *myasthurts.Package
		Imports       *CtxImports
		ModelsImports *CtxImports
		Reporter      reporters.Reporter

		Fieldables   []*Fieldable
		FieldableMap map[string]*Fieldable
	}
)

func NewFieldableContext(inputPkg, outputPkg *myasthurts.Package, reporter reporters.Reporter) *FieldableContext {
	return &FieldableContext{
		InputPkg:      inputPkg,
		ModelsImports: NewCtxImports(inputPkg),
		Imports:       NewCtxImports(outputPkg),
		Reporter:      reporter,
		Fieldables:    make([]*Fieldable, 0),
		FieldableMap:  make(map[string]*Fieldable),
	}
}

// NewFieldable returns a new instance of `Fieldable` with the given `name` set.
func NewFieldable(name string) *Fieldable {
	return &Fieldable{
		Name:              name,
		Fields:            make([]*Field, 0),
		FieldsByColumnMap: make(map[string]*Field),
	}
}

// NewField returns a new instance of `Field` with the given `name`, `fieldPath`,
// `column` set.
func NewField(name, fieldPath, column, foreignKeyColumn string, refType myasthurts.RefType) *Field {
	return &Field{
		Name:             name,
		FieldPath:        fieldPath,
		Column:           column,
		ForeignKeyColumn: foreignKeyColumn,
		RefType:          refType,
	}
}

// AddField appends the field to the list of fields and returns it.
func (f *Fieldable) AddField(field *Field) *Field {
	f.Fields = append(f.Fields, field)
	if field.IsPrimaryKey {
		f.PK = field // TODO(jota): This should support composite PKs.
	}
	if field.Column != "" {
		f.FieldsByColumnMap[field.Column] = field
	}
	return field
}

// EnsureFieldable tries to get the `Fieldable` from the map, if it does not
// exists adds creates one, adding to the list and map.
func (f *FieldableContext) EnsureFieldable(name string) *Fieldable {
	fieldable, ok := f.FieldableMap[name]
	if ok {
		return fieldable
	}
	fieldable = NewFieldable(name)
	f.Fieldables = append(f.Fieldables, fieldable)
	f.FieldableMap[fieldable.Name] = fieldable
	return fieldable
}
