// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
)

// define a struct
type Student struct {
	name  string
	age   int
	class string
}

func (s *Student) getName() string {
	return s.name
}

func (s *Student) setName(name string) {
	s.name = name
}

func (s Student) setNameVal(name string) {
	s.name = name
}

var a int = 10
var b int = 20

func main() {
	fmt.Println("Hello, 世界")
	var a int64 = 10
	fmt.Println(a)
	b := "Hello"
	fmt.Println(b)
	r := 'R'
	fmt.Printf("%T ", r)
	fmt.Printf("%T ", b)
	arrays()
	slices()

	c := 20
	fmt.Println("C value ", c)
	passByRef(&c)
	fmt.Println("C value ", c)

	maps()
	structs()
}

// arrays
// - generic arrays - fixed size
func arrays() {
	var marks [5]int // [120,0,130,0,0]
	marks[0] = 120
	marks[2] = 130
	fmt.Println(marks)

	var names [2]string
	fmt.Println(names)

	var returns [3]float64
	fmt.Println(returns)

	// len
	fmt.Println("Array length ", len(marks))

	//

}

// - slices - dynamic arrays
func slices() {
	var marks []int // 8 byte
	fmt.Println(marks)
	// slices are pass by reference.
	// var a = []int{1,2,3,4,5} ; // 5 bytes // 5 bytes empty space - occupy -> append(10 elements ) -> 15 bytes -> new location -> 15 bytes
	var sliceOf1000 = make([]int, 1000)
	fmt.Println(len(sliceOf1000))
	sliceOf1000 = append(sliceOf1000, 10)
	fmt.Println(len(sliceOf1000))
	fmt.Println("index 999", sliceOf1000[999])
	fmt.Println("index 1000", sliceOf1000[1000])

	fmt.Println("sliceOf1000 0th val ", sliceOf1000[0])
	newSlice(sliceOf1000)
	fmt.Println("sliceOf1000 new 0th val ", sliceOf1000[0])

	newSliceWithAppend(sliceOf1000)
	fmt.Println("new length after append ", len(sliceOf1000))

	// create slice with inital length
	var sliceOf10002 = make([]int, 1000)

	// len for both arrays and slices
	fmt.Println("Length ", len(sliceOf10002))
}

func newSlice(aSlice []int) {
	aSlice[0] = 1000
}

func newSliceWithAppend(aSlice []int) {
	fmt.Printf("Address of aslice %p\n", &aSlice)
	aSlice = append(aSlice, 1000)
	fmt.Println("index 1001", aSlice[1001])
	fmt.Println("length 1001", len(aSlice))
	fmt.Printf("Address of aslice new %p", &aSlice)
}

// understand how slices works

// arrays - 10 -> memory 10 element - reserve - |a=10Byetes|b=20Bytes|C=30Bytes| => |a=100Bytes| -> <40>a<30> -> 100Bytes -> [a] -> 100Bytes -> a 110,120
// dictionaries

// maps
func maps() { // dictionary {} key, value -> {key} -> value {"name" : "Python"} dict["name"] -> "python"
	// init a map
	var marks = make(map[string]int)
	// load data
	marks["english"] = 99
	marks["maths"] = 99
	marks["physics"] = 99
	// fetch data
	fmt.Println("English marks ", marks["english"])
	// check if data exists in the map
	for key, value := range marks {
		fmt.Printf("Key %v value %v \n", key, value)
	}

	chemMarks, ok := marks["chemistry"]
	if ok {
		fmt.Println("Chemistry marks ", chemMarks)
	}
}

// structs

func structs() {
	// creating objects out of strcut
	s1 := Student{name: "S1", age: 10, class: "a"}
	fmt.Println(s1)
	// get Name
	fmt.Println("Student name ", s1.getName())
	// set Name
	s1.setName("Golang")
	fmt.Println("Student name ", s1.getName())

	// set name value

	s1.setNameVal("Golang-2")
	fmt.Println("Student name ", s1.getName())

	// creating literal structs
}

// pass by value vs pass by value
// pass by value : function -> arg -> value
// a = 10 -> printValue(a) -> copy(a) -> a = 11
// a = 10 -> printValue(*a) -> 'a' address -> a = 11
// a = 11

func passByRef(c *int) {
	fmt.Println(*c)
	*c = 50
	fmt.Println(*c)
}
