// package main this is exercise 1 from chapter 12 from learning go on [oreilly].
//
// [oreilly]: https://learning.oreilly.com/library/view/learning-go-2nd/9781098139285/ch12.html#id354

// Problem 1:
// Create a function that launches three goroutines that communicate using a channel. The first two goroutines each write 10 numbers to the channel.
// The third goroutine reads all the numbers from the channel and prints them out. The function should exit when all values have been printed out.
// Make sure that none of the goroutines leak. You can create additional goroutines if needed.

package main

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

func consumer(ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for num := range ch {
		fmt.Println("Consuming", num)
	}
}

func doTheClose(wg *sync.WaitGroup, ch chan int) {
	wg.Wait() // Wait for all producers to complete
	close(ch) // Close the channel
}

func main() {
	ch := make(chan int) // Create a buffered channel
	var wg sync.WaitGroup
	wg.Add(2)

	// Start two producer goroutines
	go producer(ch, 0, &wg)
	go producer(ch, 10, &wg)

	// Close the channel after all producers are done
	go doTheClose(&wg, ch)

	var wgCons sync.WaitGroup
	wgCons.Add(1)

	// Start consumer goroutine
	go consumer(ch, &wgCons)
	wgCons.Wait() // Wait for consumer to complete
}
