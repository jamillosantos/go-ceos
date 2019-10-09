package tests

import "github.com/jamillosantos/go-ceous"

type CompositePK struct {
	UserID int
	PostID int
}

func (pk *CompositePK) ColumnAddress(column string) (interface{}, error) {
	switch column {
	case "user_id":
		return &pk.UserID, nil
	case "post_id":
		return &pk.PostID, nil
	}
	return nil, ceous.ErrFieldNotFound
}

func (pk *CompositePK) Value(column string) (interface{}, error) {
	switch column {
	case "user_id":
		return pk.UserID, nil
	case "post_id":
		return pk.PostID, nil
	}
	return nil, ceous.ErrFieldNotFound
}
