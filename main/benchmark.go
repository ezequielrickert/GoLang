package main

import (
	"fmt"
	"time"
)

// START1 OMIT
func main() {
	startTime := time.Now()

	for i := 1; i <= 1000000; i++ {
		fmt.Println(i)
	}

	duration := time.Since(startTime)
	fmt.Println("Execution time:", duration)
}

// END1 OMIT
