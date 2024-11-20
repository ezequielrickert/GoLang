package main

import "fmt"

// ExportedName START1 OMIT
var ExportedName = "Visible outside the package"       // HL
var unexportedName = "Visible only within the package" // HL

func main() {
	fmt.Println(ExportedName)
}

//END1 OMIT

// Reader START2 OMIT
type Reader interface {
	Read() int
}

//END2 OMIT
