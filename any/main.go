package main

import (
	"fmt"
	"strings"
)

type person struct {
	name string
	age  int
}

func main() {
	// Interface kosong yang dinotasikan dengan interface{} atau any,
	// merupakan tipe data yang sangat spesial karena variable bertipe ini bisa menampung
	// segala jenis data

	// Penggunaan interface{}
	// NOTE: selain menggunakan interface{} bisa juga menggunakan any
	var secret interface{}

	secret = "ethan hunt"
	fmt.Println(secret)

	secret = []string{"apple", "mango", "banana"}
	fmt.Println(secret)

	secret = 12.4
	fmt.Println(secret)

	// Contoh lain menggunakan map yang bisa menampung bermacam tipe data
	data := map[string]interface{}{
		"name":      "ethan hunt",
		"grade":     2,
		"breakfast": []string{"apple", "mango", "banana"},
	}

	fmt.Println(data)

	// var data map[string]any
	//
	// data = map[string]any{
	// 	"name":      "ethan hunt",
	// 	"grade":     2,
	// 	"breakfast": []string{"apple", "manggo", "banana"},
	// }

	// karena variable bertipa interface hanya bisa ditampilkan sebagai bentuk text
	// dari nilai aslinya, kita memerlukan casting ke tipe data aslinya

	var secret2 interface{}

	secret2 = 2
	number := secret2.(int) * 10
	fmt.Println(secret, "multiplied by 10 is :", number)

	secret2 = []string{"apple", "mango", "banana"}
	fruits := strings.Join(secret2.([]string), ",")
	fmt.Println(fruits, "is my favourite fruits")

	// castingc variable interface kosong ke objek pointer
	var man interface{} = &person{name: "Wick", age: 37}
	name := man.(*person).name
	fmt.Println(name)

	// Kombinasi slice, map, interface{}
	person := []map[string]interface{}{
		{"name": "Wick", "age": 23},
		{"name": "Ethan", "age": 23},
		{"name": "Bourne", "age": 22},
	}

	for _, each := range person {
		fmt.Println(each["name"], "age is", each["age"])
	}
}
