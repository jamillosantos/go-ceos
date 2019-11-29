package parser

import "github.com/jamillosantos/go-ceous/generator/models"

type (
	parseStoreContext struct {
		Store *models.Store
	}
)

func parseStore(ctx *parseStoreContext, model *models.Fieldable) error {
	ctx.Store.ModelRef = model.Name
	return nil
}
