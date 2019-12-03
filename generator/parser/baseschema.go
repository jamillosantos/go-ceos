package parser

import "github.com/jamillosantos/go-ceous/generator/models"

type (
	parseBaseSchemaContext struct {
		BaseSchema   *models.BaseSchema
		FieldPrefix  []string
		ColumnPrefix []string
	}

	parseBaseSchemaFieldContext struct {
		BaseSchema   *models.BaseSchema
		FieldPrefix  []string
		ColumnPrefix []string
	}
)

func parseBaseSchema(ctx *parseBaseSchemaContext, model *models.Fieldable) error {
	for _, field := range model.Fields {
		err := parseBaseSchemaField(&parseBaseSchemaFieldContext{
			BaseSchema:  ctx.BaseSchema,
			FieldPrefix: ctx.FieldPrefix,
		}, field)
		if err != nil {
			return err
		}
	}
	return nil
}

func parseBaseSchemaField(ctx *parseBaseSchemaFieldContext, field *models.Field) error {
	if field.Fieldable != nil {
		return parseBaseSchema(&parseBaseSchemaContext{
			BaseSchema:   ctx.BaseSchema,
			FieldPrefix:  append(ctx.FieldPrefix, field.Name),
			ColumnPrefix: append(ctx.ColumnPrefix, field.Column),
		}, field.Fieldable)
	}
	ctx.BaseSchema.AddField(
		memberAccess(append(ctx.FieldPrefix, field.Name)...),
		columnPrefix(append(ctx.ColumnPrefix, field.Column)),
		field.IsAutoIncrement,
	)
	return nil
}
