package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	c := make(chan interface{})
	go func() {
		time.Sleep(5 * time.Second)

		// Here we close the channel after waiting five seconds.
		close(c)
	}()

	fmt.Println("Blocking on read...")
	select {
	// Here we attempt a read on the channel.
	case <-c:
		fmt.Printf("Unblocked %v later.\n", time.Since(start))
	}
}

// We only unblock roughly five seconds after entering the select block.
// This is a simple and efficient way to block while we’re waiting for something to happen.

/*

Blocking on read...
Unblocked 5.003761973s later.

*/
