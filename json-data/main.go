package main

import (
	"encoding/json"
	"fmt"
)

// struct User ini akan digunakan untuk membuat variable baru yang menampung hasil decode JSON string
type User struct {
	// `json:"Name"`  tag ini digunakan untuk mapping informasi field json ke property struct
	FullName string `json:"Name"`
	Age      int
}

func main() {
	// JSON adalah notasi dasar penulisan data yang umum digunakan untuk komunikasi antar aplikasi/service. JSON sendiri sebenarnya merupakan subset dari javascript
	// Go menyediakan package encoding/json untuk kebutuhan operasi JSON

	// Decode JSON ke variable Objek Struct menggunakan json.Unmarshala
	// json string dikonversi menjadi bentuk Objek, entah dalam bentuk map[string]interface{} ataupun Objek Struct
	jsonString := `{"Name": "Muhammad Rangga", "Age":21}`
	jsonData := []byte(jsonString)

	var data User

	// json.Unmarshal() menerima data json berbentuk []byte, maka dari itu json string perlu di casting ke tipe []byte dahulu sebelum digunakan
	// NOTE: Perhatikan bawa argumen ke-2 harus diisi oleh variable pointer yang akan menampung hasil operasi decoding
	err := json.Unmarshal(jsonData, &data)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("user :", data.FullName)
	fmt.Println("user :", data.Age)

	// Decode JSON ke map[string]interface{} & interface{}
	var data1 map[string]any
	json.Unmarshal(jsonData, &data1)

	fmt.Println("user :", data1["Name"])
	fmt.Println("age :", data1["Age"])

	// variable bertipe interface{} juga bisa digunakan untuk menampung hasil decode
	// dengan catatan saat pengaksesan nilai property harus di casting dulu ke map[string]interface{}
	var data2 interface{}
	json.Unmarshal(jsonData, &data2)

	decodedData := data2.(map[string]interface{})
	fmt.Println("user :", decodedData["Name"])
	fmt.Println("age  :", decodedData["Age"])

	// Decode Array JSON ke Array Objek
	// Operasi decode dari Array JSON ke array/slice objek bisa dilakukan dengan cara yang sama
	jsonStringArray := `[
    {"Name": "john wick", "Age": 27},
    {"Name": "ethan hunt", "Age": 32}
]`
	var data3 []User
	err2 := json.Unmarshal([]byte(jsonStringArray), &data3)
	if err2 != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("user 1:", data3[0].FullName)
	fmt.Println("user 2:", data3[1].FullName)

	// Encode Objek ke JSON string menggunakan json.Marshal()
	// dengan hasil konversi bertipe []byte maka pastikan untuk casting dulu ke tipe string agar bisa di tampilkan bentuk json stringnya
	object := []User{{"Rangga", 21}, {"Falah", 22}}
	jsonData2, err3 := json.Marshal(object)
	if err3 != nil {
		fmt.Println(err3.Error())
		return
	}

	jsonString2 := string(jsonData2)
	fmt.Println(jsonString2)
}
