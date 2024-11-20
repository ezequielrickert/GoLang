package main

import (
	"errors"
	"fmt"
)

// START1 OMIT
func checkValue(v int) error {
	if v < 0 {
		return errors.New("value cannot be negative")
	}
	return nil
}

func main() {
	err := checkValue(-5)
	if err != nil {
		fmt.Println("Error:", err)
	}
}

// END1 OMIT

// START2 OMIT
func checkValue2(v int) error {
	if v < 0 {
		return CustomError{}
	}
	return nil
}

type CustomError struct{}

func (c CustomError) Error() string {
	return "This is a custom error"
}

// END2 OMIT
