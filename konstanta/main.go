package main

import "fmt"

func main() {
	// Konstanta merupakan nilai yang tidak bisa diubah setelah dideklarasikan
	// cara penerapan mirip dengan variable namun menggunakan keyword const bukan var
	const firstName = "John"
	fmt.Print("Halo ", firstName, "!\n")

	// Tehnik interference bisa diterapakan cukup dengan menghilangkan tipe data saat deklarasi
	const lastName = "wick"
	fmt.Print("Nice to meet you ", lastName, "!\n")
	// fmt.Print() memiliki peran yang sama dengan fmt.Println
	// namun tidak menghasilkan baris baru di akhir outputnya
	// perbedaan lainnya saat pemanggilan fungsi Print() argumen parameter yang ditulis
	// akan diprint tanpa pemisah
	fmt.Println("john wick")
	fmt.Println("john", "wick")

	fmt.Print("john wick\n")
	fmt.Print("john ", "wick\n")
	fmt.Print("john", " ", "wick\n")

	// Deklarasi multi Konstanta
	const (
		square         = "kotak"
		isToday  bool  = true
		numeric  uint8 = 1
		floatNum       = 2.2
	)

	// contoh deklarasi multi Konstanta dengan tipe data sama
	const (
		a = "Konstanta"
		b // jika tipe data tidak dituliskan maka akan saam seperti tipe data diatasnya
	)

	// contoh gabungan
	const (
		today string = "senin"
		sekarang
		isToday2 = true
	)

	// contoh deklarasi multi Konstanta dalam satu baris
	const satu, dua = 1, 2
	const three, four string = "tiga", "empat"
}
