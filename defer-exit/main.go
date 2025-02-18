package main

import (
	"fmt"
	"os"
)

func orderSomeFood(menu string) {
	defer fmt.Println("Terimakasih, silahkan tunggu")
	if menu == "pizza" {
		fmt.Print("Pilihan tepat!,", " ")
		fmt.Print("Pizza di tempat kami paling enak!", "\n")
		return
	}

	fmt.Println("Pesanan anda:", menu)
}

func main() {
	// Defer digunakan untuk mengakhirkan eksekusi sebuah statement tepat sebelum blok fungsi selesai
	// Exit digunakan untuk menghentikan program secara paksa beda dengan return yang menghentikan blok kode

	// Penerapan keyword defer
	defer fmt.Println("Hello") // defer akan dieksekusi di akhir blok fungsi lokasi defer
	fmt.Println("Welcome")

	orderSomeFood("burger")
	orderSomeFood("pizza")

	// Penerapan os.Exit()
	os.Exit(1) // akan mengakhiri program secara paksa pada saat itu juga
}
