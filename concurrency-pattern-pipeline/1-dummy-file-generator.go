package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"time"
)

const (
	totalFile     = 3000 // Total number of files to generate
	contentLength = 5000 // Length of random string content in each file
)

// tempPath is the directory where the files will be generated.
// It uses the system's TEMP directory and appends a custom folder name.
var tempPath = filepath.Join(os.Getenv("TEMP"), "chapter-A.59-pipeline-temp")

// randomString generates a random string of the specified length
func randomString(length int) string {
	// Seed the random number generator with the current time
	randomizer := rand.New(rand.NewSource(time.Now().Unix()))
	// Set of letters to randomly pick from
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	// Prepare a slice to store the result
	b := make([]rune, length)
	// Fill the slice with random letters
	for i := range b {
		b[i] = letters[randomizer.Intn(len(letters))]
	}

	return string(b) // Convert the rune slice to a string and return it
}

// generateFiles creates 3000 text files with random content
func generateFiles() {
	// Clean up the temp directory first
	os.RemoveAll(tempPath)
	// Create the directory if it doesn't exist
	os.MkdirAll(tempPath, os.ModePerm)

	// Loop to create files
	for i := range totalFile {
		// Define file name based on index
		filename := filepath.Join(tempPath, fmt.Sprintf("file-%d.txt", i))
		// Generate random content for the file
		content := randomString(contentLength)
		// Write the content to the file
		err := os.WriteFile(filename, []byte(content), os.ModePerm)
		if err != nil {
			log.Println("Error writing file", filename)
		}

		// Log progress every 100 files
		if i%100 == 0 && i > 0 {
			log.Println(i, "file created")
		}
	}

	// Log completion message
	log.Printf("%d of total files created", totalFile)
}

func main() {
	log.Println("start") // Log the start of the process
	start := time.Now()  // Record the start time

	generateFiles() // Call function to generate files

	// Calculate and log how long it took
	duration := time.Since(start)
	log.Println("done in", duration.Seconds(), "seconds")
}
