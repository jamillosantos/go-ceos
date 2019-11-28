package parser

import (
	"errors"
	"fmt"
	"strings"

	"github.com/jamillosantos/go-ceous/generator/models"
	"github.com/jamillosantos/go-ceous/generator/reporters"
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
			// return parseFK(ctx, &t, field)
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

	f := &models.ModelField{
		Name:      field.Name,
		FieldName: field.Name, // TODO(jota): Let the developer to choose its default naming convention...
		Type:      models.NewModelType(ctx.Gen, field.RefType),
		Modifiers: make([]models.ModelFieldModifier, 0), // TODO(jota): To check this..
	}

	ctx.Gen.Imports.AddRefType(field.RefType)

	if t.Value != "" {
		f.FieldName = t.Value
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
			ctx.Reporter.Linef("* PK: %s", f.Name)
		case "autoincr":
			f.Modifiers = append(f.Modifiers, fieldAutoInc)
		}
	}

	ctx.Reporter.Linef("+ %s: %s", field.Name, field.RefType.Name())

	if s, ok := field.RefType.Type().(*myasthurts.Struct); ok {
		isStructE = true

		schema := models.NewSchema(strings.Join(append(ctx.Prefix, ctx.Model.Name, field.Name, s.Name()), ""), ctx.BaseSchema)
		ctx.Reporter.Linef("Schema for %s", schema.Name)
		for _, embeddedField := range s.Fields {
			_, err := parseSchemaField(&models.SchemaFieldContext{
				Gen:            ctx.Gen,
				Reporter:       reporters.SubReporter(ctx.Reporter),
				Model:          ctx.Model,
				BaseSchema:     ctx.BaseSchema,
				Schema:         schema,
				FieldModifiers: f.Modifiers,
				ColumnPrefix:   []string{f.FieldName},
				FieldPath:      []string{field.Name},
			}, embeddedField)
			if err != nil {
				return err
			}
		}
		ctx.Gen.AddSchema(schema)
		f.SchemaType = schema.Name
		ctx.Schema.AddField(field.Name, "schema"+schema.Name, "") // TODO(jota): Remove "schema" from here and add in the schema.Name initialization.
	}

	ctx.Model.Fields = append(ctx.Model.Fields, f)
	ctx.Model.SchemaFields = append(ctx.Model.SchemaFields, f)

	if isStructE { /// If is a embedded struct, it does not need to have the field appended. Instead, only the m.PK will be set.
		return nil
	}
	cond := models.NewModelCondition(ctx.Gen, append(ctx.Prefix, field.Name), field.RefType)
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

	ctx.Schema.AddField(field.Name, "", f.FieldName)
	ctx.BaseSchema.AddField(field.Name, f.FieldName)

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

func parseField2(ctx *models.Field2Context, f *myasthurts.Field) (*models.Field, error) {
	if isRefTypeModel(f.RefType) {
		return nil, parseFieldModel(ctx, f)
	}
	if tag := f.Tag.TagParamByName("ceous"); tag != nil && tag.Value != "-" {
		return parseFieldCeous(ctx, tag, f)
	}
	return nil, Skip
}

func parseFieldModel(ctx *models.Field2Context, f *myasthurts.Field) error {
	tableName := f.Tag.TagParamByName("tableName")
	if tableName != nil {
		ctx.Fieldable.TableName = tableName.Value
	} else {
		ctx.Fieldable.TableName = ctx.Fieldable.Name // TODO(jota): Apply naming convention here.
	}
	ctx.Reporter.Linef("Table name: %s", ctx.Fieldable.TableName)
	conn := f.Tag.TagParamByName("conn")
	if conn != nil {
		ctx.Fieldable.Connection = conn.Value
	} else {
		ctx.Fieldable.Connection = "default"
	}
	ctx.Reporter.Linef("Connection: %s", ctx.Fieldable.Connection)
	return Skip
}

func parseFieldCeous(ctx *models.Field2Context, tag *myasthurts.TagParam, f *myasthurts.Field) (*models.Field, error) {
	column := tag.Value
	if column == "" {
		column = f.Name
	}

	field := &models.Field{
		Name:   f.Name,
		Column: column,
	}

	// If it is a model from the same package.
	// TODO(jota): Add this limitation to the README.
	// TODO(jota): Expand it to explore structs from other packages.
	fieldableStr := ""
	if s, ok := f.RefType.Type().(*myasthurts.Struct); ok && f.RefType.Pkg().RealPath == ctx.ModelsImports.Pkg.RealPath {
		field.Fieldable = ctx.Ctx.EnsureFieldable(s.Name())
		fieldableStr = "[*]"
	}

	optsReporter := []string{column}
	for _, opt := range tag.Options {
		switch opt {
		case "autoincr":
			field.IsAutoIncrement = true
			optsReporter = append(optsReporter, "auto increment")
		case "pk":
			field.IsPrimaryKey = true
			optsReporter = append(optsReporter, "primary key")
		default:
			return nil, fmt.Errorf("unknown tag %s", opt)
		}
	}

	var optsStr string
	if len(optsReporter) > 0 {
		optsStr = "(" + strings.Join(optsReporter, ",") + ")"
	}

	field.Type = ctx.ModelsImports.Ref(f.RefType)
	ctx.Reporter.Linef("+ %s%s: %s %s", field.Name, fieldableStr, field.Type, optsStr)
	return field, nil
}
