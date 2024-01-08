package main

import "fmt"

func main() {
	intStream := make(chan int)
	close(intStream)

	integer, ok := <-intStream
	fmt.Printf("%v: %v\n", integer, ok)
}

/*

0: false

true - channel is open
false - closed

*/
