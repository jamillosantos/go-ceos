package ceous_test

import (
	"fmt"

	"github.com/jamillosantos/go-ceous"
	"github.com/jamillosantos/go-ceous/tests"
)

func FieldAliasExample() {
	userAlias := tests.Schema.User.As("u")
	u := ceous.FieldAlias(userAlias)
	field := u(tests.Schema.User.Name)
	fmt.Println(field.Reference())
	// It will print "u.Name"
}
