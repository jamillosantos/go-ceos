package parser

import (
	"errors"
	"fmt"
	"strings"

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
