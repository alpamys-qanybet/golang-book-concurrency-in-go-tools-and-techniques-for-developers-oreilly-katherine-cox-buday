package main

import (
	"fmt"
	"sync"
)

func main() {

	// We define a type Button that contains a condition, Clicked.
	type Button struct {
		Clicked *sync.Cond
	}

	button := Button{
		Clicked: sync.NewCond(&sync.Mutex{}),
	}

	// Here we define a convenience function that will allow us
	// to register functions to handle signals from a condition.
	// Each handler is run on its own goroutine, and subscribe
	// will not exit until that goroutine is confirmed to be running.
	subscribe := func(c *sync.Cond, fn func()) {
		var goroutineRunning sync.WaitGroup
		goroutineRunning.Add(1)
		go func() {
			goroutineRunning.Done()
			c.L.Lock()
			defer c.L.Unlock()
			c.Wait()
			fn()
		}()

		goroutineRunning.Wait()
	}

	// Here we create a WaitGroup. This is done only to ensure our program doesn’t exit before our writes to stdout occur.
	var clickRegistered sync.WaitGroup
	clickRegistered.Add(3)

	// Here we register a handler that simulates maximizing
	// the button’s window when the button is clicked.
	subscribe(button.Clicked, func() {
		fmt.Println("Maximizing window.")
		clickRegistered.Done()
	})

	// Here we register a handler that simulates displaying
	// a dialog box when the mouse is clicked.
	subscribe(button.Clicked, func() {
		fmt.Println("Displaying annoying dialog box!")
		clickRegistered.Done()
	})

	// Here we set a handler for when the mouse button is raised.
	subscribe(button.Clicked, func() {
		fmt.Println("Mouse clicked.")
		clickRegistered.Done()
	})

	// Next, we simulate a user raising the mouse button
	// from having clicked the application’s button.
	// It in turn calls Broadcast on the Clicked Cond to let all handlers
	// know that the mouse button has been clicked.
	button.Clicked.Broadcast()

	clickRegistered.Wait()
}

/*
Mouse clicked.
Maximizing window.
Displaying annoying dialog box!
*/
