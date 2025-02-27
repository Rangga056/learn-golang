package main

import (
	"fmt"
	"regexp"
)

func main() {
	// Regexp atau regex atau regular expression adalah suatu tehnik yang digunakan untuk pencocokan string yang
	// memiliki pola tertentu untuk pencarian dan pebuhan data string

	text := "banana burget soup"
	// [a-z]+ bermakna semua string yang merupakan alfabet huruf kecil
	// regexp.Compile() akan mengcompile hasil ekspresi tersebut dan menyimpan dalam variable bertipe *regexp.Regexp
	regex, err := regexp.Compile(`[a-z]+`)
	if err != nil {
		fmt.Println(err.Error())
	}

	res1 := regex.FindAllString(text, 2) // hanya mengembalikan 2 data saja
	fmt.Printf("%#v \n", res1)
	// []string{"banana", "burger"}

	res2 := regex.FindAllString(text, -1) // mengembalikan semua data
	fmt.Printf("%#v \n", res2)
	// []string{"banana", "burger", "soup"}

	// Method MatchString() yang digunakan untuk mendeteksi apakah string memenuhi pola regexp
	text2 := "melon grape durian"
	regex2, _ := regexp.Compile(`[a-z]+`)

	isMatch := regex2.MatchString(text2)
	fmt.Println(isMatch)

	// Method FindString() untuk menemukan string yang memenuhi kriteria regexp
	text3 := "banana burger hotdog"
	regex3, _ := regexp.Compile(`[a-z]+`)

	str := regex3.FindString(text3)
	fmt.Println(str) // "banana" fungsi ini hanya mengembalikan 1 hasil saja atau hasil pertama

	// Method FindStringIndex() untuk mencari index string kembalian dari hasil operasi regexp
	text4 := "banana burger soup"
	regex4, _ := regexp.Compile(`[a-z]+`)

	idx := regex4.FindStringIndex(text4)
	fmt.Println(idx) // method ini sama dengan FindString() hanya saja yang dikembalikan hanya indexnya saja

	// Method FindAllString() untuk mencari banyak string yang memenuhi kriteria regexp
	text5 := "banana burger soup"
	regex5, _ := regexp.Compile(`[a-z]+`)

	str1 := regex5.FindAllString(text5, -1) // mengembalikan semua data yang sesuai
	fmt.Println(str1)
	// ["banana", "burger", "soup"]

	str2 := regex5.FindAllString(text5, 1) // mengembalikan hanya 1 data
	fmt.Println(str2)
	// ["banana"]

	// Method ReplaceAllString() untuk mereplace semua string yang memenuhi kriteria regexp denga string lain
	text6 := "banana burger soup"
	regex6, _ := regexp.Compile(`[a-z]+`)

	str3 := regex6.ReplaceAllString(text6, "potato")
	fmt.Println(str3)
	// "potato potato potato"

	// Method ReplaceAllStringFunc() untuk mereplace semua string yang memenuhi regexp dengan kondisi yang bisa ditentukan
	// untuk setiap substring yang akan direplace
	text7 := "banana burger soup"
	regex7, _ := regexp.Compile(`[a-z]+`)

	str4 := regex7.ReplaceAllStringFunc(text7, func(each string) string {
		if each == "burger" {
			return "potato"
		}
		return each
	})
	fmt.Println(str4)
	// "banana potato soup"

	// Method Split() digunakan untuk memisah string dengan pemisah
	text8 := "banana burger soup"

	// Membuat regex untuk mencocokkan karakter "a" dan/atau "b"
	regex8, _ := regexp.Compile(`[a-b]+`)

	// Menggunakan regex untuk memisahkan string berdasarkan karakter yang cocok dengan pola `[a-b]+`
	str8 := regex8.Split(text8, -1)

	// Mencetak hasil split sebagai slice string
	fmt.Printf("%#v \n", str8)
}
