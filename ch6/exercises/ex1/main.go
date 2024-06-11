package main

import (
	"errors"
	"fmt"
)

type person struct {
	FirstName string
	LastName  string
	Age       int
}

func MakePerson(firstName string, lastName string, age int) (person, error) {
	if age < 0 {
		return person{"", "", 0}, errors.New("the age must be greater then 0")
	}
	this_person := person{FirstName: firstName,
		LastName: lastName,
		Age:      age,
	}
	return this_person, nil
}

func MakePersonPointer(firstName string, lastName string, age int) (*person, error) {
	if age < 0 {
		return nil, errors.New("the age must be greater then 0")
	}

	this_person := person{
		FirstName: firstName,
		LastName:  lastName,
		Age:       age,
	}
	return &this_person, nil
}

func main() {
	// Create a struct named Person with three fields: FirstName and LastName of type string and Age of type int.
	// Write a function called MakePerson that takes in firstName, lastName, and age and returns a Person.
	// Write a second function MakePersonPointer that takes in firstName, lastName, and age and returns a *Person.
	// Call both from main. Compile your program with go build -gcflags="-m". This both compiles your code and prints out which values escape to the heap.
	// Are you surprised about what escapes?

	// I personally am supper shocked that harvey was moved to the heap really unexpected reason:
	// The reason for this is that it is passed into fmt.Println. This is because the parameter to fmt.Println are ...any.
	// The current Go compiler moves to the heap any value that is passed in to a function via a parameter that is of an interface type. (Interfaces are covered in Chapter 7.)
	firstChair, err := MakePerson("Harvey", "Specter", 40)
	if err != nil {
		fmt.Println("Unable to create first chair lawyer")
	} else {
		fmt.Println("This is the lawyer in the first chair", firstChair)
	}
	secondChair, err := MakePerson("Mike", "Ross", 33)
	if err != nil {
		fmt.Println("Unable to create second chair lawyer")
	} else {
		fmt.Println("This is the lawyer in the second chair", secondChair)
	}
}
