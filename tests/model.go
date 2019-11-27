package tests

import (
	"strconv"
	"time"

	"github.com/jamillosantos/go-ceous"
)

type (
	User struct {
		ceous.Model `tableName:"users"`
		ID          int       `ceous:"id,pk,autoincr"`
		Name        string    `ceous:"name"`
		Password    string    `ceous:"password"`
		Role        string    `ceous:"role"`
		Address     Address   `ceous:""`
		Work        Address   `ceous:"work"`
		CreatedAt   time.Time `ceous:"created_at"`
		UpdatedAt   time.Time `ceous:"updated_at"`
	}

	Group struct {
		ceous.Model `tableName:"groups"`
		ID          int    `ceous:"id,pk,autoincr"`
		Name        string `ceous:"name"`
	}

	UserGroupPK struct {
		ceous.Embedded
		UserID  int `ceous:"user_id"`
		GroupID int `ceous:"group_id"`
	}

	UserGroup struct {
		ceous.Model `tableName:"user_groups"`
		ID          UserGroupPK `ceous:",pk"`
		user        *User       `fk:"user_id"`
		Admin       bool        `ceous:"admin"`
	}

	Address struct {
		Street string `ceous:"street"`
		Number string `ceous:"number"`
		City   string `ceous:"city"`
		State  string `ceous:"state"`
	}
)

var (
	userGroupPKFields = []ceous.SchemaField{
		ceous.NewSchemaField("user_id"),
		ceous.NewSchemaField("group_id"),
	}
)

func (ugPK *UserGroupPK) Columns() []ceous.SchemaField {
	return userGroupPKFields
}

func (ugPK *UserGroupPK) String() string {
	return strconv.Itoa(ugPK.UserID) + ":" + strconv.Itoa(ugPK.GroupID)
}
