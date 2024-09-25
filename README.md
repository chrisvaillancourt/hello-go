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

## [Return a random greeting](https://go.dev/doc/tutorial/random-greeting)

Change code so we return one of three predefined greeting messages.
We'll use a [slice](https://go.dev/blog/slices-intro) to do this.
A slice is like an array, except its size changes dynamically as you add & remove items.

## [return greetings for multiple people](https://go.dev/doc/tutorial/greetings-multiple-people)

Handle multiple-value input by pairing input values with multi-value output.

In Go, you initialize a map with the following syntax:
`make(map[key-type]value-type)`.
See [Go maps in action on the Go blog](https://go.dev/blog/maps) for more info.

When looping with `for _, name := range ...`, range returns two values, the first is the index of the current item in the loop, the second is a copy of the items value. Use a `_` if you don't need the item's index. See the [blank identifier in Effective Go](https://go.dev/doc/effective_go.html#blank).

See [keeping your modules compatible](https://go.dev/blog/module-compatibility) for more info on backwards compatibility.

## [Add a test](https://go.dev/doc/tutorial/add-a-test)

Ending a file's name with `_test.go` tells the go test command that this file
contains test functions.

Test function names have the form `TestName`, where `Name` says something about
the specific test.
Test functions take a pointer to the testing package's [`testing.T` type](https://pkg.go.dev/testing/#T) as a
parameter.
You use this parameter's methods for reporting and logging from your test.

From the `greetings/` dir, run the [`go test` command](https://go.dev/cmd/go/#hdr-Test_packages) to execute the test.
Will execute functions with names beginning with `Test` in test files
(ending in `_test.go`).

Pass `-v` flag to get verbose test output i.e. `go test -v`.

## [Compile and install the app](https://go.dev/doc/tutorial/compile-install)

Running `go run .` is useful for compiling and running a program during dev but it doesn't generate a binary executable.

The `go build` command compiles packages and dependencies without installing.
The `go install` command compiles and installs packages.

### Steps

1. From the `hello/` dir, run `go build` to compile the code into an executable
1. run the hello executable with `./hello`
   a. prompt needs to be in the same directory as the hello executable or specify executable's path
1. Find the Go install path where the go command installs the current package: `go list -f '{{.Target}}'`
1. add the go install directory to your system's shell path. This allows us to run the executable without specifying where it is. Run `export PATH=$PATH:/path/to/your/install/directory` replacing `/path/to/your/install/directory` with the output from the previous step's `go list -f '{{.Target}}'` command. i.e. `export PATH=$PATH:/Users/chris/go/bin/hello`.
