# openapi-typegen
[WIP] web framework independent type generator with validation for OpenAPI 3.0 specifications.


```shell
go mod install github.com/jxsl13/openapi-typegen/cmd/openapi-typegen@latest
```


## Usage

```shell
$ openapi-typegen --help

  TYPEGEN_FILE       file path to your openapi.yaml (default: "openapi.yaml")
  TYPEGEN_OUT        out file path or 'stdout' (default: "stdout")
  TYPEGEN_PACKAGE    package name of the generated file (default: "api")

Usage:
  openapi-typegen -p api -f openapi.yaml -o types.gen.go [flags]
  openapi-typegen [command]

Available Commands:
  completion  Generate completion script
  help        Help about any command

Flags:
  -f, --file string      file path to your openapi.yaml (default "openapi.yaml")
  -h, --help             help for openapi-typegen
  -o, --out string       out file path or 'stdout' (default "stdout")
  -p, --package string   package name of the generated file (default "api")

Use "openapi-typegen [command] --help" for more information about a command.
````