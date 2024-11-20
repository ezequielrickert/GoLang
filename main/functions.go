package main

func main() {

}

// START1 OMIT
func basicFunction() {
	println("Hello, World!")
}

// END1 OMIT

// START2 OMIT
func nextInt(b []byte, i int) (int, int) {
	for ; i < len(b) && !isDigit(b[i]); i++ {
	}
	x := 0
	for ; i < len(b) && isDigit(b[i]); i++ {
		x = x*10 + int(b[i]) - '0'
	}
	return x, i
}

// END2 OMIT

// START3 OMIT
func ReadFull(r Reader, buf []byte) (n int, err error) {
	for len(buf) > 0 && err == nil {
		var nr int
		nr, err = r.Read(buf)
		n += nr
		buf = buf[nr:]
	}
	return
}

// END3 OMIT
