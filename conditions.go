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
func Not(cond Condition) Condition {
	return func(schema Schema) Sqlizer {
		return &notOperator{cond(schema)}
	}
}

// SqlCondition will return a new condition that will create a sqlCondition.
//
// See more at sqlCondition
func SqlCondition(sql string, args []interface{}) Condition {
	return func(Schema) Sqlizer {
		return &sqlCondition{sql, args}
	}
}
