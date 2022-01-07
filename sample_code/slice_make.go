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
