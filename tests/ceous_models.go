package tests

import (
	"github.com/jamillosantos/go-ceous"
	"github.com/pkg/errors"
)

// GetID returns the primary key for model `User`.
func (model *User) GetID() interface{} {
	return model.ID
}

// ColumnAddress returns the pointer address of a field given its column name.
func (model *User) ColumnAddress(name string) (interface{}, error) {
	switch name {
	case "id":
		return &model.ID, nil
	case "name":
		return &model.Name, nil
	case "password":
		return &model.Password, nil
	case "role":
		return &model.Role, nil
	case "street":
		return &model.Address.Street, nil
	case "number":
		return &model.Address.Number, nil
	case "city":
		return &model.Address.City, nil
	case "state":
		return &model.Address.State, nil
	case "work_street":
		return &model.Work.Street, nil
	case "work_number":
		return &model.Work.Number, nil
	case "work_city":
		return &model.Work.City, nil
	case "work_state":
		return &model.Work.State, nil
	case "created_at":
		return &model.CreatedAt, nil
	case "updated_at":
		return &model.UpdatedAt, nil
	default:
		return nil, errors.Wrapf(ceous.ErrFieldNotFound, "field %s not found", name)
	}
}

// Value returns the value from a field given its column name.
func (model *User) Value(name string) (interface{}, error) {
	switch name {
	case "id":
		return model.ID, nil
	case "name":
		return model.Name, nil
	case "password":
		return model.Password, nil
	case "role":
		return model.Role, nil
	case "street":
		return model.Address.Street, nil
	case "number":
		return model.Address.Number, nil
	case "city":
		return model.Address.City, nil
	case "state":
		return model.Address.State, nil
	case "work_street":
		return model.Work.Street, nil
	case "work_number":
		return model.Work.Number, nil
	case "work_city":
		return model.Work.City, nil
	case "work_state":
		return model.Work.State, nil
	case "created_at":
		return model.CreatedAt, nil
	case "updated_at":
		return model.UpdatedAt, nil
	default:
		return nil, errors.Wrapf(ceous.ErrFieldNotFound, "field %s not found", name)
	}
}

// GetID returns the primary key for model `Group`.
func (model *Group) GetID() interface{} {
	return model.ID
}

// ColumnAddress returns the pointer address of a field given its column name.
func (model *Group) ColumnAddress(name string) (interface{}, error) {
	switch name {
	case "id":
		return &model.ID, nil
	case "name":
		return &model.Name, nil
	default:
		return nil, errors.Wrapf(ceous.ErrFieldNotFound, "field %s not found", name)
	}
}

// Value returns the value from a field given its column name.
func (model *Group) Value(name string) (interface{}, error) {
	switch name {
	case "id":
		return model.ID, nil
	case "name":
		return model.Name, nil
	default:
		return nil, errors.Wrapf(ceous.ErrFieldNotFound, "field %s not found", name)
	}
}

// GetID returns the primary key for model `UserGroup`.
func (model *UserGroup) GetID() interface{} {
	return model.ID
}

// User returns the user from UserGroup.
func (model *UserGroup) User() *User {
	return model.user
}

// SetUser updates the value for the user,
// updating the user.
func (model *UserGroup) SetUser(value *User) error {
	c, err := model.ColumnAddress("user_id")
	if err != nil {
		return err
	}

	v, ok := c.(*int)
	if !ok {
		return errors.New("invalid key value") // TODO(jota): To formalize this error.
	}
	*v = value.ID
	model.user = value
	return nil
}

// ColumnAddress returns the pointer address of a field given its column name.
func (model *UserGroup) ColumnAddress(name string) (interface{}, error) {
	switch name {
	case "user_id":
		return &model.ID.UserID, nil
	case "group_id":
		return &model.ID.GroupID, nil
	case "admin":
		return &model.Admin, nil
	default:
		return nil, errors.Wrapf(ceous.ErrFieldNotFound, "field %s not found", name)
	}
}

// Value returns the value from a field given its column name.
func (model *UserGroup) Value(name string) (interface{}, error) {
	switch name {
	case "user_id":
		return model.ID.UserID, nil
	case "group_id":
		return model.ID.GroupID, nil
	case "admin":
		return model.Admin, nil
	default:
		return nil, errors.Wrapf(ceous.ErrFieldNotFound, "field %s not found", name)
	}
}
