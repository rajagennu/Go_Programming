package main

import (
	"bytes"
	"testing"
)

func TestCountWords(t *testing.T) {
	b := bytes.NewBufferString("word1 word2 word3 word4\n")
	lines := false
	bytes := false
	exp := 4
	res := count(b, lines, bytes)
	if res != exp {
		t.Errorf("Expected %d , got %d instead \n", exp, res)
	}
}

func TestCountLines(t *testing.T) {
	b := bytes.NewBufferString("word1 word2\nword3 word4")
	lines := true
	bytes := false
	exp := 2
	res := count(b, lines, bytes)
	if res != exp {
		t.Errorf("Expected %d, got %d instead \n", exp, res)
	}
}

func TestCountBytes(t *testing.T) {
	b := bytes.NewBufferString("Word1 word2 word3 word4")
	lines := false
	bytes := true
	exp := 23
	res := count(b, lines, bytes)
	if res != exp {
		t.Errorf("Expected %d got %d instead \n", exp, res)
	}
}

func TestReadBufferFunc(t *testing.T) {
	output := readBufferStr()
	expected := "Hello, I am buffer1Hello, I am buffer2"
	if output != expected {
		t.Errorf("Expected %s , instead received %s", expected, output)
	}
}
