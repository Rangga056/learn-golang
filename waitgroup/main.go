package main

import (
	"fmt"
	"runtime"
	"sync"
)

func doPrint(wg *sync.WaitGroup, message string) {
	defer wg.Done()
	fmt.Println(message)
}

// sync.WaitGroup digunakan untuk menunggu go routine, cara pengaplikasian tinggal memasukan jumlah goroutine yang dieksekusi, sebagai parameter Add() pada
// objek cetakan sync.WaitGroup, kemudian di tiap akhir goroutine pastikan untuk memanggil method Done(), lalu gunakan Wait() untuk menunggu eksekusi semua goroutine selesai
// sync.WaitGroup bersifat thread safe jadi tidak perlu khawatir terhadap potensi race condition karena variabel bertipe ini digunakan di banyak goroutine secara paralel
// perbedaan sync.WaitGroup dengan channel:
// - WaitGroup tidak perlu goroutine mana saja yang dijalankan cukup tahu jumlah goroutine yang harus selesai
// - Penerapan yang lebih mudah dibanding channel
// - Kegunaan channel adalah untuk komunikasi data antara goroutine sifatnya blocking bisa manfaatkan untuk manage goroutine, sedangkan WaitGroup khusus digunakan untuk sinkronisasi goroutine
// - Performa sync.WaitGroup lebih baik dibanding channel
func main() {
	runtime.GOMAXPROCS(2)

	var wg sync.WaitGroup

	for i := range 5 {
		data := fmt.Sprintf("data %d", i)

		wg.Add(1)
		go doPrint(&wg, data)
	}
	wg.Wait()
}
