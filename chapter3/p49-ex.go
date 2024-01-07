package main

import (
	"fmt"
	"sync"
)

func main() {
	var count int
	var lock sync.Mutex

	increment := func() {
		lock.Lock()
		defer lock.Unlock()
		count++
		fmt.Printf("Incrementing: %d\n", count)
	}

	decrement := func() {
		lock.Lock()
		defer lock.Unlock()
		count--
		fmt.Printf("Decrementing: %d\n", count)
	}
	// Increment
	var arithmetic sync.WaitGroup
	for i := 0; i <= 5; i++ {
		arithmetic.Add(1)
		go func() {
			defer arithmetic.Done()
			increment()
		}()
	}

	// Decrement
	for i := 0; i <= 5; i++ {
		arithmetic.Add(1)
		go func() {
			defer arithmetic.Done()
			decrement()
		}()
	}

	arithmetic.Wait()
	fmt.Println("Arithmetic complete.")
}

/*

results may vary

Incrementing: 1
Decrementing: 0
Decrementing: -1
Decrementing: -2
Decrementing: -3
Decrementing: -4
Incrementing: -3
Decrementing: -4
Incrementing: -3
Incrementing: -2
Incrementing: -1
Incrementing: 0
Arithmetic complete.

or

Incrementing: 1
Incrementing: 2
Incrementing: 3
Incrementing: 4
Incrementing: 5
Incrementing: 6
Decrementing: 5
Decrementing: 4
Decrementing: 3
Decrementing: 2
Decrementing: 1
Decrementing: 0
Arithmetic complete.

or

Incrementing: 1
Decrementing: 0
Decrementing: -1
Decrementing: -2
Decrementing: -3
Decrementing: -4
Decrementing: -5
Incrementing: -4
Incrementing: -3
Incrementing: -2
Incrementing: -1
Incrementing: 0
Arithmetic complete.

or etc..
*/
