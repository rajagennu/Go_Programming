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



	