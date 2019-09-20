package ceous

import (
	"database/sql/driver"
	"fmt"

	sq "github.com/elgris/sqrl"
)

// eqOperator implements the equalty operation in SQL. The preference would be
// to use the already implemented `sqrl.Eq`, however its implementation is not
// very effective.
type eqOperator [2]interface{}

const (
	equalOprStr    = "="
	inOprStr       = "IN"
	nullOprStr     = "IS"
	inEmptyExprStr = "(1=0)" // Portable FALSE

	notEqualOprStr    = "<>"
	notInOprStr       = "NOT IN"
	notNullOprStr     = "IS NOT"
	notInEmptyExprStr = "(1=1)" // Portable TRUE
)

func equalOpr(useNot bool) string {
	if useNot {
		return notEqualOprStr
	}
	return equalOprStr
}

func inOpr(useNot bool) string {
	if useNot {
		return notInOprStr
	}
	return inOprStr
}

func nullOpr(useNot bool) string {
	if useNot {
		return notNullOprStr
	}
	return nullOprStr
}

func inEmptyExpr(useNot bool) string {
	if useNot {
		return notInEmptyExprStr
	}
	return inEmptyExprStr
}

func (operator *eqOperator) toSql(useNot bool) (sql string, args []interface{}, err error) {
	key := operator[0]
	val := operator[1]

	switch v := val.(type) {
	case driver.Valuer:
		if val, err = v.Value(); err != nil {
			return
		}
	}

	if val == nil {
		sql = fmt.Sprintf("%s %s NULL", key, nullOpr(useNot))
		return
	}
	switch list := val.(type) {
	case []interface{}:
		if len(list) == 0 {
			sql = inEmptyExpr(useNot)
		} else {
			args = list
			sql = fmt.Sprintf("%s %s (%s)", key, inOpr(useNot), sq.Placeholders(len(list)))
		}
		return
	}
	sql = fmt.Sprintf("%s %s ?", key, equalOpr(useNot))
	args = []interface{}{val}
	return
}

func (operator *eqOperator) ToSql() (string, []interface{}, error) {
	return operator.toSql(false)
}

type neOperator eqOperator

func (operator *neOperator) ToSql() (string, []interface{}, error) {
	return (*eqOperator)(operator).toSql(true)
}
