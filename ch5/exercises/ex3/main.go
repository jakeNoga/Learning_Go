package main

import "fmt"

func prefixer(thePrefix string) func(string) string {
	return func(theName string) string {
		return thePrefix + " " + theName
	}
}

// This program is wild and hard to wrap my head around but essentially prefixer is returning a function that goes into the variable helloPrefix.
// From helloPrefix you can call it like a normal function and then add another string to what was originally called pretty wild.
func main() {
	// Write a function called prefixer that has an input parameter of type string and returns a function that has an input parameter of type string and returns a string.
	// The returned function should prefix its input with the string passed into prefixer. Use the following main function to test prefixer
	helloPrefix := prefixer("Hello")
	fmt.Println(helloPrefix("Bob"))   // should print Hello Bob
	fmt.Println(helloPrefix("Maria")) // should print Hello Maria
}
