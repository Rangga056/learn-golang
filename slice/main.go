package main

import "fmt"

func main() {
	// Slice merupakan reference elemen array
	// Inisialisasi Slice
	// tidak perlu mendefinisikan jumlah elemen saat deklarasi
	fruits := []string{"apple", "grape", "banana", "melon"}
	fmt.Println(fruits[0]) // "apple"

	fruitsA := []string{"apple", "grape"}     // slice
	fruitsB := [2]string{"banana", "melon"}   // array
	fruitsC := [...]string{"papaya", "grape"} // array

	// Hubungan slice dengan array & operasi slice
	// fruits := []string{"apple", "grape", "banana", "melon"}
	newFruits := fruits[0:2] // slice dari 2 elemen pertama pada array fruits

	fmt.Println(newFruits) // ["apple", "grape"]

	// Slice merupakan tipe data reference
	// artinya jika ada slice baru yang terbentuk dari slice lama
	// maka slice baru akan memiliki alamat memori yang sama dengan slice lama
	fruits := []string{"apple", "grape", "banana", "melon"}

	aFruits := fruits[0:3]
	bFruits := fruits[1:4]

	aaFruits := aFruits[1:2]
	baFruits := bFruits[0:1]

	fmt.Println(fruits)   // [apple grape banana melon]
	fmt.Println(aFruits)  // [apple grape banana]
	fmt.Println(bFruits)  // [grape banana melon]
	fmt.Println(aaFruits) // [grape]
	fmt.Println(baFruits) // [grape]

	// Buah "grape" diubah menjadi "pinnaple"
	baFruits[0] = "pinnaple"

	fmt.Println(fruits)   // [apple pinnaple banana melon]
	fmt.Println(aFruits)  // [apple pinnaple banana]
	fmt.Println(bFruits)  // [pinnaple banana melon]
	fmt.Println(aaFruits) // [pinnaple]
	fmt.Println(baFruits) // [pinnaple]

	// Fungsi
	// Fungsi len() untun menghitung julah elemen yang ada
	fruits := []string{"apple", "grape", "banana", "melon"}
	fmt.Println(len(fruits)) // 4

	// Fungsi cap() untuk menghitung lebar/kapasitas maks slice
	fruits := []string{"apple", "grape", "banana", "melon"}
	fmt.Println(len(fruits)) // len: 4
	fmt.Println(cap(fruits)) // cap: 4

	aFruits := fruits[0:3]
	fmt.Println(len(aFruits)) // len: 3
	fmt.Println(cap(aFruits)) // cap: 4

	bFruits := fruits[1:4]
	fmt.Println(len(bFruits)) // len: 3
	fmt.Println(cap(bFruits)) // cap: 3

	// Fungsi append() yang digunakan untuk menambahkan elemen pada slice
	fruits := []string{"apple", "grape", "banana"}
	cfruits := append(fruits, "papaya")

	fmt.Println(fruits)
	fmt.Println(cfruits)
	/*
			* NOTE:
		  * jika len() == cap() maka elemen baru hasil append() merupakan referensi baru
			* jika len() < cap() maka elemen baru akan masuk dalam cakupan kapasitas menjadikan
			*   semua elemen slice lain yang memiliki referensi sama berubah nilainya
	*/

	fruits := []string{"apple", "grape", "banana"}
	bFruits := fruits[0:2]

	fmt.Println(cap(bFruits)) // 3
	fmt.Println(len(bFruits)) // 2

	fmt.Println(fruits)  // ["apple", "grape", "banana"]
	fmt.Println(bFruits) // ["apple", "grape"]

	cFruits := append(bFruits, "papaya")

	fmt.Println(fruits)  // ["apple", "grape", "papaya"]
	fmt.Println(bFruits) // ["apple", "grape"]
	fmt.Println(cFruits) // ["apple", "grape", "papaya"]

	// Fungsi copy(), men-copy elemen slice pada src ke dst
	// copy(dst, src)
	// Jumlah elemen yang tercopy dari src adalah sejumlah lebar dari dst / len(dst)
	dst := make([]string, 3)
	src := []string{"watermelon", "pinnaple", "apple", "orange"}
	n := copy(dst, src)

	fmt.Println(dst) // watermelon pinnaple apple
	fmt.Println(src) // watermelon pinnaple apple orange
	fmt.Println(n)   // 3

	// contoh 2
	dst := []string{"potato", "potato", "potato"}
	src := []string{"watermelon", "pinnaple"}
	n := copy(dst, src)

	fmt.Println(dst) // watermelon pinnaple potato
	fmt.Println(src) // watermelon pinnaple
	fmt.Println(n)   // 2, nilai dst sebelumnya terganti dari hasil copy src

	// Pengaksesan elemen slice dengan 3 indeks
	// merupakan tehnik slicing yang sekaligus menentukan kapasitasnya
	// yang dimana tidak boleh lebih dari kapasitas yang akan di slicing
	fruits := []string{"apple", "grape", "banana"}
	afruits := fruits[0:2]
	bfruits := fruits[0:2:2]

	fmt.Println(fruits)      // ["apple", "grape", "banana"]
	fmt.Println(len(fruits)) // len: 3
	fmt.Println(cap(fruits)) // cap: 3

	fmt.Println(aFruits)      // ["apple", "grape"]
	fmt.Println(len(aFruits)) // len: 2
	fmt.Println(cap(aFruits)) // cap: 3

	fmt.Println(bFruits)      // ["apple", "grape"]
	fmt.Println(len(bFruits)) // len: 2
	fmt.Println(cap(bFruits)) // cap: 2
}
