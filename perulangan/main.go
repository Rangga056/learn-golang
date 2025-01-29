package main

import "fmt"

func main() {
	// Perulangan menggunakan keyword for
	for i := 0; i < 5; i++ {
		fmt.Println("Angka", i)
	}
	// Perulangan for dengan argumen hanya kondisi
	// i := 0
	//
	// for i < 5 {
	// 	fmt.Println("Angka", i)
	// 	i++
	// }

	// Perulangan for tanpaa argumen
	i := 0

	for {
		fmt.Println("Angka", i)
		i++

		if i == 5 {
			break
		}
	}

	// Perulangan for - range
	xs := "123"

	for i, v := range xs {
		fmt.Println("Index:=", i, "Value=", v)
	}

	ys := [5]int{10, 20, 30, 40, 50} // Array
	for _, v := range ys {
		fmt.Println("Value=", v)
	}

	zs := ys[0:2] // Slice
	for _, v := range zs {
		fmt.Println("Value=", v)
	}

	kvs := map[byte]int{'a': 0, 'b': 1, 'c': 2} // Map
	for k, v := range kvs {
		fmt.Println("Keys=", k, "Value=", v)
	}

	// boleh juga k dan v diabaikan
	for range kvs {
		fmt.Println("Done")
	}

	// Selain itu, bisa juga dengan cukup memnentukan  nilai numerik Perulangan
	for i := range 5 {
		fmt.Println(i) // 01234
	}

	// Penggunaan break dan continue
	// break digunakan untuk menghentikan secara paksa sebuah perulangan
	// continua digunakan untuk memaksa maju ke perulangan atau iterasi selanjutnya
	for i := 1; i <= 10; i++ { // perulangan ini hanya akan menampilkan angka genap kurang dari 10
		if i%2 == 1 {
			continue
		}

		if i > 8 {
			break
		}

		fmt.Println("Angka", i)
	}

	// Perulangan bersarang
	for i := 0; i < 5; i++ {
		for j := i; j < 5; j++ {
			fmt.Print(j, " ")
		}
		fmt.Println()
	}

	// Pemanfaatan label dalam perulangan
outerloop:
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if i == 3 {
				break outerloop
			}
			fmt.Print("matriks [", i, "][", j, "]", "\n")
		}
	}
}
