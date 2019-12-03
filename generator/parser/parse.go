package parser

import (
	"github.com/jamillosantos/go-ceous/generator/models"
	"github.com/jamillosantos/go-ceous/generator/reporters"
	myasthurts "github.com/lab259/go-my-ast-hurts"
)

func isRefTypeModel(refType myasthurts.RefType) bool {
	return refType.Pkg().Name == "ceous" && refType.Name() == "Model"
}

func isEmbedded(refType myasthurts.RefType) bool {
	return refType.Pkg().Name == "ceous" && refType.Name() == "Embedded"
}

func isStructModel(s *myasthurts.Struct) bool {
	for _, field := range s.Fields {
		if isRefTypeModel(field.RefType) {
			return true
		}
	}
	return false
}

func isStructEmbedded(s *myasthurts.Struct) bool {
	for _, field := range s.Fields {
		if isEmbedded(field.RefType) {
			return true
		}
	}
	return false
}

func Parse(ctx *models.FieldableContext) error {
	for _, s := range ctx.InputPkg.Structs {
		fieldable := ctx.EnsureFieldable(s.Name())
		ctx.Reporter.Linef("Model %s", fieldable.Name)
		fieldable.IsModel = isStructModel(s)
		fieldable.IsEmbedded = isStructEmbedded(s)
		for _, f := range s.Fields {
			field, err := parseField(&parseFieldContext{
				Ctx:           ctx,
				Fieldable:     fieldable,
				Reporter:      reporters.SubReporter(ctx.Reporter),
				ModelsImports: ctx.ModelsImports,
				Imports:       ctx.Imports,
			}, f)
			if err == Skip {
				continue
			}
			if err != nil {
				return err
			}
			fieldable.AddField(field)
		}
	}
	return nil
}
