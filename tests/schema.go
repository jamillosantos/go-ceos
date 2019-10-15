package tests

import (
	sq "github.com/elgris/sqrl"
	"github.com/jamillosantos/go-ceous"
)

type schemaUser struct {
	*ceous.BaseSchema
	ID        ceous.SchemaField
	Name      ceous.SchemaField
	Password  ceous.SchemaField
	Role      ceous.SchemaField
	CreatedAt ceous.SchemaField
	UpdatedAt ceous.SchemaField
}

type schemaUserGroupPK struct {
	UserID  ceous.SchemaField
	GroupID ceous.SchemaField
}

type schemaUserGroup struct {
	*ceous.BaseSchema
	ID    *schemaUserGroupPK
	Admin ceous.SchemaField
}

func init() {
	// TODO(jota): This should be gone in the when finishing the library.
	sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
}

type schema struct {
	User      *schemaUser
	UserGroup *schemaUserGroup
}

func (schema *schemaUser) PrimaryKey() ceous.SchemaField {
	return schema.ID
}

var (
	userBaseSchema = ceous.NewBaseSchema(
		"users",
		"",
		ceous.NewSchemaField("id", ceous.FieldPK, ceous.FieldAutoIncrement),
		ceous.NewSchemaField("name"),
		ceous.NewSchemaField("password"),
		ceous.NewSchemaField("role"),
		ceous.NewSchemaField("created_at"),
		ceous.NewSchemaField("updated_at"),
	)

	userGroupBaseSchema = ceous.NewBaseSchema(
		"user_groups",
		"",
		ceous.NewSchemaField("user_id", ceous.FieldPK),
		ceous.NewSchemaField("group_id", ceous.FieldPK),
		ceous.NewSchemaField("admin"),
	)

	userGroupPKSchema = &schemaUserGroupPK{
		UserID:  userGroupBaseSchema.ColumnsArr[0],
		GroupID: userGroupBaseSchema.ColumnsArr[1],
	}

	Schema = schema{
		User: &schemaUser{
			BaseSchema: userBaseSchema,
			ID:         userBaseSchema.ColumnsArr[0],
			Name:       userBaseSchema.ColumnsArr[1],
			Password:   userBaseSchema.ColumnsArr[2],
			Role:       userBaseSchema.ColumnsArr[3],
			CreatedAt:  userBaseSchema.ColumnsArr[4],
			UpdatedAt:  userBaseSchema.ColumnsArr[5],
		},
		UserGroup: &schemaUserGroup{
			BaseSchema: userGroupBaseSchema,
			ID:         userGroupPKSchema,
			Admin:      userGroupBaseSchema.ColumnsArr[2],
		},
	}
)
