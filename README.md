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
* [Methods](#methods)
  * [Pointer receivers](#pointer-receivers)
* [Interfaces](#interfaces)
  * [Interface values](#interface-values)
  * [Interface values with nil underlying values](#interface-values-with-nil-underlying-values)
  * [Nil interface values](#nil-interface-values)
  * [The empty interface values](#the-empty-interface)
  * [Type assertions](#type-assertions)
  * [Type switches](#type-switches)
  * [Stringers](#stringers)
    * [*Exercise: Stringers*](#exercise-stringers)
* [Errors](#errors)
* [Concurrency](#concurrency)
  * [Goroutines](#goroutines)
  * [Channels](#channels)
    * [Buffered channels](#buffered-channels)
    * [Range and close](#range-and-close)
  * [Select](#select)
    * [Default selection](#default-selection)
  * [*Exercise: Equivalent binary trees*](#exercise-equivalent-binary-trees)
  * [sync.Mutex](#sync-mutex)
  * [*Exercise: Web Crawler*](#exercise-web-crawler)
* [Where to *Go* from here...](#where-to-go-from-here)
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

#### *Exercise: Fibonacci closure*

Let's have some fun with functions.

Implement a `fibonacci` function that returns a function (a closure) that returns successive [Fibonacci numbers](https://en.wikipedia.org/wiki/Fibonacci_number) (0, 1, 1, 2, 3, 5, ...).

You can find a possible solution in [exercise-fibonacci-closure.go](https://github.com/lopecillo/golang/blob/master/exercise-fibonacci-closure.go).

## Methods

Go does not have classes. However, you can define methods on types.

A method is a function with a special *receiver* argument.

The receiver appears in its own argument list between the `func` keyword and the method name.

In this example, the `Abs` method has a receiver of type `Vertex` named `v`.

```golang
type Vertex struct {
   X, Y float64
}

func (v Vertex) Abs() float64 {
   return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
```

Remember: a method is just a function with a receiver argument.

Here's `Abs` written as a regular function with no change in functionality.

```golang
func Abs(v Vertex) float64 {
   return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
```

You can declare a method on non-struct types, too.

In this example we see a numeric type `MyFloat` with an `Abs` method.

```golang
type MyFloat float64

func (f MyFloat) Abs() float64 {
   if f < 0 {
      return float64(-f)
   }
      return float64(f)
}
```

You can only declare a method with a receiver whose type is defined in the same package as the method. You cannot declare a method with a receiver whose type is defined in another package (which includes the built-in types such as `int`).

### Pointer receivers

You can declare methods with pointer receivers.

This means the receiver type has the literal syntax `*T` for some type `T`. (Also, `T` cannot itself be a pointer such as `*int`.)

For example, the `Scale` method here is defined on `*Vertex`.

```golang
type Vertex struct {
   X, Y float64
}

func (v Vertex) Abs() float64 {
   return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vertex) Scale(f float64) {
   v.X = v.X * f
   v.Y = v.Y * f
}
```

Methods with pointer receivers can modify the value to which the receiver points (as `Scale` does here). Since methods often need to modify their receiver, pointer receivers are more common than value receivers.

Try removing the `*` from the declaration of the `Scale` function on [methods.go](https://github.com/lopecillo/golang/blob/master/methods.go) and observe how the program's behavior changes.

With a value receiver, the `Scale` method operates on a copy of the original `Vertex` value. (This is the same behavior as for any other function argument.) The `Scale` method must have a pointer receiver to change the Vertex value declared in the `main` function.

Here we see the `Abs` and `Scale` methods rewritten as functions.

```golang
type Vertex struct {
   X, Y float64
}

func Abs(v Vertex) float64 {
   return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func Scale(v *Vertex, f float64) {
   v.X = v.X * f
   v.Y = v.Y * f
}
```

In this example the `*` from `func Scale(v *Vertex, f float64)` is necessary for the function to access and modify the original `Vertex` value.

Comparing the previous two programs, you might notice that functions with a pointer argument must take a pointer:

```golang
var v Vertex
ScaleFunc(v, 5)  // Compile error!
ScaleFunc(&v, 5) // OK
```

while methods with pointer receivers take either a value or a pointer as the receiver when they are called:

```golang
var v Vertex
v.Scale(5)  // OK
p := &v
p.Scale(10) // OK
```

For the statement `v.Scale(5)`, even though `v` is a value and not a pointer, the method with the pointer receiver is called automatically. That is, as a convenience, Go interprets the statement `v.Scale(5)` as `(&v).Scale(5)` since the `Scale` method has a pointer receiver.

The equivalent thing happens in the reverse direction.

Functions that take a value argument must take a value of that specific type:

```golang
var v Vertex
fmt.Println(AbsFunc(v))  // OK
fmt.Println(AbsFunc(&v)) // Compile error!
```

while methods with value receivers take either a value or a pointer as the receiver when they are called:

```golang
var v Vertex
fmt.Println(v.Abs()) // OK
p := &v
fmt.Println(p.Abs()) // OK
```

In this case, the method call `p.Abs()` is interpreted as `(*p).Abs()`.

There are two reasons to use a pointer receiver.

The first is so that the method can modify the value that its receiver points to.

The second is to avoid copying the value on each method call. This can be more efficient if the receiver is a large struct, for example.

In this example, both `Scale` and `Abs` are with receiver type `*Vertex`, even though the `Abs` method needn't modify its receiver.

```golang
func (v *Vertex) Scale(f float64) {
   v.X = v.X * f
   v.Y = v.Y * f
}

func (v *Vertex) Abs() float64 {
   return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
```

In general, all methods on a given type should have either value or pointer receivers, but not a mixture of both. (We'll see why over the next few pages.)

## Interfaces

An *interface type* is defined as a set of method signatures.

A value of interface type can hold any value that implements those methods.

Note that in the following example Vertex (the value type) doesn't implement Abser because the Abs method is defined only on *Vertex (the pointer type).

```golang
type Abser interface {
   Abs() float64
}

type Vertex struct {
   X, Y float64
}

func (v *Vertex) Abs() float64 {
   return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
```

A type implements an interface by implementing its methods. There is no explicit declaration of intent, no "implements" keyword.

```golang
import "fmt"

type I interface {
   M()
}

type T struct {
   S string
}

// This method means type T implements the interface I,
// but we don't need to explicitly declare that it does so.
func (t T) M() {
   fmt.Println(t.S)
}

func main() {
   var i I = T{"hello"}
   i.M()
}
```

Implicit interfaces decouple the definition of an interface from its implementation, which could then appear in any package without prearrangement.

### Interface values

Under the hood, interface values can be thought of as a tuple of a value and a concrete type:

```golang
(value, type)
```

An interface value holds a value of a specific underlying concrete type.

Calling a method on an interface value executes the method of the same name on its underlying type.

```golang
func main() {
   var i I

   i = &T{"Hello"}
   describe(i)
   i.M()

   i = F(math.Pi)
   describe(i)
   i.M()
}

func describe(i I) {
   fmt.Printf("(%v, %T)\n", i, i)
}
```

### Interface values with nil underlying values

If the concrete value inside the interface itself is nil, the method will be called with a nil receiver.

In some languages this would trigger a null pointer exception, but in Go it is common to write methods that gracefully handle being called with a nil receiver (as with the method `M` in this example.)

```golang
func main() {
   var i I

   var t *T
   i = t
   describe(i)
   i.M()

   i = &T{"hello"}
   describe(i)
   i.M()
}
```

Note that an interface value that holds a nil concrete value is itself non-nil.

### Nil interface values

A nil interface value holds neither value nor concrete type.

Calling a method on a nil interface is a run-time error because there is no type inside the interface tuple to indicate which `concrete` method to call.

### The empty interface

The interface type that specifies zero methods is known as the *empty interface*:

```golang
interface{}
```

An empty interface may hold values of any type. (Every type implements at least zero methods.)

Empty interfaces are used by code that handles values of unknown type. For example, `fmt.Print` takes any number of arguments of type `interface{}`.

### Type assertions

A type assertion provides access to an interface value's underlying concrete value.

```golang
t := i.(T)
```

This statement asserts that the interface value `i` holds the concrete type `T` and assigns the underlying `T` value to the variable `t`.

If `i` does not hold a `T`, the statement will trigger a panic.

To test whether an interface value holds a specific type, a type assertion can return two values: the underlying value and a boolean value that reports whether the assertion succeeded.

```golang
t, ok := i.(T)
```

If `i` holds a `T`, then `t` will be the underlying value and `ok` will be true.

If not, `ok` will be false and `t` will be the zero value of type `T`, and no panic occurs.

Note the similarity between this syntax and that of reading from a map.

### Type switches

A *type switch* is a construct that permits several type assertions in series.

A type switch is like a regular switch statement, but the cases in a type switch specify types (not values), and those values are compared against the type of the value held by the given interface value.

```golang
switch v := i.(type) {
case T:
   // here v has type T
case S:
   // here v has type S
default:
   // no match; here v has the same type as i
}
```

The declaration in a type switch has the same syntax as a type assertion `i.(T)`, but the specific type `T` is replaced with the keyword `type`.

This switch statement tests whether the interface value `i` holds a value of type `T` or `S`. In each of the `T` and `S` cases, the variable `v` will be of type `T` or `S` respectively and hold the value held by `i`. In the default case (where there is no match), the variable `v` is of the same interface type and value as `i`.

### Stringers

One of the most ubiquitous interfaces is `Stringer` defined by the `fmt` package.

```golang
type Stringer interface {
   String() string
}
```

A `Stringer` is a type that can describe itself as a string. The `fmt` package (and many others) look for this interface to print values.

#### Exercise: Stringers

Make the `IPAddr` type implement `fmt.Stringer` to print the address as a dotted quad.

For instance, `IPAddr{1, 2, 3, 4}` should print as `"1.2.3.4"`.

You can find a possible solution in [exercise-stringer.go](https://github.com/lopecillo/golang/blob/master/exercise-stringer.go).

## Errors

Go programs express error state with `error` values.

The `error` type is a built-in interface similar to `fmt.Stringer`:

```golang
type error interface {
   Error() string
}
```

(As with `fmt.Stringer`, the `fmt` package looks for the `error` interface when printing values.)

Functions often return an `error` value, and calling code should handle errors by testing whether the error equals `nil`.

```golang
i, err := strconv.Atoi("42")
if err != nil {
   fmt.Printf("couldn't convert number: %v\n", err)
   return
}
fmt.Println("Converted integer:", i)
```

A `nil` error denotes success; a non-nil `error` denotes failure.

## Concurrency

### Goroutines

A *goroutine* is a lightweight thread managed by the Go runtime.

```golang
go f(x, y, z)
```

starts a new goroutine running

```golang
f(x, y, z)
```

The evaluation of `f`, `x`, `y`, and `z` happens in the current goroutine and the execution of `f` happens in the new goroutine.

Goroutines run in the same address space, so access to shared memory must be synchronized. The `sync` package provides useful primitives, although you won't need them much in Go as there are other primitives.

### Channels

Channels are a typed conduit through which you can send and receive values with the channel operator, `<-`.

```golang
ch <- v    // Send v to channel ch.
v := <-ch  // Receive from ch, and
           // assign value to v.
```

(The data flows in the direction of the arrow.)

Like maps and slices, channels must be created before use:

```golang
ch := make(chan int)
```

By default, sends and receives block until the other side is ready. This allows goroutines to synchronize without explicit locks or condition variables.

The example code sums the numbers in a slice, distributing the work between two goroutines. Once both goroutines have completed their computation, it calculates the final result.

```golang
func sum(s []int, c chan int) {
   sum := 0
   for _, v := range s {
      sum += v
   }
   c <- sum // send sum to c
}

func main() {
   s := []int{7, 2, 8, -9, 4, 0}

   c := make(chan int)
   go sum(s[:len(s)/2], c)
   go sum(s[len(s)/2:], c)
   x, y := <-c, <-c // receive from c

   fmt.Println(x, y, x+y)
}
```

#### Buffered channels

Channels can be *buffered*. Provide the buffer length as the second argument to `make` to initialize a buffered channel:

```golang
ch := make(chan int, 100)
```

Sends to a buffered channel block only when the buffer is full. Receives block when the buffer is empty.

Trying to overfill the buffer will trigger a fatal error (deadlock).

#### Range and close

A sender can `close` a channel to indicate that no more values will be sent. Receivers can test whether a channel has been closed by assigning a second parameter to the receive expression: after

```golang
v, ok := <-ch
```

`ok` is `false` if there are no more values to receive and the channel is closed.

The loop `for i := range c` receives values from the channel repeatedly until it is closed.

**Note:** Only the sender should close a channel, never the receiver. Sending on a closed channel will cause a panic.

**Another note:** Channels aren't like files; you don't usually need to close them. Closing is only necessary when the receiver must be told there are no more values coming, such as to terminate a `range` loop.

### Select

The `select` statement lets a goroutine wait on multiple communication operations.

A `select` blocks until one of its cases can run, then it executes that case. It chooses one at random if multiple are ready.

#### Default selection

The `default` case in a `select` is run if no other case is ready.

Use a `default` case to try a send or receive without blocking:

```golang
select {
case i := <-c:
   // use i
default:
   // receiving from c would block
}
```

### *Exercise: Equivalent binary trees*

There can be many different binary trees with the same sequence of values stored in it. For example, here are two binary trees storing the sequence 1, 1, 2, 3, 5, 8, 13.

[Example binary trees](https://tour.golang.org/content/img/tree.png)

A function to check whether two binary trees store the same sequence is quite complex in most languages. We'll use Go's concurrency and channels to write a simple solution.

This example uses the `tree` package, which defines the type:

```golang
type Tree struct {
   Left  *Tree
   Value int
   Right *Tree
}
```

In the following piece of code:

```golang
package main

import "golang.org/x/tour/tree"

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int)

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool

func main() {
}
```

1. Implement the `Walk` function.

2. Test the `Walk` function.

The function `tree.New(k)` constructs a randomly-structured (but always sorted) binary tree holding the values `k`, `2k`, `3k`, ..., `10k`.

Create a new channel `ch` and kick off the walker:

```golang
go Walk(tree.New(1), ch)
```

Then read and print 10 values from the channel. It should be the numbers 1, 2, 3, ..., 10.

3. Implement the `Same` function using `Walk` to determine whether `t1` and `t2` store the same values.

4. Test the `Same` function.

`Same(tree.New(1), tree.New(1))` should return `true`, and `Same(tree.New(1), tree.New(2))` should return `false`.

The documentation for `Tree` can be found [here](https://godoc.org/golang.org/x/tour/tree#Tree).

You can find a possible solution in [exercise-equivalent-binary-trees.go](https://github.com/lopecillo/golang/blob/master/exercise-equivalent-binary-trees.go).

### sync.Mutex

We've seen how channels are great for communication among goroutines.

But what if we don't need communication? What if we just want to make sure only one goroutine can access a variable at a time to avoid conflicts?

This concept is called *mutual exclusion*, and the conventional name for the data structure that provides it is *mutex*.

Go's standard library provides mutual exclusion with `sync.Mutex` and its two methods:

```golang
   Lock
   Unlock
```

We can define a block of code to be executed in mutual exclusion by surrounding it with a call to `Lock` and `Unlock` as shown on the `Inc` method.

We can also use `defer` to ensure the mutex will be unlocked as in the `Value` method.

```golang
package main

import (
   "fmt"
   "sync"
   "time"
)

// SafeCounter is safe to use concurrently.
type SafeCounter struct {
   v   map[string]int
   mux sync.Mutex
}

// Inc increments the counter for the given key.
func (c *SafeCounter) Inc(key string) {
   c.mux.Lock()
   // Lock so only one goroutine at a time can access the map c.v.
   c.v[key]++
   c.mux.Unlock()
}

// Value returns the current value of the counter for the given key.
func (c *SafeCounter) Value(key string) int {
   c.mux.Lock()
   // Lock so only one goroutine at a time can access the map c.v.
   defer c.mux.Unlock()
   return c.v[key]
}

func main() {
   c := SafeCounter{v: make(map[string]int)}
   for i := 0; i < 1000; i++ {
      go c.Inc("somekey")
   }

   time.Sleep(time.Second)
   fmt.Println(c.Value("somekey"))
}
```

### *Exercise: Web Crawler*

In this exercise you'll use Go's concurrency features to parallelize a web crawler.

Modify the `Crawl` function to fetch URLs in parallel without fetching the same URL twice.

*Hint:* you can keep a cache of the URLs that have been fetched on a map, but maps alone are not safe for concurrent use!

```golang
package main

import (
   "fmt"
)

type Fetcher interface {
   // Fetch returns the body of URL and
   // a slice of URLs found on that page.
   Fetch(url string) (body string, urls []string, err error)
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher) {
   // TODO: Fetch URLs in parallel.
   // TODO: Don't fetch the same URL twice.
   // This implementation doesn't do either:
   if depth <= 0 {
      return
   }
   body, urls, err := fetcher.Fetch(url)
   if err != nil {
      fmt.Println(err)
      return
   }
   fmt.Printf("found: %s %q\n", url, body)
   for _, u := range urls {
      Crawl(u, depth-1, fetcher)
   }
   return
}

func main() {
   Crawl("https://golang.org/", 4, fetcher)
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
   body string
   urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
   if res, ok := f[url]; ok {
      return res.body, res.urls, nil
   }
   return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
   "https://golang.org/": &fakeResult{
      "The Go Programming Language",
      []string{
         "https://golang.org/pkg/",
         "https://golang.org/cmd/",
      },
   },
   "https://golang.org/pkg/": &fakeResult{
      "Packages",
      []string{
         "https://golang.org/",
         "https://golang.org/cmd/",
         "https://golang.org/pkg/fmt/",
         "https://golang.org/pkg/os/",
      },
   },
   "https://golang.org/pkg/fmt/": &fakeResult{
      "Package fmt",
      []string{
         "https://golang.org/",
         "https://golang.org/pkg/",
      },
   },
   "https://golang.org/pkg/os/": &fakeResult{
      "Package os",
      []string{
         "https://golang.org/",
         "https://golang.org/pkg/",
      },
   },
}
```

You can find a possible solution in [exercise-web-crawler.go](https://github.com/lopecillo/golang/blob/master/exercise-web-crawler.go).

## Where to *Go* from here...

The [Go Documentation](https://golang.org/doc/) is a great place to continue. It contains references, tutorials, videos, and more.

To learn how to organize and work with Go code, read [How to Write Go Code](https://golang.org/doc/code.html).

If you need help with the standard library, see the [package reference](https://golang.org/pkg/). For help with the language itself, you might be surprised to find the [Language Spec](https://golang.org/ref/spec) is quite readable.

To further explore Go's concurrency model, watch [Go Concurrency Patterns](https://www.youtube.com/watch?v=f6kdp27TYZs) ([slides](https://talks.golang.org/2012/concurrency.slide)) and [Advanced Go Concurrency Patterns](https://www.youtube.com/watch?v=QDDwwePbDtw) ([slides](https://talks.golang.org/2013/advconc.slide)) and read the [Share Memory by Communicating](https://golang.org/doc/codewalk/sharemem/) codewalk.

To get started writing web applications, watch [A simple programming environment](https://vimeo.com/53221558) ([slides](https://talks.golang.org/2012/simple.slide)) and read the [Writing Web Applications](https://golang.org/doc/articles/wiki/) tutorial.

The [First Class Functions in Go](https://golang.org/doc/codewalk/functions/) codewalk gives an interesting perspective on Go's function types.

The [Go Blog](https://blog.golang.org/) has a large archive of informative Go articles.

Visit [golang.org](https://golang.org/) for more.
