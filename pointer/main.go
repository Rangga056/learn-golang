package main

import "fmt"

func main() {
	// Pointer adalah reference atau alamat memori
	// Contoh penerapan
	// var number *int
	// var name *string
	// NOTE:
	// - Variable biasa bisa diambil nilai pointernya dengan menambahkan & biasa disebut referencing
	// - NIlai asli variable pointer juga bisa diambil dengan menambahkan * biasa disebut dereferencing

	// var numberA int = 4
	// var numberB *int = &numberA
	// dideklarasikan bertipe pointer int dengan nilai awal referensi variable numberA

	// fmt.Println("numberA (value):", numberA)     // 4
	// fmt.Println("numberA (address) :", &numberA) // 0xc20800a220

	// fmt.Println("numberB (value)   :", *numberB) // 4
	// fmt.Println("numberB (address) :", numberB)  // 0xc20800a220

	// NOTE: jika salah satu variable pointer diubah nilainya maka variable dengan referensi memori
	// yang sama akan ikut berubah
	var numberA int = 4
	var numberB *int = &numberA

	fmt.Println("numberA (value)   :", numberA)
	fmt.Println("numberA (address) :", &numberA)
	fmt.Println("numberB (value)   :", *numberB)
	fmt.Println("numberB (address) :", numberB)

	fmt.Println("")

	numberA = 5

	fmt.Println("numberA (value)   :", numberA)
	fmt.Println("numberA (address) :", &numberA)
	fmt.Println("numberB (value)   :", *numberB)
	fmt.Println("numberB (address) :", numberB)

	// Parameter Pointer
	// Parameter juga bisa dirancang sebagai pointer dengan mendeklarasikan Parameter sebagai pointer
	number := 4
	fmt.Println("before", number) // 4

	change(&number, 10)
	fmt.Println("after", number) // 10
}

func change(original *int, value int) {
	*original = value
}
