package main

import (
	"bytes"
	"fmt"
	"sync"
)

func main() {
	printData := func(wg *sync.WaitGroup, data []byte) {
		defer wg.Done()
		var buff bytes.Buffer

		for _, b := range data {
			fmt.Fprintf(&buff, "%c", b)
		}

		fmt.Println(buff.String())
	}

	var wg sync.WaitGroup
	wg.Add(2)
	data := []byte("golang")

	// Here we pass in a slice containing the first three bytes in the data structure.
	go printData(&wg, data[:3])

	// Here we pass in a slice containing the last three bytes in the data structure.
	go printData(&wg, data[3:])
	wg.Wait()
}

/*

results may vary

ang
gol

or

gol
ang

*/
