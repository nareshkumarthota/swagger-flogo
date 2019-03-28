# swagger-flogo
Mini cli which performs swagger implementation for the given swagger spec underlined with flogo/microgateway app.

## Test
Run the below command.
```sh
go run main.go -input swagger.json -convertto flogoapiapp
```
You can see `flogoapiapp.go` file gets generated.