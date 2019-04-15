# swagger-flogo
Swagger to flogo app converter tool converts given swagger spec to its implementation based on flogo api/descriptor model.

Currently this tool accepts below arguments.
```sh
Usage of swagger-flogo:
  -convertto string
        conversion type like flogoapiapp or flogodescriptor (default "flogoapiapp")
  -input string
        input swagger file (default "swagger.json")
  -output string
        path to store generated file (default ".")
  -port string
        flogo app running port (default "8080")
```

## Usage
1. Flogo api app model.
```sh
swagger-flogo -input swagger.json -convertto flogoapiapp -port 9090 -output .
```
The resulting output is two files get created `flogoapiapp.go` and `support.go`.  Use both files to build/install app. User defined code can be written in `support.go`.
```sh
mkdir flogoapp
cp support.go flogoapiapp.go ./flogoapp
cd flogoapp
go build
./flogoapp
```
2. Flogo descriptor app model.
```sh
swagger-flogo -input swagger.json -convertto flogodescriptor -port 9090 -output .
```
The resulting output is two files get created `flogodescriptor.json` and `support.go`. Use `flogodescriptor.json` to create flogo app.

```sh
flogo create -f flogodescriptor.json flogoapp
cp support.go flogoapp/src
cd flogoapp
flogo build
./bin/flogoapp
```

## Support File
User defined code can be written in `support.go` file. Activity `methdinvoker` is used so that user can write his own logic for the implementation. Methods are stored based on operation id given under swagger file. Sample code is given below.
```go
package main

import (
	"github.com/nareshkumarthota/flogocomponents/activity/methodinvoker"
)

func init() {
    methodinvoker.RegisterMethods("loginUser", loginUser_mthd)
}

func loginUser_mthd(input interface{}) (map[string]interface{}, error) {
	//
	// User Implementation area
	//

	output := make(map[string]interface{})
	output["response"] = "Auto generated code response for loginUser method"
	output["inputs"] = input

	return output, nil
}
```

## Flogo Plugin Support
This tool can be integrated into [flogocli](https://github.com/project-flogo/cli).
```sh
# Install your plugin
$ flogo plugin install github.com/nareshkumarthota/swagger-flogo/cmd

# Run your new plugin command
$ flogo appgen
```