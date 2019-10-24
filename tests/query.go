package tests

import (
	"github.com/jamillosantos/go-ceous"
)

type userQuery struct {
	*ceous.BaseQuery
}

func NewUserQuery(options ...ceous.QueryOption) *userQuery {
	bq := ceous.NewBaseQuery(options...)
	if bq.Schema == nil {
		bq.Schema = Schema.User.BaseSchema
	}
	return &userQuery{
		BaseQuery: bq,
	}
}

func (q *userQuery) ByID(id int) *userQuery {
	q.BaseQuery.Where(ceous.Eq(Schema.User.ID, id))
	return q
}

func (q *userQuery) ByName(name string) *userQuery {
	q.BaseQuery.Where(ceous.Eq(Schema.User.Name, name))
	return q
}

func (q *userQuery) One() (u User, err error) {
	q.Limit(1).Offset(1)
	err = NewUserResultSet(q.RawQuery()).ToModel(&u)
	return
}

func (q *userQuery) Select(fields ...ceous.SchemaField) *userQuery {
	q.BaseQuery.Select(fields...)
	return q
}

func (q *userQuery) ExcludeFields(fields ...ceous.SchemaField) *userQuery {
	q.BaseQuery.ExcludeFields(fields...)
	return q
}

func (q *userQuery) Where(pred interface{}, args ...interface{}) *userQuery {
	q.BaseQuery.Where(pred, args...)
	return q
}

func (q *userQuery) Limit(limit uint64) *userQuery {
	q.BaseQuery.Limit(limit)
	return q
}

func (q *userQuery) Offset(offset uint64) *userQuery {
	q.BaseQuery.Offset(offset)
	return q
}
