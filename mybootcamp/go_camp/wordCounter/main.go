package main

import (
	"bufio" // IO Buffer read/write : NewScanner -> io.Reader , read text delimiters line, word, byte(char)
	"flag"
	"fmt"
	"io"
	"os" // for Stdin
)

func main() {
	//countWordsV1(os.Stdin)
	line := flag.Bool("l", false, "Read line")
	flag.Parse()
	fmt.Println(countWordsV2(os.Stdin, *line)) // *(&line)

}

// func countWordsV1(read io.Reader) {
// 	scanner := bufio.NewScanner(read)
// 	scanner.Split(bufio.ScanWords)
// 	for scanner.Scan() {
// 		fmt.Println("I received ", scanner.Text())

// 	}
// }

func countWordsV2(r io.Reader, countLines bool) int {
	scanner := bufio.NewScanner(r)

	if !countLines {
		scanner.Split(bufio.ScanWords)
	}
	wc := 0
	for scanner.Scan() {
		wc++
	}
	return wc
}
