package p45

import (
	"sync"
	"testing"
)

func BenchmarkContextSwitch(b *testing.B) {
	var wg sync.WaitGroup
	begin := make(chan struct{})
	c := make(chan struct{})

	var token struct{}
	sender := func() {
		defer wg.Done()
		<-begin // Here we wait until we’re told to begin.
		// We don’t want the cost of setting up and starting each goroutine
		// to factor into the measurement of context switching.
		for i := 0; i < b.N; i++ {
			c <- token // Here we send messages to the receiver goroutine.
			// A struct{}{} is called an empty struct and takes up no memory;
			// thus, we are only measuring the time it takes to signal a message.
		}
	}

	receiver := func() {
		defer wg.Done()
		<-begin // Here we wait until we’re told to begin.
		for i := 0; i < b.N; i++ {
			<-c // Here we receive a message but do nothing with it.
		}
	}

	wg.Add(2)
	go sender()
	go receiver()
	b.StartTimer() // Here we begin the performance timer.
	close(begin)   // Here we tell the two goroutines to begin.
	wg.Wait()
}

// go test -v --bench=. p45-ex_test.go
// go test --bench=. p45-ex_test.go
// 1000000              1456 ns/op

// We run the benchmark specifying that we only want to utilize one CPU
// so that it’s a similar test to the Linux benchmark.

// go test --bench=. --cpu=1 p45-ex_test.go
// 1243086               965.4 ns/op
