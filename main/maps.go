package main

// START OMIT
var timeZone = map[string]int{
	"UTC": 0 * 60 * 60,
	"EST": -5 * 60 * 60,
	"CST": -6 * 60 * 60,
	"MST": -7 * 60 * 60,
	"PST": -8 * 60 * 60,
}

func main() {
	seconds, ok := timeZone["UTC"]
	println(seconds, ok) // 0 true
	seconds, ok = timeZone["Invalid"]
	println(seconds, ok) // 0 false
}

// END OMIT
