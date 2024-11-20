package main

import "fmt"

func main() {
	value()
	fmt.Println("----------")
}

// START1 OMIT
func value() [3]int {
	a := [3]int{1, 2, 3} // Array of 3 integers
	b := a               // Copy `a` to `b`
	b[0] = 10            // Modify `b`, does not affect `a`

	fmt.Println("a:", a) // Output: a: [1 2 3]
	fmt.Println("b:", b) // Output: b: [10 2 3]
	return a
	// return [4]int{1, 2, 3, 4} // Error!!
}

// END1 OMIT
