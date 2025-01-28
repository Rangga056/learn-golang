package main

import "fmt"

func main() {
	// Digunakan untuk mengontrol alur eksekusi flow program
	// menggunakan nilai bertipe boolean sebagai acuan
	// Terdapat 2 macam keyword if else dan switch

	// if, else if & else
	point := 8

	if point == 10 {
		fmt.Println("Lulus dengan sempurna")
	} else if point > 5 {
		fmt.Println("Lulus")
	} else if point == 4 {
		fmt.Println("Hampir lulus")
	} else {
		fmt.Printf("Tidak lulus. Nilai anda %d\n", point)
	}

	// Variable temporary pada if-else
	// Variable yang hanya bisa digunakan dalam blok/scope seleksi kondisi
	point := 8840.0

	// variable percent hanya dapat diakses didalam blok seleksi kondisi
	if percent := point / 100; percent >= 100 {
		fmt.Printf("%.1f%s perfect!\n", percent, "%")
	} else if percent >= 70 {
		fmt.Printf("%.1f%s good\n", percent, "%")
	} else {
		fmt.Printf("%.1f%s not bad\n", percent, "%")
	}

	// Seleksi Kondisi menggunakan keyword switch - case
	point := 6

	switch point {
	case 8:
		fmt.Println("perfect")
	case 7:
		fmt.Println("awesome")
	default:
		fmt.Println("not bad")
	}

	// Pemanfaatan case untuk banyak kondisi
	// case dapat menampung banyak kondisi dipisahkan dengan tanda koma (,)
	point := 6

	switch point {
	case 8:
		fmt.Println("Perfect")
	case 7, 6, 5, 4:
		fmt.Println("awesome")
	default:
		fmt.Println("not bad")
	}

	// tanda kurung kurawal {} dapat digunakan pada case & default namun optional
	// bagus dipakai jika di dalamnya ada banyak statement
	point := 6

	switch point {
	case 8:
		fmt.Println("perfect")
	case 7, 6, 5, 4:
		fmt.Println("awesome")
	default:
		{
			fmt.Println("not bad")
			fmt.Println("you can be better!")
		}
	}

	// switch dengan gaya if else
	point := 6

	switch {
	case point == 8:
		fmt.Println("Perfect")
	case (point < 8) && (point > 3):
		fmt.Println("awesome")
	default:
		{
			fmt.Println("not bad")
			fmt.Println("you need to learn more")
		}
	}

	// Penggunaan keyword fallthrough dalam switch
	// digunakan untuk memaksa proses pengecekan untuk tetap diteruskan ke case selanjutnya
	// tanpa menghiraukan nilai kondisinya
	point := 8

	switch {
	case point == 8:
		fmt.Println("perfect")
	case (point < 8) && (point > 3):
		fmt.Println("awesome")
		fallthrough
	case point < 5:
		fmt.Println("you need to learn more")
	default:
		{
			fmt.Println("not bad")
			fmt.Println("you need to learn more")
		}
	}

	// Seleksi kondisi bersarang
	// adalah seleksi kondisi dalam seleksi kondisi lain
	point := 10

	if point > 7 {
		switch point {
		case 10:
			fmt.Println("perfect")
		default:
			fmt.Println("nice!")
		}
	} else {
		if point == 5 {
			fmt.Println("not bad")
		} else if point == 3 {
			fmt.Println("keep trying")
		} else {
			fmt.Println("you can do it")
			if point == 0 {
				fmt.Println("try harde!")
			}
		}
	}
}
