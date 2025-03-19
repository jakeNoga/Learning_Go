package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	// Define the first set of flags
	fs1 := flag.NewFlagSet("set1", flag.ExitOnError)
	var node1 string
	fs1.StringVar(&node1, "node", "", "Node name for set 1")
	var showHidden bool
	fs1.BoolVar(&showHidden, "hidden", false, "Show hidden values from set 2")

	// Define the second set of flags
	fs2 := flag.NewFlagSet("set2", flag.ContinueOnError)
	var node2 string
	fs2.StringVar(&node2, "node", "", "Node name for set 2")

	// Override the Usage function for set2 to customize output
	fs2.Usage = func() {
		fmt.Fprintf(fs2.Output(), "Usage of %s:\n", fs2.Name())
		fmt.Fprintf(fs2.Output(), "  -node string\n\tNode name for set 2\n")
	}

	// Parse command-line arguments for set 1
	fs1.Parse(os.Args[1:])

	// If --hidden is specified, parse command-line arguments for set 2
	if showHidden {
		fs2.Parse(os.Args[1:])
	}

	// Print the parsed values from each flag set
	fmt.Println("Node name for set 1:", node1)
	if showHidden {
		fmt.Println("Node name for set 2:", node2)
	}
}
