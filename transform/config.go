package transform

// Config swagger config details
type Config struct {
	SwaggerData    map[string]interface{}
	FileName       string
	ConversionType string
}

// APIAppData holds app data retrieved from swagger file
type APIAppData struct {
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
