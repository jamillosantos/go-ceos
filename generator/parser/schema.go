package parser

import (
	"strings"

	. "github.com/jamillosantos/go-ceous/generator/helpers"
	"github.com/jamillosantos/go-ceous/generator/models"
	"github.com/jamillosantos/go-ceous/generator/reporters"
)

type (
	SchemaFieldModifier func() string

	parseSchemaContext struct {
		Env            *models.Environment
		Reporter       reporters.Reporter
		BaseSchema     *models.BaseSchema
		Schema         *models.Schema
		ColumnPrefix   []string
		FieldPath      []string
		FieldModifiers []SchemaFieldModifier
	}

	parseSchemaFieldContext struct {
		Env            *models.Environment
		Reporter       reporters.Reporter
		BaseSchema     *models.BaseSchema
		Schema         *models.Schema
		FieldPath      []string
		ColumnPrefix   []string
		FieldModifiers []SchemaFieldModifier
	}
)

// ParseSchema returns a new instance of `Schema` based on the given `s`.
//
// The function registers the schema in the `models.GenContext`.
func parseSchema(ctx *parseSchemaContext, model *models.Fieldable) error {
	ctx.Reporter.Linef("Schema for %s", strings.Join(append(ctx.FieldPath, model.Name), "."))
	for _, modelField := range model.Fields {
		err := parseSchemaField(&parseSchemaFieldContext{
			Env:            ctx.Env,
			Reporter:       reporters.SubReporter(ctx.Reporter),
			Schema:         ctx.Schema,
			BaseSchema:     ctx.BaseSchema,
			ColumnPrefix:   ctx.ColumnPrefix,
			FieldPath:      ctx.FieldPath,
			FieldModifiers: ctx.FieldModifiers,
		}, modelField)
		if err == Skip {
			continue
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// parseSchemaField returns a new instance of the `SchemaField` based on the
// given `field`.
//
// If the field type is a struct, it creates a new `Schema` with a prefix with
// the field that is being parsed.
func parseSchemaField(ctx *parseSchemaFieldContext, field *models.Field) error {
	columnName := columnPrefix(ctx.ColumnPrefix) + field.Column

	// If the field is a model type, it has to be parsed as a subschema.
	if field.Fieldable != nil {
		if field.ForeignKeyColumn == "" {
			schema := models.NewSchema(namePrefix(append(ctx.FieldPath, ctx.Schema.Name, field.Name)), ctx.BaseSchema)
			ctx.Env.AddSchema(schema)
			err := parseSchema(&parseSchemaContext{
				Env:          ctx.Env,
				BaseSchema:   ctx.BaseSchema,
				Schema:       schema,
				Reporter:     ctx.Reporter,
				ColumnPrefix: AppendStringIfNotEmpty(ctx.ColumnPrefix, field.Column),
				FieldPath:    append(ctx.FieldPath, field.Name),
			}, field.Fieldable)
			ctx.Schema.AddField(field.Name, field.Type, schema.FullName, columnName)
			if err != nil {
				return err
			}
			return nil
		}
		return Skip
	}
	ctx.BaseSchema.AddField(field.Name, columnName)
	ctx.Schema.AddField(field.Name, "", "", columnName)

	return nil
}
