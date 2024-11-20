package main

import "fmt"

// START OMIT
type Shape interface {
	Area() float64
	//Perimeter() float64 <-- Uncomment this line to see the error, as Circle does not implement this method
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

func main() {
	var s Shape = Circle{Radius: 5} // This compiles fine :)
	fmt.Println("Area:", s.Area())
}

// END OMIT
