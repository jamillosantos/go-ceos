package parser

import (
	"fmt"

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
		if err == Skip {
			continue
		}
		if err != nil {
			return err
		}
		ctx.Query.AddField(qField)
	}
	return nil
}

func parseQueryField(ctx *parseQueryFieldContext, field *models.Field) (*models.QueryField, error) {
	if field.ForeignKeyColumn != "" {
		localField, ok := ctx.Query.BaseSchema.FieldsMap[field.ForeignKeyColumn]
		if !ok {
			return nil, fmt.Errorf("local field %s not found", field.ForeignKeyColumn) // TODO(jota): To formalize this error.
		}

		relation := models.NewRelation(
			ctx.Query,
			helpers.PascalCase(field.Name),
			field.Name,
			localField.Name,
			ctx.Model.Name,
			field.ForeignKeyColumn,
			field.Fieldable.Name,
			field.Fieldable.PK.Name,
			field.Fieldable.PK.Column,
			field.Fieldable.PK.Type,
		)
		ctx.Query.AddRelation(relation)
		ctx.Model.AddRelation(relation)
		return nil, Skip
	} else if field.Fieldable != nil {
		for _, f := range field.Fieldable.Fields {
			qField, err := parseQueryField(&parseQueryFieldContext{
				Query:       ctx.Query,
				Model:       ctx.Model,
				FieldPrefix: append(ctx.FieldPrefix, field.Name),
			}, f)
			if err == Skip {
				continue
			}
			if err != nil {
				return nil, err
			}
			ctx.Query.AddField(qField)
		}
		return nil, Skip
	}
	qField := models.NewQueryField(field.Name, append(ctx.FieldPrefix, field.Name))
	qField.Type = field.Type
	return qField, nil
}
