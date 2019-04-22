# swagger-flogo
Swagger to flogo app converter tool converts given swagger spec to its implementation based on flogo api/descriptor model.

Currently this tool accepts below arguments.
```sh
Usage of swagger-flogo:
  -type string
        conversion type like flogoapiapp or flogodescriptor (default "flogoapiapp")
  -input string
        input swagger file (default "swagger.json")
  -output string
        path to store generated file (default ".")
  -port string
        flogo app running port (default "8080")
```
## Setup
To install the tool, simply open a terminal and enter the below command
```sh
go get github.com/nareshkumarthota/swagger-flogo/...
```

## Usage
### Flogo app api model.
```sh
cd $GOPATH/src/github.com/nareshkumarthota/swagger-flogo/examples

swagger-flogo -input swagger.json -type flogoapiapp -port 9090 -output flogoApiApp
```
The resulting output is two files get created `flogoapiapp.go` and `support.go` under flogoApiApp folder.  Use both files to build/install app. User defined code can be written in `support.go`.
```sh
cd flogoApiApp
go build
./flogoApiApp
```
### Flogo app descriptor model.
```sh
cd $GOPATH/src/github.com/nareshkumarthota/swagger-flogo/examples

swagger-flogo -input swagger.json -type flogodescriptor -port 9090 -output flogoDescriptorApp
```
The resulting output is two files get created `flogodescriptor.json` and `support.go` under flogoDescriptorApp folder. Use `flogodescriptor.json` to create flogo app. If you like to add any operation on incoming requests you can edit `support.go` file.

```sh
cd flogoDescriptorApp
flogo create -f flogodescriptor.json flogoapp
cp support.go flogoapp/src
cd flogoapp
flogo build
./bin/flogoapp
```

## Support File
User defined code can be written in `support.go` file. Activity `methodinvoker` is used to invoke user defined logic for each `operationId`. Methods are stored based on `operationId` present in swagger file. Sample swagger content for a path is given below.
```json
"/app/{id}": {
            "get":{
                    "operationId": "getApp_id"
                  }
             }      
```
Sample autogenerated code is given below.
```go
package main

import (
	"github.com/nareshkumarthota/flogocomponents/activity/methodinvoker"
)

func init() {
    methodinvoker.RegisterMethods("getApp_id", loginUser_mthd)
}

func getApp_id_mthd(input interface{}) (map[string]interface{}, error) {
	//
	// User Implementation area
	//

	output := make(map[string]interface{})
	output["response"] = "Auto generated code response for getApp_id method"
	output["inputs"] = input

	return output, nil
}
```
For incoming requests on `/app/{id}` method `getApp_id_mthd` will get invoked.
## Flogo Plugin Support
This tool can be integrated into [flogocli](https://github.com/project-flogo/cli).
```sh
# Install your plugin
$ flogo plugin install github.com/nareshkumarthota/swagger-flogo/cmd

# Run your new plugin command for api app model
$ flogo appgen -i swagger.json -t flogoapiapp -p 9090 -o flogoApiApp

# Run your new plugin command for descriptor app model
$ flogo appgen -i swagger.json -t flogodescriptor -p 9090 -o flogoDescriptorApp
```