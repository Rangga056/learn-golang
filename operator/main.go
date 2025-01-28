package main

import "fmt"

func main() {
	// Operator Aritmatika
	// contoh
	// var value int = (((2+6)%3)*4 - 2) / 3

	// Operator Perbandingan yang memiliki haisl boolean
	// contoh

	var value int = (((2+6)%3)*4 - 2) / 3
	isEqual := (value == 2)

	fmt.Printf("nilai %d (%t) \n", value, isEqual)

	// Operator Logika
	// digunakan untuk mencari benar tidaknya kombinasi data bertipe bool
	left := false
	right := true

	leftAndRight := left && right
	fmt.Printf("left && right \t(%t) \n", leftAndRight)

	leftOrRight := left || right
	fmt.Printf("left || right \t(%t) \n", leftOrRight)

	leftReverse := !left
	fmt.Printf("!left \t(%t) \n", leftReverse)
}
