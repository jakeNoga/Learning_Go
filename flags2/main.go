package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	// Define flag aliases and their corresponding help texts
	flagAliases := map[string]string{
		"-n":    "node name (alias for -node)",
		"-node": "node name",
	}

	// Parse arguments
	for i := 1; i < len(os.Args); i++ {
		if alias, ok := flagAliases[os.Args[i]]; ok {
			os.Args[i] = alias
		}
	}

	// Define flags with help texts
	var nodeName string
	flag.StringVar(&nodeName, "node", "", "Node name")
	flag.StringVar(&nodeName, "n", "", "Node name (alias for -node)")

	// Parse flags
	flag.Parse()

	// Output the parsed node name
	fmt.Println("Node name:", nodeName)
}
