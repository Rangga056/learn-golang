package main

import (
	"fmt"
	"math"
)

func main() {
	var diameter float64 = 15
	area, circumference := calculate(diameter) // karena fungsi return 2 nilai

	fmt.Printf("luas lingkaran\t\t: %.2f \n", area)
	fmt.Printf("keliling lingkaran\t: %.2f \n", circumference)
}

// Penerapan
func calculate(d float64) (float64, float64) {
	// hitung luas
	area := math.Pi * math.Pow(d/2, 2)
	// hitung keliling
	circumference := math.Pi * d

	// kembalikan 2 nilai
	return area, circumference
}

// Fungsi dengan predefined return value
func calculate(d float64) (area float64, circumference float64) {
	area = math.Pi * math.Pow(d/2, 2)
	circumference = math.Pi * d

	return // karena variable balik sudah ditentukan di awal maka hanya cukup return saja
}
