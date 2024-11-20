package main

import "fmt"

func main() {
	// START1 OMIT
	for i := 0; i < 5; i++ {
		defer fmt.Printf("%d ", i)
	}
	// END1 OMIT
	defer fmt.Println("----------------------")
	a()
}

// START2 OMIT
func trace(s string)   { fmt.Println("entering:", s) }
func untrace(s string) { fmt.Println("leaving:", s) }

// Use them like this:
func a() {
	trace("a")
	defer untrace("a")
	fmt.Println("in a")
}

// END2 OMIT
