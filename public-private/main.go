package main

import (
	// Pemanfaatan alias saat import package

	f "fmt"

	// Import dengan prefix tanda titik
	// Di Go, komponen yang berada di package lain yang di-import bisa dijadikan
	// se-level dengan komponen package peng-import,
	// caranya dengan menambahkan tanda titik (.) sebelum nama package yang akan diimport
	"public-private/library"
	// . "public-private/library"
)

func main() {
	// NOTE: Pada Go terdapat 2 jenis level akses
	// 1. Level akses Public atau Exported dimana package bisa digunakan di package  lain
	// 2. Level akses Private atau Unexported dimana hanya bisa diakses di package yang sama
	// Cara menentukan level akses mengacu pada character case huruf pertama nama fungsi,
	// variable, struct, atau lainnya.
	// Huruf awal Kapital => public dan Huruf awal kecil => private

	// library.SayHello()
	// library.introduce("ethan") // akan muncul error karena mencoba mengakses fungsi private

	// Mengubah cara pemanggilan fungsi private melalui fungsi public
	// library.SayHello("ethan")

	// s1 := library.student{"ethan", 21}
	// akan muncul error karena student merupakan struct private
	// s1 := library.Student{"ethan", 21}
	// fmt.Println("name", s1.Name)
	// fmt.Println("grade", s1.Grade)
	// akan error karena property grade pada struct merupakan private

	// Penggunakan import dengan prefix "."
	// s1 := Student{Name: "ethan", Grade: 21}
	// f.Println("name", s1.Name)
	// f.Println("grade", s1.Grade)

	// Mengakses property dalam file di package yang sama
	sayHello("ethan")

	f.Printf("Name : %s\n", library.Student.Name)
	f.Printf("Grade : %d\n", library.Student.Grade)
}
