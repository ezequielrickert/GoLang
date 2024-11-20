package main

import "fmt"

func main() {
	typeSwitch()
}

func shouldEscape(c byte) bool {
	// START1 OMIT
	switch c {
	case ' ', '?', '&', '=', '#', '+', '%': // <- Multiple values in a single case
		return true
	}
	return false
	// END1 OMIT
}

var t interface{}

func typeSwitch() {
	t = shouldEscape(' ')
	// START2 OMIT
	switch t := t.(type) {
	default:
		fmt.Printf("unexpected type %T\n", t) // %T prints whatever type t has
	case bool:
		fmt.Printf("boolean %t\n", t) // t has type bool
	case int:
		fmt.Printf("integer %d\n", t) // t has type int
	case *bool:
		fmt.Printf("pointer to boolean %t\n", *t) // t has type *bool
	case *int:
		fmt.Printf("pointer to integer %d\n", *t) // t has type *int
	}
	// END2 OMIT
}
