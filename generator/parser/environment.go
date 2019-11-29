package parser

import (
	"github.com/jamillosantos/go-ceous/generator/models"
	"github.com/jamillosantos/go-ceous/generator/reporters"
	myasthurts "github.com/lab259/go-my-ast-hurts"
)

type EnvironmentContext struct {
	Reporter      reporters.Reporter
	InputPkg      *myasthurts.Package
	OutputPkg     *myasthurts.Package
	Fieldables    []*models.Fieldable
	FieldableMap  map[string]*models.Fieldable
	Imports       *models.CtxImports
	ModelsImports *models.CtxImports
}

func ParseEnvironment(ctx *EnvironmentContext) (*models.Environment, error) {
	env := models.NewEnvironment(ctx.InputPkg, ctx.OutputPkg, ctx.Imports, ctx.ModelsImports)
	for _, model := range ctx.Fieldables {
		if model.IsModel {
			baseSchema := models.NewBaseSchema(model.Name, model.TableName)
			env.EnsureConnection(model.Connection)
			schema := models.NewSchema(model.Name, baseSchema)
			err := parseSchema(&parseSchemaContext{
				Env:          env,
				BaseSchema:   baseSchema,
				Schema:       schema,
				Reporter:     ctx.Reporter,
				FieldPath:    []string{},
				ColumnPrefix: []string{},
			}, model)
			if err != nil {
				return nil, err
			}
			env.AddSchema(schema)
			env.AddBaseSchema(baseSchema)

			query := models.NewQuery(model.Name)
			err = parseQuery(&parseQueryContext{
				Query:  query,
				Prefix: []string{},
			}, model)
			if err != nil {
				return nil, err
			}
			env.AddQuery(query)

			store := models.NewStore(model.Name)
			err = parseStore(&parseStoreContext{
				Store: store,
			}, model)
			if err != nil {
				return nil, err
			}
			env.AddStore(store)
		}
	}

	ctx.Reporter.Line("Connections")
	for _, conn := range env.Connections {
		ctx.Reporter.Linef("+ %s", conn.Name)
	}

	ctx.Reporter.Line("Queries")
	for _, conn := range env.Queries {
		ctx.Reporter.Linef("+ %s", conn.FullName)
	}

	return env, nil
}
