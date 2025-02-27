package main

import (
	"crypto/sha1"
	"fmt"
	"time"
)

// Salting pada hash
func doHashUsingSalt(text string) (string, string) {
	salt := fmt.Sprintf("%d", time.Now().UnixNano())
	saltedText := fmt.Sprintf("textL '%x', salt: %s", text, salt)
	sha := sha1.New()
	sha.Write([]byte(saltedText))
	encrypted := sha.Sum(nil)

	return fmt.Sprintf("%x", encrypted), salt
}

func main() {
	// Hash adalah algoritma enkripsi satu arah untuk mengubah text menjadi deretan karakter acak.
	// Jumlah karakter acak selalu sama
	// NOTE: hasil hash tidak bisa dikembalikan ke text asli

	// SHA1 atau Secure Hash Algorithm 1 merupakan salah satu algoritma hashing yang sering digunakan
	// dengan hasil data yang memiliki lebar 20 byte atau 160 bit,
	// biasanya ditampilkan dalam bentuk hexadesimal 40 digit

	// Penerapan Hash SHA1 dengan package crypto/sha1
	// text := "this is secret"
	// sha := sha1.New()
	// Write() untuk set data yang akan di hash dalam bentuk []byte
	// sha.Write([]byte(text))
	// Sum() untuk eksekusi proses hash data yang sudah di hash dalam bentuk []byte
	// encrypted := sha.Sum(nil)
	// encryptedString := fmt.Sprintf("%x", encrypted) // gunakan Sprintf untuk mengambil bentuk hexadesimal string
	// fmt.Println(encryptedString)

	// Metode salting pada Has SHA1 merupakan data acak yang digabung pada data asli sebelum proses hash
	// salt digunakan untuk mencegah serangan menggunakan metode pencocokan data-data yang hasil hashnya sama
	// (dictionary attack)

	text := "this is secret"
	fmt.Printf("original: %s\n\n", text)

	hashed1, salt1 := doHashUsingSalt(text)
	fmt.Printf("hashed 1: %s\n salt1: %s \n\n", hashed1, salt1)

	hashed2, salt2 := doHashUsingSalt(text)
	fmt.Printf("hashed 2: %s\n salt2: %s \n\n", hashed2, salt2)

	hashed3, salt3 := doHashUsingSalt(text)
	fmt.Printf("hashed 3: %s\n salt3: %s \n\n", hashed3, salt3)
}
