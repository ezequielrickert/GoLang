package main

import "fmt"

// START OMIT
type Person struct {
	Name string
	Age  int
}

func main() {
	n := new(Person)
	fmt.Printf("%+v\n", n)
	m := make([]Person, 16)
	fmt.Printf("%+v\n", m)
}

// END OMIT
