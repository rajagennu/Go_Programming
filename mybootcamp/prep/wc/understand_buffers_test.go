package main

import (
	"testing"
)

func TestReadBuffer(t *testing.T) {
	output := readBuffer()
	//fmt.Println(output)
	expected := "Hello Members of Linux plantThis is How buffer works"
	if output != expected {
		t.Errorf("Expected %s , received %s", expected, output)
	}
}
