package main

import (
	"flag"
	"fmt"
)

// Define constants for bit masks
type Flags uint8

const (
	FlagA Flags = 1 << iota // 1 << 0 = 1
	FlagB                   // 1 << 1 = 2
	FlagC                   // 1 << 2 = 4
	FlagD                   // 1 << 3 = 8
)

func (f *Flags) Set(value string) error {
	switch value {
	case "A":
		*f |= FlagA
	case "B":
		*f |= FlagB
	case "C":
		*f |= FlagC
	case "D":
		*f |= FlagD
	default:
		return fmt.Errorf("unknown flag: %s", value)
	}
	return nil
}

func (f *Flags) String() string {
	return fmt.Sprintf("%08b", *f)
}

// Custom function to set FlagB
func setFlagB(f *Flags) {
	*f |= FlagB
}

func main() {
	var flags Flags
	var enableB bool

	// Register custom flag -FlagsB
	flag.Func("FlagsB", "Enable FlagB", func(value string) error {
		setFlagB(&flags)
		return nil
	})

	flag.Var(&flags, "flags", "Bit flags (A, B, C, D)")
	flag.BoolVar(&enableB, "enableB", false, "Enable FlagB")

	flag.Parse()

	fmt.Printf("Flags set: %08b\n", flags)
	fmt.Printf("FlagB enabled: %v\n", enableB)
}
