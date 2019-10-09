package tests

import (
	"fmt"
	"time"

	"github.com/jamillosantos/go-ceous"
	"github.com/pkg/errors"
)

type User struct {
	ceous.Model
	ID        int
	Name      string
	Password  string
	Role      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

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
