package main

func main() {
	// START OMIT
	a := []int{1, 2, 3}
	b := a
	b[1] = 5
	println(a[1])    // 5
	println(b[1])    // 5
	a = append(a, 4) // Appends creates a new slice
	println(len(a))  // 4 <- Slice is dynamic
	println(len(b))  // 3 <- b is a copy of a
	b[1] = 6
	println(a[1]) // 5
	println(b[1]) // 6
	// END OMIT
}
