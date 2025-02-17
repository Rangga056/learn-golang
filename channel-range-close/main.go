package main

import (
	"fmt"
	"runtime"
)

// kita buat func sendMessage() untuk mengirim data via channel
// NOTE: chan<- menerapkan level akses channel tersebut hanya bisa digunakan untuk send data
func sendMessage(ch chan<- string) {
	for i := 0; i < 20; i++ {
		ch <- fmt.Sprintf("data %d", i)
	}
	close(ch)
}

// NOTE: <-chan menerapkan level akses channel tersebut hanya bisa digunakan untuk receive data
func printMessage(ch <-chan string) {
	for message := range ch {
		fmt.Println(message)
	}
}

func main() {
	// Proses transfer data dari banyak channel bisa dipermudah dengan
	// kombinasi keyword for - range

	// Penerapan for - range - close
	runtime.GOMAXPROCS(2)

	messages := make(chan string)
	go sendMessage(messages)
	printMessage(messages)
}
