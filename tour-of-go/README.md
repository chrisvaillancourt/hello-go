# [Tour of Go](https://go.dev/tour/list)

## Packages, variables, and functions

### Naming packages

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

### Zero values

Variables declared without an explicit initial value are given _their zero value_.

Zero values:

- `0` for numeric types
- `false` for booleans
- `""` for strings

```go
package main

import "fmt"

func main() {
  var i int // 0
  var f float64 // 0
  var b bool // false
  var s string // ""
  fmt.Printf("%v %v %v %q\n", i, f, b, s)
}
```

### Type conversions

The expression `T(v)` converts the `v` value to the type `T`. There's no implicit conversion of types in Go, everything must be explicitly converted.

```go
var i int = 42
var f float64 = float64(i)
var u uint = uint(f)
// simple version

i := 42
f := float64(i)
u := uint(f)
```

```go
package main

import (
  "fmt"
  "math"
)

func main() {
  var x, y int = 3, 4
  var f float64 = math.Sqrt(float64(x*x + y*y))
  var z uint = uint(f)
  fmt.Println(x, y, z)
}
```

### Type inference

Typed declarations create variables of that type:

```go
var i int
j := i // j is an int

i := 42           // int
f := 3.142        // float64
g := 0.867 + 0.5i // complex128
```

### Constants

Declared with the `const` keyword; can't be declared with the `:=` syntax.

```go
package main

import "fmt"

const Pi = 3.14

func main() {
  const World = "世界"
  fmt.Println("Hello", World)
  fmt.Println("Happy", Pi, "Day")

  const Truth = true
  fmt.Println("Go rules?", Truth)
}
```

### Numeric constants

Numeric constants are high-precision values. Untyped constants take the type
from it's context.

```go
package main

import "fmt"

const (
  // Create a huge number by shifting a 1 bit left 100 places.
  // In other words, the binary number that is 1 followed by 100 zeroes.
  Big = 1 << 100
  // Shift it right again 99 places, so we end up with 1<<1, or 2.
  Small = Big >> 99
)

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
  return x * 0.1
}

func main() {
  fmt.Println(needInt(Small)) // 21
  fmt.Println(needFloat(Small)) // 0.2
  fmt.Println(needFloat(Big)) // 1.2676506002282295e+29
}
```

## Flow control: for, if, else, switch, and defer

### for

The only looping construct is the `for` loop.
The `for` loop has three components, each separated with a `;`:

1. The init statement: executed before the first loop iteration
1. the condition expression: evaluated before every iteration
1. the post statement: executed at the end of every iteration

The init statement is usually a short variable declaration. The variables declared here are only in scope of the `for` statement.

The loop continues until the boolean condition evaluates to `false`.

```go
package main

import "fmt"

func main() {
  sum := 0
  for i := 0; i < 10; i++ {
    sum += i
  }
  fmt.Println(sum) // 45
}
```

The init and post statements are optional:

```go
package main

import "fmt"

func main() {
  sum := 1
  for ; sum < 1000; {
    sum += sum
  }
  fmt.Println(sum) // 1024
}
```

We can write a while loop like:

```go
package main

import "fmt"

func main() {
  sum := 1
  for sum < 1000 {
    sum += sum
  }
  fmt.Println(sum) // 1024
}
```

You can write an infinite loop by dropping the loop condition:

```go
package main

func main() {
  for {
    // runs forever
  }
}
```

### if

Parentheses aren't needed by braces `{ }` are required.

```go
package main

import (
	"fmt"
)

func isPositive(x int) bool {
	if x < 0 {
		return false
	}
	return true
}

func main() {
	fmt.Println(isPositive(5)) // true
}
```

### if with a short statement

The `if` statement can start with a short statement that's executed before the
conditional.
Variables declared by the statement are only in scope within the `if`
statement's curly braces.

```go
package main

import (
	"fmt"
)

func isOverLimit(val int) string {
	const limit = 10
	if valPlusOne := val + 1; valPlusOne < limit {
		return "below limit"
	}
	return "above limit"
}

func main() {
	fmt.Println(isOverLimit(1))  // below limit
	fmt.Println(isOverLimit(15)) // above limit
}
```

### if and else

Variables declared inside an `if`'s short statement are available in the else
blocks too.

```go
package main

import (
	"fmt"
)

func foo(val string) string {
	if isEmpty := val == ""; isEmpty == true {
		return "empty string"
	} else {
		fmt.Println(isEmpty)
	}
	return "bar"
}

func main() {
	fmt.Println(foo("")) // empty string
}
```

### switch

More succinct version of an `if else` statements. Only the selected case is
executed so there's no need for a `break` statement like in JavaScript.
Switch cases are evaluated from top to bottom, stopping when the first case
matches.

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}
}
```

You can write a switch statement without a condition. It's a good alternative
to long `if else` chains.

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}
```

### Defer

The `defer` statement defers the execution of a function until the surrounding
function returns.
The function's arguments are evaluated immediately but the function isn't
isn't executed until the surrounding function returns.

```go
package main

import "fmt"

func main() {
	defer fmt.Println("world")
	fmt.Println("hello")
}
```

### Stacked defers

Deferred function calls are pushed onto a stack.
When a function returns, its deferred calls are executed in a last in, first
out order.

```go
package main

import "fmt"

func main() {
	fmt.Println("counting")

	defer fmt.Println("A")
	defer fmt.Println("B")
	defer fmt.Println("C")

	fmt.Println("done")
}
// counting
// done
// C
// B
// A
```

## More types: structs, slices, and maps

### Pointers

A pointer holds the memory address of a value.
The type `*T` is a pointer to a `T` value.
The `&` operator generates a pointer to its operand (?).
The `*` operator denotes the pointer's underlying value.

```go
package main

import "fmt"

func main() {
	i, j := 42, 2701

	p := &i         // point to i
	fmt.Println(*p) // read i through the pointer
	*p = 21         // set i through the pointer
	fmt.Println(i)  // see the new value of i

	p = &j         // point to j
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of j
}
```

### Structs

A `struct` is a collection of fields.

```go
package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	fmt.Println(Vertex{1, 2}) // {1 2}
}
```

Access fields using a dot.

```go
package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	v.X = 4
	fmt.Println(v.X) // 4
}
```

### Pointers to structs

Fields can be accessed through a struct pointer.
To access the field `X` of a struct when we have the struct pointer `p` we could
write `(*p).X` or the short version `p.x`

```go
package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	p := &v
	p.X = 1e9
	fmt.Println(v) // {1000000000 2}
}
```

### Struct literals

A struct literal denotes a newly allocated struct value by listing the values
of its fields.
You can set a subset of the struct's fields with the `Name:` syntax. Unnamed
fields will have the default zero value for that type. You can provide multiple fields by separating each with a comma:

```go
func main() {
	type Vertex struct {
		X, Y, Z int
	}
	var v2 = Vertex{X: 1, Y: 2}
	fmt.Println(v2)
}
```

Adding `&` as a prefix will return a pointer to the struct value.

```go
package main

import "fmt"

type Vertex struct {
	X, Y int
}

var (
	v1 = Vertex{1, 2}  // has type Vertex
	v2 = Vertex{X: 1}  // Y:0 is implicit
	v3 = Vertex{}      // X:0 and Y:0
	p  = &Vertex{1, 2} // has type *Vertex
)

func main() {
	fmt.Println(v1, p, v2, v3)
}
```

### Arrays

The type `[n]T` is an array of `n` values of type `T`.
`var a [10]int` declares a variable `a` as an array of 10 integers.
The length of an array is fixed (it's a part of its type).

```go
package main

import "fmt"

func main() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
}
```

### Slices

A dynamically sized view into an array.
The type of a slice is `[]T` with `T` representing the element's type.
Create a slice by specifying two indices, a low and high bound, separated by a
colon: `a[low : high]`. The low bound is inclusive while the upper bound is
exclusive.
For example, `a[1:4]` creates a slice of `a` with elements 1 through 3.

```go
package main

import "fmt"

func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13}

	var s []int = primes[1:4]
	fmt.Println(s) // [3 5 7]
}
```

#### slices are like references to arrays

A slice doesn't store any data. A slice only describes a section of an
underlying array.
Modifying the elements of a slice will modify the corresponding elements in the
underlying array.

```go
package main

import "fmt"

func main() {
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	// 	fmt.Println(names)

	a := names[0:2] // [John Paul]
	b := names[1:3] // [Paul George]

	b[0] = "XXX"
	fmt.Println(a)     // [John XXX]
	fmt.Println(b)     // [XXX George ]
	fmt.Println(names) // [John XXX Geaorge Ringo]
}
```

#### Slice literals

Like an array literal without the length.
An array literal: `[3]bool{true, true, false}`.
A slice literal: `[]bool{true, true, false}`. This creates an array and then
builds a slice that references it.

```go
package main

import "fmt"

func main() {
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q) // [2 3 5 7 11 13]
}
```

#### Slice defaults

You can omit the high & low bounds when creating a slice to instead use the
default high and low bounds.
The default low bound is 0 and default high bound is the length of the slice.
Given an array `var a [10] int`, all of the following are equivalent:

```go
a[0:10]
a[:10]
a[0:]
a[:]
```

```go
package main

import "fmt"

func main() {
	s := []int{2, 3, 5, 7, 11, 13}

	s = s[1:4] // [3 5 7]
	fmt.Println(s)

	s = s[:2]
	fmt.Println(s) // [3 5]

	s = s[1:]
	fmt.Println(s) //[5]
}
```

#### Slice length and capacity

Slices have a length (the number of elements) and a capacity (the number of
elements in the underlying array, counting from the first element in the slice).
Get the length of a slice `s` with `len(s)`. Get the capacity with `cap(s)`.
You can extend a slice's length by re-slicing it, provided it has sufficient
capacity.

#### Nil slices

The zero value of a slice is `nil`. A nil slice has a length and capacity of 0
with no underlying array.

```go
package main

import "fmt"

func main() {
	var s []int
	fmt.Println(s, len(s), cap(s)) // [] 0 0
	if s == nil {
		fmt.Println("nil!") // nil!
	}
}
```

#### Creating a slice with make

We can use the built-in `make()` function to create slices. This is how we can
make dynamically sized arrays. For example: `a := make([]int, 5)  // len(a)=5`

`make()` allocates a zeroed array and returns a slice that refers to that array.
You can specify the capacity by passing a third argument to `make()`:
`b := make([]int, 0, 5) // len(b)=0, cap(b)=5`.

#### Appending to a slice

Add elements to a slice with built-in `append()`.
The first parameter is the slice to append to.
The second parameter is the value to add to the slice.

```go
var s []int
// append works on nil slices.
s = append(s, 0)
```

More than one value can be appended at a time: `s = append(s, 2, 3, 4)`.

See [this blog on using go slices](https://go.dev/blog/go-slices-usage-and-internals) for more info.

### Range

Use `range` with a `for` loop to iterate over a slice or map.
When using `range` on a slice, two values are returned on each iteration.
The first is the index and the second is a copy of the value at that index.

```go
package main

import "fmt"

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
}
```

You can skip the index or value by assigning `_` to it: `for i, _ := range pow`
or `for _, value := range pow`.
You can leave out the second value to get only the index:
`for i := range pow`.

### Maps

A map holds keys and values. The zero value is `nil`.
A nil map has no keys and we can't add keys to it.
Use the `make()` function to create a map of a given type. The returned map is
initialized and ready for use.

```go
package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

func main() {
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"]) // {40.68433 -74.39967}
}

```
