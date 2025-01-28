package main

import "fmt"

func main() {
	// Tipe data numerik non Desimal
	var positiveNumber uint64 = 89
	negativeNumber := -1243423644 // otomatis akan menentukan tipe data variable menjadi int32

	// String format %d digunakan untuk memformat data numerik non-Desimal
	fmt.Printf("bilangan positif: %d\n", positiveNumber)
	fmt.Printf("bilangan positif: %d\n", negativeNumber)

	// Tipe data numerik Desimal
	// ada dua yaitu float32 dan float64 yang memiliki perbedaan pada lebar cakupan nilai Desimal
	// yang dapat ditamnpung
	decimalNumber := 2.62

	fmt.Printf("bilangan desimal: %f", decimalNumber)   // %f digunakan untuk memformat desimal 6 digit
	fmt.Printf("bilangan desimal: %.3f", decimalNumber) // %.3f format desimal 3 digit

	// Tipe data Boolean terdiri dari 2 nilai true dan false
	var exist bool = true
	fmt.Printf("exist?: %t \n", exist) // %t untuk format tipe data bool

	// Tipe data String
	var message string = "Halo"
	fmt.Println("message: %s", message)
	// selain menggunakan tanda quote,
	// deklarasi string juga bisa dengan tanda grave accent/backtick (``)
	// dengan keistimewaan string yang dideklarasi adalah
	// di dalam backtick membuat semua karakter di dalamnya tidak escape, termasuk \n
	message := `Nama saya "John Wick".
    Salam kenal.
    Mari belajar "Golang"`

	fmt.Println(message)

	// NIlai nil & Zero value
	// nil berarti memiliki nilai kosong
	// Semua tipe data di atas memiliki zero value atau nilai default tipe data
	// artinya meskipun variable tidak dideklarasi dengan nilai, tetap akan  ada nilianya
}
