package cmd

import (
	"fmt"

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
		for _, pkgName := range args {
			_, err := env.Parse(pkgName)
			if err != nil {
				panic(err) // TODO(jota): Decide how critical errors will be reported.
			}
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

		for _, pkgName := range args {
			models := make([]*generatorModels.Model, 0)

			pkg, ok := env.PackageByImportPath(pkgName)
			if !ok {
				panic(pkgName + " not found") // TODO(jota): Decide how critical errors will be reported.
			}

			for _, s := range pkg.Structs {
				for _, f := range s.Fields {
					if f.RefType == ceousModel {
						models = append(models, generatorModels.NewModel(s))
					}
				}
			}

			fmt.Println("#", pkg.Name)
			fmt.Println("")
			fmt.Println(tpl.Schema(pkg, models))
		}

	},
}

func init() {
	rootCmd.AddCommand(genCmd)

	genCmd.PersistentFlags().StringArrayVarP(&excludedFiles, "recursive", "r", []string{}, "exclude files")
	genCmd.PersistentFlags().StringArrayVarP(&excludedFiles, "exclude-files", "e", []string{}, "exclude files")
	//
}
