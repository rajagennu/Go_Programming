package main

import (
	"fmt"
)

func main() {
	fmt.Println("For loop init variable is return value of the function")
	for i := init_func(); i < 5; i++ {
		fmt.Println("i value ", i)
	}
	forv2()
}

func init_func() int {
	return 1
}

func forv2() {

	var a = [5]int{1, 2, 3, 4, 5}
	for _, value := range a {
		fmt.Println("Value : ", value)
	}
}
