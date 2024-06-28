// package main this is exercise 2 from chapter 12 from learning go on [oreilly].
//
// [oreilly]: https://learning.oreilly.com/library/view/learning-go-2nd/9781098139285/ch12.html#id354
package main

// Problem 2:
// Create a function that launches two goroutines. Each goroutine writes 10 numbers to its own channel.
// Use a for-select loop to read from both channels, printing out the number and the goroutine that wrote the value.
// Make sure that your function exits after all values are read and that none of your goroutines leak.

import (
	"fmt"
	"sync"
)

func producer(ch chan<- int, base int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		ch <- (i + base) // Send data to channel
	}
}

func doTheClose(wg *sync.WaitGroup, ch chan int) {
	wg.Wait() // Wait for all producers to complete
	close(ch) // Close the channel
}

func main() {
	ch := make(chan int)  // Create a buffered channel
	ch2 := make(chan int) // Create a buffered channel
	var wg sync.WaitGroup
	wg.Add(1)
	var wg2 sync.WaitGroup
	wg2.Add(1)

	// Start two producer goroutines
	go producer(ch, 0, &wg)
	go producer(ch2, 10, &wg2)

	// Close the channel after all producers are done
	go doTheClose(&wg, ch)
	go doTheClose(&wg2, ch2)

	for count := 0; count < 2; {
		select {
		case v, ok := <-ch:
			if !ok {
				ch = nil // the case will never succeed again!
				count++
				continue
			}
			fmt.Println("Channel 1", v)
		case v, ok := <-ch2:
			if !ok {
				ch2 = nil // the case will never succeed again!
				count++
				continue
			}
			fmt.Println("Channel 2", v)
		}
	}
}
