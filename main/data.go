package main

import "fmt"

// START OMIT
func main() {
	n := new(map[string]int)
	fmt.Println(n)
	m := make(map[string]int)
	fmt.Println(m)
	m["one"] = 1 // <- Works fine, initialized
	fmt.Println(m)
	(*n)["one"] = 1 // <- This will panic, not initialized
}

// END OMIT
