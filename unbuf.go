package main

import (
	"time"
	"fmt"
)
// START OMIT
func addNumbersSlowly(done chan int, a, b int){
	time.Sleep(time.Duration(2*time.Second))
	done <- a+b
}

func main(){
	doneChannel := make(chan int)
	go addNumbersSlowly(doneChannel, 1, 1)
	fmt.Println("Waiting...")
	fmt.Println(<-doneChannel)
	fmt.Println("Done!")

}
// END OMIT
