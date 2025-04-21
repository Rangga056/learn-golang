package main

import (
	"crypto/md5"    // to compute MD5 hashes
	"fmt"           // for string formatting
	"log"           // for logging progress and errors
	"os"            // for file I/O and env vars
	"path/filepath" // for building and walking file paths
	"sync"          // to synchronize multiple goroutines
	"time"          // to measure elapsed time
)

// tempPath is the directory where we read from and rename files.
var tempPath = filepath.Join(os.Getenv("TEMP"), "chapter-A.59-pipeline-temp")

// FileInfo carries everything we need through the pipeline:
// - FilePath: where the file lives
// - Content: raw bytes we read
// - Sum:    computed MD5 hash of Content
// - IsRenamed: flag indicating whether the rename succeeded
type FileInfo struct {
	FilePath  string
	Content   []byte
	Sum       string
	IsRenamed bool
}

// readFiles kicks off a goroutine that walks tempPath, reads each file’s bytes,
// and emits a FileInfo{FilePath, Content} down the returned channel.
// When done, it closes the channel.
func readFiles() <-chan FileInfo {
	chanOut := make(chan FileInfo)

	go func() {
		// Walk the directory tree
		err := filepath.Walk(tempPath, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				// Abort on filesystem error
				return err
			}
			if info.IsDir() {
				// Skip directories
				return nil
			}

			// Read entire file into memory
			buf, err := os.ReadFile(path)
			if err != nil {
				return err
			}

			// Send the raw FileInfo downstream
			chanOut <- FileInfo{
				FilePath: path,
				Content:  buf,
			}
			return nil
		})
		if err != nil {
			log.Println("ERROR:", err)
		}

		// Signal that no more files will be sent
		close(chanOut)
	}()

	return chanOut
}

// getSum reads FileInfo structs from chanIn, computes MD5(Content),
// writes Sum back into the struct, then sends it on chanOut.
// Closes chanOut when chanIn is drained.
func getSum(chanIn <-chan FileInfo) <-chan FileInfo {
	chanOut := make(chan FileInfo)

	go func() {
		for fileInfo := range chanIn {
			// Compute MD5 hash of file content
			fileInfo.Sum = fmt.Sprintf("%x", md5.Sum(fileInfo.Content))
			// Pass along enriched FileInfo
			chanOut <- fileInfo
		}
		close(chanOut)
	}()

	return chanOut
}

// mergeChanFileInfo merges multiple <-chan FileInfo inputs into a single output channel.
// It spins up one goroutine per input, copies all data into chanOut, and uses a WaitGroup
// to close chanOut only after all inputs are drained.
func mergeChanFileInfo(chanInMany ...<-chan FileInfo) <-chan FileInfo {
	wg := new(sync.WaitGroup)
	chanOut := make(chan FileInfo)

	wg.Add(len(chanInMany))
	for _, eachChan := range chanInMany {
		go func(c <-chan FileInfo) {
			// Drain this channel
			for fi := range c {
				chanOut <- fi
			}
			// Signal this input is done
			wg.Done()
		}(eachChan) // insert eachChan as self argument
	}

	// Once all readers finish, close the merged channel
	go func() {
		wg.Wait()
		close(chanOut)
	}()

	return chanOut
}

// rename reads FileInfo from chanIn, computes the new filename based on Sum,
// attempts to rename the file on disk, marks IsRenamed accordingly,
// then emits the updated FileInfo on chanOut. Closes chanOut at end.
func rename(chanIn <-chan FileInfo) <-chan FileInfo {
	chanOut := make(chan FileInfo)

	go func() {
		for fileInfo := range chanIn {
			// Build new path: chapter-A.59-pipeline-temp/file-<md5sum>.txt
			newPath := filepath.Join(tempPath, fmt.Sprintf("file-%s.txt", fileInfo.Sum))
			// Attempt filesystem rename
			err := os.Rename(fileInfo.FilePath, newPath)
			// Mark success or failure
			fileInfo.IsRenamed = (err == nil)
			// Send result downstream
			chanOut <- fileInfo
		}
		close(chanOut)
	}()

	return chanOut
}

func main() {
	log.Println("start")
	start := time.Now()

	// Pipeline stage 1: read raw files
	chanFileContent := readFiles()

	// Pipeline stage 2: compute MD5 sums in parallel
	// WARNING: reading from chanFileContent multiple times like this means only
	// the first getSum will actually see all files; the others will compete and
	// drain the channel. Proper fan-out requires a tee mechanism.
	chanFileSum1 := getSum(chanFileContent)
	chanFileSum2 := getSum(chanFileContent)
	chanFileSum3 := getSum(chanFileContent)
	// Merge the three sum streams back into one
	chanFileSum := mergeChanFileInfo(chanFileSum1, chanFileSum2, chanFileSum3)

	// Pipeline stage 3: rename files in parallel
	chanRename1 := rename(chanFileSum)
	chanRename2 := rename(chanFileSum)
	chanRename3 := rename(chanFileSum)
	chanRename4 := rename(chanFileSum)
	// Merge the four renaming streams
	chanRename := mergeChanFileInfo(chanRename1, chanRename2, chanRename3, chanRename4)

	// Final: consume all rename results and tally successes
	counterRenamed := 0
	counterTotal := 0
	for fileInfo := range chanRename {
		if fileInfo.IsRenamed {
			counterRenamed++
		}
		counterTotal++
	}

	log.Printf("%d/%d files renamed", counterRenamed, counterTotal)

	// Report total elapsed time
	duration := time.Since(start)
	log.Println("done in", duration.Seconds(), "seconds")
}
