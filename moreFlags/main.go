package main

import "fmt"

// Define constants for bit masks
const (
	FlagA uint8 = 1 << iota // 1 << 0 = 1
	FlagB                   // 1 << 1 = 2
	FlagC                   // 1 << 2 = 4
	FlagD                   // 1 << 3 = 8
)

func main() {
	// Example usage of bit fields
	var flags uint8 = FlagA | FlagC // Setting FlagA and FlagC

	flags & FlagB = 1

	// Check if FlagB is set
	if flags&FlagB != 0 {
		fmt.Println("FlagB is set")
	} else {
		fmt.Println("FlagB is not set")
	}

	// Check if FlagD is set
	if flags&FlagD != 0 {
		fmt.Println("FlagD is set")
	} else {
		fmt.Println("FlagD is not set")
	}

	// Clear FlagA
	flags &^= FlagA

	fmt.Printf("Remaining flags: %08b\n", flags)
}
