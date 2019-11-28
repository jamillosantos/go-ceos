package parser

import (
	"strings"

	"github.com/jamillosantos/go-ceous/generator/models"
	"github.com/jamillosantos/go-ceous/generator/reporters"
	myasthurts "github.com/lab259/go-my-ast-hurts"
)

// ParseSchema returns a new instance of `Schema` based on the given `s`.
//
// The function registers the schema in the `models.GenContext`.
func ParseSchema(ctx *models.SchemaContext, s *myasthurts.Struct) error {
	ctx.Reporter.Linef("Schema for %s", strings.Join(append(ctx.FieldPath, s.Name()), "."))
	for _, structField := range s.Fields {
		_, err := parseSchemaField(&models.SchemaFieldContext{
			Gen:            ctx.Gen,
			Reporter:       reporters.SubReporter(ctx.Reporter),
			Model:          ctx.Model,
			Schema:         ctx.Schema,
			ColumnPrefix:   ctx.ColumnPrefix,
			FieldPath:      ctx.FieldPath,
			FieldModifiers: ctx.FieldModifiers,
		}, structField)
		if err == Skip {
			continue
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// parseSchemaField returns a new instance of the `SchemaField` based on the
// given `field`.
//
// If the field type is a struct, it creates a new `Schema` with a prefix with
// the field that is being parsed.
func parseSchemaField(ctx *models.SchemaFieldContext, field *myasthurts.Field) (*models.SchemaField, error) {
	cond := models.NewModelCondition(ctx.Gen, append(ctx.FieldPath, field.Name), field.RefType)

	tag := field.Tag.TagParamByName("ceous")
	if tag == nil || tag.Value == "-" {
		return nil, Skip
	}

	columnNameOrig := tag.Value
	if columnNameOrig == "" {
		columnNameOrig = field.Name
	}

	columnName := columnPrefix(ctx.ColumnPrefix) + columnNameOrig

	// If the field we did encounter is a struct, we have to use it as a subschema.
	if s, ok := field.RefType.Type().(*myasthurts.Struct); ok {
		err := ParseSchema(&models.SchemaContext{
			Gen:          ctx.Gen,
			BaseSchema:   ctx.BaseSchema,
			Schema:       ctx.Schema,
			Reporter:     ctx.Reporter,
			ColumnPrefix: append(ctx.ColumnPrefix, columnNameOrig),
			FieldPath:    append(ctx.FieldPath, field.Name),
		}, s)
		if err != nil {
			return nil, err
		}
		return nil, Skip
	}

	column := &models.ModelColumn{
		Column:    columnName,
		Type:      models.NewModelType(ctx.Gen, field.RefType),
		FullField: strings.Join(append(ctx.FieldPath, field.Name), "."),
		Modifiers: ctx.FieldModifiers,
	}
	ctx.Model.Columns = append(ctx.Model.Columns, column)
	ctx.Model.ColumnsMap[column.Column] = len(ctx.Model.Columns) - 1

	cond.Conditions = append(cond.Conditions, &models.ModelConditionCondition{
		FullField:   strings.Join(append(ctx.FieldPath, field.Name), "."),
		Field:       field.Name,
		SchemaField: ctx.Model.Name + "." + strings.Join(append(ctx.FieldPath, field.Name), "."),
	})

	ctx.Model.Conditions = append(ctx.Model.Conditions, cond)

	ctx.Reporter.Linef("+ %s: %s", field.Name, field.RefType.Name())

	ctx.BaseSchema.AddField(field.Name, columnName)

	return ctx.Schema.AddField(field.Name, "", columnName), nil
}
