package main

import (
	"fmt"
	"sync"
)

func main() {
	var numCalcsCreated int
	calcPool := &sync.Pool{
		New: func() interface{} {
			numCalcsCreated += 1
			mem := make([]byte, 1024)

			// Notice that we are storing the address of the slice of bytes.
			return &mem
		},
	}

	// Seed the pool with 4KB
	calcPool.Put(calcPool.New())
	calcPool.Put(calcPool.New())
	calcPool.Put(calcPool.New())
	calcPool.Put(calcPool.New())
	const numWorkers = 1024 * 1024
	var wg sync.WaitGroup
	wg.Add(numWorkers)

	for i := numWorkers; i > 0; i-- {
		go func() {
			defer wg.Done()
			// And here we are asserting the type is a pointer to a slice of bytes.
			mem := calcPool.Get().(*[]byte)

			defer calcPool.Put(mem)
			// Assume something interesting, but quick is being done with
			// this memory.
		}()
	}
	wg.Wait()
	fmt.Printf("%d calculators were created.\n", numCalcsCreated)
}

/*
results may vary

4 calculators were created.

or

6 calculators were created.

or

8calculators were created.

but not 1024*1024
*/
