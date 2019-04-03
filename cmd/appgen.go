package cmd

import (
	"github.com/nareshkumarthota/swagger-flogo/app"
	"github.com/nareshkumarthota/swagger-flogo/transform"
	"github.com/project-flogo/cli/common" // Flogo CLI support code
	"github.com/spf13/cobra"
)

func init() {
	appgen.Flags().StringVarP(&input, "input", "i", "swagger.json", "path to input swagger file")
	appgen.Flags().StringVarP(&port, "port", "p", "8080", "flogo app running port")
	appgen.Flags().StringVarP(&conversionType, "type", "t", "flogoapiapp", "conversion type like flogoapiapp or flogodescriptor")
	appgen.Flags().StringVarP(&output, "output", "o", ".", "path to generated file")
	common.RegisterPlugin(appgen)
}

var input, port, conversionType, output string
var appgen = &cobra.Command{
	Use:   "appgen",
	Short: "generates flogo/microgateway app",
	Long:  "This plugin command generates supplied spec to flogo/microgateway app",
	Run: func(cmd *cobra.Command, args []string) {
		config := &transform.Config{}

		config.FileName = input
		config.ConversionType = conversionType
		config.Port = port
		config.OutFilePath = output

		app.Transform(config)
	},
}
