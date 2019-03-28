package app

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/nareshkumarthota/swagger-flogo/transform"
)

// Transform transforms swagger
func Transform(config *transform.Config) {
	// Read the swagger file
	swaggerData, err := ioutil.ReadFile(config.FileName)
	if err != nil {
		log.Fatal("error occured in reading template: ", err)
	}

	var swagger map[string]interface{}

	json.Unmarshal(swaggerData, &swagger)

	config.SwaggerData = swagger

	switch config.ConversionType {
	case "flogoapiapp":
		transform.SwaggerToFlogoAPI(config)
	default:
		fmt.Println("ConversionType not specified")
	}
}
