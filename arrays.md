### Arrays
- all the elements must be of the type that specified. 
- reading elements out of the index will create a panic
- access elements via negative index is not supported.
- Go arrays are rarely used explicitly, this is because, Go considers the size of the array to be part of the type of the array. this makes an array thats delcared to be [3]int a different type from an array thats delcared as [4]int.
- There is no type conversion to convert arrays of different sizes to identical types. 
- you cant write a function that works with arrays of any size and you cant assign arrays of different sizes to the same variable.

```go
igo> var x[3]int
igo> x
{0, 0, 0}
igo> 

igo> var y = [3]int{10,20,30}
igo> y
{10, 20, 30}
```
- sparse array: an array where most elements are set to their zero value

```go
igo> var z = [12]int{1, 5: 4, 6, 10: 100, 15}
igo> z
{1, 4, 6, 100, 15, 0, 0, 0, 0, 0, 0, 0}
igo> 
```

- we can also declare an array using

```go

var x = [...]int{10,20,30}
```

### Array Comparision 

- use == and != operators.

### 2-D Array

```go
igo> var x [2][3]int
igo> x
{{0, 0, 0}, {0, 0, 0}}
igo> x[0]
{0, 0, 0}
igo> x[0][0]
0
igo> 
```
### Built-in Functions

- To calculate length : `len(x)`

### Other links for more information
[2-D Arrays and Slices](https://golangbyexample.com/two-dimensional-array-slice-golang/)
[]




