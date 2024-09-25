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
