package tests

import (
	"fmt"
	"time"

	"github.com/jamillosantos/go-ceous"
	"github.com/pkg/errors"
)

type (
	User struct {
		ceous.Model `tableName:"users"`
		ID          int       `ceous:"id,pk"`
		Name        string    `ceous:"name"`
		Password    string    `ceous:"password"`
		Role        string    `ceous:"role"`
		CreatedAt   time.Time `ceous:"created_at"`
		UpdatedAt   time.Time `ceous:"updated_at"`
	}

	Group struct {
		ceous.Model
		ID   int
		Name string
	}

	UserGroupPK struct {
		UserID  int
		GroupID int
	}

	UserGroup struct {
		ceous.Model
		ID    UserGroupPK
		Admin bool
	}
)

var (
	userGroupPKFields = []ceous.SchemaField{
		ceous.NewSchemaField("user_id"),
		ceous.NewSchemaField("group_id"),
	}
)

func (u *User) GetID() interface{} {
	return u.ID
}

func (g *Group) GetID() interface{} {
	return g.ID
}

func (ug *UserGroup) GetID() interface{} {
	return ug.ID
}

func (ugPK *UserGroupPK) ColumnAddress(name string) (interface{}, error) {
	switch name {
	case "user_id":
		return &ugPK.UserID, nil
	case "group_id":
		return &ugPK.GroupID, nil
	default:
		return nil, errors.Wrap(ceous.ErrFieldNotFound, fmt.Sprintf("field %s not found", name))
	}
}

func (ugPK *UserGroupPK) Value(name string) (interface{}, error) {
	switch name {
	case "user_id":
		return ugPK.UserID, nil
	case "group_id":
		return ugPK.GroupID, nil
	default:
		return nil, errors.Wrap(ceous.ErrFieldNotFound, fmt.Sprintf("field %s not found", name))
	}
}

func (ugPK *UserGroupPK) Columns() []ceous.SchemaField {
	return userGroupPKFields
}
