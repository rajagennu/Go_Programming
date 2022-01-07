// time pkg

package main

import (
	"fmt"
	"time"
)

func main() {
	start_time := time.Now()
	fmt.Println("Started ")
	time.Sleep(10 * time.Second) // sleep for 10 seconds
	end_time := time.Now()
	elapsed_time := end_time.Sub(start_time)
	fmt.Println("Elapsed time ", elapsed_time)
}
