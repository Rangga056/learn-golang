package main

import "fmt"

type student struct {
	name        string
	height      float64
	age         int
	isGraduated bool
	hobbies     []string
}

func main() {
	data := student{
		name:        "wick",
		height:      182.5,
		age:         26,
		isGraduated: false,
		hobbies:     []string{"eating", "sleeping"},
	}

	// Untuk memformat data numerik, menjadi bentuk string numerik berbasis 2 (biner).
	fmt.Printf("%b\n", data.age)

	// Untuk memformat data numerik yang merupakan kode unicode, menjadi bentuk string karakter unicode-nya.
	fmt.Printf("%c\n", 1400)
	fmt.Printf("%c\n", 1235)

	// untuk memformat data numerik, menjadi bentuk string numerik berbasis 10 (basis bilangan yang kita gunakan).
	fmt.Printf("%d\n", data.age)

	// Untuk memformat data numerik desimal ke dalam bentuk notasi numerik standar Scientific notation.
	fmt.Printf("%e\n", data.height)
	fmt.Printf("%E\n", data.height)

	// Untuk memformat data numerik desimal, dengan lebar desimal bisa ditentukan. Secara default lebar digit desimal adalah 6 digit.
	fmt.Printf("%f\n", data.height)
	fmt.Printf("%.9f\n", data.height)
	fmt.Printf("%.2f\n", data.height)
	fmt.Printf("%.f\n", data.height)

	// Untuk memformat data numerik desimal, dengan lebar desimal bisa ditentukan.
	// pada %g lebar digit desimal adalah sesuai dengan datanya, tidak bisa dicustom seperti pada %f.
	fmt.Printf("%e\n", 0.123123123123)
	fmt.Printf("%f\n", 0.123123123123)
	fmt.Printf("%g\n", 0.123123123123)

	// Untuk memformat data numerik, menjadi bentuk string numerik berbasis 8 (oktal).
	fmt.Printf("%o\n", data.age)

	// Digunakan untuk memformat data pointer, mengembalikan alamat pointer referensi variabel.
	fmt.Printf("%p\n", &data.name) // Contoh output: 0x2081be0c0

	// Digunakan untuk escape string.
	fmt.Printf("%q\n", `" name \ height "`) // Contoh output: "\" name \\ height \""

	// Digunakan untuk memformat data string.
	fmt.Printf("%s\n", data.name) // Contoh output: wick

	// Digunakan untuk memformat data boolean.
	fmt.Printf("%t\n", data.isGraduated) // Contoh output: false

	// Berfungsi untuk mengambil tipe variabel yang akan diformat.
	fmt.Printf("%T\n", data.name)        // Contoh output: string
	fmt.Printf("%T\n", data.height)      // Contoh output: float64
	fmt.Printf("%T\n", data.age)         // Contoh output: int32
	fmt.Printf("%T\n", data.isGraduated) // Contoh output: bool
	fmt.Printf("%T\n", data.hobbies)     // Contoh output: []string

	// Digunakan untuk memformat data apa saja (termasuk interface{}).
	fmt.Printf("%v\n", data) // Contoh output: {wick 182.5 26 false [eating sleeping]}

	// Digunakan untuk memformat struct, menampilkan nama dan nilai property.
	fmt.Printf("%+v\n", data) // Contoh output: {name:wick height:182.5 age:26 isGraduated:false hobbies:[eating sleeping]}

	// Digunakan untuk memformat struct, menampilkan struktur dan nilai property.
	fmt.Printf("%#v\n", data) // Contoh output: main.student{name:"wick", height:182.5, age:26, isGraduated:false, hobbies:[]string{"eating", "sleeping"}}

	// Contoh anonymous struct
	data2 := struct {
		name   string
		height float64
	}{
		name:   "wick",
		height: 182.5,
	}
	fmt.Printf("%#v\n", data2) // Contoh output: struct { name string; height float64 }{name:"wick", height:182.5}

	// Digunakan untuk memformat data numerik ke heksadesimal.
	fmt.Printf("%x\n", data.age) // Contoh output: 1a
}
