package main

func main() {
	// START1 OMIT
	for i := 0; i < 10; i++ {
		println(i)
	}
	// END1 OMIT

	// START2 OMIT
	i := 0
	for i < 10 {
		println(i)
		i++
	}
	// END2 OMIT

	// START3 OMIT
	i = 0
	for {
		if i >= 10 {
			break
		}
		println(i)
		i++
	}
	// END3 OMIT
}
