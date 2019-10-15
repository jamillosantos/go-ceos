package tests

import (
	"github.com/jamillosantos/go-ceous"
)

type (
	userQuery struct {
		*ceous.BaseQuery
	}

	userGroupQuery struct {
		*ceous.BaseQuery
	}
)

func NewUserQuery(options ...ceous.CeousOption) *userQuery {
	bq := ceous.NewBaseQuery(options...)
	if bq.Schema == nil {
		bq.Schema = Schema.User.BaseSchema
	}
	return &userQuery{
		BaseQuery: bq,
	}
}

func NewUserGroupQuery(options ...ceous.CeousOption) *userGroupQuery {
	bq := ceous.NewBaseQuery(options...)
	if bq.Schema == nil {
		bq.Schema = Schema.UserGroup.BaseSchema
	}
	return &userGroupQuery{
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
	q.Limit(1).Offset(0)

	query, err := q.RawQuery()
	if err != nil {
		return User{}, err
	}

	rs := NewUserResultSet(query, nil)
	defer rs.Close()

	if rs.Next() {
		err = rs.ToModel(&u)
	} else {
		err = ceous.ErrNotFound
	}
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

/**
 * UserGroupQuery
 */

func (q *userGroupQuery) ByID(id UserGroupPK) *userGroupQuery {
	q.BaseQuery.Where(ceous.Eq(Schema.UserGroup.ID.UserID, id.UserID))
	q.BaseQuery.Where(ceous.Eq(Schema.UserGroup.ID.GroupID, id.GroupID))
	return q
}

func (q *userGroupQuery) One() (u UserGroup, err error) {
	q.Limit(1).Offset(0)

	query, err := q.RawQuery()
	if err != nil {
		return UserGroup{}, err
	}

	rs := NewUserGroupResultSet(query, nil)
	defer rs.Close()

	if rs.Next() {
		err = rs.ToModel(&u)
	} else {
		err = ceous.ErrNotFound
	}
	return
}

func (q *userGroupQuery) Select(fields ...ceous.SchemaField) *userGroupQuery {
	q.BaseQuery.Select(fields...)
	return q
}

func (q *userGroupQuery) ExcludeFields(fields ...ceous.SchemaField) *userGroupQuery {
	q.BaseQuery.ExcludeFields(fields...)
	return q
}

func (q *userGroupQuery) Where(pred interface{}, args ...interface{}) *userGroupQuery {
	q.BaseQuery.Where(pred, args...)
	return q
}

func (q *userGroupQuery) Limit(limit uint64) *userGroupQuery {
	q.BaseQuery.Limit(limit)
	return q
}

func (q *userGroupQuery) Offset(offset uint64) *userGroupQuery {
	q.BaseQuery.Offset(offset)
	return q
}
