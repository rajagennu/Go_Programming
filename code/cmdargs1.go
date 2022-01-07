package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	var s, sep string
	start_time := time.Now()
	for i := 1; i < len(os.Args); i++ {
		s = s + sep + os.Args[i]
		sep = " "
	}
	end_time := time.Now()
	fmt.Println("F1 Elapsed time ", end_time.Sub(start_time))
	fmt.Println(s)
	v2()
	v3()
	v4()
	ex1_print_arg0()
	ex2_print_index_value()
}

func v2() {
	s, sep := "", ""
	start_time := time.Now()
	for _, args := range os.Args[1:] {
		s += sep + args
		sep = " "
	}
	end_time := time.Now()
	fmt.Println("F2 Elapsed time ", end_time.Sub(start_time))
	fmt.Println(s)
}

func v3() {
	start_time := time.Now()
	fmt.Println(strings.Join(os.Args[1:], " "))
	end_time := time.Now()
	fmt.Println("F3 Elapsed time ", end_time.Sub(start_time))

}

func v4() {
	fmt.Println(os.Args[1:])

}

func ex1_print_arg0() {
	fmt.Println("name of the command invoked this ", os.Args[0])
	fmt.Println("And args are", strings.Join(os.Args[1:], " "))
}

func ex2_print_index_value() {
	for index, value := range os.Args[1:] {
		fmt.Println("Index ", index, " Value ", value)
	}
}
