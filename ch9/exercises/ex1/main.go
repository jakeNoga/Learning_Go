package main

import (
	"errors"
	"fmt"
)

var ErrYouFoundMe = errors.New("You found me!")

func didYouFindMe(spotted bool) error {
	if spotted {
		return ErrYouFoundMe
	}
	return nil
}

func main() {
	// Create a sentinel error to represent an invalid ID. In main, use errors.Is to check for the sentinel error, and print a message when it is found.
	amIFound := false
	err := didYouFindMe(amIFound)
	if errors.Is(err, ErrYouFoundMe) {
		fmt.Println("They won.", err)
	} else {
		fmt.Println("Game on!")
	}
	amIFound = true
	err = didYouFindMe(amIFound)
	if errors.Is(err, ErrYouFoundMe) {
		fmt.Println("They won.", err)
	} else {
		fmt.Println("Game on!")
	}
}
