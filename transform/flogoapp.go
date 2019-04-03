package transform

import (
	"strings"

	"github.com/nareshkumarthota/swagger-flogo/util"
)

// SwaggerToFlogoAPI transforms given swagger to flogo api application
func SwaggerToFlogoAPI(config *Config) error {
	util.ExecuteTemplate(config.ConversionType, config.OutFilePath, appData(config))
	return nil
}

// SwaggerToFlogoDescriptor transforms given swagger to flogo api application
func SwaggerToFlogoDescriptor(config *Config) error {
	util.ExecuteTemplate(config.ConversionType, config.OutFilePath, appData(config))
	return nil
}

func appData(config *Config) AppData {

	data := AppData{}

	data.Port = config.Port

	// retrieve details from paths map
	paths := config.SwaggerData["paths"].(map[string]interface{})

	// initialise PathData
	data.PathData = make([]Path, len(paths))
	index := 0
	for key, value := range paths {

		data.PathData[index].PathURL = util.ModifyPathSymbols(key)

		// retrieving method details for the paths value component
		methods := value.(map[string]interface{})
		data.PathData[index].MethodData = make([]Method, len(methods))

		mIndex := 0
		for mk, mv := range methods {
			data.PathData[index].MethodData[mIndex].MethodType = strings.ToUpper(mk)
			data.PathData[index].MethodData[mIndex].HandlerName = mv.(map[string]interface{})["operationId"].(string)

			// increment method index
			mIndex++
		}

		// increment path index
		index++
	}
	return data
}
