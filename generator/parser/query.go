package parser

import (
	"github.com/jamillosantos/go-ceous/generator/helpers"
	"github.com/jamillosantos/go-ceous/generator/models"
)

type (
	parseQueryContext struct {
		Query        *models.Query
		Model        *models.Model
		Prefix       []string
		ColumnPrefix []string
	}

	parseQueryFieldContext struct {
		Query       *models.Query
		Model       *models.Model
		FieldPrefix []string
	}
)

func parseQuery(ctx *parseQueryContext, model *models.Fieldable) error {
	for _, field := range model.Fields {
		qField, err := parseQueryField(&parseQueryFieldContext{
			Model:       ctx.Model,
			Query:       ctx.Query,
			FieldPrefix: ctx.Prefix,
		}, field)
		if err != nil {
			return err
		}
		ctx.Query.AddField(qField)
	}
	return nil
}

func parseQueryField(ctx *parseQueryFieldContext, field *models.Field) (*models.QueryField, error) {
	qField := models.NewQueryField(field.Name, append(ctx.FieldPrefix))
	qField.FieldPath = append(ctx.FieldPrefix, field.Name)
	qField.Type = field.Type
	if field.ForeignKeyColumn != "" {
		relation := models.NewRelation(
			ctx.Query,
			helpers.PascalCase(field.Name),
			memberAccess(qField.FieldPath...),
			field.Type,
			field.Column,
			field.Fieldable.Name,
			field.Fieldable.PK.Name,
			field.Fieldable.PK.Column,
			field.Fieldable.PK.Type,
		)
		ctx.Query.AddRelation(relation)
		ctx.Model.AddRelation(relation)
	}
	return qField, nil
}
