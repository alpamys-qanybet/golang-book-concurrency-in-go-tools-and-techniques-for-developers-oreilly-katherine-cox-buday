package main

import (
	"fmt"
	"time"
)

func main() {

	start := time.Now()
	var c1, c2 <-chan int
	select {
	case <-c1:
		// do something
	case <-c2:
		// do something
	default:
		fmt.Printf("In default after %v\n\n", time.Since(start))
	}
}

// You can see that it ran the default statement almost instantaneously.
// This allows you to exit a select block without blocking.

/*

In default after 2.76µs

*/
