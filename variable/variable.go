package main

import "fmt"

func main() {
	var firstName string = "john"
	// Deklarasi variable dengan tipe data atau manifest typing
	// var <nama-variabel> <tipe-data>
	// var <nama-variabel> <tipe-data> = <nilai>

	// var lastName string
	// lastName = "Wick"
	// Deklarasi variable tanpa tipe data atau type inference
	lastName := "Wick"
	// tanda := hanya digunakan di awal saat deklarasi untuk assigment nilai tetap menggunakan =
	lastName = "ethan"
	lastName = "bourne"

	// Deklarasi multi variable, deklarasi banyak variable sekaligus
	var first, second, third string
	first, sesecond, third = "satu", "dua", "tiga"

	// Pengisian nilai juga bisa dilakukan bersamaan saat deklarasi
	var fourth, fifth, sixth string = "empat", "lima", "enam"

	// atau lebih ringkas
	seventh, eight, ninth := "tujuh",
		"delapan", "sembilan"

	// Multi variable juga bisa untuk tipe data yang berbeda beda
	one, isFriday, twoPointTwo, say := 1, true, 2.2, "hello"

	// Go memiliki aturan tidak boleh ada variable yang menganggur jika ada maka akan muncul error saat kompilasi
	// (_) undescore variable adalah reserved variable yang bisa dimanfaatkan untuk menampung nilai yang tidak terpakai
	// biasa digunakan untuk menampung nilai balik fungsi yang tidak digunakan
	_ = "belajar Golang"
	_ = "Golang itu mudah"
	name, _ := "John", "wick"

	// Deklarasi variable dengan menggunakan keyword new
	// variable name menampung data bertipe pointer string sehingga saat di print akan menampilkan alamat memori
	// untuk menampilkan nilai asli perlu di dereference terlebih dahulu dengan menambahkan *
	name := new(String)
	fmt.Println(name)  // 0x20818a220
	fmt.Println(*name) // ""

	/*
		* 	Deklarasi variable dengan menggunakan keyword make
		* fungsi make() hanya bisa digunakan untuk beberapa variable saja yaitu:
				* channel
				* slice
				* map
	*/

	fmt.Printf("halo %s %s!\n", firstName, lastName)

	// ketiga baris kode dibawah akan memberikan hasil yang sama
	fmt.Printf("halo john wick!\n")
	fmt.Printf("halo %s %s!\n", firstName, lastName)
	fmt.Println("halo", firstName, lastName+"!")
}
