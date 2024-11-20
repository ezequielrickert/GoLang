package main

import (
	"fmt"
	"os"
)

func main() {
	// START1 OMIT
	if x := 10; x > 5 {
		fmt.Println("x is greater than 5")
	}
	// END1 OMIT

	// START2 OMIT
	name := "invalid file"
	f, err := os.Open(name)
	if err != nil {
		fmt.Println("Error opening file:", err)
	}
	d, err := f.Stat()
	if err != nil {
		f.Close()
		fmt.Println("Error getting file info:", err)
	}
	fmt.Println(d)
	// END2 OMIT
}
