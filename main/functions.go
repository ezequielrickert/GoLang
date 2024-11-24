package main

import "fmt"

func main() {

}

// START1 OMIT
func basicFunction() {
	println("Hello, World!")
}

// END1 OMIT

// START2 OMIT
// Function to calculate the quotient and remainder of a division
func divide(dividend int, divisor int) (int, int, error) {
	if divisor == 0 {
		return 0, 0, fmt.Errorf("division by zero")
	}
	quotient := dividend / divisor
	remainder := dividend % divisor
	return quotient, remainder, nil
}

// END2 OMIT

// START3 OMIT
func namedDivide(dividend int, divisor int) (quotient int, remainder int) {
	if divisor == 0 {
		return 0, 0
	}
	quotient = dividend / divisor
	remainder = dividend % divisor
	return
}

// END3 OMIT
