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
    * [Exercise: Loops and functions](#exercise-loops-and-functions)
    * [Switch](#switch)
    * [Defer](#defer)
<!--te-->

## Packages

Every Go program is made up of packages.

Programs start running in package `main`.

The [random.go](https://github.com/lopecillo/golang/blob/master/random.go) example is using the packages with import paths `"fmt"`, `"time"` and `"math/rand"`.

By convention, the package name is the same as the last element of the import path. For instance, the `"math/rand"` package comprises files that begin with the statement `package rand`.

### Imports

Thie [random.go](https://github.com/lopecillo/golang/blob/master/random.go) example groups the imports into a parenthesized, "factored" import statement.

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

In the [math.go](https://github.com/lopecillo/golang/blob/master/math.go) example, `math.Pi` works, but `math.pi` wouldn't.

## Functions

A function can take zero or more arguments.

In the [functions.go](https://github.com/lopecillo/golang/blob/master/functions.go) example, `add` takes two parameters of type `int`.

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

The `swap` function returns two strings.

### Named return values

Go's return values may be named. If so, they are treated as variables defined at the top of the function.

These names should be used to document the meaning of the return values.

A `return` statement without arguments returns the named return values. This is known as a "naked" return.

Naked return statements should be used only in short functions, as with the `split` function in the [functions.go](https://github.com/lopecillo/golang/blob/master/functions.go) example. They can harm readability in longer functions.

## Variables

The `var` statement declares a list of variables; as in function argument lists, the type is last.

A `var` statement can be at package or function level. We can see both in the [variables.go](https://github.com/lopecillo/golang/blob/master/variables.go) example.

### Initializers

A `var` declaration can include initializers, one per variable.

If an initializer is present, the type can be omitted; the variable will take the type of the initializer.

### Short variable declarations

Inside a function, the `:=` short assignment statement can be used in place of a `var` declaration with implicit type.

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

The [variables.go](https://github.com/lopecillo/golang/blob/master/variables.go) example shows variables of several types, and also that variable declarations may be "factored" into blocks, as with import statements.

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

The init statement will often be a short variable declaration, and the variables declared there are visible only in the scope of the `for` statement.

The loop will stop iterating once the boolean condition evaluates to `false`.

**Note:** Unlike other languages like C, Java, or JavaScript there are no parentheses surrounding the three components of the `for` statement and the braces `{ }` are always required.

The init and post statements are optional. At that point you can drop the semicolons: C's `while` is spelled `for` in Go.

If you omit the loop condition it loops forever, so an infinite loop is compactly expressed.

### If

Go's `if` statements are like its `for` loops; the expression need not be surrounded by parentheses `( )` but the braces `{ }` are required.

Like `for`, the `if` statement can start with a short statement to execute before the condition.

Variables declared by the statement are only in scope until the end of the `if`.

In the [if.go](https://github.com/lopecillo/golang/blob/master/if.go) example, `lim` must be used because `v` is not visible outside the scope of the `if` statement.

Variables declared inside an `if` short statement are also available inside any of the `else` blocks.

### Exercise: Loops and functions

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

Go's switch is like the one in C, C++, Java, JavaScript, and PHP, except that Go only runs the selected case, not all the cases that follow. In effect, the `break` statement that is needed at the end of each case in those languages is provided automatically in Go. Another important difference is that Go's switch cases need not be constants, and the values involved need not be integers.

Switch cases evaluate cases from top to bottom, stopping when a case succeeds.

(For example, in [switch.go](https://github.com/lopecillo/golang/blob/master/switch.go)

```golang
switch i {
case 0:
case f():
}
```

does not call `f` if `i==0`.)

A switch without a condition is the same as `switch true`.

This construct can be a clean way to write long if-then-else chains.

### Defer

A `defer` statement defers the execution of a function until the surrounding function returns.

The deferred call's arguments are evaluated immediately, but the function call is not executed until the surrounding function returns.

Deferred function calls are pushed onto a stack. When a function returns, its deferred calls are executed in last-in-first-out order.

To learn more about defer statements read this [blog post](https://blog.golang.org/defer-panic-and-recover).
