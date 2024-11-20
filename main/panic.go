package main

import "fmt"

// START OMIT
func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered:", r)
		}
	}()

	fmt.Println("Before panic")
	panic("Something went wrong")
}

// END OMIT
