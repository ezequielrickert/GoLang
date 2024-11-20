package main

import "fmt"

// START OMIT
type Rectangle struct {
	Width, Height float64
}

func main() {
	r := Rectangle{10, 20}
	fmt.Printf("%+v\n", r)
	r.Redimension() // should not change the original shape
	fmt.Printf("%+v\n", r)
	r.RedimensionPointer() // should change the original shape
	fmt.Printf("%+v\n", r)
}

func (r Rectangle) Redimension() {
	r.Width = r.Width * 2
	r.Height = r.Height * 2
}

func (r *Rectangle) RedimensionPointer() {
	r.Width = r.Width * 2
	r.Height = r.Height * 2
}

// END OMIT
