package tests

import "github.com/jamillosantos/go-ceous"

type schemaUser struct {
	*ceous.BaseSchema
	ID   ceous.SchemaField
	Name ceous.SchemaField
}

type schema struct {
	User *schemaUser
}

var Schema = schema{
	User: &schemaUser{
		BaseSchema: ceous.NewBaseSchema(
			"users",
			"",
			ceous.NewSchemaField("id"),
			ceous.NewSchemaField("name"),
		),
		ID:   ceous.NewSchemaField("id"),
		Name: ceous.NewSchemaField("name"),
	},
}
