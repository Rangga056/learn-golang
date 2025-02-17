package main

import (
	"fmt"
	"reflect"
)

type student struct {
	Name  string
	Grade int
}

func (s *student) getPropertyInfo() {
	reflectValue := reflect.ValueOf(s)

	if reflectValue.Kind() == reflect.Ptr {
		// Check jika tipe data == pointer dan gunakan .Elem() untuk mengambil data object reflect aslinya
		reflectValue = reflectValue.Elem()
	}

	reflectType := reflectValue.Type()

	for i := 0; i < reflectValue.NumField(); i++ {
		fmt.Println("nama      :", reflectType.Field(i).Name)
		fmt.Println("tipe data :", reflectType.Field(i).Type)
		fmt.Println("nilai     :", reflectValue.Field(i).Interface())
		fmt.Println("")
	}
}

func (s *student) SetName(name string) {
	s.Name = name
}

func main() {
	// Reflection adalah tehnik untuk inspeksi variabel
	// Digunakan melalui package reflect
	// dari banyak fungsi dari package reflect ada 2 yang sering Digunakan
	// - reflect.ValueO() yang mengembalikan object berisi informasi yang berhubungan dengan nilai/data yang diinspeksi
	// - reflect.TypeOf() yang mengembalikan object berisi informasi tipe data variabel yang diinspeksi

	// reflect.ValueOf()
	var number int = 23
	reflectValue := reflect.ValueOf(number)

	fmt.Println("tipe variable : ", reflectValue.Type())

	// fungsi yang digunakan harus sesuai dengan tipe data nilai yang ditampung variabel agar tidak error
	if reflectValue.Kind() == reflect.Int {
		fmt.Println("nilai variable : ", reflectValue.Int())
	}

	number2 := 12
	reflectValue2 := reflect.ValueOf(number2)

	fmt.Println("Tipe data : ", reflectValue2.Type())
	// Jika nilai hanya perlu ditampilkan sebagai output, bisa menggunakan .Interface()
	fmt.Println("nilai variable : ", reflectValue2.Interface())

	// Pengaksesan informasi property variable object
	s1 := &student{Name: "john wick", Grade: 2}
	s1.getPropertyInfo()

	fmt.Println("Nama :", s1.Name)

	reflectValue3 := reflect.ValueOf(s1)
	// untuk mengambil reflection method kita gunakan .MethodByName()
	// dan .Call() untuk memanggil method tersebut
	method := reflectValue3.MethodByName("SetName")
	// method .Call() harus menerima parameter bertipe []reflect.value karena
	// method ini memiliki banyak parameter
	// gunakan reflect.ValueOf() untuk konversi nilai "wick" menjadi object reflect.Value()
	method.Call([]reflect.Value{
		reflect.ValueOf("wick"),
	})

	fmt.Println("name :", s1.Name)
}
