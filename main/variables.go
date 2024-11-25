package main

// START1 OMIT
import "fmt"

// Global variables must be declared using `var`
var globalMessage string
var globalCount int
var assignedGlobalCount = 42 // <- Type can be omitted

func main() {
	globalMessage = "Hello from the global scope!" // <- Global variables can be assigned anywhere
	globalCount = 42
	//globalCount = "hello" // <- This will not compile, non-dynamic typing
	fmt.Println(globalMessage)
	fmt.Println("Global count is:", globalCount)
}

// END1 OMIT

// START2 OMIT
func inline() {
	// Local variables can be declared using `:=`
	localMessage := "Hello from the local scope!"
	localCount := 42
	fmt.Println(localMessage)
	fmt.Println("Local count is:", localCount)
}

// END2 OMIT
