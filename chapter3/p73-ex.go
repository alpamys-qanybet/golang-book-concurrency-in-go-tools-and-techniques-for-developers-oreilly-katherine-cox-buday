package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	// Here we create an in-memory buffer to help mitigate the nondeterministic nature of the output.
	// It doesn’t give us any guarantees, but it’s a little faster than writing to stdout directly.
	var stdoutBuff bytes.Buffer

	// Here we ensure that the buffer is written out to stdout before the process exits.
	defer stdoutBuff.WriteTo(os.Stdout)

	// Here we create a buffered channel with a capacity of one.
	intStream := make(chan int, 5)

	go func() {
		defer close(intStream)
		defer fmt.Fprintln(&stdoutBuff, "Producer Done.")

		for i := 0; i < 5; i++ {
			fmt.Fprintf(&stdoutBuff, "Sending: %d\n", i)
			intStream <- i
		}
	}()

	for integer := range intStream {
		fmt.Fprintf(&stdoutBuff, "Received %v.\n", integer)
	}
}

/*
if bufferred channel size is 4,5
Sending: 0
Sending: 1
Sending: 2
Sending: 3
Sending: 4
Producer Done.
Received 0.
Received 1.
Received 2.
Received 3.
Received 4.





if bufferred channel size is 1
Sending: 0
Sending: 1
Sending: 2
Received 0.
Received 1.
Received 2.
Sending: 3
Sending: 4
Producer Done.
Received 3.
Received 4.
*/
