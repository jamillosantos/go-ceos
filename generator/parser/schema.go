package parser

import (
	"strings"

	"github.com/jamillosantos/go-ceous/generator/helpers"
	"github.com/jamillosantos/go-ceous/generator/models"
	"github.com/jamillosantos/go-ceous/generator/reporters"
	myasthurts "github.com/lab259/go-my-ast-hurts"
)

type ()

// columnPrefix serializes the prefix array into a string by concating the
// strings dividing the fields by _.
func columnPrefix(prefix []string) string {
	r := ""
	for _, s := range prefix {
		r = r + s + "_"
	}
	return r
}

// namePrefix serializes the prefix array into a string by concating the
// strings pascal casing all names.
func namePrefix(prefix []string) string {
	r := ""
	for _, s := range prefix {
		r = r + helpers.PascalCase(s)
	}
	return r
}

// ParseSchema returns a new instance of `Schema` based on the given `s`.
//
// The function registers the schema in the `models.GenContext`.
func ParseSchema(ctx *models.SchemaContext, s *myasthurts.Struct) (*models.Schema, error) {
	schema := models.NewSchema(namePrefix(ctx.Prefix) + s.Name())
	ctx.Reporter.Linef("Schema for %s", strings.Join(append(ctx.Prefix, s.Name()), "."))
	for _, structField := range s.Fields {
		_, err := parseSchemaField(&models.SchemaFieldContext{
			Gen:      ctx.Gen,
			Schema:   schema,
			Prefix:   append(ctx.Prefix, s.Name()),
			Reporter: reporters.SubReporter(ctx.Reporter),
		}, structField)
		if err == Skip {
			continue
		}
		if err != nil {
			return nil, err
		}
	}
	return ctx.Gen.AddSchema(schema), nil
}

// parseSchemaField returns a new instance of the `SchemaField` based on the
// given `field`.
//
// If the field type is a struct, it creates a new `Schema` with a prefix with
// the field that is being parsed.
func parseSchemaField(ctx *models.SchemaFieldContext, field *myasthurts.Field) (*models.SchemaField, error) {
	tag := field.Tag.TagParamByName("ceous")
	if tag == nil || tag.Value == "-" {
		return nil, Skip
	}

	columnNameOrig := tag.Value
	if columnNameOrig == "" {
		columnNameOrig = field.Name
	}

	columnName := columnPrefix(ctx.Prefix) + columnNameOrig

	// If the field we did encounter is a struct, we have to use it as a subschema.
	if s, ok := field.RefType.Type().(*myasthurts.Struct); ok {
		schema, err := ParseSchema(&models.SchemaContext{
			Gen:      ctx.Gen,
			Reporter: reporters.SubReporter(ctx.Reporter),
			Prefix:   append(ctx.Prefix, field.Name),
		}, s)
		if err != nil {
			return nil, err
		}
		field := ctx.Schema.AddField(field.Name, "")
		field.Type = schema.Name
		return field, nil
	}

	return ctx.Schema.AddField(field.Name, columnName), nil
}
