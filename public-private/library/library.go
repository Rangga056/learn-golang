package library

import "fmt"

func SayHello(name string) { // Merupakan fungsi public
	fmt.Println("Hello")
	introduce(name) // untuk bisa mengakses fungsi private melalui fungsi public
}

func introduce(name string) { // merupakan fungsi private
	fmt.Println("nama saya", name)
}

// Penggunaan hak akses pada struct dan propertynya
// struct private
type student struct {
	Name  string
	grade int
}

// struct public
// type Student struct {
// 	Name  string
// 	Grade int
// }

// Selain fungsi main(), terdapat juga fungsi spesial yaitu init().
// Fungsi ini otomatis dipanggil saat pertama kali program dijalankan.
// Jika fungsi ini ditulis di package-package lain yang di-import di main,
// maka semua fungsi init() tersebut dipanggil lebih dulu sebelum fungsi main().

var Student = struct {
	Name  string
	Grade int
}{}

func init() {
	Student.Name = "John Wick"
	Student.Grade = 2

	fmt.Println("--> library/library.go imported")
}
