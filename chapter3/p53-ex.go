package main

import (
	"fmt"
	"sync"
	"time"
)

func conditionTrue() bool {
	return true // it has to be true and for has to !conditionTrue(), do not why?
}

func main() {

	/*
		// this would consume all cycles of one core.
		for conditionTrue() == false {

		}
	*/

	/*
		Still inefficient, and you have to figure out how long to sleep for:
		too long, and you’re artificially degrading performance;
		too short, and you’re unnecessarily consuming too much CPU time.
		for conditionTrue() == false {
			time.Sleep(1*time.Millisecond)
		}
	*/

	// Efficiently sleep until it was signaled to wake and check its condition.
	c := sync.NewCond(&sync.Mutex{}) // The NewCond function takes in a type that satisfies
	// the sync.Locker interface. This is what allows the Cond type to facilitate coordination
	// with other goroutines in a concurrent-safe way.
	c.L.Lock()
	for !conditionTrue() {
		fmt.Println("inside loop")
		c.Wait() // Here we wait to be notified that the condition has occurred.
		// This is a blocking call and the goroutine will be suspended.
	}
	c.L.Unlock() // Here we unlock the Locker for this condition.
	// This is necessary because when the call to Wait exits,
	// it calls Lock on the Locker for the condition.

	time.Sleep(1 * time.Second)
	c.Signal()
}

// c.Wait documentation says this is common pattern of implementing the code:

//	c.L.Lock()
//	for !condition() {
//	    c.Wait()
//	}
//	... make use of condition ...
//	c.L.Unlock()

// fatal error: all goroutines are asleep - deadlock!
// ????
// sometimes just passes without any err
