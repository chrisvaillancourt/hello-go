# go tutorial

## Hello Steps

1. run `mkdir hello`
1. enable dependency tracking: `go mod init example/hello`
1. add module requirements and checksum: `go mod tidy`
1. run `go run .`

## Create a go module

### Modules

A module is a collection of one or more related packages.
Go code is grouped into packages; packages are grouped into modules.

The module specifies:

1. your dependencies
1. required go version
1. other required modules

### tutorial steps

1. `mkdir greetings`
1. `cd greetings`
1. initialize a new go module with `go mod init example.com/greetings`. If publishing a module, the path must be the location to download the module i.e. the code's repository. See [managing dependencies](https://go.dev/doc/modules/managing-dependencies#naming_module) for more info.

### notes

- functions with a capital name can be imported by other functions outside of the package it's defined in. See [exported names](https://go.dev/tour/basics/3) for more info.
- the `:=` operator is a way to declare and initialize a variable in a single expression. The variable's type is inferred.

  - the long form is:

  ```go
  var message string
  message = "something"
  ```
