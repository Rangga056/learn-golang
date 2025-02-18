package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func validate(input string) (bool, error) {
	if strings.TrimSpace(input) == "" {
		// Pembuatan custom Error menggunakan errors.New()
		return false, errors.New("cannot be empty")
	}
	return true, nil
}

// Penerapan recover
func catch() {
	if r := recover(); r != nil {
		fmt.Println("Error occured", r)
	} else {
		fmt.Println("Application running perfectly")
	}
}

func main() {
	// Error merupakan sebuah tipe data yang memiliki beberapa property yang salah satunya adalah
	// method Error() yang mengembalikan detail pesan error dalam string
	// Error juga termasuk tipe data yang bisa memiliki value nil

	// Penerapan recover
	// NOTE: Untuk menggunakan recover fungsi/closure/IIFE letak recover harus dieksekusi dengan defer
	defer catch()

	var input string
	fmt.Print("Type some number: ")
	fmt.Scanln(&input)

	var number int
	var err error
	number, err = strconv.Atoi(input)

	if err == nil {
		fmt.Println(number, "is number")
	} else {
		fmt.Println(input, "is not number")
		fmt.Println(err.Error())
	}

	// Pembuatan custom Error menggunakan errors.New()
	var name string
	fmt.Print("Type your name: ")
	fmt.Scanln(&name)

	// Penggunaan panic
	// digunakan untuk menampilkan stack trace error sekaligus stop flow dari goroutine
	// NOTE: apapun setelah panic tidak di eksekusi kecuali proses yang sudah didefer sebelumnya

	if valid, err := validate(name); valid {
		fmt.Println("Halo", name)
	} else {
		// contoh penerapan panic
		panic(err.Error())
		// fmt.Println(err.Error())
	}

	// Penggunaan recover yang berguna untuk menghandle panic error
	// pada saat panic error muncl revocer akan take over goroutine yang sedang panic dah side effectnya
	// pesan panic tidak muncul dan eksekusi program tidak error

	// recover dalam blok fungsi spesifikg yang akan handle panic error
	// tanpe menghentikan perulangan itu sendiri
	data := []string{"superman", "aquaman", "wonder woman"}

	for _, each := range data {
		func() {
			// recover untuk IIFE dalam perulangan
			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Panic occured on looping", each, "| message:", r)
				} else {
					fmt.Println("Application running perfectly!")
				}
			}()
			if each == "superman" {
				panic("Some error happen")
			}
		}()
	}
}
