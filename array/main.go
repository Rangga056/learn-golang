package main

import "fmt"

func main() {
	// Array merupakan kumpulan data dengan tipe yang sama yang disimpan dalam sebuah variable
	var names [4]string

	names[0] = "trafalgar"
	names[1] = "d"
	names[2] = "water"
	names[3] = "law"

	fmt.Println(names[0], names[1], names[2], names[3])

	// Pengisian elemen array yang melebihi alokasi awal akan memberikan error
	/*
		  var names [4]string
			names[0] = "trafalgar"
			names[1] = "d"
			names[2] = "water"
			names[3] = "law"
			names[4] = "ez" // baris kode ini menghasilkan error
	*/

	// Inisialisasi nilai awal array
	// Pengisian elemen array bisa dilakukan saat deklarasi variable
	fruits := [4]string{"apple", "grape", "banana", "melon"}

	fmt.Println("Jumlah element \t\t", len(fruits))
	fmt.Println("Isi semua element \t", fruits)

	// Inisialisasi nilai dengan dengan gaya vertikal
	// var fruits [4]string

	// cara horizontal
	// fruits = [4]string{"apple", "grape", "banana", "melon"}
	// cara vertikal
	fruits = [4]string{
		"apple",
		"grape",
		"banana",
		"melon",
	}

	fmt.Println(fruits)

	// Inisialisasi nilai awal array tanpa Jumlah elemen
	numbers := [...]int{2, 3, 2, 4, 3}

	fmt.Println("data array \t:", numbers)
	fmt.Println("jumlah elemen \t:", len(numbers))

	// Array multidimensi
	// array 2 dimensi dengan lenght 3
	numbers1 := [2][3]int{{3, 2, 3}, {3, 4, 5}}
	numbers2 := [2][3]int{{3, 2, 3}, {3, 4, 5}}

	fmt.Println("numbers1", numbers1)
	fmt.Println("numbers2", numbers2)

	// Perulangan elemen array menggunakan for loop
	// fruits := [4]string{"apple", "grape", "banana", "melon"}

	for i := 0; i < len(fruits); i++ {
		fmt.Printf("elemen %d : %s\n", i, fruits[i])
	}

	// Perulangan elemen array menggunakan for - range
	// fruits := [4]string{"apple", "grape", "banana", "melon"}

	for i, fruit := range fruits {
		fmt.Printf("elemen %d : %s\n", i, fruit)
	}

	// Perulangan variable underscore ("_") dalam for - range
	// mengabaikan indeks
	for _, fruit := range fruits {
		fmt.Printf("nama buah: %s\n", fruit)
	}

	// hanya mengambil indeks
	// for i, _ := range fruits { }
	// atau
	// for i := range fruits { }

	// Alokasi elemen array menggunakan make ()
	// digunakan untuk mendeklarasi sekaligus alokasi kapasistas array
	fruits := make([]string, 2) // make([]<tipe-data>, <jumlah-elemen>)
	fruits[0] = "apple"
	fruits[1] = "mangga"

	fmt.Println(fruits)
}
