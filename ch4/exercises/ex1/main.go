package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// Write a for loop that puts 100 random numbers between 0 and 100 into an int slice
	thisSlice := make([]int, 0, 100)
	fmt.Println("Capacity of the slice:", cap(thisSlice))
	for i := 0; i < cap(thisSlice); i++ {
		thisSlice = append(thisSlice, rand.Intn(100))
	}
	fmt.Println("Capacity of the slice:", cap(thisSlice))
	fmt.Println("100 random numbers in a slice:", thisSlice)
}
