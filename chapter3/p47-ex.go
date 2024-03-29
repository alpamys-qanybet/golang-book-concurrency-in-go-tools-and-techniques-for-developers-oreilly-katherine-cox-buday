package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(1) // Here we call Add with an argument of 1 to indicate that one goroutine is beginning.
	go func() {
		defer wg.Done() // Here we call Done using the defer keyword to ensure that before we exit
		// the goroutine’s closure, we indicate to the WaitGroup that we’ve exited.
		fmt.Println("1st goroutine sleeping...")
		time.Sleep(1)
	}()

	wg.Add(1) // Here we call Add with an argument of 1 to indicate that one goroutine is beginning.
	go func() {
		defer wg.Done() // Here we call Done using the defer keyword to ensure that before we exit
		// the goroutine’s closure, we indicate to the WaitGroup that we’ve exited.
		fmt.Println("2nd goroutine sleeping...")
		time.Sleep(2)
	}()

	wg.Wait() // Here we call Wait, which will block the main goroutine until all goroutines
	// have indicated they have exited.
	fmt.Println("All goroutines complete.")
}

/*
results may vary

1st goroutine sleeping...
2nd goroutine sleeping...
All goroutines complete.

or

2nd goroutine sleeping...
1st goroutine sleeping...
All goroutines complete.
*/
