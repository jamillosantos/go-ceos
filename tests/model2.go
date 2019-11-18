package tests

import (
	"time"

	"github.com/jamillosantos/go-ceous"
)

type (
	UserIgnored struct {
		ceous.Model `tableName:"users"`
		ID          int       `ceous:"id,pk,autoincr"`
		Name        string    `ceous:"name"`
		Password    string    `ceous:"password"`
		Role        string    `ceous:"role"`
		CreatedAt   time.Time `ceous:"created_at"`
		UpdatedAt   time.Time `ceous:"updated_at"`
	}
)
