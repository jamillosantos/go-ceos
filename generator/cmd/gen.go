package cmd

import (
	"os"

	generatorModels "github.com/jamillosantos/go-ceous/generator/models"
	"github.com/jamillosantos/go-ceous/generator/tpl"
	myasthurts "github.com/lab259/go-my-ast-hurts"
	"github.com/spf13/cobra"
)

var (
	recursive     bool
	excludedFiles []string
)

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
			panic(err) // TODO(jota): Decide how critical errors will be reported.
		}

		ceousPkg, ok := env.PackageByImportPath("github.com/jamillosantos/go-ceous")
		if !ok {
			panic("ceous package not found") // TODO(jota): To formalize this error
		}

		ceousModel, ok := ceousPkg.RefTypeByName("Model")
		if !ok {
			panic("ceous.Model definition not found") // TODO(jota): To formalize this error
		}

		// Models will be a list of the structs that implement Model

		models := make([]*generatorModels.Model, 0)

		for _, s := range pkg.Structs {
			for _, f := range s.Fields {
				if f.RefType == ceousModel {
					models = append(models, generatorModels.NewModel(s))
				}
			}
		}

		f, err := os.OpenFile("ceous.go", os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0755)
		if err != nil {
			panic(err) // TODO(jota): To formalize this error
		}
		defer f.Close()

		tpl.RenderSchema(f, pkg, models)
		if err != nil {
			panic(err) // TODO(jota): To formalize this error
		}
	},
}

func init() {
	rootCmd.AddCommand(genCmd)

	genCmd.PersistentFlags().StringArrayVarP(&excludedFiles, "recursive", "r", []string{}, "exclude files")
	genCmd.PersistentFlags().StringArrayVarP(&excludedFiles, "exclude-files", "e", []string{}, "exclude files")
	//
}
