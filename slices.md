# Slices

- In Arrays while declaring, we have to define array size and it will include as part of type of the array.
- In Slices we dont have to define size while defining. It can be used for same purpose as Arrays with all its limitations addressed.

### Declare
- No need to specify the size while declaring the slices

```go
igo> var slice_1 = []int{10,20,30}
igo> slice_1
{10, 20, 30}
```

- multidimentional slices

    - Declartion : `var x [][]int`
    ```go
    igo> var x []int
    igo> x
    <nil>
    ```
    - empty slices would returl `nil` value.
    - nil has no type, so it can be assigned to any type.

### Compare slices

- Comparing slices is not possible with `==` and `!=` like we do for arrays. Only compariosn thats work with slices is whether its equals to nil or not. 
```go
x == nil
```

**Note**: package: reflect has `DeepEqual` that can compare anything , including slices.

<hr>

- `len` can be used to calculate length of slice

### append

- append function takes 2 arguments, slice and value needs to append. one more values can be append at a time. 
```go
package main

import "fmt"

func main() {
   var slice1 = []int{1,2,3}
   var slice2 = []int{5,6,7}
   fmt.Println(append(slice1, slice2...))
 }

 // Output

 ➜  sample_code git:(main) ✗ go run append_slices.go 
[1 2 3 5 6 7]
➜  sample_code git:(main) ✗  

```
- apped returns a new slice with value appended, so to save the appended value to existing slice, we can use that slice as assignment variable. 
- we can append two slices as well by using `...` operator
- go is pass by value language, it copies existing value and perform the action on that copied value and returns the same. Never touch original value. 

<hr>

### Capacity

- each element in slice assigned to consecutive memory locations, makes it easy for accessing or writing to slice.
- every slice has a capacity, which is the number of consecutive memory locations reserved. 
- when ever a new element added, the capacity will be increase by 1. 
- Remember, slice is a sequential list. if no enough capacity left sequentially then go runtime allocate a new slice and whole slice is copied to and next sequence element will be added. ((( This is going to cost good amount of memory and CPU, as we have to move entire slice duing no space cases.)))
- 



### Other useful links
[Golang Slices Intro](https://go.dev/blog/slices-intro)

