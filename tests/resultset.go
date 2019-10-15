package tests

import "github.com/jamillosantos/go-ceous"

type userResultSet struct {
	*ceous.RecordResultSet
	lastErr error
}

func NewUserResultSet(rs ceous.ResultSet, err error) *userResultSet {
	return &userResultSet{
		RecordResultSet: ceous.NewRecordResultSet(rs, err),
	}
}

type userGroupResultSet struct {
	*ceous.RecordResultSet
	lastErr error
}

func NewUserGroupResultSet(rs ceous.ResultSet, err error) *userGroupResultSet {
	return &userGroupResultSet{
		RecordResultSet: ceous.NewRecordResultSet(rs, err),
	}
}
