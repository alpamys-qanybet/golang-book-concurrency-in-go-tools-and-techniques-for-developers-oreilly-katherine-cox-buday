package main

import (
	"fmt"
	"sync"
)

func main() {
	begin := make(chan interface{})

	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()

			// Here the goroutine waits until it is told it can continue.
			<-begin

			fmt.Printf("%v has begun\n", i)
		}(i)
	}

	fmt.Println("Unblocking goroutines...")

	// Here we close the channel, thus unblocking all the goroutines simultaneously.
	close(begin)

	wg.Wait()
}

/*

Unblocking goroutines...
0 has begun
3 has begun
1 has begun
2 has begun
4 has begun

*/
