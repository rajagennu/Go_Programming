### Make
- To create slice with specific size, creating slice with `make` is the only way

```go
x := make([]int, 5)
// a slice with length 5 and capacity 5 created, as you declared the slice length already
```
- the slice that got created above is different from a regular slice. At this moment of initialization the regular slice containes, `nil` as its value but this slice created via make will have default value as type of the slice. So if you use append now, your slice creates a new element in the sequence. 

```go
package main

import "fmt"

func main() {

	var slice1 = make([]int, 5)
	var slice2 = make([]string, 5)

	fmt.Println("Slice 1 Integer type: ", slice1)
	fmt.Println("Slice 2 String type: ", slice2)

	slice1 = append(slice1, 10)
	slice2 = append(slice2, "Hello")

	fmt.Println("Slice 1 Integer type: ", slice1)
	fmt.Println("Slice 2 String type: ", slice2)
}

'''
➜  Go_Programming git:(main) ✗ go run sample_code/slice_make.go
Slice 1 Integer type:  [0 0 0 0 0]
Slice 2 String type:  [    ]
Slice 1 Integer type:  [0 0 0 0 0 10]
Slice 2 String type:  [     Hello]
➜  Go_Programming git:(main) ✗ 
'''
```

Syntax: `var slice1 = make([]int, length_of_slice, capacity_of_slice)`

Dont give less capacity than length, always make sure capacity > length of slice else you might expect panic errors during runtime. 

### slicing slices

By use indexing we can slice the slices

var x = []int{1,2,3,4}