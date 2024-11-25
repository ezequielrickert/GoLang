package main

import (
	"fmt"
	"sync"
)

// START1 OMIT
func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done() // <- Practical use of defer,
		// decrementing the counter when the goroutine is done
		for i := 0; i < 10; i++ {
			fmt.Println(i)
		}
	}()
	wg.Wait() // <- Wait for the goroutine to finish
	// otherwise the program will exit before the goroutine is done
	fmt.Println("Goroutine finished")
}

// END1 OMIT
