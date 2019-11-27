package cmd

import (
	"bytes"
	"go/format"
	"os"
	"path"

	generatorModels "github.com/jamillosantos/go-ceous/generator/models"
	"github.com/jamillosantos/go-ceous/generator/parser"
	"github.com/jamillosantos/go-ceous/generator/reporters"
	"github.com/jamillosantos/go-ceous/generator/tpl"
	myasthurts "github.com/lab259/go-my-ast-hurts"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var (
	recursive               bool
	excludedFiles           []string
	outputModelsFileName    string
	outputGeneratedFileName string
	inputPackage            string
	outputPackage           string
	verbose                 bool
	reporter                reporters.Reporter
)

type fileIgnorer struct{}

// BeforeFile will ignore all files defined by the command line.
func (*fileIgnorer) BeforeFile(_ *myasthurts.ParsePackageContext, filePath string) error {
	cwd, err := os.Getwd()
	if err != nil {
		return errors.Wrap(err, "could not get the working directory")
	}

	for _, f := range excludedFiles {
		if !path.IsAbs(f) {
			f = path.Clean(path.Join(cwd, f))
		}
		if f == filePath || filePath == outputGeneratedFileName || filePath == outputModelsFileName {
			if verbose {
				reporter.Linef("Ignoring file %s", filePath)
			}
			return myasthurts.Skip
		}
	}
	return nil
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
		if verbose {
			reporter = &reporters.Verbose{}
		} else {
			reporter = &reporters.Quiet{}
		}

		env, err := myasthurts.NewEnvironmentWithListener(&fileIgnorer{})
		if err != nil {
			panic(err) // TODO(jota): Decide how critical errors will be reported.
		}

		inputPkg, err := env.Parse(inputPackage)
		if err != nil {
			panic(errors.Wrapf(err, "could not parse the input package %s", inputPackage)) // TODO(jota): Decide how critical errors will be reported.
		}

		reporter.Linef("Input package: %s (%s)", inputPkg.Name, inputPkg.ImportPath)

		var outputPkg *myasthurts.Package
		if outputPackage == "" || outputPackage == inputPackage {
			outputPkg = inputPkg
		} else {
			outputPkg, err = env.Parse(outputPackage)
			if err != nil {
				panic(errors.Wrap(err, "could not parse the output package"))
			}
		}

		reporter.Linef("Output package: %s (%s)", outputPkg.Name, outputPkg.ImportPath)

		// Models will be a list of the structs that implement Model

		genCtx := generatorModels.NewGenContext(reporter, inputPkg, outputPkg, env.BuiltIn)

		err = parser.Parse(genCtx)
		if err != nil {
			panic(errors.Wrap(err, "could not parse information"))
		}

		buffCeous := bytes.NewBuffer(nil)
		buffModels := bytes.NewBuffer(nil)

		reporter.Line("Generating code ...")
		tpl.RenderCeous(buffCeous, genCtx)
		tpl.RenderModels(buffModels, genCtx)

		reporter.Line("Formatting code ...")

		formattedCeousCode, err := format.Source(buffCeous.Bytes())
		if err != nil {
			/*
				TODO(jota): In case of error, report the excerpt of the code.
				if serr, ok := err.(scanner.ErrorList); ok {
					for _, serri := range serr {
						// reporter.Line(reflect.TypeOf(serri))
					}
				}
				reporter.Line(buffCeous.String())
			*/
			panic(errors.Wrapf(err, "could not format the ceous code"))
		}
		formattedModelsCode, err := format.Source(buffModels.Bytes())
		if err != nil {
			panic(errors.Wrapf(err, "could not format the models code"))
		}

		reporter.Line("Creating files ...")

		ceousFile, err := os.OpenFile(path.Join(outputPkg.RealPath, outputGeneratedFileName), os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0755)
		if err != nil {
			panic(errors.Wrapf(err, "error opening %s", outputGeneratedFileName)) // TODO(jota): Decide how critical errors will be reported.
		}
		defer ceousFile.Close()

		modelsFile, err := os.OpenFile(path.Join(inputPkg.RealPath, outputModelsFileName), os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0755)
		if err != nil {
			panic(errors.Wrapf(err, "error opening %s", outputModelsFileName)) // TODO(jota): Decide how critical errors will be reported.
		}
		defer modelsFile.Close()

		_, err = ceousFile.Write(formattedCeousCode)
		if err != nil {
			panic(errors.Wrapf(err, "error writing %s", outputGeneratedFileName)) // TODO(jota): Decide how critical errors will be reported.
		}

		_, err = modelsFile.Write(formattedModelsCode)
		if err != nil {
			panic(errors.Wrapf(err, "error writing %s", outputModelsFileName)) // TODO(jota): Decide how critical errors will be reported.
		}
		reporter.Line("Done ...")
	},
}

func init() {
	rootCmd.AddCommand(genCmd)

	genCmd.PersistentFlags().StringVarP(&inputPackage, "input-package", "p", ".", "the package where the ceous will find the models for generating the code.")
	genCmd.PersistentFlags().StringArrayVarP(&excludedFiles, "exclude-files", "e", []string{}, "list of file names that will be ignored in the input package.")
	genCmd.PersistentFlags().StringVarP(&outputGeneratedFileName, "output-filename", "o", "ceous.go", "the file name that will contain the connections, queries and stores implemented.")
	genCmd.PersistentFlags().StringVarP(&outputModelsFileName, "output-model-filename", "m", "ceous_models.go", "the file name that will contain the models extensions.")
	genCmd.PersistentFlags().StringVarP(&outputPackage, "output-package", "d", "", "the package were the connections, queries and stores will be placed.")
	genCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "enable verbose mode.")
}
