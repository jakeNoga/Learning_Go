package main

import (
	"errors"
	"fmt"
)

type Status int

const (
	NoName Status = iota + 1
	NotAValidName
)

type EmployeeNameErr struct {
	Status Status
	Msg    string
}

type Employee struct {
	FirstName string
	LastName  string
}

func (e EmployeeNameErr) Error() string {
	return e.Msg
}

func checkForNameErr(name string, whichName string) error {
	if name == "" {
		return EmployeeNameErr{
			Status: NoName,
			Msg:    fmt.Sprintf("Invalid %s. Please fill out your %s. Got \"%s\"", whichName, whichName, name),
		}
	} else if name == "Jake" && whichName == "First name" {
		return EmployeeNameErr{
			Status: NotAValidName,
			Msg:    fmt.Sprintf("That is an invalid %s. %s is taken", whichName, name),
		}
	}
	return nil
}

func fillYourNameOut(firstName string, lastName string) (Employee, error) {
	var errs []error
	emp := Employee{}
	err := checkForNameErr(firstName, "First name")
	if err != nil {
		errs = append(errs, err)
	} else {
		emp.FirstName = firstName
	}
	err = checkForNameErr(lastName, "Last name")
	if err != nil {
		errs = append(errs, err)
	} else {
		emp.LastName = lastName
	}
	if len(errs) > 0 {
		return emp, errors.Join(errs...)
	}

	return emp, nil
}

func checkForError(theEmployee Employee, thisError error) bool {

	if thisError != nil {
		fmt.Println("There was an error:\n", thisError)
		return true
	}
	fmt.Printf("Hello %s\n", theEmployee.FirstName)
	return false
}

func main() {
	// Define a custom error type to represent an empty field error.
	// This error should include the name of the empty Employee field.
	// In main, use errors.As to check for this error. Print out a message that includes the field name.

	emp1, err := fillYourNameOut("Kacie", "Noga")
	checkForError(emp1, err)
	emp2, err := fillYourNameOut("", "Noga")
	checkForError(emp2, err)
	emp3, err := fillYourNameOut("", "")
	checkForError(emp3, err)
	emp4, err := fillYourNameOut("Kacie", "")
	checkForError(emp4, err)
	emp5, err := fillYourNameOut("Jake", "Noga")
	checkForError(emp5, err)
}
