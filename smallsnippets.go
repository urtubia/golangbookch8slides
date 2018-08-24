package main

import "fmt"

func f(){

}

func main() {
	fmt.Println("Hi")
	// START1 OMIT
	f()   		// synchronously starts a function
	go f()		// kick a goroutine, don't wait
	// END1 OMIT
}

// START2 OMIT
func mirroredQuery() string {
	responses := make(chan string, 3)
	go func() { responses <- request("asia.com")}()
	go func() { responses <- request("europe.com")}()
	go func() { responses <- request("africa.com")}()
	return <- responses
}
// END2 OMIT
