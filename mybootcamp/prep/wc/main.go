package main

import (
	"bufio" //to read text
	"bytes"
	"flag" // to manage command line options
	"fmt"  //to print format output
	"io"   // io.Reader interface
	"os"   // to use OS resources
)

func main() {
	lines := flag.Bool("l", false, "Count lines")
	bytes := flag.Bool("b", false, "Count Bytes")

	flag.Parse()
	fmt.Println(count(os.Stdin, *lines, *bytes))

}

func count(r io.Reader, countLines bool, countBytes bool) int {
	scanner := bufio.NewScanner(r) // scanner is a way to read data delimited by space or new lines.

	if countBytes {
		scanner.Split(bufio.ScanBytes)
	} else if !countLines {
		scanner.Split(bufio.ScanWords)
	}

	wc := 0
	for scanner.Scan() {
		wc++
	}
	return wc
}

// about buffer
func readBufferStr() string {
	var buffer bytes.Buffer
	buffer.WriteString("Hello, I am buffer1")
	buffer.WriteString("Hello, I am buffer2")
	// reading from buffer
	return buffer.String()
}
