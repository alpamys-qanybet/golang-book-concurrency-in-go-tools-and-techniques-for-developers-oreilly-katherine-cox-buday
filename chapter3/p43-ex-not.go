package main

import "fmt"

func main() {
	go func() {
		// <operation that will block forever>
		for { // program will exit and the goroutine also will die and the loop will break
			fmt.Println("inside forever loop")
		}

		fmt.Println("goroutine job done")
	}()

	// Do work
	fmt.Println("main")
}

// The goroutine here will hang around until the process exits.
