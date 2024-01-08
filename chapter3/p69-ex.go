package main

import "fmt"

func main() {
	intStream := make(chan int)

	go func() {
		defer close(intStream)

		for i := 1; i <= 5; i++ {
			intStream <- i
		}
	}()

	for integer := range intStream {
		fmt.Printf("%v ", integer)
	}
}

// 1 2 3 4 5

// it will read from even closed channel
