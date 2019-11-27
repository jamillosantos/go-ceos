package parser

import (
	"github.com/jamillosantos/go-ceous/generator/models"
	"github.com/jamillosantos/go-ceous/generator/reporters"
	myasthurts "github.com/lab259/go-my-ast-hurts"
	"github.com/pkg/errors"
)

func isModel(refType myasthurts.RefType) bool {
	return refType.Pkg().Name == "ceous" && refType.Name() == "Model"
}

func isEmbedded(refType myasthurts.RefType) bool {
	return refType.Pkg().Name == "ceous" && refType.Name() == "Embedded"
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
		if err == Skip {
			reporter.Line("Skipping")
			continue
		} else if err != nil {
			return errors.Wrapf(err, "error parsing model %s", s.Name()) // TODO(jota): Decide how critical errors will be reported.
		}
		_, err = ParseSchema(&models.SchemaContext{
			Gen:      ctx,
			Reporter: reporter,
			Prefix:   []string{},
		}, s)
		if err != nil {
			return errors.Wrapf(err, "error parsing schema %s", s.Name()) // TODO(jota): Decide how critical errors will be reported.
		}

		if _, ok := connectionsMap[model.Connection]; !ok {
			conn := models.NewConnection(model.Connection)
			connectionsMap[model.Connection] = conn
			ctx.Connections = append(ctx.Connections, conn)
		}
	}

	return nil
}
