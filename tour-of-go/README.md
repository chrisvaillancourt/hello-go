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

### variables

`var` statement declares a list of variables. The type comes last.
Can be scoped to a package or function.

```go
package main

import "fmt"

var c, python, java bool

func main() {
  var i int
  fmt.Println(i, c, python, java)
}
```

#### Variables with initializers

`var` declarations can include initializers, one per variable.
We can omit the type if we provide an initial value.

```go
package main

import "fmt"

var i, j int = 1, 2

func main() {
  var c, python, java = true, false, "no!"
  fmt.Println(i, j, c, python, java)
}
```

#### Short variable declarations

We can use `:=` inside a function instead of a `var` declaration. The type of the value is inferred from the initial value.

```go
package main

import "fmt"

func main() {
  var i, j int = 1, 2
  k := 3
  c, python, java := true, false, "no!"

  fmt.Println(i, j, k, c, python, java) // 1 2 3 true false no!
}
```

### Basic types

bool

string

int int8 int16 int32 int64
uint uint8 uint16 uint32 uint64 uintptr

byte // alias for uint8

rune // alias for int32
// represents a Unicode code point

float32 float64

complex64 complex128

Variable declarations can be factored into blocks (like import statements).

```go
package main

import (
  "fmt"
  "math/cmplx"
)

var (
  ToBe   bool       = false
  MaxInt uint64     = 1<<64 - 1
  z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
  fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe) // Type: bool Value: false
  fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt) // Type: uint64 Value: 18446744073709551615
  fmt.Printf("Type: %T Value: %v\n", z, z) // Type: complex128 Value: (2+3i)
}
```
