package main

import "fmt"

func main() {
	type employee struct {
		firstName string
		lastName  string
		id        int
	}
	// Initialize the first one using the struct literal style without names
	darthVader := employee{
		"Anakin",
		"Skywalker",
		1,
	}
	fmt.Println("Darth vader:", darthVader)
	// the second using the struct literal style with names
	joker := employee{
		firstName: "Jack",
		lastName:  "White",
		id:        2,
	}
	fmt.Println("Joker:", joker)
	// the third with a var declaration. Use dot notation to populate the fields in the third struct.
	var captain employee
	captain.firstName = "Jack"
	captain.lastName = "Sparrow"
	captain.id = 3
	fmt.Println("Captain:", captain)
}
