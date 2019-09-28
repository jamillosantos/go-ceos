package tests

import "fmt"

type User struct {
	ID   int
	Name string
}

func (u *User) ColumnAddress(name string) (interface{}, error) {
	switch name {
	case "id":
		return &u.ID, nil
	case "name":
		return &u.Name, nil
	default:
		return nil, fmt.Errorf("field %s not found", name)
	}
}
