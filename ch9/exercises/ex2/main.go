package main

import "fmt"

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
	Name string
}

func (e EmployeeNameErr) Error() string {
	return e.Msg
}

func checkForNameErr(name string) error {
	if name == "" {
		return EmployeeNameErr{
			Status: NoName,
			Msg:    fmt.Sprintf("Please fill out your name. Got \"%s\"", name),
		}
	} else if name == "Jake" {
		return EmployeeNameErr{
			Status: NotAValidName,
			Msg:    fmt.Sprintf("That is an invalid name. %s is taken", name),
		}
	}
	return nil
}

func fillYourNameOut(name string) (Employee, error) {
	err := checkForNameErr(name)
	if err == nil {
		return Employee{
			Name: name,
		}, nil
	}
	return Employee{}, err
}

func checkForError(theEmployee Employee, thisError error) bool {

	if thisError != nil {
		fmt.Println("There was an error:", thisError)
		return true
	}
	fmt.Printf("Hello %s\n", theEmployee.Name)
	return false
}

func main() {
	// Define a custom error type to represent an empty field error.
	// This error should include the name of the empty Employee field.
	// In main, use errors.As to check for this error. Print out a message that includes the field name.

	emp1, err := fillYourNameOut("Kacie")
	checkForError(emp1, err)
	emp2, err := fillYourNameOut("")
	checkForError(emp2, err)
	emp3, err := fillYourNameOut("Jake")
	checkForError(emp3, err)
}
