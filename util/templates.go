package util

// flogoAPITemplate template for flogo api app
const flogoAPITemplate = `
package main

import (
	"context"

	"github.com/TIBCOSoftware/flogo-contrib/activity/log"
	rt "github.com/TIBCOSoftware/flogo-contrib/trigger/rest"
	"github.com/TIBCOSoftware/flogo-lib/core/data"
	"github.com/TIBCOSoftware/flogo-lib/engine"
	"github.com/TIBCOSoftware/flogo-lib/flogo"
	"github.com/TIBCOSoftware/flogo-lib/logger"
)

func main() {
	// Create a new Flogo app
	app := appBuilder()

	e, err := flogo.NewEngine(app)

	if err != nil {
		logger.Error(err)
		return
	}

	engine.RunEngine(e)
}

func appBuilder() *flogo.App {
	app := flogo.NewApp()

	// Register the HTTP trigger
	trg := app.NewTrigger(&rt.RestTrigger{}, map[string]interface{}{"port": {{.Port}}})
	{{- range .PathData }}
	{{$pathURL := .PathURL}}
	{{- range .MethodData }}
	trg.NewFuncHandler(map[string]interface{}{"method": "{{.MethodType}}", "path": "{{$pathURL}}"}, {{.HandlerName}})
	{{- end }}
	{{- end }}
	return app
}


{{- range .PathData }}
{{- range .MethodData }}
func {{.HandlerName}}(ctx context.Context, inputs map[string]*data.Attribute) (map[string]*data.Attribute, error) {

	// Execute the log activity
	in := map[string]interface{}{"message": "logmessage from operationID:{{.HandlerName}}", "flowInfo": "true", "addToFlow": "true"}
	_, err := flogo.EvalActivity(&log.LogActivity{}, in)
	if err != nil {
		return nil, err
	}

	/*********
	//
	//
	//User implementation area
	//
	//
	***********/

	// The return message is a map[string]*data.Attribute which we'll have to construct
	response := make(map[string]interface{})
	response["response"] = "success response"

	ret := make(map[string]*data.Attribute)
	ret["code"], _ = data.NewAttribute("code", data.TypeInteger, 200)
	ret["data"], _ = data.NewAttribute("data", data.TypeAny, response)

	return ret, nil
}
{{- end }}
{{- end }}
`

const flogoAppDescriptor = `
{
	"name": "SampleApp",
	"type": "flogo:app",
	"version": "0.0.1",
	"appModel": "1.1.0",
	"imports": [
		"github.com/project-flogo/contrib/trigger/rest",
		"github.com/project-flogo/flow",
		"github.com/project-flogo/contrib/activity/log"
	],
	"triggers": [
	  {
		"id": "receive_http_message",
		"ref": "#rest",
		"name": "Receive HTTP Message",
		"description": "Simple REST Trigger",
		"settings": {
		  "port": 9233
		},
		"handlers": [
		{{- range .PathData }}
		{{$pathURL := .PathURL}}
		{{- range .MethodData }}
		  {
			"settings": {
			  "method": "{{.MethodType}}",
			  "path": "{{$pathURL}}"
			},
			"action": {
			  "ref": "#flow",
			  "settings": {
				"flowURI": "res://flow:sample_flow"
			  }
			}
		  },
		{{- end }}
		{{- end }}
		]
	  }
	],
	"resources": [
	  {
		"id": "flow:sample_flow",
		"data": {
		  "name": "SampleFlow",
		  "tasks": [
			{
			  "id": "log_message",
			  "name": "Log Message",
			  "description": "Simple Log Activity",
			  "activity": {
				"ref": "#log",
				"input": {
				  "message": "Simple Log",
				  "addDetails": "false"
				}
			  }
			}
		  ]
		}
	  }
	]
  }
`
