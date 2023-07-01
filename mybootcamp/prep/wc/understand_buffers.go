package main

import (
	"bytes"
	"fmt"
)

func readBuffer() string {
	var x bytes.Buffer
	x.WriteString("Hello Members of Linux plant")
	x.WriteString("This is How buffer works")
	fmt.Println(x.String())
	return x.String()
}
