package main

import "fmt"

type person struct {
	FirstName string
	LastName  string
	Age       int
}

func makePerson(firstName string, lastName string, age int) person {
	return person{
		FirstName: firstName,
		LastName:  lastName,
		Age:       age,
	}
}

func main() {
	// Write a program that builds a []Person with 10,000,000 entries (they could all be the same names and ages).
	// See how long it takes to run. Change the value of GOGC and see how that affects the time it takes for the program to complete.
	// Set the environment variable GODEBUG=gctrace=1 to see when garbage collections happen and see how changing GOGC changes the number of garbage collections.
	// What happens if you create the slice with a capacity of 10,000,000?

	// ALotOfVs := make([]person, 10_000_000)
	var ALotOfVs []person

	// Change len(ALotOfVs for 10_000_000
	for i := 0; i < 10_000_000; i++ {
		// Change ALotOfVs[i] to append
		ALotOfVs = append(ALotOfVs, makePerson("V", "Vendetta", 34))
	}
	fmt.Println(len(ALotOfVs))

	// 10000000
	// 0.01user 0.09system 0:01.07elapsed 10%CPU (0avgtext+0avgdata 5256maxresident)

	// 10000000
	// 0.00user 0.07system 0:03.41elapsed 2%CPU (0avgtext+0avgdata 5260maxresident)k
}
