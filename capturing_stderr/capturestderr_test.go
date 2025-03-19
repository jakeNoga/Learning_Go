package main

import (
	"fmt"
	"io"
	"os"
	"testing"
)

func PrintError() {
	fmt.Fprintln(os.Stderr, "This is an error message.")
}

func TestCaptureStderr(t *testing.T) {
	// Backup the original stderr
	rescueStderr := os.Stderr

	// Create a pipe to capture stderr
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatalf("os.Pipe() failed: %v", err)
	}
	os.Stderr = w

	// Defer restoring the original stderr and closing the pipe
	defer func() {
		os.Stderr = rescueStderr
		w.Close()
	}()

	// Call the function that prints to stderr
	PrintError()

	// Close the write end of the pipe to signal that no more data will be written
	w.Close()

	// Read from the pipe
	out, err := io.ReadAll(r)
	if err != nil {
		t.Fatalf("io.ReadAll() failed: %v", err)
	}

	// Convert byte slice to string
	stderrOutput := string(out)

	// Print the captured output for debugging
	fmt.Println("Captured stderr:", stderrOutput)

	// Check the captured output
	expected := "This is an error message.\n"
	if stderrOutput != expected {
		t.Errorf("PrintError() = %q; want %q", stderrOutput, expected)
	}
}
