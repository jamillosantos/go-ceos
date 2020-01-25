package parser

import (
	"errors"
	"fmt"
	"strings"

	. "github.com/jamillosantos/go-ceous/generator/helpers"
	"github.com/jamillosantos/go-ceous/generator/models"
	"github.com/jamillosantos/go-ceous/generator/reporters"
	myasthurts "github.com/lab259/go-my-ast-hurts"
)

type parseFieldContext struct {
	Ctx           *models.FieldableContext
	Fieldable     *models.Fieldable
	Reporter      reporters.Reporter
	ModelsImports *models.CtxImports
	Imports       *models.CtxImports
}

func fieldPK() string {
	return "ceous.FieldPK"
}

func fieldAutoInc() string {
	return "ceous.FieldAutoIncrement"
}

// Skip is an error that is returned by the ParseField for ignoring a
// field.
var Skip = errors.New("field skipped")

func parseField(ctx *parseFieldContext, f *myasthurts.Field) (*models.Field, error) {
	if isRefTypeModel(f.RefType) {
		return nil, parseFieldModel(ctx, f)
	}
	tagCeous := f.Tag.TagParamByName("ceous")
	tagFK := f.Tag.TagParamByName("fk")
	return parseFieldCeous(ctx, tagCeous, tagFK, f)
}

func parseFieldModel(ctx *parseFieldContext, f *myasthurts.Field) error {
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

func parseFieldCeous(ctx *parseFieldContext, tagCeous *myasthurts.TagParam, tagFK *myasthurts.TagParam, f *myasthurts.Field) (*models.Field, error) {
	if tagCeous == nil && tagFK == nil {
		ctx.Reporter.Linef("Ignoring %s due to non ceous or fk tag found", f.Name)
		return nil, Skip
	}

	if tagCeous != nil && tagCeous.Value == "-" && tagFK == nil {
		ctx.Reporter.Linef("Ignoring %s due to ceous tag is '-'", f.Name)
		return nil, Skip
	}

	var column string
	if tagCeous != nil {
		column = tagCeous.Value
	}

	var foreignKeyColumn string
	if tagFK != nil {
		foreignKeyColumn = tagFK.Value
	}

	field := models.NewField(f.Name, f.Name, column, foreignKeyColumn, f.RefType)

	// If it is a model from the same package.
	// TODO(jota): Add this limitation to the README.
	// TODO(jota): Expand it to explore structs from other packages.
	fieldableStr := ""
	if s, ok := f.RefType.Type().(*myasthurts.Struct); ok && f.RefType.Pkg().RealPath == ctx.ModelsImports.Pkg.RealPath {
		field.Fieldable = ctx.Ctx.EnsureFieldable(s.Name())
		fieldableStr = "[*]"
	}
	if foreignKeyColumn != "" {
		fieldableStr += "[FK:" + field.Fieldable.TableName + "." + field.ForeignKeyColumn + "]"
	}
	optsReporter := []string{}
	if column != "" {
		optsReporter = AppendStringIfNotEmpty(optsReporter, column)
	}
	if tagCeous != nil {
		for _, opt := range tagCeous.Options {
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
	}

	var optsStr string
	if len(optsReporter) > 0 {
		optsStr = "(" + strings.Join(optsReporter, ",") + ")"
	}

	// ctx.ModelsImports.AddRefType(f.RefType)
	ctx.Imports.AddRefType(f.RefType)
	field.Type = ctx.ModelsImports.Ref(f.RefType)
	ctx.Reporter.Linef("+ %s%s: %s %s", field.Name, fieldableStr, field.Type, optsStr)
	return field, nil
}
