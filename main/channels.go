package main

import "fmt"

// START1 OMIT
func main() {
	ch := make(chan int) // <- Create a channel of type int
	go func() {
		var sum = 0
		for i := 0; i < 10; i++ {
			sum += i
		}
		ch <- sum // <- Send the sum to the channel
	}()
	output := <-ch // <- Receive the sum from the channel
	fmt.Println(output)
}

// END1 OMIT
