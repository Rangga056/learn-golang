package main

import (
	"fmt"
	"strings"
)

func main() {
	// Fungsi variadic adalah fungsi dengan parameter bisa menampung niali sejenis yang
	// tidak terbatas jumlahnya

	// Penerapan
	avg := calculate(2, 4, 3, 5, 4, 3, 3, 5, 5, 3)
	msg := fmt.Sprintf("Rata-rata : %.2f", avg)
	// Sprintf() sama seperti Printf namun hanya mengembalikan nilai dalam bentuk string
	fmt.Println(msg)

	// slice juga bisa digunakan sebagai argument pada fungsi variadic
	// numbers := []int{2, 4, 3, 5, 4, 3, 3, 5, 5, 3}
	// avg := calculate(numbers...) // NOTE: tambahkan ... pada akhir
	// msg := fmt.Sprintf("Rata-rata : %.2f", avg)

	fmt.Println(msg)

	yourHobbies("wick", "sleeping", "eating")
	// atau
	hobbies := []string{"sleeping", "eating"}
	yourHobbies("wick", hobbies...)
}

func calculate(numbers ...int) float64 { // ...int menandakan parameter variadic tipe int
	var total int = 0

	for _, number := range numbers {
		total += number
	}

	avg := float64(total) / float64(len(numbers))

	return avg
}

// Fungsi dengan paramter biasa denga variadic
func yourHobbies(name string, hobbies ...string) {
	hobbiesAsString := strings.Join(hobbies, ",")

	fmt.Printf("Hello my name is: %s\n", name)
	fmt.Printf("My hobbies are: %s\n", hobbiesAsString)
}
