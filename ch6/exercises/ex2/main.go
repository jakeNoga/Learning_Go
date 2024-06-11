package main

import (
	"fmt"
)

func UpdateSlice(sliceStr []string, str string) {
	sliceStr[len(sliceStr)-1] = str
	fmt.Println("In UpdateSlice:", sliceStr)
}

func GrowSlice(sliceStr []string, str string) {
	sliceStr = append(sliceStr, str)
	fmt.Println("In GrowSlice:", sliceStr)
}

func main() {
	// Write two functions. The UpdateSlice function takes in a []string and a string. It sets the last position in the passed-in slice to the passed-in string.
	// At the end of UpdateSlice, print the slice after making the change.
	// The GrowSlice function also takes in a []string and a string. It appends the string onto the slice. At the end of GrowSlice, print the slice after making the change.
	// Call these functions from main.
	// Print out the slice before each function is called and after each function is called.
	// Do you understand why some changes are visible in main and why some changes are not?

	// I expected the behavior because you can't change the size of a slice in a function.
	s := []string{"first", "second", "third"}
	fmt.Println(s, len(s))

	firstSlice := []string{"I", "am", "feeling", "really", "un-cleaver"}
	fmt.Println("firstSlice", firstSlice)
	UpdateSlice(firstSlice, "cleaver but what if the string is long")
	fmt.Println("firstSlice", firstSlice)

	magicTrick := []string{"This", "is", "a", "magic", "trick."}
	fmt.Println(magicTrick)
	GrowSlice(magicTrick, "Now you see me!")
	fmt.Println(magicTrick)
}
