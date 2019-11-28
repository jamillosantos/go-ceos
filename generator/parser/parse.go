package parser

import (
	"github.com/jamillosantos/go-ceous/generator/models"
	"github.com/jamillosantos/go-ceous/generator/reporters"
	myasthurts "github.com/lab259/go-my-ast-hurts"
	"github.com/pkg/errors"
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

func Parse(ctx *models.GenContext) error {
	connectionsMap := make(map[string]*models.Connection, 0)

	var (
		model *models.Model
		err   error
	)

	for _, s := range ctx.InputPkg.Structs {
		ctx.Reporter.Line("Analysing ", s.Name())
		reporter := reporters.SubReporter(ctx.Reporter)
		model, err = ParseModel(&models.ModelContext{
			Gen:      ctx,
			Reporter: reporter,
		}, s)
		if err != Skip && err != nil {
			return errors.Wrapf(err, "error parsing model %s", s.Name()) // TODO(jota): Decide how critical errors will be reported.
		}

		if _, ok := connectionsMap[model.Connection]; !ok {
			conn := models.NewConnection(model.Connection)
			connectionsMap[model.Connection] = conn
			ctx.Connections = append(ctx.Connections, conn)
		}
	}

	return nil
}

func Parse2(ctx *models.FieldableContext) error {
	for _, s := range ctx.InputPkg.Structs {
		fieldable := ctx.EnsureFieldable(s.Name())
		ctx.Reporter.Linef("Model %s", fieldable.Name)
		fieldable.IsModel = isStructModel(s)
		fieldable.IsEmbedded = isStructEmbedded(s)
		for _, f := range s.Fields {
			fieldContext := models.NewField2Context(ctx, fieldable, reporters.SubReporter(ctx.Reporter), ctx.ModelsImports)
			field, err := parseField2(fieldContext, f)
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
