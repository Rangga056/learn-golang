package main

import "fmt"

func main() {
	// Closure adalah anynomous function tanpa nama yang disimpan dalam variable
	// Closure disimpan sebagai variable
	// getMinMax := func(n []int) (int, int) {
	// 	var min, max int
	// 	for i, e := range n {
	// 		switch {
	// 		case i == 0:
	// 			max, min = e, e
	// 		case e > max:
	// 			max = e
	// 		case e < min:
	// 			min = e
	// 		}
	// 	}
	// 	return min, max
	// }

	// numbers := []int{2, 3, 4, 3, 4, 2, 3}
	// min, max := getMinMax(numbers)

	// fmt.Printf("data: %v\nmin : %v\nmax : %v\n", numbers, min, max)
	// %v digunakan untuk menampilkan data tanpa melihat tipe datanya

	// Immediately- Invoked Function Expression
	// numbers := []int{2, 3, 0, 4, 3, 2, 0, 4, 2, 0, 3}
	//
	// newNumbers := func(min int) []int {
	// 	var r []int
	// 	for _, e := range numbers {
	// 		if e < min {
	// 			continue
	// 		}
	// 		r = append(r, e)
	// 	}
	// 	return r
	// }(3) // eksekusi langsung fungsi saat deklarasi dengan parameter 3
	//
	// fmt.Println("original number :", numbers)
	// fmt.Println("filtered number :", newNumbers)

	//	Closure bisa juga dengan gaya manifest typing, caranya dengan menuliskan skema closure-nya sebagai tipe data. Contoh:
	//
	// var closure (func (string, int, []string) int)
	//
	//	closure = func (a string, b int, c []string) int {
	//	    // ..
	//	}

	// Contoh penggunaan closure sebagai kembalian dari fungsi
	max := 3
	numbers := []int{2, 3, 0, 4, 3, 2, 0, 4, 2, 0, 3}
	howMany, getNumbers := findMax(numbers, max)
	theNumbers := getNumbers()

	fmt.Println("numbers\t:", numbers)
	fmt.Printf("find \t: %d\n\n", max)

	fmt.Println("found \t:", howMany)    // 9
	fmt.Println("value \t:", theNumbers) // [2 3 0 3 2 0 2 0 3]
}

// Closure sebagai nilai kembalian

func findMax(numbers []int, max int) (int, func() []int) {
	var res []int
	for _, e := range numbers {
		if e <= max {
			res = append(res, e)
		}
	}
	return len(res), func() []int { // returning the length and a function
		return res
	}
}
