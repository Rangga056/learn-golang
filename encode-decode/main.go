package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	// Menggunakan package encoding/base64 untuk kebutuhan encode dan decode data dari bentuk base64 dan sebaliknya

	// Fungsi EncodeToString() untuk encode data dari string ke base64 dan DecodeString() untuk sebaliknya
	// data := "john wick"

	// data bertipe string harus di-casting terlebih dahulu ke dalam bentuk []byte sebelum di-encode dengan hasil data base64 bertipe string
	// encodedString := base64.StdEncoding.EncodeToString([]byte(data))
	// fmt.Println("encoded:", encodedString)

	// decodedByte, _ := base64.StdEncoding.DecodeString(encodedString)
	// decodedString := string(decodedByte)
	// fmt.Println("decoded:", decodedString)

	// Fungsi Encode() dan Decode() sama dengan fungsi sebelumnya hanya data yang dihasilkan bertipe []byte
	// Penggunaan cara ini cukup panjang karena variable penyimpan hasil encode dan decode harus disiapkan dulu, dan
	// harus memiliki lebar data sesuai dengan hasil yang akan ditampung
	// data := "Jason Bourne"

	// encoded := make([]byte, base64.StdEncoding.EncodedLen(len(data)))
	// base64.StdEncoding.Encode(encoded, []byte(data))
	// encodedString := string(encoded)
	// fmt.Println(encodedString)

	// decoded := make([]byte, base64.StdEncoding.DecodedLen(len(encoded)))
	// if _, err := base64.StdEncoding.Decode(decoded, encoded); err != nil {
	// 	fmt.Println(err.Error())
	// }

	// decodedString := string(decoded)
	// fmt.Println(decodedString)

	// Encode dan Decode data URL dengan URLEncoding dibandingkan StdEncoding
	data := "https://kalipara.com/"

	encodedString := base64.URLEncoding.EncodeToString([]byte(data))
	fmt.Println(encodedString)

	decodedByte, _ := base64.URLEncoding.DecodeString(encodedString)
	decodedString := string(decodedByte)
	fmt.Println(decodedString)
}
