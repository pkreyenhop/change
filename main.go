// Package main is the main package for the change program.
package main

import (
	"fmt"
	"os"
	"strings"
)

// main is the entry point of the program.
func main() {
	// Check if enough arguments are provided
	if len(os.Args) < 4 {
		fmt.Println("Usage: change <filename> <string_to_replace> <replacement>")
		return
	}

	// Extract command-line arguments
	filename := os.Args[1]
	stringToReplace := os.Args[2]
	replacement := os.Args[3]

	// Read the contents of the file
	fileContents, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Perform the replacement in memory
	updatedContents := strings.ReplaceAll(string(fileContents), stringToReplace, replacement)

	// Write the modified contents back to the file
	err = os.WriteFile(filename, []byte(updatedContents), 0644)
	if err != nil {
		fmt.Println("Error writing file:", err)
		return
	}

	// Print a success message
	fmt.Printf("Occurrences of '%s' changed to '%s' in %s\n", stringToReplace, replacement, filename)
}
