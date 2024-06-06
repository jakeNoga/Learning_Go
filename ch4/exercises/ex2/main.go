package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// Loop over the slice you created in exercise 1. For each value in the slice, apply the following rules:
	// If the value is divisible by 2, print “Two!”
	// If the value is divisible by 3, print “Three!”
	// IIf the value is divisible by 2 and 3, print “Six!”. Don’t print anything else.
	// Otherwise, print “Never mind”.

	thisSlice := make([]int, 0, 100)
	for i := 0; i < cap(thisSlice); i++ {
		thisSlice = append(thisSlice, rand.Intn(100))
	}
	// Theoretically this could be done in the loop above but I wanted to try other type of looping
	for i, val := range thisSlice {
		fmt.Println("The value at index:", i, "is", val)
		switch {
		case val%2 == 0 && val%3 == 0:
			fmt.Println("Six!")
		case val%2 == 0:
			fmt.Println("Two!")
		case val%3 == 0:
			fmt.Println("Three!")
		default:
			fmt.Println("Never mind")
		}
	}
}
