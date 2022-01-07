package main

import "fmt"

func main() {
	var x []int
	fmt.Println(x, len(x), cap(x))
	x = append(x, 10)
	fmt.Println(x, len(x), cap(x))
	x = append(x, 20)
	fmt.Println(x, len(x), cap(x))
	x = append(x, 30)
	fmt.Println(x, len(x), cap(x)) // here capacity will be 4 instead of 3, as go add space 25% automatically for slices.
	x = append(x, 40)
	fmt.Println(x, len(x), cap(x))
}
