package main

import (
	"fmt"
	"os"
)

// My function was more simplistic then the answer key. I know I could have done it that way it just seemed like a waste of resources...
func fileLen(fileName string) (int, error) {
	f, err := os.Open(fileName)
	if err != nil {
		return 0, err
	}
	defer f.Close()
	fileStats, err := f.Stat()
	if err != nil {
		return 0, err
	}

	return int(fileStats.Size()), nil
}

func main() {
	// Write a function called fileLen that has an input parameter of type string and returns an int and an error.
	// The function takes in a filename and returns the number of bytes in the file. If there is an error reading the file, return the error.
	// Use defer to make sure the file is closed properly.
	testFiles := []string{
		"testFile.txt",         // This file exists
		"fileDoesNotExist.txt", // This file does not exist
		"thisIsADirectory",     //
	}

	for _, thisFile := range testFiles {
		thisInt, thisError := fileLen(thisFile)
		if thisError != nil {
			fmt.Println(thisError)
			continue
		}
		fmt.Println("This file", thisFile, "has the size:", thisInt)
	}

}
