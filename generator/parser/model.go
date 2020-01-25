package parser

import "github.com/jamillosantos/go-ceous/generator/models"

type (
	parseModelContext struct{}

	parseMOdelFieldContext struct {
		Model *models.Model
	}
)

func parseModel(ctx *parseModelContext, fieldable *models.Fieldable) (*models.Model, error) {
	model := models.NewModel(fieldable.Name)
	for _, f := range fieldable.Fields {
		field, err := parseModelField(&parseMOdelFieldContext{
			Model: model,
		}, f)
		if err != nil {
			return nil, err
		}
		model.AddField(field)
		if f.IsPrimaryKey { // TODO(jota): Support for multiple PKs...
			model.PK = field
		}
	}
	return model, nil
}

func parseModelField(ctx *parseMOdelFieldContext, f *models.Field) (*models.ModelField, error) {
	field := models.NewModelField(f.Name)
	return field, nil
}
