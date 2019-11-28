package parser

import (
	"errors"

	"github.com/jamillosantos/go-ceous/generator/models"
	"github.com/jamillosantos/go-ceous/generator/reporters"
	myasthurts "github.com/lab259/go-my-ast-hurts"
)

var (
	ErrTableNameNotDefined = errors.New("tablename was not defined")
)

func ParseModel(ctx *models.ModelContext, s *myasthurts.Struct) (*models.Model, error) {
	var modelInfo *myasthurts.Field
	for _, field := range s.Fields {
		if isRefTypeModel(field.RefType) {
			modelInfo = field
		}
	}

	if modelInfo == nil { // If no modelInfo found, just skip this struct.
		return nil, Skip
	}

	m, ok := ctx.Gen.AddModel(s)
	if ok {
		ctx.Reporter.Line("model found among the previous analyzed, skipping.")
		return m, nil
	}

	ctx.Reporter.Linef("Model for %s", s.Name())

	reporter := reporters.SubReporter(ctx.Reporter)

	ctx.Model = m

	// Process table name
	// TODO(jota): If it is empty, add a way to "calculate" the table name by the struct name.
	tableNameTag := modelInfo.Tag.TagParamByName("tableName")
	if tableNameTag != nil {
		m.TableName = tableNameTag.Value
	}
	reporter.Line("TableName: ", m.TableName)
	if m.TableName == "" {
		return nil, ErrTableNameNotDefined
	}

	baseSchema := models.NewBaseSchema(s.Name(), m.TableName)
	ctx.Gen.AddBaseSchema(baseSchema)

	schema := models.NewSchema(s.Name(), baseSchema)
	schema.IsModel = true
	ctx.Gen.AddSchema(schema)

	// Find the connection name
	connectionTag := modelInfo.Tag.TagParamByName("conn")
	if connectionTag != nil {
		m.Connection = connectionTag.Value
	}
	reporter.Line("Connection: ", m.Connection)

	for _, field := range s.Fields {
		if field == modelInfo { // If the field is the model info, just ignore it.
			continue
		}
		if field.Name == "" {
			continue
		}
		err := ParseField(&models.FieldContext{
			BaseSchema: baseSchema,
			Schema:     schema,
			Gen:        ctx.Gen,
			Model:      m,
			Reporter:   reporter,
		}, field)
		if err == Skip {
			reporter.Linef("skipping field %s", field.Name)
			continue
		} else if err != nil {
			return nil, err
		}
	}
	return m, nil
}
