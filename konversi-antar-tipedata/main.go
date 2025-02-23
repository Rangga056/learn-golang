package main

import (
	"fmt"
	"strconv"
)

func main() {
	// Konversi menggunakan strconv
	// strconv.Atoi() untuk konversi dari string => int, return hasil konversi dan err
	str := "123"
	num, err := strconv.Atoi(str)

	if err == nil {
		fmt.Println(num)
	}

	// strconv.Itoa() yang merupakan kebalikan dari strconv.Atoi() dari int => string
	num2 := 123
	str2 := strconv.Itoa(num2)
	fmt.Println(str2)

	// strconv.ParseInt() digunakan untuk konversi string ke tipe numerik non-desimal dengan lebar data yang
	// bisa ditentukan
	// Base 10 (10): Standard decimal numbers.
	// Base 2 (2): Binary numbers (e.g., "101" → 5).
	// Base 8 (8): Octal numbers (e.g., "12" → 10 in decimal).
	// Base 16 (16): Hexadecimal numbers (e.g., "A" → 10 in decimal).
	// Base 0 (0): Auto-detects the base from the string format
	str3 := "456"
	if num3, err2 := strconv.ParseInt(str3, 10, 64); err2 == nil {
		fmt.Println(num3)
	}

	// strconv.FormatInt() yang digunakan untuk konversi int64 => string dengan basis yang bisa ditentukan
	num4 := int64(655)
	str4 := strconv.FormatInt(num4, 8)
	fmt.Println(str4)

	// strconv.ParseFloat() konversi dari string => float
	str5 := "12.34"
	if num5, err4 := strconv.ParseFloat(str5, 32); err4 == nil {
		println(num5)
	}

	// strconv.FormatFloat() konversi dari float64 => string dengan format eksponen, lebar digit desima, dan bitsize
	num6 := 123.123
	str6 := strconv.FormatFloat(num6, 'f', 6, 64)
	fmt.Println(str6)

	// strconv.ParseBool() konversi string => bool
	str7 := "true"
	if bul, err5 := strconv.ParseBool(str7); err5 == nil {
		fmt.Println(bul)
	}

	// strconv.FormatBool konversi bool => string
	bul2 := true
	str8 := strconv.FormatBool(bul2)
	fmt.Println(str8)

	// Konversi data menggunakan tehnik casting
	a := float64(24) // int => float64
	fmt.Println(a)   // 24

	b := int32(24.00) // float32 => int32
	fmt.Println(b)    // 24

	// Casting string <-> byte untuk mendapatkan slice byte dari sebuah data string
	text1 := "halo"
	c := []byte(text1)

	fmt.Printf("%d %d %d %d \n", c[0], c[1], c[2], c[3])
	// 104 97 108 111

	byte1 := []byte{104, 97, 108, 111}
	s := string(byte1)

	fmt.Printf("%s \n", s)
	// halo

	// Type assertions pada type any atau interface{}

	data := map[string]interface{}{
		"nama":    "john wick",
		"grade":   2,
		"height":  156.5,
		"isMale":  true,
		"hobbies": []string{"eating", "sleeping"},
	}

	fmt.Println(data["nama"].(string))
	fmt.Println(data["grade"].(int))
	fmt.Println(data["height"].(float64))
	fmt.Println(data["isMale"].(bool))
	fmt.Println(data["hobbies"].([]string))

	// Tipe asli data pada variabel interface{} bisa diketahui dengan cara meng-casting ke tipe type,\
	// namun casting ini hanya bisa dilakukan pada switch
	for _, val := range data {
		switch val := val.(type) {
		case string:
			fmt.Println(val)
		case int:
			fmt.Println(val)
		case float64:
			fmt.Println(val)
		case bool:
			fmt.Println(val)
		case []string:
			fmt.Println(val)
		default:
			fmt.Println(val.(int))
		}
	}
}
