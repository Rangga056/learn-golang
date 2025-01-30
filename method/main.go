package main

import (
	"fmt"
	"strings"
)

type student struct {
	name  string
	grade int
}

// Penerapan Method sayHello() dan getNameAt() untuk struct student
func (s student) sayHello() {
	fmt.Println("Hallo", s.name)
}

func (s student) getNameAt(i int) string {
	return strings.Split(s.name, " ")[i-1]
}

// Method Pointer
// NOTE: Ketika kita melakukan manipulasi nilai pada property lain yang masih satu struct,
// nilai pada property tersebut bisa diubah di-level reference-nya.
func (s student) changeName1(name string) { // tidak akan mengubah nilai dari struct
	fmt.Println("--> on changeName1, name changed to", name)
	s.name = name
}

func (s *student) changeName2(name string) { // menggunaan reference memori dapat mengubah nilai dari data tipe struct
	fmt.Println("--> on changeName2, name changed to", name)
	s.name = name
}

func main() {
	// Method merupakan suatu fungsi yang menempel pada suatu tipe data
	// s1 := student{"John Wick", 21}
	// s1.sayHello()

	// name := s1.getNameAt(2)
	// fmt.Println("nama panggilan :", name)

	s1 := student{"John Wick", 21}
	fmt.Println("s1 before", s1.name)
	// John Wick

	s1.changeName1("Jason Bourne")
	fmt.Println("s1 after changeName1", s1.name)
	// John Wick

	s1.changeName2("ethan hunt") // tidak perlu &s1 golang akan secara otomatis mengkonversi
	fmt.Println("s1 after changeName2", s1.name)
	// ethan hunt
}
