package main

import (
	"fmt"
	"runtime"
)

// Channel sebagai tipe data parameter
func printMessage(what chan string) {
	fmt.Println(<-what) // gunakan <- untuk mengambil dan print data dari channel
}

func main() {
	// Channel digunakan untuk menghubungkan goroutine satu dengan goroutine lain
	// dengan mekanisme serah terima data
	// Pengiriman dan penerimaan data pada channel bersifat blocking atau synchronous
	// Artinya, statement di-bawah syntax pengiriman dan penerimaan data via channel hanya akan dieksekusi
	// setelah proses serah terima berlangsung dan selesai.

	// Penerapan channel dibuat menggunakan kombinasi keyword make dan chan
	runtime.GOMAXPROCS(2)

	// Buat channel pada variable messages bertipe string
	messages := make(chan string)

	// buat closure untuk store data ke channel messages dengan keyword <-
	sayHelloTo := func(who string) {
		data := fmt.Sprintf("Hello %s", who)
		messages <- data
	}

	go sayHelloTo("John Wick")
	go sayHelloTo("ethan hunt")
	go sayHelloTo("jason bourne")

	// simpah proses goroutine pertama yang selesai ke message1
	// dan lanjut jika proses serah terima data sudah selesai
	message1 := <-messages
	fmt.Println(message1)

	message2 := <-messages
	fmt.Println(message2)

	message3 := <-messages
	fmt.Println(message3)

	fmt.Println("=======================================")

	for _, each := range []string{"wick", "ethan", "bourne"} {
		// eksekusi go routine tidak harus pada fungsi atau closure yang terdefinisi
		go func(who string) {
			data := fmt.Sprintf("hello %s", who)
			messages <- data
		}(each) // each sebagai nilai yang di pass in ke parameter who
	}

	for i := 0; i < 3; i++ {
		printMessage(messages)
	}
}
