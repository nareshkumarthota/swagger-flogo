package transform

// Config swagger config details
type Config struct {
	SwaggerData    map[string]interface{}
	FileName       string
	ConversionType string
}

// APIAppData holds app data retrieved from swagger file
type APIAppData struct {
	Port        string
	Path        string
	Method      string
	HandlerName string
}
