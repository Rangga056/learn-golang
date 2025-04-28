package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"sync"
	"time"
)

// totalFile is the number of files to generate
const (
	totalFile     = 3000
	contentLength = 5000 // each file will contain this many random characters
)

// tempPath is the directory where files will be written
var tempPath = filepath.Join(os.Getenv("TEMP"), "chapter-A.60-worker-pool")

// FileInfo carries metadata for each file job and its result
// Index: the sequence number of the job
// FileName: the name of the file to create
// WorkerIndex: which worker handled this job
// Err: any error encountered during file creation

type FileInfo struct {
	Index       int
	FileName    string
	WorkerIndex int
	Err         error
}

// randomString returns a random string of the specified length
func randomString(length int) string {
	randomizer := rand.New(rand.NewSource(time.Now().Unix())) // seed per call
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	b := make([]rune, length)
	for i := range b {
		b[i] = letters[randomizer.Intn(len(letters))]
	}

	return string(b)
}

// generateFiles orchestrates the cleanup, pipelines, and tracking of results
func generateFiles() {
	// ensure a clean temp directory
	os.RemoveAll(tempPath)
	os.MkdirAll(tempPath, os.ModePerm)

	// pipeline 1: generate a channel of file jobs
	chanFileIndex := generateFileIndexes()

	// pipeline 2: launch workers to create files concurrently
	createFilesWorker := 100
	chanFileResult := createFiles(chanFileIndex, createFilesWorker)

	// collect and log results
	var (
		counterTotal   int
		counterSuccess int
	)
	for fileResult := range chanFileResult {
		if fileResult.Err != nil {
			log.Printf("error creating file %s. stack trace: %s", fileResult.FileName, fileResult.Err)
		} else {
			counterSuccess++
		}
		counterTotal++
	}
	log.Printf("%d/%d of total files created", counterSuccess, counterTotal)
}

// generateFileIndexes creates a channel and sends FileInfo jobs into it, then closes it
func generateFileIndexes() <-chan FileInfo {
	chanOut := make(chan FileInfo)
	go func() {
		for i := 0; i < totalFile; i++ {
			chanOut <- FileInfo{
				Index:    i,
				FileName: fmt.Sprintf("file-%d.txt", i),
			}
		}
		close(chanOut) // signal no more jobs
	}()
	return chanOut
}

// createFiles launches a fixed number of workers to process incoming jobs from chanIn
// and returns a channel that emits the results for each job
func createFiles(chanIn <-chan FileInfo, numberOfWorkers int) <-chan FileInfo {
	chanOut := make(chan FileInfo)
	wg := new(sync.WaitGroup)

	// add one for each worker to the wait group
	wg.Add(numberOfWorkers)

	// start workers
	for workerIndex := 0; workerIndex < numberOfWorkers; workerIndex++ {
		// each worker is a goroutine listening on chanIn
		go func(workerIndex int) {
			defer wg.Done() // mark this worker done when function returns
			for job := range chanIn {
				// perform the file creation
				filePath := filepath.Join(tempPath, job.FileName)
				content := randomString(contentLength)
				err := os.WriteFile(filePath, []byte(content), os.ModePerm)

				log.Println("worker", workerIndex, "working on", job.FileName)

				// send result
				chanOut <- FileInfo{
					FileName:    job.FileName,
					WorkerIndex: workerIndex,
					Err:         err,
				}
			}
		}(workerIndex)
	}

	// close chanOut once all workers finish
	go func() {
		wg.Wait()
		close(chanOut)
	}()

	return chanOut
}

func main() {
	log.Println("start")
	start := time.Now()

	generateFiles() // run the pipeline

	duration := time.Since(start)
	log.Println("done in", duration.Seconds(), "seconds")
}
