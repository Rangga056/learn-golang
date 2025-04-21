package main

import (
	"crypto/md5"    // Provides MD5 hashing
	"fmt"           // Formatting strings
	"log"           // Logging errors and info
	"os"            // File and environment operations
	"path/filepath" // Building and walking file paths
	"time"          // Measuring elapsed time
)

// tempPath is where your files live—same temp folder you used before.
var tempPath = filepath.Join(os.Getenv("TEMP"), "chapter-A.59-pipeline-temp")

// proceed walks through every file under tempPath, computes its MD5 hash,
// and renames the file to include that hash (de-duplicating by content).
func proceed() {
	counterTotal := 0   // How many files we’ve seen
	counterRenamed := 0 // How many we successfully renamed

	// Walk the directory tree rooted at tempPath
	err := filepath.Walk(tempPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			// Propagate any filesystem error
			return err
		}
		if info.IsDir() {
			// Skip directories
			return nil
		}

		counterTotal++ // Found a file

		// Read the file’s entire content into memory
		buf, err := os.ReadFile(path)
		if err != nil {
			return err
		}

		// Compute its MD5 checksum, render as hex string
		sum := fmt.Sprintf("%x", md5.Sum(buf))

		// Build the new filename, e.g. file‑<md5sum>.txt
		destinationPath := filepath.Join(tempPath, fmt.Sprintf("file-%s.txt", sum))

		// Rename (move) the file to the hashed name
		err = os.Rename(path, destinationPath)
		if err != nil {
			return err
		}

		counterRenamed++ // Successfully renamed
		return nil       // Continue walking
	})
	if err != nil {
		// If anything bubbled up, log it
		log.Println("ERROR!", err.Error())
	}

	// Report how many files got renamed out of how many scanned
	log.Printf("%d/%d files Renamed", counterRenamed, counterTotal)
}

func main() {
	log.Println("start") // Kick off logging
	start := time.Now()  // Timestamp the start

	proceed() // Run the renaming pipeline

	// Measure and log total duration
	duration := time.Since(start)
	log.Println("done in", duration.Seconds(), "Seconds")
}
