package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	buffer := bufio.NewScanner(os.Stdin)
	for buffer.Scan() {
		fmt.Println(buffer.Text())
	}

}
