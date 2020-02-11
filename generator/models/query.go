package models

import (
	"strings"

	"github.com/jamillosantos/go-ceous/generator/naming"
)

type (
	Query struct {
		Name       string
		FullName   string
		BaseSchema *BaseSchema
		Fields     []*QueryField
		Relations  []*Relation
	}

	QueryField struct {
		Name       string
		FieldPath  string
		MethodName string
		Type       string
	}
)

// NewQuery returns a new instance of `Query` with the given `name` set.
func NewQuery(baseSchema *BaseSchema, name string) *Query {
	return &Query{
		BaseSchema: baseSchema,
		Name:       name,
		FullName:   naming.PascalCase.Do(name) + "Query",
		Fields:     make([]*QueryField, 0),
		Relations:  make([]*Relation, 0),
	}
}

// NewQueryField returns a new instance of `QueryField` with the given params
// set.
func NewQueryField(name string, fieldPath []string) *QueryField {
	return &QueryField{
		Name:       name,
		FieldPath:  strings.Join(fieldPath, "."),
		MethodName: strings.Join(fieldPath, ""),
	}
}

// AddField appends the given `field` to the fields list.
func (q *Query) AddField(field *QueryField) *QueryField {
	q.Fields = append(q.Fields, field)
	return field
}

// AddRelation appends the given `relation` to the relation list.
func (q *Query) AddRelation(relation *Relation) *Relation {
	q.Relations = append(q.Relations, relation)
	return relation
}
