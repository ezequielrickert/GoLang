package main

import "fmt"

func main() {
	// START1 OMIT
	fmt.Println("Hello, Go!") // Valid
	//Incorrect: Cannot have two statements on the same line
	//fmt.Println("Hello, Go!") fmt.Println("Hello, Go!")
	//fmt.Println("Another line"  <- Incorrect: Auto-semicolon cannot resolve this
	//END1 OMIT

	// START2 OMIT
	//if i < f()  // wrong!
	//{           // wrong!
	//	g()
	//}
	//END2 OMIT
}
