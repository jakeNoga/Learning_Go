package main

import (
	"flag"
	"fmt"
	"os"
)

type DebugFlags struct {
	Verbose bool
	LogFile string
}

type ConfigFlags struct {
	Port int
	Host string
}

func main() {
	// Create two flag sets
	debugFlags := flag.NewFlagSet("debug", flag.ExitOnError)
	configFlags := flag.NewFlagSet("config", flag.ExitOnError)

	// Define flags for the debug flag set
	var debug DebugFlags
	debugFlags.BoolVar(&debug.Verbose, "verbose", false, "Enable verbose logging")
	debugFlags.StringVar(&debug.LogFile, "logfile", "app.log", "Log file path")

	// Define flags for the config flag set
	var config ConfigFlags
	configFlags.IntVar(&config.Port, "port", 8080, "Port to listen on")
	configFlags.StringVar(&config.Host, "host", "localhost", "Host to bind to")

	// Check if we have enough arguments
	if len(os.Args) < 2 {
		fmt.Println("Usage: [debug|config] [options]")
		os.Exit(1)
	}

	// Parse the appropriate flag set based on the first argument
	switch os.Args[1] {
	case "debug":
		debugFlags.Parse(os.Args[2:])
		fmt.Println("Debug Flags:")
		fmt.Printf("Verbose: %v\n", debug.Verbose)
		fmt.Printf("LogFile: %s\n", debug.LogFile)
	case "config":
		configFlags.Parse(os.Args[2:])
		fmt.Println("Config Flags:")
		fmt.Printf("Port: %d\n", config.Port)
		fmt.Printf("Host: %s\n", config.Host)
	default:
		fmt.Println("Unknown command. Use 'debug' or 'config'.")
		os.Exit(1)
	}
}
