// 1. RNG (Random Number Generator)
//   - Menghasilkan deret angka acak berdasarkan seed.
//   - Pseudo-RNG: algoritma deterministik, sehingga dengan seed yang sama,
//     deret angka yang dihasilkan pun sama.
//
// 2. Menggunakan Package math/rand
//   - Fungsi rand.NewSource(seed) untuk menentukan seed.
//   - rand.New(source) mengembalikan object randomizer.
//   - Contoh: dengan seed tetap (misal 10) akan selalu menghasilkan angka yang sama.
//
// 3. Seed Unik
//   - Agar hasil acak berbeda tiap eksekusi, gunakan seed berbasis waktu, misalnya time.Now().UnixNano().
//
// 4. Tipe Data Random Lainnya
//   - Selain Int(), ada Float32(), Uint32(), dll.
//   - Untuk range tertentu, gunakan Intn(n) untuk mendapatkan angka antara 0 dan n-1.
//
// 5. Membuat Random String
//   - Pilih karakter dari slice rune (contoh: alfabet) secara acak untuk membentuk string.
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// Contoh dengan seed tetap (angka yang sama setiap kali dijalankan)
	fixedRNG := rand.New(rand.NewSource(10))
	fmt.Println("Seed tetap:")
	fmt.Println("random ke-1:", fixedRNG.Int()) // 5221277731205826435
	fmt.Println("random ke-2:", fixedRNG.Int()) // 3852159813000522384
	fmt.Println("random ke-3:", fixedRNG.Int()) // 8532807521486154107

	// Contoh dengan seed unik (angka acak berbeda tiap eksekusi)
	uniqueRNG := rand.New(rand.NewSource(time.Now().UnixNano()))
	fmt.Println("\nSeed unik:")
	fmt.Println("random ke-1:", uniqueRNG.Int())
	fmt.Println("random ke-2:", uniqueRNG.Int())
	fmt.Println("random ke-3:", uniqueRNG.Int())

	// Generate tipe data lain
	fmt.Println("\nTipe data lain:")
	fmt.Println("random int:", uniqueRNG.Int())
	fmt.Println("random float32:", uniqueRNG.Float32())
	fmt.Println("random uint:", uniqueRNG.Uint32())

	// Mendapatkan angka acak dalam range 0-99
	fmt.Println("\nRandom indeks (0-99):", uniqueRNG.Intn(100))

	// Membuat random string dengan panjang tertentu
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	randomString := func(length int) string {
		s := make([]rune, length)
		for i := range s {
			s[i] = letters[uniqueRNG.Intn(len(letters))]
		}
		return string(s)
	}
	fmt.Println("\nRandom string (5 karakter):", randomString(5))
}
