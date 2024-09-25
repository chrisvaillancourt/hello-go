# [Tour of Go](https://go.dev/tour/list)

## Basics

### Packages, variables, and functions

#### Naming packages

Run an app with a `main` package.
By convention, the package name is the same as the last element of the import
path. i.e. the `"math/rand"` package comprises files that begin with the statement `package rand`.

#### Exported names

When importing a package, only the exported names are available (names starting with a capital letter). Lower case names aren't available outside the package.

### Functions

#### Types come after names

```go
func add(x int, y int) int {
  return x + y
}
```

[See blog on go's declaration syntax](https://go.dev/blog/declaration-syntax).

If two or more consecutive function parameters share a type, we can omit the
type from each param except the last.

For example:

```go
func add(x int, y int) int {
  return x + y
}
```

Becomes:

```go
func add(x, y int) int {
  return x + y
}
```

#### A function can return multiple values

```go
package main

import "fmt"

func swap(x, y string) (string, string) {
  return y, x
}

func main() {
  a, b := swap("hello", "world")
  fmt.Println(a, b)
}

```

#### Named return values

We can name the values returned from a function. When used, their names are
treated as variables defined at the top of the function.
It's best to use the names to document the meaning of the returned values.

A `return` statement without arguments returns the named return values (aka a naked return). It's best to only use naked returns in short functions.

```go
package main

import "fmt"

func split(sum int) (x, y int) {
  x = sum * 4 / 9
  y = sum - x
  return
}

func main() {
  fmt.Println(split(17))
}
```
