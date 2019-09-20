package tests

import (
	"github.com/jamillosantos/go-ceous"
)

type userQuery struct {
	*ceous.BaseQuery
}

func NewUserQuery() *userQuery {
	return &userQuery{
		BaseQuery: ceous.NewBaseQuery(Schema.User.BaseSchema),
	}
}

func (q *userQuery) ByID() {
	// q.BaseQuery.Where()
}
