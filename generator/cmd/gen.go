package cmd

import (
	"os"

	generatorModels "github.com/jamillosantos/go-ceous/generator/models"
	"github.com/jamillosantos/go-ceous/generator/tpl"
	myasthurts "github.com/lab259/go-my-ast-hurts"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var (
	recursive     bool
	excludedFiles []string
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

		// Models will be a list of the structs that implement Model

		models := make([]*generatorModels.Model, 0)
		embeddeds := make([]*generatorModels.Model, 0)

		for _, s := range pkg.Structs {
			for _, f := range s.Fields {
				if isModel(f.RefType) {
					model, err := generatorModels.NewModel(s)
					if err != nil {
						panic(errors.Wrapf(err, "error parsing model %s", s.Name())) // TODO(jota): Decide how critical errors will be reported.
					}
					models = append(models, model)
				} else if isEmbedded(f.RefType) {
					model, err := generatorModels.NewModel(s)
					if err != nil {
						panic(errors.Wrapf(err, "error parsing embedded %s", s.Name())) // TODO(jota): Decide how critical errors will be reported.
					}
					embeddeds = append(embeddeds, model)
				}
			}
		}

		f, err := os.OpenFile("ceous.go", os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0755)
		if err != nil {
			panic(errors.Wrapf(err, "error generating ceous.go")) // TODO(jota): Decide how critical errors will be reported.
		}
		defer f.Close()

		tpl.RenderSchema(f, pkg, models, embeddeds)
	},
}

func init() {
	rootCmd.AddCommand(genCmd)

	genCmd.PersistentFlags().StringArrayVarP(&excludedFiles, "recursive", "r", []string{}, "exclude files")
	genCmd.PersistentFlags().StringArrayVarP(&excludedFiles, "exclude-files", "e", []string{}, "exclude files")
	//
}
