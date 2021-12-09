# Primitive types & declarations

## built in types

### zero values

- all variables will having a default zero vale if they dont get assigned. zero is decided per type.

### literals

- integers
	- 0b -> binary
	- 0o -> octa
	- 0x -> hexa decimal
- float
	- decimal
	- exponential
		- +ve 6.3e3
		- -ve 6.e-3
- Strings
	- "string"
	- 'string'
	- `stringg "literal"`
- Integers can be used in Floating expressions and float variables can be used in integer expressions.
- 
### Booleans
```go
var flag bool;
var isAwesome = true;
```

### Integer types
int8 -> -128 to 127
int16 -> -32768 to 32767
int32 -> -2147483648 to 2147483647
int64 -> -9223372036854775808 to 9223372036854775807
unit8 -> 0 - 255
unit16 -> 0 - 65536
unit32 -> 0 to 4294967295
unit64 -> 0 to 18446744073709551615

Special int types
- byte
- int (32bit - int32, 64bit - int64)
- uint

### Operators 

Integer
	- Arithmetic: +, -, *, /, %
	- Integer division results int, floating division results in fload. use type conversion if different type needed
	- dont divide something by 0, it causes panic

	- conditional : ==, !=, >, <, >=, <=
	- logical : &, |, ^, &^

Floating point types
	- type: float32, float64
	- most suited for graphics and scientific operations.
	- dividing a non-zero floating with 0 returns +Inf or -Inf , by sign of the number
	- NaN : Divide a float 0 by another float 0 
	```go
	package main

	import (
		"fmt"
	)

	func main() {
		var myvar1 float64  = 0;
		var myvar2 float64 = 0 ;
		fmt.Println(myvar1/myvar2);
	}
	```
Complex Types

	- declaration `x := complex(2.5, 3.1)`
	- extraction `real(x)`
	- imaginary `img(x)`
	- package: math/cmplx
	- In future version support for complex numbers might be removed from Go
	- matrix support isnt there in the go by default.
	- 3rd party library `gonum` provides more functionalities towards Go numerical, floating and complex types.

A Taste of Strings and Runes
	- supports ==, != to find the difference
	- Ordering with >, >=, <, <=
	- concatenation with +

- Go doesnt support automatic type promotion, developer manually have to convert the type. 
```
var x int = 10
var y float64 = 20.20
var z = x + int(y)
var d = y + float64(x)
```
- No type conversion for boolean type
- No zero number and non empty number cant be called as boolean false. 
- 

### var and `:=` variable declaration

- use `var` to declare variables.
```go
var x int = 10
var x = 10
var x, y int = 10, 20
var x, y = 10, 20
var x, y = 10, "hello"
var (
	a int
	b = 20
	c = "string"
	d string
)
```
all are valid. Once type assign it cant be changed.

go also has another variable declaration format `:=`, its called as short declaration format. It can be used as equal to regular variable declaration format, but this only works in function level. In package level short declaration format is invalid and only `var` declaration format is allowed.

`:=` create the variable if its not exist. So if you are thinking that you are reusing the variable you already declared, better confirm once or always define the variables which you might wanna reuse later with `var [type] [variable_name]`

**Declare package level variable which wont change as your program executes. its hard to track changes of package level variables**


### Using const

`const` use to define immutable variables. syntax of `var` can be followed but you cant change the variable value once its defined with `const`

```go
const x int64 = 10
const (
	idKey = "id"
	nameKey = "name"
)
```

### typed and untyped constants

if you define a constat untyped, this its give you flexiability to be used multiple places with different types. 
for example

```go
const x = 10;

var y int = x
var z float64 = x
var d byte = x

```

if you define a typed constant like
```go

const x int = 10
// it can assigned to only integer type varaibles

var y int = x // valid
var z float64 = x // invalid
var p byte = x // invalid
```

### Unused variables

- unused variables will be catched and warned by Go compiler
- unused constatns will be eliminated in the binary. 
[Debugging Go in Runtime ](https://chetan177.medium.com/runtime-debugging-in-golang-b8a065d0fb5e)





	