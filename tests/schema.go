package tests

import "github.com/jamillosantos/go-ceous"

type schemaUser struct {
	*ceous.BaseSchema
	ID        ceous.SchemaField
	Name      ceous.SchemaField
	Password  ceous.SchemaField
	Role      ceous.SchemaField
	CreatedAt ceous.SchemaField
	UpdatedAt ceous.SchemaField
}

type schema struct {
	User *schemaUser
}

var userBaseSchema = ceous.NewBaseSchema(
	"users",
	"",
	ceous.NewSchemaField("id"),
	ceous.NewSchemaField("name"),
	ceous.NewSchemaField("password"),
	ceous.NewSchemaField("role"),
	ceous.NewSchemaField("created_at"),
	ceous.NewSchemaField("updated_at"),
)

var Schema = schema{
	User: &schemaUser{
		BaseSchema: userBaseSchema,
		ID:         userBaseSchema.ColumnsArr[0],
		Name:       userBaseSchema.ColumnsArr[1],
		Password:   userBaseSchema.ColumnsArr[2],
		Role:       userBaseSchema.ColumnsArr[3],
		CreatedAt:  userBaseSchema.ColumnsArr[4],
		UpdatedAt:  userBaseSchema.ColumnsArr[5],
	},
}