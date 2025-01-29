package main

import (
	"fmt"
	"strings"
)

func main() {
	// NOTE: Di Go fungsi bisa dijadikan sebagai tipe data variable,
	// maka sangat memungkinkan untuk dijadikan parameter dari fungsi
	data := []string{"wick", "jason", "ethan"}

	dataContainsO := filter(data, func(each string) bool {
		return strings.Contains(each, "o")
	}) // cari elemen dari data yang memiliki huruf o

	dataLenght5 := filter(data, func(each string) bool {
		return len(each) == 5
	}) // cek elemen dari data yang memiliki 5 huruf

	fmt.Println("data asli \t\t:", data)
	// data asli : [wick jason ethan]

	fmt.Println("filter ada huruf \"o\"\t:", dataContainsO)
	// filter ada huruf "o" : [jason]

	fmt.Println("filter jumlah huruf \"5\"\t:", dataLenght5)
	// filter jumlah huruf "5" : [jason ethan]
}

// Penerapan
func filter(data []string, callback func(string) bool) []string {
	var result []string
	for _, each := range data {
		if filtered := callback(each); filtered {
			result = append(result, each)
		}
	}
	return result
}

// Alias Skema Closure
// kita bisa membuat alias fungsi untuk tipe data baru menggunakan keyword type
// type FilterCallback func(string)bool
//
// func filter(data []string, callback FilterCallback) []string {
//   // ....
// }
