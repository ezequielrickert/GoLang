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
	m := make([]Person, 2, 4) // <- 2 more elements can be added
	fmt.Printf("%+v\n", m)
	// Less expensive append, declared capacity is not exceeded
	a := append(m, Person{"Alice", 25})
	fmt.Printf("%+v\n", m)
	fmt.Printf("%+v\n", a)
}

// END OMIT
