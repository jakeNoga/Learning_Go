package main

import (
	"fmt"
	"hello/anotherstrings"
	"hello/morestrings"
)

func main() {
	fmt.Println(morestrings.ReverseRunes("Hello, go!"))
	fmt.Println(anotherstrings.AnotherReverseRunes("Really!"))
}
