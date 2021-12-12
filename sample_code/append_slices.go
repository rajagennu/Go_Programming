package main

import "fmt"

func main() {
   var slice1 = []int{1,2,3}
   var slice2 = []int{5,6,7}
   fmt.Println(append(slice1, slice2...))
 }
