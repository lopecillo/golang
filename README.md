# Golang

## Table of contents

<!--ts-->
* [Golang](#golang)
* [Table of contents](#table-of-contents)
* [Packages](#packages)
  * [Imports](#imports)
  * [Exported names](#exported-names)
* [Functions](#functions)
  * [Multiple results](#multiple-results)
  * [Named return values](#named-return-values)
* [Variables](#variables)
  * [Initializers](#initializers)
  * [Short variable declarations](#short-variable-declarations)
  * [Basic types](#basic-types)
  * [Zero values](#zero-values)
  * [Type conversions](#type-conversions)
  * [Type inference](#type-inference)
  * [Constants](#constants)
    * [Numeric constants](#numeric-constants)
* [Flow control](#flow-control)
  * [For](#for)
  * [If](#if)
  * [*Exercise: Loops and functions*](#exercise-loops-and-functions)
  * [Switch](#switch)
  * [Defer](#defer)
* [More types](#more-types)
  * [Pointers](#pointers)
  * [Structs](#structs)
    * [Struct literals](#struct-literals)
  * [Arrays](#arrays)
  * [Slices](#slices)
    * [Slice literals](#slice-literals)
    * [Slice defaults](#slice-defaults)
    * [Slice lengths and capacity](#slice-lengths-and-capacity)
    * [Nil slices](#nil-slices)
    * [Creating a slice with make](#creating-a-slice-with-make)
    * [Appending to a slice](#appending-to-a-slice)
    * [Range](#range)
    * [*Exercise: Slices*](#exercise-slices)
  * [Maps](#maps)
    * [Map literals](#map-literals)
    * [Mutating maps](#mutating-maps)
  * [Function values](#function-values)
  * [Function closures](#function-closures)
  * [*Exercise: Fibonacci closure*](#exercise-fibonacci-closure)
<!--te-->

## Packages

Every Go program is made up of packages.

Programs start running in package `main`.

The following example is using the packages with import paths `"fmt"` and `"math/rand"`:

```golang
package main

import (
   "fmt"
   "math/rand"
)
```

By convention, the package name is the same as the last element of the import path. For instance, the `"math/rand"` package comprises files that begin with the statement `package rand`.

### Imports

The [random.go](https://github.com/lopecillo/golang/blob/master/random.go) example groups the imports into a parenthesized, "factored" import statement:

```golang
import (
   "fmt"
   "time"
   "math/rand"
)
```

You can also write multiple import statements, like:

```golang
import "fmt"
import "math"
```

But it is good style to use the factored import statement.

### Exported names

In Go, a name is exported if it begins with a capital letter. For example, `Pizza` is an exported name, as is `Pi`, which is exported from the `math` package.

`pizza` and `pi` do not start with a capital letter, so they are not exported.

When importing a package, you can refer only to its exported names. Any "unexported" names are not accessible from outside the package.

In the [math.go](https://github.com/lopecillo/golang/blob/master/math.go) example, `math.Pi` works, but `math.pi` wouldn't:

```golang
fmt.Println(math.Pi)
```

## Functions

A function can take zero or more arguments.

In this example, `add` takes two parameters of type `int`:

```golang
func add(x int, y int) int {
   return x + y
}
```

Notice that the type comes *after* the variable name.

(For more about why types look the way they do, see the [article on Go's declaration syntax](https://blog.golang.org/declaration-syntax).)

When two or more consecutive named function parameters share a type, you can omit the type from all but the last.

In the [functions.go](https://github.com/lopecillo/golang/blob/master/functions.go) example, we shortened

```golang
x int, y int
```

to

```golang
x, y int
```

### Multiple results

A function can return any number of results.

The `swap` function returns two strings:

```golang
func swap(x, y string) (string, string) {
   return y, x
}
```

### Named return values

Go's return values may be named. If so, they are treated as variables defined at the top of the function.

These names should be used to document the meaning of the return values.

A `return` statement without arguments returns the named return values. This is known as a "naked" return.

Naked return statements should be used only in short functions, as with this example:

```golang
func split(sum int) (x, y int) {
   x = sum * 4 / 9
   y = sum - x
   return
}
```

They can harm readability in longer functions.

## Variables

The `var` statement declares a list of variables; as in function argument lists, the type is last.

A `var` statement can be at package or function level. We can see both in this example:

```golang
var c, python, java bool

func main() {
   var i int
   fmt.Println(i, c, python, java)
}

```

### Initializers

A `var` declaration can include initializers, one per variable.

If an initializer is present, the type can be omitted; the variable will take the type of the initializer.

### Short variable declarations

Inside a function, the `:=` short assignment statement can be used in place of a `var` declaration with implicit type:

```golang
func main() {
   var i, j int = 1, 2
   k := 3
   c, python, java := true, false, "no!"

   fmt.Println(i, j, k, c, python, java)
}
```

Outside a function, every statement begins with a keyword (`var`, `func`, and so on) and so the `:=` construct is not available.

### Basic types

Go's basic types are

```golang
bool

string

int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr

byte // alias for uint8

rune // alias for int32
     // represents a Unicode code point

float32 float64

complex64 complex128
```

This example shows variables of several types, and also that variable declarations may be "factored" into blocks, as with import statements:

```golang
var (
   ToBe   bool       = false
   MaxInt uint64     = 1<<64 - 1
   z      complex128 = cmplx.Sqrt(-5 + 12i)
)
```

The `int`, `uint`, and `uintptr` types are usually 32 bits wide on 32-bit systems and 64 bits wide on 64-bit systems. When you need an integer value you should use `int` unless you have a specific reason to use a sized or unsigned integer type.

### Zero values

Variables declared without an explicit initial value are given their *zero* value.

The zero value is:

* `0` for numeric types
* `false` for the boolean type
* `""` (the empty string) for strings

### Type conversions

The expression `T(v)` converts the value `v` to the type `T`.

Some numeric conversions:

```golang
var i int = 42
var f float64 = float64(i)
var u uint = uint(f)
```

Or, put more simply:

```golang
i := 42
f := float64(i)
u := uint(f)
```

Unlike in C, in Go assignment between items of different type requires an explicit conversion. Try removing the `float64` or `uint` conversions in the [type-conversions.go](https://github.com/lopecillo/golang/blob/master/type-conversions.go) example and see what happens.

### Type inference

When declaring a variable without specifying an explicit type (either by using the `:=` syntax or `var =` expression syntax), the variable's type is inferred from the value on the right hand side.

When the right hand side of the declaration is typed, the new variable is of that same type:

```golang
var i int
j := i // j is an int
```

But when the right hand side contains an untyped numeric constant, the new variable may be an `int`, `float64`, or `complex128` depending on the precision of the constant:

```golang
i := 42           // int
f := 3.142        // float64
g := 0.867 + 0.5i // complex128
```

### Constants

Constants are declared like variables, but with the `const` keyword.

Constants can be character, string, boolean, or numeric values.

Constants cannot be declared using the `:=` syntax.

#### Numeric constants

Numeric constants are high-precision *values*.

An untyped constant takes the type needed by its context.

In the [numeric-constants.go](https://github.com/lopecillo/golang/blob/master/numeric-constants.go) example, try printing `needInt(Big)` too.

(An `int` can store at maximum a 64-bit integer, and sometimes less.)

## Flow control

### For

Go has only one looping construct, the `for` loop.

The basic `for` loop has three components separated by semicolons:

* the init statement: executed before the first iteration
* the condition expression: evaluated before every iteration
* the post statement: executed at the end of every iteration

```golang
sum := 0
for i := 0; i < 10; i++ {
   sum += i
}
```

The init statement will often be a short variable declaration, and the variables declared there are visible only in the scope of the `for` statement.

The loop will stop iterating once the boolean condition evaluates to `false`.

**Note:** Unlike other languages like C, Java, or JavaScript there are no parentheses surrounding the three components of the `for` statement and the braces `{ }` are always required.

The init and post statements are optional. At that point you can drop the semicolons: C's `while` is spelled `for` in Go.

```golang
sum := 1
for sum < 1000 {
   sum += sum
}
```

If you omit the loop condition it loops forever, so an infinite loop is compactly expressed.

```golang
for {
}
```

### If

Go's `if` statements are like its `for` loops; the expression need not be surrounded by parentheses `( )` but the braces `{ }` are required.

```golang
if x < 0 {
   return -x
}
```

Like `for`, the `if` statement can start with a short statement to execute before the condition:

```golang
if v := math.Pow(x, n); v < lim {
   return v
} else {
   fmt.Printf("%g >= %g\n", v, lim)
}
return lim
```

Variables declared by the statement are only in scope until the end of the `if`.

In the example above, `lim` must be used because `v` is not visible outside the scope of the `if` statement.

Variables declared inside an `if` short statement are also available inside any of the `else` blocks.

### *Exercise: Loops and functions*

As a way to play with functions and loops, let's implement a square root function: given a number x, we want to find the number z for which z² is most nearly x.

Computers typically compute the square root of x using a loop. Starting with some guess z, we can adjust z based on how close z² is to x, producing a better guess:

```golang
z -= (z*z - x) / (2*z)
```

Repeating this adjustment makes the guess better and better until we reach an answer that is as close to the actual square root as can be.

Implement this in the `func Sqrt` provided. A decent starting guess for z is 1, no matter what the input. To begin with, repeat the calculation 10 times and print each z along the way. See how close you get to the answer for various values of x (1, 2, 3, ...) and how quickly the guess improves.

Hint: To declare and initialize a floating point value, give it floating point syntax or use a conversion:

```golang
z := 1.0
z := float64(1)
```

Next, change the loop condition to stop once the value has stopped changing (or only changes by a very small amount). See if that's more or fewer than 10 iterations. Try other initial guesses for z, like x, or x/2. How close are your function's results to the [math.Sqrt](https://golang.org/pkg/math/#Sqrt) in the standard library?

(**Note:** If you are interested in the details of the algorithm, the z² − x above is how far away z² is from where it needs to be (x), and the division by 2z is the derivative of z², to scale how much we adjust z by how quickly z² is changing. This general approach is called [Newton's method](https://en.wikipedia.org/wiki/Newton%27s_method). It works well for many functions but especially well for square root.)

You can find a possible solution in [exercise-loops.go](https://github.com/lopecillo/golang/blob/master/exercise-loops.go).

### Switch

A `switch` statement is a shorter way to write a sequence of `if - else` statements. It runs the first case whose value is equal to the condition expression.

```golang
fmt.Print("Go runs on ")
switch os := runtime.GOOS; os {
case "darwin":
   fmt.Println("OS X.")
case "linux":
   fmt.Println("Linux.")
default:
   // freebsd, openbsd,
   // plan9, windows...
   fmt.Printf("%s.\n", os)
}
```

Go's switch is like the one in C, C++, Java, JavaScript, and PHP, except that Go only runs the selected case, not all the cases that follow. In effect, the `break` statement that is needed at the end of each case in those languages is provided automatically in Go. Another important difference is that Go's switch cases need not be constants, and the values involved need not be integers.

Switch cases evaluate cases from top to bottom, stopping when a case succeeds.

(For example,

```golang
switch i {
case 0:
case f():
}
```

does not call `f` if `i==0`.)

A switch without a condition is the same as `switch true`.

This construct can be a clean way to write long if-then-else chains:

```golang
t := time.Now()
switch {
case t.Hour() < 12:
   fmt.Println("Good morning!")
case t.Hour() < 17:
   fmt.Println("Good afternoon.")
default:
   fmt.Println("Good evening.")
}
```

### Defer

A `defer` statement defers the execution of a function until the surrounding function returns.

The deferred call's arguments are evaluated immediately, but the function call is not executed until the surrounding function returns:

```golang
func main() {
   defer fmt.Println("world")
   fmt.Println("hello")
}
```

Deferred function calls are pushed onto a stack. When a function returns, its deferred calls are executed in last-in-first-out order.

To learn more about defer statements read this [blog post](https://blog.golang.org/defer-panic-and-recover).

## More types

### Pointers

Go has pointers. A pointer holds the memory address of a value.

The type `*T` is a pointer to a `T` value. Its zero value is `nil`.

```golang
var p *int
```

The `&` operator generates a pointer to its operand.

```golang
i := 42
p = &i
```

The `*` operator denotes the pointer's underlying value.

```golang
fmt.Println(*p) // read i through the pointer p
*p = 21         // set i through the pointer p
```

This is known as "dereferencing" or "indirecting".

Unlike C, Go has no pointer arithmetic.

### Structs

A `struct` is a collection of fields:

```golang
type Vertex struct {
   X int
   Y int
}
```

Struct fields are accessed using a dot:

```golang
v.X = 4
```

Struct fields can be accessed through a struct pointer.

To access the field `X` of a struct when we have the struct pointer `p` we could write `(*p).X`. However, that notation is cumbersome, so the language permits us instead to write just `p.X`, without the explicit dereference:

```golang
v := Vertex{1, 2}
p := &v
p.X = 1e9
```

#### Struct literals

A struct literal denotes a newly allocated struct value by listing the values of its fields.

You can list just a subset of fields by using the `Name:` syntax. (And the order of named fields is irrelevant.)

The special prefix `&` returns a pointer to the struct value.

```golang
var (
   v1 = Vertex{1, 2}  // has type Vertex
   v2 = Vertex{X: 1}  // Y:0 is implicit
   v3 = Vertex{}      // X:0 and Y:0
   p  = &Vertex{1, 2} // has type *Vertex
)
```

### Arrays

The type `[n]T` is an array of `n` values of type `T`.

The expression

```golang
var a [10]int
```

declares a variable `a` as an array of ten integers.

An array's length is part of its type, so arrays cannot be resized. This seems limiting, but don't worry; Go provides a convenient way of working with arrays.

### Slices

An array has a fixed size. A slice, on the other hand, is a dynamically-sized, flexible view into the elements of an array. In practice, slices are much more common than arrays.

The type `[]T` is a slice with elements of type `T`.

A slice is formed by specifying two indices, a low and high bound, separated by a colon:

```golang
a[low : high]
```

This selects a half-open range which includes the first element, but excludes the last one.

The following expression creates a slice which includes elements 1 through 3 of `a`:

```golang
a[1:4]
```

A slice does not store any data, it just describes a section of an underlying array.

Changing the elements of a slice modifies the corresponding elements of its underlying array.

Other slices that share the same underlying array will see those changes.

Slices can contain any type, including other slices.

#### Slice literals

A slice literal is like an array literal without the length.

This is an array literal:

```golang
[3]bool{true, true, false}
```

And this creates the same array as above, then builds a slice that references it:

```golang
[]bool{true, true, false}
```

#### Slice defaults

When slicing, you may omit the high or low bounds to use their defaults instead.

The default is zero for the low bound and the length of the slice for the high bound.

For the array

```golang
var a [10]int
```

these slice expressions are equivalent:

```golang
a[0:10]
a[:10]
a[0:]
a[:]
```

#### Slice length and capacity

A slice has both a *length* and a *capacity*.

The length of a slice is the number of elements it contains.

The capacity of a slice is the number of elements in the underlying array, counting from the first element in the slice.

The length and capacity of a slice `s` can be obtained using the expressions `len(s)` and `cap(s)`.

You can extend a slice's length by re-slicing it, provided it has sufficient capacity. Try changing one of the slice operations in the [array.go](https://github.com/lopecillo/golang/blob/master/array.go) example to extend it beyond its capacity and see what happens.

#### Nil slices

The zero value of a slice is `nil`.

A nil slice has a length and capacity of 0 and has no underlying array.

#### Creating a slice with make

Slices can be created with the built-in `make` function; this is how you create dynamically-sized arrays.

The `make` function allocates a zeroed array and returns a slice that refers to that array:

```golang
a := make([]int, 5)  // len(a)=5
```

To specify a capacity, pass a third argument to `make`:

```golang
b := make([]int, 0, 5) // len(b)=0, cap(b)=5

b = b[:cap(b)] // len(b)=5, cap(b)=5
b = b[1:]      // len(b)=4, cap(b)=4
```

#### Appending to a slice

It is common to append new elements to a slice, and so Go provides a built-in `append` function. The [documentation](https://golang.org/pkg/builtin/#append) of the built-in package describes `append`.

```golang
func append(s []T, vs ...T) []T
```

The first parameter `s` of `append` is a slice of type `T`, and the rest are `T` values to append to the slice.

The resulting value of `append` is a slice containing all the elements of the original slice plus the provided values.

If the backing array of `s` is too small to fit all the given values a bigger array will be allocated. The returned slice will point to the newly allocated array.

To learn more about slices, read the [Slices: usage and internals](https://blog.golang.org/go-slices-usage-and-internals) article.

#### Range

The `range` form of the `for` loop iterates over a slice or map.

When ranging over a slice, two values are returned for each iteration. The first is the index, and the second is a copy of the element at that index.

```golang
var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}
for i, v := range pow {
   fmt.Printf("2**%d = %d\n", i, v)
}
```

You can skip the index or value by assigning to `_`.

```golang
for i, _ := range pow
for _, value := range pow
```

If you only want the index, you can omit the second variable.

```golang
for i := range pow
```

#### *Exercise: Slices*

Implement `Pic`. It should return a slice of length `dy`, each element of which is a slice of `dx` 8-bit unsigned integers. When you run the program, it will display your picture, interpreting the integers as grayscale (well, bluescale) values.

The choice of image is up to you. Interesting functions include `(x+y)/2`, `x*y`, and `x^y`.

(You need to use a loop to allocate each `[]uint8` inside the `[][]uint8`.)

(Use `uint8(intValue)` to convert between types.)

You can find a possible solution in [exercise-slices.go](https://github.com/lopecillo/golang/blob/master/exercise-slices.go).
You can also find the images generated for:

* [`(x+y)/2`](https://github.com/lopecillo/golang/blob/master/x_y_average.png)
* [`x*y`](https://github.com/lopecillo/golang/blob/master/x_times_y.png)
* [`x^y`](https://github.com/lopecillo/golang/blob/master/x_power_y.png)

### Maps

A map maps keys to values.

```golang
type Vertex struct {
   Lat, Long float64
}

var m map[string]Vertex
```

The zero value of a map is `nil`. A `nil` map has no keys, nor can keys be added.

The `make` function returns a map of the given type, initialized and ready for use.

```golang
m = make(map[string]Vertex)
```

#### Map literals

Map literals are like struct literals, but the keys are required.

```golang
var m = map[string]Vertex{
   "Bell Labs": Vertex{
      40.68433, -74.39967,
   },
   "Google": Vertex{
      37.42202, -122.08408,
   },
}
```

If the top-level type is just a type name, you can omit it from the elements of the literal.

```golang
var m = map[string]Vertex{
   "Bell Labs": {40.68433, -74.39967},
   "Google":    {37.42202, -122.08408},
}
```

#### Mutating maps

Insert or update an element in map `m`:

```golang
m[key] = elem
```

Retrieve an element:

```golang
elem = m[key]
```

Delete an element:

```golang
delete(m, key)
```

Test that a key is present with a two-value assignment:

```golang
elem, ok = m[key]
```

If `key` is in `m`, `ok` is `true`. If not, `ok` is `false`.

If `key` is not in the map, then `elem` is the zero value for the map's element type.

**Note:** If `elem` or `ok` have not yet been declared you could use a short declaration form:

```golang
elem, ok := m[key]
```

#### *Exercise: Maps*

Implement WordCount in [exercise-maps.go](https://github.com/lopecillo/golang/blob/master/exercise-maps.go). It should return a map of the counts of each “word” in the string `s`. The `wc.Test` function runs a test suite against the provided function and prints success or failure.

You might find [strings.Fields](https://golang.org/pkg/strings/#Fields) helpful.

You can find a possible solution in [exercise-maps.go](https://github.com/lopecillo/golang/blob/master/exercise-maps.go).

### Function values

Functions are values too. They can be passed around just like other values.

Function values may be used as function arguments and return values.

```golang
func compute(fn func(float64, float64) float64) float64 {
   return fn(3, 4)
}
```

### Function closures

Go functions may be closures. A closure is a function value that references variables from outside its body. The function may access and assign to the referenced variables; in this sense the function is "bound" to the variables.

For example, the `adder` function returns a closure. Each closure is bound to its own `sum` variable.

```golang
func adder() func(int) int {
   sum := 0
   return func(x int) int {
      sum += x
      return sum
   }
}
```

### *Exercise: Fibonacci closure*

Let's have some fun with functions.

Implement a `fibonacci` function that returns a function (a closure) that returns successive [Fibonacci numbers](https://en.wikipedia.org/wiki/Fibonacci_number) (0, 1, 1, 2, 3, 5, ...).

You can find a possible solution in [exercise-fibonacci-closure.go](https://github.com/lopecillo/golang/blob/master/exercise-fibonacci-closure.go).
