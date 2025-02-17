package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	// Unbuffered channel: Pengiriman dan penerimaan data bersifat blocking.
	// Harus ada goroutine penerima untuk setiap pengiriman data.

	// Buffered channel: Memiliki kapasitas buffer yang ditentukan.
	// Pengiriman data bersifat asynchronous (tidak blocking) selama buffer belum penuh.
	// Jika buffer penuh, pengiriman baru bisa dilakukan setelah ada data yang diambil oleh penerima.
	// NOTE: Proses transfer data pada buffered channel adalah asynchronous ketika jumlah data yang dikirim tidak melebihi batas buffer. Namun pada bagian channel penerimaan data selalu bersifat synchronous atau blocking.

	// Penerapan buffered channel mirip dengan channel namun ditambahkan angka buffer pada argumen make()
	runtime.GOMAXPROCS(2)

	messages := make(chan int, 3) // karena dimulai dari 0 jadi jumlah buffer 4

	go func() {
		for {
			i := <-messages
			fmt.Println("receive data", i)
		}
	}()

	for i := 0; i < 5; i++ {
		fmt.Println("send data", i)
		messages <- i
	}

	time.Sleep(1 * time.Second)
}
