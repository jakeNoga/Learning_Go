package main

import (
	"fmt"
	"time"
)

func producer(ch chan<- int) {
	for i := 0; i < 5; i++ {
		fmt.Println("Producing", i)
		ch <- i                 // Send data to channel
		time.Sleep(time.Second) // Simulate some work
	}
	close(ch) // Close the channel when done
}

func consumer(ch <-chan int) {
	for num := range ch {
		fmt.Println("Consuming", num)
		time.Sleep(2 * time.Second) // Simulate some work
	}
}

func main() {
	ch := make(chan int) // Create a channel

	go producer(ch) // Start producer goroutine
	go consumer(ch) // Start consumer goroutine

	time.Sleep(10 * time.Second) // Allow time for goroutines to finish
}
