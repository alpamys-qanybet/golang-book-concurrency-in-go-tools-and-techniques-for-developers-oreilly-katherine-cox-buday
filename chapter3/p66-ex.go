package main

/*
import (
	"fmt"


func main() {
	stringStream := make(chan string)
	go func() {
		stringStream <- "Hello channels!"
	}()

	fmt.Println(<-stringStream)
}
*/

import (
	"fmt"
	"time"
)

func main() {
	stringStream := make(chan string)
	go func() {
		time.Sleep(2 * time.Second)
		stringStream <- "Hello channels!"
	}()

	fmt.Println(<-stringStream)
}

// Hello channels!

// Just because a goroutine was scheduled,
// there was no guarantee that it would run before the process exited;
// But channels are said to be blocking.
// This means that any goroutine that attempts to write to a channel that is full will wait until the channel has been emptied,
// and any goroutine that attempts to read from a channel that is empty will wait until at least one item is placed on it.

// works even with timer, it waits for goroutine because of channel

// Likewise, the anonymous goroutine is attempting to place a string literal on the stringStream,
// and so the goroutine will not exit until the write is successful.
