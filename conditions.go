package ceous

import (
	sq "github.com/elgris/sqrl"
)

type Sqlizer interface {
	sq.Sqlizer
}

type Condition func(Schema) Sqlizer

func Eq(field SchemaField, value interface{}) Condition {
	return func(Schema) Sqlizer {
		return &eqOperator{field, value}
	}
}

func Ne(field SchemaField, value interface{}) Condition {
	return func(Schema) Sqlizer {
		return &neOperator{field, value}
	}
}
