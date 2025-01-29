package main

import (
	"fmt"
)

// func main() {
// 	// Fungsi adalah sekumpulan blok kode yang dibungkus dengan nama tertentu.
// 	// Penerapan fungsi
// 	// Fungsi main() sendiri merupakan fungsi utama pada Go yang dieksekusi saat program berjalan
// 	names := []string{"John", "Wick"}
// 	printMessage("halo", names)
// }
//
// func printMessage(message string, arr []string) {
// 	nameString := strings.Join(arr, "")
// 	fmt.Println(message, nameString)
// }

// Fungsi dengan return value

// var randomizer = rand.New(rand.NewSource(time.Now().Unix()))
//
// func main() {
// 	var randomValue int
//
// 	randomValue = randomWithRanga(2, 10)
// 	fmt.Println("random number:", randomValue)
//
// 	randomValue = randomWithRanga(2, 10)
// 	fmt.Println("random number:", randomValue)
//
// 	randomValue = randomWithRanga(2, 10)
// 	fmt.Println("random number:", randomValue)
// }
//
// func randomWithRanga(min, max int) int { // return a value of int data type
// 	value := randomizer.Int()%(max-min+1) + min
// 	return value
// }

// Penggunaan fungsi rand.new()
// var randomizer = rand.new(rand.NewSource(time.Now().Unix()))

// Deklarasi parameter bertipe data sama
// func nameOfFunc(paramA type, paramB type, paramC type) returnType
// func nameOfFunc(paramA, paramB, paramC type) returnType
//
// func randomWithRange(min int, max int) int
// func randomWithRange(min, max int) int

// Penggunaan keyword return untuk menghentikan proeses dalam fungsi

func main() {
	divideNumber(10, 2)
	divideNumber(4, 0)
	divideNumber(8, -4)
}

func divideNumber(m, n int) {
	if n == 0 {
		fmt.Printf("Invalid divider. %d cannot divided by %0\n", m, n)
		return
	}

	res := m / n
	fmt.Printf("%d / %d = %d\n", m, n, res)
}
