package ceous_test

import "fmt"

type UserTestModel struct {
	ID   int
	Name string
}

func (u *UserTestModel) ColumnAddress(name string) (interface{}, error) {
	switch name {
	case "id":
		return &u.ID, nil
	case "name":
		return &u.Name, nil
	default:
		return nil, fmt.Errorf("field %s not found", name)
	}
}
