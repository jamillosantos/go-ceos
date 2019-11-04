package cmd

import (
	"bytes"
	"io"
	"os"

	generatorModels "github.com/jamillosantos/go-ceous/generator/models"
	"github.com/jamillosantos/go-ceous/generator/reporters"
	"github.com/jamillosantos/go-ceous/generator/tpl"
	myasthurts "github.com/lab259/go-my-ast-hurts"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var (
	recursive     bool
	excludedFiles []string
	verbose       bool
)

func isModel(refType myasthurts.RefType) bool {
	return refType.Pkg().Name == "ceous" && refType.Name() == "Model"
}

func isEmbedded(refType myasthurts.RefType) bool {
	return refType.Pkg().Name == "ceous" && refType.Name() == "Embedded"
}

// genCmd represents the gen command
var genCmd = &cobra.Command{
	Use:   "gen",
	Short: "the generation tool for ceous",
	// TODO(jota): Describe the generation tool.
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args: cobra.ArbitraryArgs,
	Run: func(cmd *cobra.Command, args []string) {
		env, err := myasthurts.NewEnvironment()
		if err != nil {
			panic(err) // TODO(jota): Decide how critical errors will be reported.
		}

		pkg, err := env.Parse(".")
		if err != nil {
			panic(errors.Wrap(err, "could not parse the package")) // TODO(jota): Decide how critical errors will be reported.
		}

		var reporter reporters.Reporter

		if verbose {
			reporter = &reporters.Verbose{}
		} else {
			reporter = &reporters.Quiet{}
		}

		// Models will be a list of the structs that implement Model

		models := make([]*generatorModels.Model, 0)
		embeddeds := make([]*generatorModels.Model, 0)
		connections := make([]*generatorModels.Connection, 0)
		connectionsMap := make(map[string]*generatorModels.Connection, 0)

		ctx := generatorModels.NewCtx(reporter, pkg, env.BuiltIn)

		var (
			model *generatorModels.Model
		)

		for _, s := range pkg.Structs {
			reporter.Line("Analysing", s.Name())
			for _, f := range s.Fields {
				if f.RefType.Pkg() == nil {
					continue
				}
				if isModel(f.RefType) {
					model, err = generatorModels.ParseModel(ctx, s)
					if err != nil {
						panic(errors.Wrapf(err, "error parsing model %s", s.Name())) // TODO(jota): Decide how critical errors will be reported.
					}
					models = append(models, model)
				} else if isEmbedded(f.RefType) {
					model, err = generatorModels.ParseModel(ctx, s)
					if err != nil {
						panic(errors.Wrapf(err, "error parsing embedded %s", s.Name())) // TODO(jota): Decide how critical errors will be reported.
					}
					embeddeds = append(embeddeds, model)
				} else {
					continue
				}

				if _, ok := connectionsMap[model.Connection]; !ok {
					conn := generatorModels.NewConnection(model.Connection)
					connectionsMap[model.Connection] = conn
					connections = append(connections, conn)
				}
			}
		}

		buff := bytes.NewBuffer(nil)

		reporter.Line("Generating code ...")
		tpl.RenderCeous(buff, ctx, models, embeddeds, connections)

		f, err := os.OpenFile("ceous.go", os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0755)
		if err != nil {
			panic(errors.Wrapf(err, "error opening ceous.go")) // TODO(jota): Decide how critical errors will be reported.
		}
		defer f.Close()

		reporter.Line("Generating file ...")
		_, err = io.Copy(f, buff)
		if err != nil {
			panic(errors.Wrapf(err, "error copying buffer to ceous.go")) // TODO(jota): Decide how critical errors will be reported.
		}
		reporter.Line("Done ...")
	},
}

func init() {
	rootCmd.AddCommand(genCmd)

	genCmd.PersistentFlags().StringArrayVarP(&excludedFiles, "exclude-files", "e", []string{}, "exclude files")
	genCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "enable verbose mode")
	//
}
