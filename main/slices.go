package main

func main() {
	// START OMIT
	a := []int{1, 2, 3}
	b := a
	b[1] = 5
	println(a[1])    // 5
	println(b[1])    // 5
	a = append(a, 4) // Append creates a new slice, previous slice now owned by b
	println(len(a))  // 4
	println(len(b))  // 3 <- b was not modified
	b[1] = 6
	println(a[1]) // 5
	println(b[1]) // 6
	// END OMIT
}
