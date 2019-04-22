package util

// flogoAPITemplate template for flogo api app
const flogoAPITemplate = `
package main

import (
	"context"
	"log"

	"github.com/nareshkumarthota/flogocomponents/activity/methodinvoker"
	"github.com/project-flogo/contrib/trigger/rest"
	"github.com/project-flogo/core/activity"
	"github.com/project-flogo/core/api"
	"github.com/project-flogo/core/engine"
)

func main() {

	app := myApp()

	e, err := api.NewEngine(app)

	if err != nil {
		log.Fatal("Error:", err)
		return
	}

	engine.RunEngine(e)
}

func myApp() *api.App {
	app := api.NewApp()

	trg := app.NewTrigger(&rest.Trigger{}, &rest.Settings{Port: {{.Port}}})

	var handler *api.Handler

	{{- range  .PathData }} {{$pathURL := .PathURL}}  {{- range  .MethodData }} 
	handler, _ = trg.NewHandler(&rest.HandlerSettings{Method: "{{.MethodType}}", Path: "{{$pathURL}}"})
	handler.NewAction({{.HandlerName}})
	{{- end }} {{- end }}

	//store in map to avoid activity instance recreation
	mtdInvkr, _ := api.NewActivity(&methodinvoker.Activity{})
	activities = map[string]activity.Activity{"methodinvoker": mtdInvkr}

	return app
}

var activities map[string]activity.Activity

{{- range  .PathData }} {{- range  .MethodData }} 
func {{.HandlerName}}(ctx context.Context, inputs map[string]interface{}) (map[string]interface{}, error) {
	return methodInvokerActivity(ctx, inputs, "{{.HandlerName}}")
}
{{- end }} {{- end }}

func methodInvokerActivity(ctx context.Context, inputs map[string]interface{}, methodName string) (map[string]interface{}, error) {

	trgOut := &rest.Output{}
	trgOut.FromMap(inputs)

	out, err := api.EvalActivity(activities["methodinvoker"], &methodinvoker.Input{MethodName: methodName, InputData: inputs})
	if err != nil {
		return nil, err
	}

	reply := &rest.Reply{Code: 200, Data: out}
	return reply.ToMap(), nil
}
`

const flogoAppDescriptor = `
{
	"name": "SampleApp",
	"type": "flogo:app",
	"version": "0.0.1",
	"description": "",
	"appModel": "1.1.0",
	"imports": [
		"github.com/nareshkumarthota/flogocomponents/activity/methodinvoker",
		"github.com/project-flogo/contrib/activity/mapper",
		"github.com/project-flogo/contrib/trigger/rest",
		"github.com/project-flogo/flow"
	],
	"triggers": [
		{
			"id": "receive_http_message",
			"ref": "#rest",
			"settings": {
				"port": {{.Port}}
			},
			"handlers": [
			{{ $pathsLen := len .PathData }} {{- range $index , $element := .PathData }} {{ $pathIndex := increment $index }}	{{$pathURL := .PathURL}}	{{ $mthdLen := len .MethodData }} {{- range $mindex , $melement := .MethodData }} {{ $methodIndex := increment $mindex }}
				{
					"settings": {
						"method": "{{.MethodType}}",
						"path": "{{$pathURL}}"
					},
					"actions": [
						{
							"ref": "#flow",
							"settings": {
								"flowURI": "res://flow:sample_flow"
							},
							"input": {
								"inputData": {
									"mapping" : {
										"pathParams": "=$.pathParams",
										"headers": "=$.headers",
										"queryParams": "=$.queryParams",
										"content":"=$.content"
									}                  
								},
								"methodName": "{{.HandlerName}}"
							},
							"output": {
								"data": "=$.outputData"
							}
						}
          			]
				}{{ if eq $pathsLen $pathIndex }}{{ if eq $mthdLen $methodIndex }}{{else}},{{end}}{{else}},{{end}}
			{{- end }} {{- end }}
			]
	   }
	],
	"resources": [
    {
      "id": "flow:sample_flow",
      "data": {
        "name": "SampleFlow",
        "metadata": {
          "output": [
            {
              "name": "outputData",
              "type": "any"
            }
          ],
          "input": [
            {
              "name": "inputData",
              "type": "object"
            },
            {
              "name": "methodName",
              "type": "string"
            }
          ]
        },
        "tasks": [
          {
            "id": "mthdInvk_activity",
            "name": "method Invoke activity",
            "description": "Simple Method Invoke Activity",
            "activity": {
              "ref": "#methodinvoker",
              "input": {
                "methodName": "=$.methodName",
                "inputData": "=$.inputData"
              }
            }
          },
          {
            "id": "mapperAct",
            "name": "Mapper",
            "description": "Mapper",
            "activity": {
              "ref": "#mapper",
              "settings": {
                "mappings": {
                  "outputData": "=$activity[mthdInvk_activity].outputData"
                }
              }
            }
          }
        ],
        "links": [
          {
            "from": "mthdInvk_activity",
            "to": "mapperAct"
          }
        ]
      }
    }
  ]
  }
`
const supportFile = `
package main

import (
	"log"

	"github.com/nareshkumarthota/flogocomponents/activity/methodinvoker"
)

func init() {
	{{- range  .PathData }} {{- range  .MethodData }}
	methodinvoker.RegisterMethods("{{.HandlerName}}", {{.HandlerName}}_mthd)
	{{- end }} {{- end }}
}

{{- range  .PathData }} {{- range  .MethodData }}
func {{.HandlerName}}_mthd(input interface{}) (map[string]interface{}, error) {

	log.Println("inputs ", input)

	//
	// User Implementation area
	//

	output := make(map[string]interface{})
	output["response"] = "Auto generated code response for {{.HandlerName}} method"
	output["inputs"] = input

	return output, nil
}
{{- end }} {{- end }}
`
