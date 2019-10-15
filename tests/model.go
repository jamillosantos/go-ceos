package tests

import (
	"fmt"
	"time"

	"github.com/jamillosantos/go-ceous"
	"github.com/pkg/errors"
)

type (
	User struct {
		ceous.Model
		ID        int
		Name      string
		Password  string
		Role      string
		CreatedAt time.Time
		UpdatedAt time.Time
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

func (u *User) ColumnAddress(name string) (interface{}, error) {
	switch name {
	case "id":
		return &u.ID, nil
	case "name":
		return &u.Name, nil
	case "password":
		return &u.Password, nil
	case "role":
		return &u.Role, nil
	case "created_at":
		return &u.CreatedAt, nil
	case "updated_at":
		return &u.UpdatedAt, nil
	default:
		return nil, errors.Wrap(ceous.ErrFieldNotFound, fmt.Sprintf("field %s not found", name))
	}
}

func (u *User) Value(name string) (interface{}, error) {
	switch name {
	case "id":
		return u.ID, nil
	case "name":
		return u.Name, nil
	case "password":
		return u.Password, nil
	case "role":
		return u.Role, nil
	case "created_at":
		return u.CreatedAt, nil
	case "updated_at":
		return u.UpdatedAt, nil
	default:
		return nil, errors.Wrap(ceous.ErrFieldNotFound, fmt.Sprintf("field %s not found", name))
	}
}

func (g *Group) GetID() interface{} {
	return g.ID
}

func (g *Group) ColumnAddress(name string) (interface{}, error) {
	switch name {
	case "id":
		return &g.ID, nil
	case "name":
		return &g.Name, nil
	default:
		return nil, errors.Wrap(ceous.ErrFieldNotFound, fmt.Sprintf("field %s not found", name))
	}
}

func (g *Group) Value(name string) (interface{}, error) {
	switch name {
	case "id":
		return g.ID, nil
	case "name":
		return g.Name, nil
	default:
		return nil, errors.Wrap(ceous.ErrFieldNotFound, fmt.Sprintf("field %s not found", name))
	}
}

func (ug *UserGroup) GetID() interface{} {
	return ug.ID
}

func (ug *UserGroup) ColumnAddress(name string) (interface{}, error) {
	switch name {
	case "user_id":
		return &ug.ID.UserID, nil
	case "group_id":
		return &ug.ID.GroupID, nil
	case "admin":
		return &ug.Admin, nil
	default:
		return nil, errors.Wrap(ceous.ErrFieldNotFound, fmt.Sprintf("field %s not found", name))
	}
}

func (ug *UserGroup) Value(name string) (interface{}, error) {
	switch name {
	case "user_id":
		return ug.ID.UserID, nil
	case "group_id":
		return ug.ID.GroupID, nil
	case "admin":
		return ug.Admin, nil
	default:
		return nil, errors.Wrap(ceous.ErrFieldNotFound, fmt.Sprintf("field %s not found", name))
	}
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
