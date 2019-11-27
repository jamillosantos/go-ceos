package parser

import (
	"github.com/jamillosantos/go-ceous/generator/models"
	"github.com/jamillosantos/go-ceous/generator/reporters"
	myasthurts "github.com/lab259/go-my-ast-hurts"
)

func ParseModel(ctx *models.ModelContext, s *myasthurts.Struct) (*models.Model, error) {
	var modelInfo *myasthurts.Field
	for _, field := range s.Fields {
		if isModel(field.RefType) {
			modelInfo = field
		}
	}
	if modelInfo == nil {
		return nil, Skip
	}
	m, ok := ctx.Gen.AddModel(s)
	if ok {
		ctx.Reporter.Line("model found among the previous analyzed, skipping.")
		return m, nil
	}

	ctx.Reporter.Linef("Model for %s", s.Name())

	reporter := reporters.WithPrefix(ctx.Reporter, "    ")

	ctx.Model = m

	// Process table name
	tableNameTag := modelInfo.Tag.TagParamByName("tableName")
	if tableNameTag != nil {
		m.TableName = tableNameTag.Value
	}
	reporter.Line("TableName: ", m.TableName)

	// Find the connection name
	connectionTag := modelInfo.Tag.TagParamByName("conn")
	if connectionTag != nil {
		m.Connection = connectionTag.Value
	}
	reporter.Line("Connection: ", m.Connection)

	for _, field := range s.Fields {
		if field == modelInfo {
			continue
		}
		if field.Name == "" {
			continue
		}
		err := ParseField(&models.FieldContext{
			Schema:   ctx.Schema,
			Gen:      ctx.Gen,
			Model:    m,
			Reporter: reporters.WithPrefix(reporter, "    "),
		}, field)
		if err == Skip {
			continue
		} else if err != nil {
			return nil, err
		}
	}
	return m, nil
}
