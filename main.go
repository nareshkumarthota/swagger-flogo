package main

import (
	"flag"

	"github.com/nareshkumarthota/swagger-flogo/app"
	"github.com/nareshkumarthota/swagger-flogo/transform"
)

func main() {
	input := flag.String("input", "swagger.json", "input swagger file")
	conversionType := flag.String("convertto", "flogoapiapp", "conversion type like flogoapiapp or flogodescriptor")

	flag.Parse()

	config := &transform.Config{}

	config.FileName = *input
	config.ConversionType = *conversionType

	app.Transform(config)
}
