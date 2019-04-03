package transform

// Config swagger config details
type Config struct {
	SwaggerData    map[string]interface{}
	FileName       string
	ConversionType string
	Port           string
	OutFilePath    string
}

// AppData holds app data retrieved from swagger file
type AppData struct {
	Port     string
	PathData []Path
}

// Method holds method data
type Method struct {
	MethodType  string
	HandlerName string // operationID treated as handler name
}

// Path holds all path details from swagger
type Path struct {
	PathURL    string
	MethodData []Method
}
