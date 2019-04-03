package main

import (
	"flag"

	"github.com/nareshkumarthota/swagger-flogo/app"
	"github.com/nareshkumarthota/swagger-flogo/transform"
)

func main() {
	input := flag.String("input", "swagger.json", "input swagger file")
	port := flag.String("port", "8080", "flogo app running port")
	conversionType := flag.String("convertto", "flogoapiapp", "conversion type like flogoapiapp or flogodescriptor")
	output := flag.String("output", ".", "path to store generated file")

	flag.Parse()

	config := &transform.Config{}

	config.FileName = *input
	config.ConversionType = *conversionType
	config.Port = *port
	config.OutFilePath = *output

	app.Transform(config)
}
