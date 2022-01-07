package main

import "fmt"

func main() {
	var slice1 = []int{1, 2, 3, 4, 5}
	fmt.Println(slice1) // [1 2 3 4 5]
	var new_slice = make([]int, 2)
	var copied_num = copy(new_slice, slice1)
	fmt.Println("copied nummebrs", copied_num)
	fmt.Println("new slice ", new_slice)

	/*
		[1 2 3 4 5]
		copied nummebrs 2
		new slice  [1 2]
	*/

	// slice share memory
	var new_slice2 = []int{}
	new_slice2 = slice1[0:2]
	new_slice2 = append(new_slice2, 10)
	fmt.Println("New slice post append ", new_slice2)
	fmt.Println("Orig slice post append ", slice1)
	/*
		New slice post append  [1 2 10]
		Orig slice post append  [1 2 10 4 5]
		check the position 3 index 2 value 10 inserted. Because in Go child slcies the same memory as parent slices
		to avoid this you have to use make while creating the chid
	*/

}
