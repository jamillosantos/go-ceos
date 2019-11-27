package parser

import (
	"errors"

	"github.com/jamillosantos/go-ceous/generator/models"
	myasthurts "github.com/lab259/go-my-ast-hurts"
)

func fieldPK() string {
	return "ceous.FieldPK"
}

func fieldAutoInc() string {
	return "ceous.FieldAutoIncrement"
}

// Skip is an error that is returned by the ParseField for ignoring a
// field.
var Skip = errors.New("field skipped")

func ParseField(ctx *models.FieldContext, field *myasthurts.Field) error {
	for _, t := range field.Tag.Params {
		switch t.Name {
		case "ceous":
			return parseField(ctx, &t, field)
		case "fk":
			return parseFK(ctx, &t, field)
		default:
			// If there is nothing is detected. No problem, just continue.
		}
	}
	// If nothing is found, just skip the field.
	return Skip
}

func parseField(ctx *models.FieldContext, t *myasthurts.TagParam, field *myasthurts.Field) error {
	if t.Value == "-" {
		return Skip
	}

	isStructE := false

	fieldColumnName := ""

	f := &models.ModelField{
		Name:      field.Name,
		FieldName: field.Name, // TODO(jota): Let the developer to choose its default naming convention...
		Type:      models.NewModelType(ctx.Gen, field.RefType),
		Modifiers: make([]models.ModelFieldModifier, 0), // TODO(jota): To check this..
	}

	ctx.Gen.Imports.AddRefType(field.RefType)

	if t.Value != "" {
		f.FieldName = t.Value
		fieldColumnName = t.Value
	}

	// To support multiple options...
	for _, o := range t.Options {
		switch o {
		case "pk":
			if ctx.Model.PK != nil {
				return errors.New("PK already defined") // TODO(jota): To formalize this error.
			}
			f.Modifiers = append(f.Modifiers, fieldPK)
			ctx.Model.PK = f
			ctx.Reporter.Linef(" * PK: %s", f.Name)
		case "autoincr":
			f.Modifiers = append(f.Modifiers, fieldAutoInc)
		}
	}

	if s, ok := field.RefType.Type().(*myasthurts.Struct); ok {
		isStructE = true
		cond := models.NewModelCondition(ctx.Gen, field.Name, field.RefType)
		for _, embeddedField := range s.Fields {
			ceousTag := embeddedField.Tag.TagParamByName("ceous")
			if ceousTag == nil || ceousTag.Value == "" {
				continue
			}
			columnName := ceousTag.Value
			if columnName == "" {
				columnName = embeddedField.Name // TODO(jota): Apply the default naming convention here.
			}
			if fieldColumnName != "" {
				columnName = fieldColumnName + "_" + columnName
			}
			column := &models.ModelColumn{
				Column:    columnName,
				Type:      models.NewModelType(ctx.Gen, embeddedField.RefType),
				FullField: field.Name + "." + embeddedField.Name,
				Modifiers: f.Modifiers,
			}
			ctx.Model.Columns = append(ctx.Model.Columns, column)
			ctx.Model.ColumnsMap[column.Column] = len(ctx.Model.Columns) - 1

			cond.Conditions = append(cond.Conditions, &models.ModelConditionCondition{
				FullField:   field.Name + "." + embeddedField.Name,
				Field:       embeddedField.Name,
				SchemaField: ctx.Model.Name + "." + field.Name + "." + embeddedField.Name,
			})
		}
		ctx.Model.Conditions = append(ctx.Model.Conditions, cond)
		f.SchemaType = field.RefType.Name()
	}

	ctx.Reporter.Linef(" + %s: %s", field.Name, field.RefType.Name())
	ctx.Model.Fields = append(ctx.Model.Fields, f)
	ctx.Model.SchemaFields = append(ctx.Model.SchemaFields, f)

	if isStructE { /// If is a embedded struct, it does not need to have the field appended. Instead, only the m.PK will be set.
		return nil
	}
	cond := models.NewModelCondition(ctx.Gen, f.Name, field.RefType)
	cond.Conditions = append(cond.Conditions, &models.ModelConditionCondition{
		SchemaField: ctx.Model.Name + "." + field.Name,
		FullField:   field.Name,
	})
	ctx.Model.Conditions = append(ctx.Model.Conditions, cond)

	column := &models.ModelColumn{
		Column:    f.FieldName,
		FullField: field.Name,
		Modifiers: f.Modifiers,
	}
	ctx.Model.Columns = append(ctx.Model.Columns, column)
	ctx.Model.ColumnsMap[f.FieldName] = len(ctx.Model.Columns) - 1

	return nil
}

func parseFK(ctx *models.FieldContext, t *myasthurts.TagParam, field *myasthurts.Field) error {
	s, ok := field.RefType.Type().(*myasthurts.Struct)
	if !ok {
		return Skip
	}
	model, _ := ctx.Gen.AddModel(s)
	relation := &models.ModelRelation{
		FromModel:      ctx.Model,
		FromField:      field.Name,
		FromColumnType: models.NewModelType(ctx.Gen, field.RefType),
		ToModel:        model,
		ToColumn:       t.Value,
	}
	ctx.Model.Relations = append(ctx.Model.Relations, relation)
	ctx.Gen.ModelsImports.AddRefType(field.RefType)
	ctx.Reporter.Linef("FK: %s(%s): %s", field.Name, relation.ToColumn, relation.FromColumnType.String())
	return nil
}
