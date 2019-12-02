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
	ctx.Imports.AddImportPkg(ctx.InputPkg)
	env := models.NewEnvironment(ctx.InputPkg, ctx.OutputPkg, ctx.Imports, ctx.ModelsImports)
	for _, model := range ctx.Fieldables {
		if model.IsModel {
			baseSchema := models.NewBaseSchema(model.Name, model.TableName)
			env.EnsureConnection(model.Connection)
			schema := models.NewSchema(model.Name, baseSchema)
			schema.IsModel = true
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
	subReporter := reporters.SubReporter(ctx.Reporter)
	for _, conn := range env.Connections {
		subReporter.Linef("+ %s", conn.Name)
	}
	ctx.Reporter.Line()
	ctx.Reporter.Line("BaseSchema")
	for _, baseSchema := range env.BaseSchemas {
		subReporter.Linef("+ %s", baseSchema.FullName)
		subSubReporter := reporters.SubReporter(subReporter)
		for _, field := range baseSchema.Fields {
			subSubReporter.Linef("- %s (%s)", field.Name, field.ColumnName)
		}
	}
	ctx.Reporter.Line()
	ctx.Reporter.Line("Schema")
	for _, schema := range env.Schemas {
		subReporter.Linef("+ %s", schema.FullName)
		subSubReporter := reporters.SubReporter(subReporter)
		for _, field := range schema.Fields {
			if field.Type == "" {
				subSubReporter.Linef("- %s", field.Name)
			} else {
				subSubReporter.Linef("- %s: %s(%s)", field.Name, field.Type, field.SchemaName)
			}
		}

	}
	ctx.Reporter.Line()
	ctx.Reporter.Line("Queries")
	for _, query := range env.Queries {
		subReporter.Linef("+ %s", query.FullName)
	}
	ctx.Reporter.Line()
	ctx.Reporter.Line("Stores")
	for _, store := range env.Stores {
		subReporter.Linef("+ %s", store.FullName)
	}

	return env, nil
}