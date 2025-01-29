package main

import "fmt"

// Embedded Struct
// adalah mekanisme untuk menempelkan sebuah struct sebagai property struct lain
// type person struct {
// 	name string
// 	age  int
// }
//
// type student struct {
// 	grade int
// 	person
// }

// Embedded Struct dengan nama property yang sama
type person struct {
	name string
	age  int
}

type student struct {
	person
	age   int
	grade int
}

// Nested Struct
// type student struct {
// 	person struct {
// 		name string
// 		age  int
// 	}
// 	grade   int
// 	hobbies []string
// }

// Deklarasi dan inisialisasi struct secara horizontal
// type person struct{name string; age int;  hobbies []string}

// Tag property dalam struct
// merupakan informasi tambahan opsional yang bisa ditambahkan property struct
// biasa dimanfaatkan untuk keperluan encode/decode data
// type person struct {
// 	name string `tag1`
// 	age  int    `tag2`
// }

// Pembuatan struct baru juga bisa dilakukan lewat teknik type alias.
type People1 struct {
	name string
	age  int
}

type People2 = struct { // merupakan alias dari anonymous struct
	name string
	age  int
}

// Contoh
// type student struct {
// 	// Default value 0
// 	name  string
// 	grade int
// }

func main() {
	// Struct adalah kumpulan definisi variable dan atau fungsi yang
	// di bungkus sebagai tipe data baru dengan nama tertentu.
	// Mirip seperti map hanya saja key nya sudah didefinisikan di awal
	// NOTE: konsep Struct di golang mirip dengan konsel class pada OOP

	// var s1 student
	// s1.name = "John Wick"
	// s1.grade = 2

	// fmt.Println("name: ", s1.name)
	// fmt.Println("grade: ", s1.grade)

	// Inisialisasi Object Struct
	// var s1 student
	// atau
	// s1 := student{}
	// s1.name = "wick"
	// s1.grade = 2
	//
	// s2 := student{"ethan", 2}
	//
	// s3 := student{name: "jason"}
	//
	// fmt.Println("student 1 :", s1.name)
	// fmt.Println("student 2 :", s2.name)
	// fmt.Println("student 3 :", s3.name)
	//
	// Penentuan nilai property bisa dilakukan tidak berurutan
	// s4 := student{name: "wayne", grade: 2}
	// s5 := student{grade: 2, name: "bruce"}

	// Variable Object pointer
	// Object yang dibuat dengan tipe data struct bisa diambil nilai pointernya
	// dan disimpan pada variable Object yang bertipe struct pointer
	// s1 := student{name: "wick", grade: 2}
	//
	// var s2 *student = &s1
	// merupakan variable pointer cetakan struct student dengan nilai referensi s1
	// fmt.Println("student 1, name :", s1.name)
	// fmt.Println("student 4, name :", s2.name)
	//
	// s2.name = "ethan"
	// fmt.Println("student 1, name :", s1.name)
	// fmt.Println("student 4, name :", s2.name)

	// Embedded Struct
	// Embedded struct adalah mutable, nilai property-nya bisa diubah
	// s1 := student{}
	// s1.name = "wick"
	// s1.age = 21
	// s1.grade = 2
	//
	// fmt.Println("name  :", s1.name)
	// fmt.Println("age   :", s1.age)
	// fmt.Println("age   :", s1.person.age)
	// fmt.Println("grade :", s1.grade)

	// Embedded Struct dengan nama property yang sama
	// s1 := student{}
	// s1.name = "wick"
	// s1.age = 21        // age of student
	// s1.person.age = 22 // age of person
	//
	// fmt.Println(s1.name)
	// fmt.Println(s1.age)
	// fmt.Println(s1.person.age)

	// Pengisian nilai sub-Struct
	// bisa dilakukan dengan langsung memasukkan variable object yang tercetak dari struct yang sama
	// p1 := person{name: "wick", age: 21}
	// s1 := student{person: p1, grade: 2}
	//
	// fmt.Println("name  :", s1.name)
	// fmt.Println("age   :", s1.age)
	// fmt.Println("grade :", s1.grade)
	//
	// Anonymous Struct
	// adalah struct yang tidak dideklarasikan di awal sebagai tipe data baru,
	// namun langsung ketika pembuatan object (cocok untuk variable object sekali pakai)
	// s1 := struct {
	// 	person
	// 	grade int
	// }{}
	// s1.person = person{"wick", 21}
	// s1.grade = 2
	//
	// fmt.Println("name  :", s1.person.name)
	// fmt.Println("age   :", s1.person.age)
	// fmt.Println("grade :", s1.grade)

	// dalam pembuatan anonymous struct adalah, deklarasi harus diikuti dengan inisialisasi {}.
	// anonymous struct tanpa pengisian property
	// s1 := struct {
	// 	person
	// 	grade int
	// }{}
	//
	// // anonymous struct dengan pengisian property
	// s2 := struct {
	// 	person
	// 	grade int
	// }{
	// 	person: person{"wick", 21},
	// 	grade:  2,
	// }

	// Kombinasi Slice & Struct
	// allStudents := []person{
	// 	{name: "Wick", age: 23},
	// 	{name: "Ethan", age: 23},
	// 	{name: "Bourne", age: 22},
	// }
	//
	// for _, student := range allStudents {
	// 	fmt.Println(student)
	// }

	// Inisialisasi Slice anonymous struct
	allStudents := []struct {
		person
		grade int
	}{
		{person: person{"wick", 21}, grade: 2},
		{person: person{"ethan", 22}, grade: 3},
		{person: person{"bond", 21}, grade: 3},
	}

	for _, student := range allStudents {
		fmt.Println(student)
	}

	// Deklarasi anonymous struct menggunakan keyword var
	// var student struct {
	// 	person
	// 	grade int
	// }
	//
	// student.person = person{"wick", 21}
	// student.grade = 2

	// Deklatrasi  anonymous struct juga bisa dilakukan disertai inisialisasi data
	// hanya deklarasi
	// var student struct {
	// 	grade int
	// }

	// deklarasi sekaligus inisialisasi
	// student := struct {
	// 	grade int
	// }{
	// 	12,
	// }

	// Deklarasi dan inisialisasi struct secara horizontal
	// var p1 = struct { name string; age int } { age: 22, name: "wick" }
	// var p2 = struct { name string; age int } { "ethan", 23 }

	// Type Alias
	// type namaAlias = targetAlias
	type people = person

	p1 := person{"wick", 21}
	fmt.Println(p1)

	p2 := people{"wick", 21}
	fmt.Println(p2)

	People := people{"wick", 21}
	fmt.Println(person(People))

	Person := person{"wick", 21}
	fmt.Println(people(Person))
}
