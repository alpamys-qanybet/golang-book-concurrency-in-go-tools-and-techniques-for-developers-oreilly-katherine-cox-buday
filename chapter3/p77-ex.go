package main

import "fmt"

func main() {

	// Notice how the lifecycle of the resultStream channel
	// is encapsulated within the chan Owner function.
	// It’s very clear that the writes will not happen on a nil or closed channel,
	// and that the close will always happen once.
	chanOwner := func() <-chan int {

		// Here we instantiate a buffered channel.
		// Since we know we’ll produce six results,
		// we create a buffered channel of five so that
		// the goroutine can complete as quickly as possible.
		resultStream := make(chan int, 5)

		// Here we start an anonymous goroutine that performs writes on resultStream.
		// Notice that we’ve inverted how we create goroutines.
		// It is now encapsulated within the surrounding function.
		go func() {
			// Here we ensure resultStream is closed once we’re finished with it.
			// As the channel owner, this is our responsibility.
			defer close(resultStream)

			for i := 0; i <= 5; i++ {
				resultStream <- i
			}
		}()

		// Here we return the channel.
		// Since the return value is declared as a read-only channel,
		// resultStream will implicitly be converted to read-only for consumers.
		return resultStream
	}

	resultStream := chanOwner()

	// Here we range over resultStream.
	// As a consumer, we are only concerned with blocking and closed channels.
	for result := range resultStream {
		fmt.Printf("Received: %d\n", result)
	}

	fmt.Println("Done receiving!")
}

/*
Received: 0
Received: 1
Received: 2
Received: 3
Received: 4
Received: 5
Done receiving!
*/
