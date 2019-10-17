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

		for _, rel := range q.BaseQuery.Relations {
			err = rel.Aggregate(&u)
			if err != nil {
				return User{}, err // TODO(jota): Shall this error be wrapped? At first, yes.
			}
		}
	} else {
		err = ceous.ErrNotFound
	}
	if err == nil {
		for _, rel := range q.BaseQuery.Relations {
			err = rel.Realize()
			if err != nil {
				return User{}, err // TODO(jota): Shall this error be wrapped? At first, yes.
			}
		}
	}
	return
}

func (q *userQuery) All() (records []*User, err error) {
	query, err := q.RawQuery()
	if err != nil {
		return nil, err
	}

	rs := NewUserGroupResultSet(query, nil)
	defer rs.Close()

	records = make([]*User, 0)
	for rs.Next() {
		user := &User{}
		err = rs.ToModel(user)
		if err != nil {
			return nil, err // TODO(jota): Shall this error be wrapped? At first, yes.
		}

		for _, rel := range q.BaseQuery.Relations {
			err = rel.Aggregate(user)
			if err != nil {
				return nil, err // TODO(jota): Shall this error be wrapped? At first, yes.
			}
		}

		records = append(records, user)
	}

	for _, rel := range q.BaseQuery.Relations {
		err = rel.Realize()
		if err != nil {
			return nil, err // TODO(jota): Shall this error be wrapped? At first, yes.
		}
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

func (q *userGroupQuery) WithUser() *userGroupQuery {
	q.BaseQuery.Relations = append(q.BaseQuery.Relations, NewUserGroupModelUserRelation())
	return q
}

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

		for _, rel := range q.BaseQuery.Relations {
			err = rel.Aggregate(&u)
			if err != nil {
				return UserGroup{}, err // TODO(jota): Shall this error be wrapped? At first, yes.
			}
		}
	} else {
		err = ceous.ErrNotFound
	}
	if err == nil {
		for _, rel := range q.BaseQuery.Relations {
			err = rel.Realize()
			if err != nil {
				return UserGroup{}, err // TODO(jota): Shall this error be wrapped? At first, yes.
			}
		}
	}
	return
}

func (q *userGroupQuery) All() (records []*UserGroup, err error) {
	query, err := q.RawQuery()
	if err != nil {
		return nil, err
	}

	rs := NewUserGroupResultSet(query, nil)
	defer rs.Close()

	records = make([]*UserGroup, 0)
	for rs.Next() {
		record := &UserGroup{}
		err = rs.ToModel(record)
		if err != nil {
			return nil, err // TODO(jota): Shall this error be wrapped? At first, yes.
		}

		for _, rel := range q.BaseQuery.Relations {
			err = rel.Aggregate(record)
			if err != nil {
				return nil, err // TODO(jota): Shall this error be wrapped? At first, yes.
			}
		}
		records = append(records, record)

	}

	for _, rel := range q.BaseQuery.Relations {
		err = rel.Realize()
		if err != nil {
			return nil, err // TODO(jota): Shall this error be wrapped? At first, yes.
		}
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

func (q *userGroupQuery) OrderBy(fields ...interface{}) *userGroupQuery {
	q.BaseQuery.OrderBy(fields...)
	return q
}
