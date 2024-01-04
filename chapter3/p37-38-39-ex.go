package main

import "fmt"

func main() {
	ex1()
	ex2()
	ex3()

	// program will terminate and will not wait for goroutines to finish(may even no goroutine start or even scheduled)
}

// example 1
func ex1() {
	go sayHello()
	// continue doing other things
}

func sayHello() {
	fmt.Println("hello 1")
}

// example 2
func ex2() {
	go func() {
		fmt.Println("hello 2")
	}()
}

// example 3
func ex3() {
	sayHello3 := func() {
		fmt.Println("hello 3")
	}

	go sayHello3()
}

// this is just example of way of writing code, but no one of these examples working code
