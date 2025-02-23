package main

import (
	"fmt"
	"strings"
)

func main() {
	// strings.Contains() untuk deteksi apakah string pada paremeter kedua merupakan bagian dari parameter pertama
	ifExist := strings.Contains("john wick", "wick")
	fmt.Println(ifExist)

	// strings.HasPrefix() digunakan untuk deteksi apakah string parameter pertama diawali string parameter kedua
	isPrefix1 := strings.HasPrefix("john wick", "jo")
	fmt.Println(isPrefix1) // true

	isPrefix2 := strings.HasPrefix("john wick", "wi")
	fmt.Println(isPrefix2) // false

	// Fungsi strings.HasSuffix() untuk deteksi apakah string parameter pertama diakhiri string parameter kedua
	isSuffix1 := strings.HasSuffix("john wick", "ic")
	fmt.Println(isSuffix1) // false

	isSuffix2 := strings.HasSuffix("john wick", "ck")
	fmt.Println(isSuffix2) // true

	// Fungsi strings.Count() untuk menghitung jumlah karakter tertentu
	howMany := strings.Count("ethan hunt", "t")
	fmt.Println(howMany) // 2

	// Fungsi strings.Index() untuk mencari posisi indeks sebuah string
	index1 := strings.Index("ethan hunt", "ha")
	fmt.Println(index1) // 2

	// Fungsi strings.Replace() untuk mengganti bagian atau keseluruhan string
	text := "banana"
	find := "a"
	replaceWith := "o"

	newText1 := strings.Replace(text, find, replaceWith, 1)
	fmt.Println(newText1) // "bonana"

	newText2 := strings.Replace(text, find, replaceWith, 2)
	fmt.Println(newText2) // "bonona"

	newText3 := strings.Replace(text, find, replaceWith, -1)
	// or
	newText4 := strings.ReplaceAll(text, find, replaceWith)

	fmt.Println(newText3, newText4) // "bonono"

	// Fungsi strings.Repeat() untuk mengulang string parameter pertama sesuai jumlah yang ditentukan
	str := strings.Repeat("na", 4)
	fmt.Println(str) // "nananana"

	// Fungsi strings.Split() untuk memisah string pada tanda pemisah di parameter kedua
	string1 := strings.Split("the dark knight", " ")
	fmt.Println(string1) // output: ["the", "dark", "knight"]

	string2 := strings.Split("batman", "")
	fmt.Println(string2) // output: ["b", "a", "t", "m", "a", "n"]

	// Fungsi strings.Join() untuk mengembalikan slice string menjadi sebuah string dengan pemisah
	data := []string{"banana", "papaya", "tomato"}
	str2 := strings.Join(data, "-")
	fmt.Println(str2) // "banana-papaya-tomato"

	// Fungsi strings.ToLower() dan strings.ToUpper untuk mengubah kapitalisasi string
	str3 := strings.ToLower("aLAy")
	fmt.Println(str3) // "alay"
	str4 := strings.ToUpper("eat!")
	fmt.Println(str4) // "EAT!"
}
