package main

import (
	"fmt"
	"runtime"
)

func print(till int, message string) {
	for i := 0; i < till; i++ {
		fmt.Println((i + 1), message)
	}
}

func main() {
	// Goroutine mirip seperti thread, tetapi lebih ringan (hanya ~2kB memori per goroutine).
	// Goroutine bersifat asynchronous, sehingga tidak saling menunggu dengan goroutine lain.
	// Meskipun goroutine ringan, eksekusi banyak goroutine (misal 1 juta) dapat membebani RAM.
	// Selain ukuran goroutine, kompleksitas kode/logika di dalamnya juga memengaruhi konsumsi resource.
	// Goroutine adalah bagian penting dalam concurrent programming di Go dan dapat dijalankan di multi-core processor,
	// di mana semakin banyak core yang aktif, semakin cepat eksekusinya.

	// Mendapatkan jumlah core yang tersedia
	maxCores := runtime.NumCPU()
	fmt.Println("Jumlah core yang tersedia:", maxCores)

	// Penerapan Goroutine dengan menambahkan keyword go
	runtime.GOMAXPROCS(2) // digunakan untuk menentukan jumlah core yang diaktifkan untuk mengeksekusi program

	go print(5, "halo")
	print(5, "apa kabar")

	var input string
	fmt.Scanln(&input) // mengakibatkan proses jalannya app berhenti sampai user menekan enter
}
