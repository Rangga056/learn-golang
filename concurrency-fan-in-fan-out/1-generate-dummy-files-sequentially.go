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
	totalFile     = 3000
	contentLength = 5000
)

var tempPath = filepath.Join(os.Getenv("TEMP"), "chapter-A.60-worker-pool")

func randomString(length int) string {
	randomizer := rand.New(rand.NewSource(time.Now().Unix()))
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	b := make([]rune, length)
	for i := range b {
		b[i] = letters[randomizer.Intn(len(letters))]
	}

	return string(b)
}

func generateFile() {
	os.RemoveAll(tempPath)
	os.MkdirAll(tempPath, os.ModePerm)

	for i := range totalFile {
		filename := filepath.Join(tempPath, fmt.Sprintf("file-%d.txt", i))
		content := randomString(contentLength)
		err := os.WriteFile(filename, []byte(content), os.ModePerm)
		if err != nil {
			log.Println("Error writing file", filename)
		}
		log.Println(i, "files created")
	}
	log.Printf("%d of total files created", totalFile)
}

func main() {
	log.Println("Start")
	start := time.Now()

	generateFile()

	duration := time.Since(start)
	log.Println("done in", duration.Seconds(), "Seconds")
}
