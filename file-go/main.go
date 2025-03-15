package main

import (
	"fmt"
	"io"
	"os"
)

var path = "/home/rangga/Documents/test/test.txt"

func isError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}

	return (err != nil)
}

// Untuk Create File
func createFile() {
	// cek apakah file sudah ada
	_, err := os.Stat(path) // os.Stat() mengembalikan 2 data yaitu iinformasi path dan error

	// buat file baru jika belum ada
	if os.IsNotExist(err) {
		file, err := os.Create(path)
		// os.Create() mengembalikan objek bertipe *os.file dimana statusnya otomatis openn sehingga setelah operasi file selesai
		// file harus di close dengan methd file.Close()
		if isError(err) {
			return
		}
		defer file.Close()
	}

	fmt.Println("==> file berhasil dibuat", path)
}

// Untuk Edit File
func writeFile() {
	// Buka file dengan level akses READ & WRITE
	file, err := os.OpenFile(path, os.O_RDWR, 0644)
	if isError(err) {
		return
	}
	defer file.Close()

	// Tulis data ke dalam file
	_, err = file.WriteString("halo\n")
	if isError(err) {
		return
	}
	_, err = file.WriteString("nama saya muhammad rangga\n")
	if isError(err) {
		return
	}

	// Simpan perubahan pada file
	err = file.Sync()
	if isError(err) {
		return
	}
	fmt.Println("==> file berhasil di isi")
}

// Untuk Read File
func readFile() {
	file, err := os.OpenFile(path, os.O_RDONLY, 0644)
	if isError(err) {
		return
	}
	defer file.Close()

	text := make([]byte, 1024)
	// variable text disiapkan bertipe slice []byte dengan alokasi elemmen 1024 yang bertugas menampung data hasil file.Read()
	// proses pembacaa dilakukan terus menerus hingga akhir
	// error yang muncul ketike eksekusi file.Read() akan difilter, ketike errr selain io.EOF maka proses baca file akan lanjut
	// NOTE: io.EOF sendiri menandakan bahwa file yang sedang dibaca adalah baris terakhir isi atau end of file
	for {
		n, err := file.Read(text)
		if err != io.EOF {
			if isError(err) {
				break
			}
		}
		if n == 0 {
			break
		}
	}

	if isError(err) {
		return
	}

	fmt.Println("==> file berhasil dibaca")
	fmt.Println(string(text))
}

// Untuk Delete File
func deleteFile() {
	err := os.Remove(path)
	if isError(err) {
		return
	}

	println("==> file berhasil dihapus")
}

func main() {
	// Untuk membuat file di Go dapat dilakukan dengan menggunakan fungsi os.Create() disertasi path file sebagai argument
	// jika ternyata file akan dibuat sudah ada sebelumnya maka operasi os.Create() akan menimpa file yang sudah ada dengan yang baru
	// Untuk menghindari penimpaan file bisa digunakan fungsi os.IsNotExist()
	createFile()
	writeFile()
	readFile()
	deleteFile()
}
