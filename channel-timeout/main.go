package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

func sendData(ch chan<- int) {
	randomizer := rand.New(rand.NewSource(time.Now().Unix()))

	for i := 0; true; i++ {
		ch <- i
		time.Sleep(time.Duration(randomizer.Int()%10+1) * time.Second)
	}
}

func receiveData(ch <-chan int) {
loop:
	for {
		select {
		case data := <-ch: // case saat ada transfer data pada channel ch
			fmt.Print(`receiveData"`, data, `"`, "\n")
		case <-time.After(time.Second * 5): // case saat tidak ada transfer data dari channel ch
			fmt.Println("timout. no activation under 5 seconds")
			break loop
		}
	}
}

func main() {
	// Tehnik channel timeout biasa digunakan untuk kontrol waktu penerimaan data pada channel
	// Penerapaan channel timeout

	runtime.GOMAXPROCS(2)

	messages := make(chan int)

	go sendData(messages)
	receiveData(messages)
}
