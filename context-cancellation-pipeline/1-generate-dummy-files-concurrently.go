package main

import (
	"context"
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
	totalFile       = 3000
	contentLength   = 5000 // each file will contain this many random characters
	timeoutDuration = 3 * time.Second
)

// tempPath is the directory where files will be written
var tempPath = filepath.Join(os.Getenv("TEMP"), "chapter-A.61-pipeline-cancellation-context")

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
	generateFilesWithContext(context.Background())
}

func generateFilesWithContext(ctx context.Context) {
	os.RemoveAll(tempPath)
	os.MkdirAll(tempPath, os.ModePerm)

	done := make(chan int)

	go func() {
		chanFileIndex := generateFileIndexes(ctx)

		createFilesWorker := 1000
		chanFileResult := createFiles(ctx, chanFileIndex, createFilesWorker)

		counterSuccess := 0
		for fileResult := range chanFileResult {
			if fileResult.Err != nil {
				log.Printf("error creating files %s. stack trace: %s", fileResult.FileName, fileResult.Err)
			} else {
				counterSuccess++
			}
		}

		done <- counterSuccess
	}()

	select {
	case <-ctx.Done():
		log.Printf("generation process stopped. %s", ctx.Err())
	case counterSuccess := <-done:
		log.Printf("%d/%d of total files created", counterSuccess, totalFile)
	}
}

// generateFileIndexes creates a channel and sends FileInfo jobs into it, then closes it
func generateFileIndexes(ctx context.Context) <-chan FileInfo {
	chanOut := make(chan FileInfo)
	go func() {
		for i := 0; i < totalFile; i++ {
			select {
			case <-ctx.Done():
				break
			default:
				chanOut <- FileInfo{
					Index:    i,
					FileName: fmt.Sprintf("file-%d.txt", i),
				}
			}
		}
		close(chanOut) // signal no more jobs
	}()
	return chanOut
}

// createFiles launches a fixed number of workers to process incoming jobs from chanIn
// and returns a channel that emits the results for each job
func createFiles(ctx context.Context, chanIn <-chan FileInfo, numberOfWorkers int) <-chan FileInfo {
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
				select {
				case <-ctx.Done():
					break
				default:
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
	ctx, cancel := context.WithTimeout(context.Background(), timeoutDuration)
	defer cancel()
	time.AfterFunc(timeoutDuration, cancel)
	generateFilesWithContext(ctx)

	duration := time.Since(start)
	log.Println("done in", duration.Seconds(), "seconds")
}
