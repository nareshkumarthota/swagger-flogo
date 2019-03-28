package transform

import (
	"fmt"

	"github.com/nareshkumarthota/swagger-flogo/util"
)

// SwaggerToFlogoAPI transforms given swagger to flogo api application
func SwaggerToFlogoAPI(config *Config) error {
	paths := config.SwaggerData["paths"].(map[string]interface{})
	for key, value := range paths {
		fmt.Println(key, value)
	}

	data := APIAppData{}

	data.Port = "8080"
	data.Path = "/api/invoices/:id"
	data.Method = "GET"
	data.HandlerName = "invoiceget"

	util.ExecuteTemplate(config.ConversionType, data)

	return nil
}
