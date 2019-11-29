package parser

import "github.com/jamillosantos/go-ceous/generator/models"

type (
	parseQueryContext struct {
		Query        *models.Query
		Prefix       []string
		ColumnPrefix []string
	}

	parseQueryFieldContext struct {
		Query       *models.Query
		FieldPrefix []string
	}
)

func parseQuery(ctx *parseQueryContext, model *models.Fieldable) error {
	for _, field := range model.Fields {
		qField, err := parseQueryField(&parseQueryFieldContext{
			Query:       ctx.Query,
			FieldPrefix: append(ctx.Prefix, field.Name),
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
		ctx.Query.AddRelation(&models.Relation{
			Name:             field.Name + "Relation",
			LocalModelType:   field.Type,
			LocalColumn:      field.Column,
			ForeignModelType: field.Fieldable.Name,
			ForeignColumn:    field.Fieldable.PK.Column,
		})
	}
	return qField, nil
}
