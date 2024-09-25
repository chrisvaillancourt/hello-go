# go tutorial

## Hello Steps

1. run `mkdir hello`
1. enable dependency tracking: `go mod init example/hello`
1. add module requirements and checksum: `go mod tidy`
1. run `go run .`

## [Create a go module](https://go.dev/doc/tutorial/create-module)

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

## [Call code from another module](https://go.dev/doc/tutorial/call-module-code)

### Steps

1. create a hello directory (location of the caller): `mkdir hello`
   a. trash old directory from first part of the tutorial
1. `cd hello`
1. enable dependency tracking with `go mod init example.com/hello`
1. add a main function that calls the `Hello()` function
   a. code executed as an application must be in a main package.
1. edit example.com/hello module to use local greetings module: `go mod edit -replace example.com/greetings=../greetings`
   a. allows us to redirect go tools from the module path (fake example.com/greetings) to the actual location (local directory)
1. From the hello directory, run `go mode tidy`
   a. syncs the example.com/hello module's dependencies
   a. If we were referencing a real published module, the `go.mod` file would look like `require example.com/greetings v1.1.0`
1. run the code with `go run .` from the hello directory

## [Return and handle an error](https://go.dev/doc/tutorial/handle-errors)

1. edit `greetings/greetings.go` to return an error if the name is empty
