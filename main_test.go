package main

import (
	"io/ioutil"
	"os"
	"strings"
	"testing"
	"time"	
)

// Test main function with a successful replacement
func TestMain_SuccessfulReplacement(t *testing.T) {
	// Create a temporary file
	file, err := ioutil.TempFile("", "testfile*.txt")
	if err != nil {
		t.Fatalf("Error creating temporary file: %v", err)
	}
	defer os.Remove(file.Name())

	// Write initial content to the file
	initialContent := "This is Peter's test file. Peter is here!"
	if _, err := file.Write([]byte(initialContent)); err != nil {
		t.Fatalf("Error writing to the temporary file: %v", err)
	}

	// Close the file before running the main function
	file.Close()

	// Call the main function to perform replacement
	os.Args = []string{"change", file.Name(), "Peter", "Daniel"}
	main()

	// Read the contents of the modified file
	updatedContent, err := ioutil.ReadFile(file.Name())
	if err != nil {
		t.Fatalf("Error reading the modified file: %v", err)
	}

	// Check if the replacement was successful
	expectedContent := "This is Daniel's test file. Daniel is here!"
	if string(updatedContent) != expectedContent {
		t.Errorf("Expected: %s, but got: %s", expectedContent, string(updatedContent))
	}
}

// Test main function with an empty file
func TestMain_EmptyFile(t *testing.T) {
	// Create a temporary empty file
	file, err := ioutil.TempFile("", "testfile*.txt")
	if err != nil {
		t.Fatalf("Error creating temporary file: %v", err)
	}
	defer os.Remove(file.Name())

	// Close the file before running the main function
	file.Close()

	// Call the main function to perform replacement
	os.Args = []string{"change", file.Name(), "Peter", "Daniel"}
	main()

	// Read the contents of the modified file
	updatedContent, err := ioutil.ReadFile(file.Name())
	if err != nil {
		t.Fatalf("Error reading the modified file: %v", err)
	}

	// An empty file should remain unchanged
	if string(updatedContent) != "" {
		t.Errorf("Expected an empty file, but got: %s", string(updatedContent))
	}
}

// Test main function with a non-existent file
func TestMain_NonExistentFile(t *testing.T) {
	// Call the main function with a non-existent file
	os.Args = []string{"change", "non_existent.txt", "Peter", "Daniel"}
	main()

	// The main function should handle the non-existent file gracefully
	// No error should be raised, and the program should exit normally.
}

// Test main function with insufficient command-line arguments
func TestMain_InsufficientArguments(t *testing.T) {
	// Call the main function with insufficient arguments
	os.Args = []string{"change"}
	main()

	// The main function should print the usage message and exit normally.
}

// Test main function with file read error
func TestMain_FileReadError(t *testing.T) {
	// Create a temporary file
	file, err := ioutil.TempFile("", "testfile*.txt")
	if err != nil {
		t.Fatalf("Error creating temporary file: %v", err)
	}
	defer os.Remove(file.Name())

	// Close the file before running the main function
	file.Close()

	// Make the file read-only to cause a read error
	if err := os.Chmod(file.Name(), 0444); err != nil {
		t.Fatalf("Error changing file permissions: %v", err)
	}

	// Call the main function to perform replacement
	os.Args = []string{"change", file.Name(), "Peter", "Daniel"}
	main()

	// The main function should handle the file read error gracefully
	// The error message should be printed, and the program should exit normally.
}

// Test main function with file write error
func TestMain_FileWriteError(t *testing.T) {
	// Create a temporary file
	file, err := ioutil.TempFile("", "testfile*.txt")
	if err != nil {
		t.Fatalf("Error creating temporary file: %v", err)
	}
	defer os.Remove(file.Name())

	// Close the file before running the main function
	file.Close()

	// Make the file read-only to cause a write error
	if err := os.Chmod(file.Name(), 0222); err != nil {
		t.Fatalf("Error changing file permissions: %v", err)
	}

	// Call the main function to perform replacement
	os.Args = []string{"change", file.Name(), "Peter", "Daniel"}
	main()

	// The main function should handle the file write error gracefully
	// The error message should be printed, and the program should exit normally.
}

// Test main function with a large file for performance
func TestMain_PerformanceWithLargeFile(t *testing.T) {
	// Create a large test file with repetitive occurrences of the string to be replaced
	file, err := ioutil.TempFile("", "large_testfile*.txt")
	if err != nil {
		t.Fatalf("Error creating temporary file: %v", err)
	}
	defer os.Remove(file.Name())

	// Generate a large content with 10,000,000 occurrences of "Peter" and write to the file
	largeContent := strings.Repeat("This is Peter's test file. Peter is here!\n", 10000000)
	if _, err := file.Write([]byte(largeContent)); err != nil {
		t.Fatalf("Error writing to the temporary file: %v", err)
	}

	// Close the file before running the main function
	file.Close()

	// Measure the time it takes to perform the replacement
	startTime := time.Now()

	// Call the main function to perform replacement
	os.Args = []string{"change", file.Name(), "Peter", "Daniel"}
	main()

	// Calculate the execution time
	executionTime := time.Since(startTime)

	// Print the execution time (optional, for informational purposes)
	t.Logf("Execution time: %v", executionTime)

	// Add your desired performance criteria here (e.g., must complete within 1 second)
	// Adjust this criteria based on the expected performance of your application.
	maxExecutionTime := 2 * time.Second
	if executionTime > maxExecutionTime {
		t.Errorf("Performance test failed: Execution time exceeds the maximum limit.")
	}
}
