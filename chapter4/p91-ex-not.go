package main

import "fmt"

func main() {
	doWork := func(strings <-chan string) <-chan interface{} {
		fmt.Println("#2") // printed
		completed := make(chan interface{})
		go func() {
			fmt.Println("#1") // not printed. why?
			defer fmt.Println("doWork exited.")
			defer close(completed)

			for s := range strings {
				fmt.Println("a") // not printed
				// Do something interesting
				fmt.Println(s)
			}
		}()

		return completed
	}

	doWork(nil)
	// a := doWork(nil)
	// Perhaps more work is done here

	// <-a // fatal error: all goroutines are asleep - deadlock!
	fmt.Println("Done.")
}

// main goroutine passes a nil channel into doWork.
// Therefore, the strings channel will never actually gets any strings written onto it,
// and the goroutine containing doWork will remain in memory for the lifetime of this process
// (we would even deadlock if we joined the goroutine within doWork and the main goroutine).

// ???????
/*

Done

*/
