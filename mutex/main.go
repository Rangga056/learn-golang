// =====================
// A.60 - Mutex & Race Condition
// =====================
//
// Race condition adalah kondisi ketika dua atau lebih goroutine
// mengakses dan mengubah data yang sama secara bersamaan tanpa mekanisme pengaman.
// Hal ini menyebabkan data menjadi tidak konsisten atau error yang sulit dilacak.
//
// Contoh umum:
// - Misalnya struct `counter` punya properti `val` yang diakses banyak goroutine.
// - Setiap goroutine menambah nilai `val` secara bersamaan tanpa proteksi.
// - Hasil akhir bisa acak atau lebih kecil dari seharusnya.
//
// Deteksi race condition bisa dilakukan dengan flag -race:
//   go run -race main.go
//
// Jika terjadi race, Go akan memberikan warning di terminal,
// termasuk lokasi kode yang menyebabkan konflik akses.
//
// -----------------------------
// Solusi: sync.Mutex
// -----------------------------
// Untuk mengatasi masalah race condition, kita bisa menggunakan sync.Mutex.
// Mutex adalah pengunci yang memastikan hanya satu goroutine
// yang bisa mengakses bagian kode tertentu dalam satu waktu.
//
// Dua fungsi utama dari Mutex:
// - Lock()   → mengunci akses
// - Unlock() → membuka kunci agar goroutine lain bisa melanjutkan
//
// Saat satu goroutine memanggil Lock(), goroutine lain akan menunggu
// sampai Unlock() dipanggil. Ini mencegah modifikasi data terjadi bersamaan.
//
// -----------------------------
// Cara Penggunaan Mutex
// -----------------------------
//
// 1. Dengan embed Mutex ke dalam struct:
//
//     type counter struct {
//         sync.Mutex
//         val int
//     }
//
//     func (c *counter) Add() {
//         c.Lock()
//         defer c.Unlock() // Gunakan defer agar kunci dilepas otomatis meskipun terjadi error
//         c.val++
//     }
//
//     func (c *counter) Value() int {
//         return c.val
//     }
//
// Cara ini lebih idiomatik di Go karena mutex menjadi bagian langsung dari objek.
//
// 2. Dengan mutex eksternal:
//
//     type counter struct {
//         val int
//     }
//
//     func (c *counter) Add() {
//         c.val++
//     }
//
//     func (c *counter) Value() int {
//         return c.val
//     }
//
//     // Di dalam main():
//     var mtx sync.Mutex
//     mtx.Lock()
//     meter.Add()
//     mtx.Unlock()
//
// Mutex dideklarasikan terpisah, dan dikendalikan manual di luar objek.
//
// -----------------------------
// Catatan Penting:
// -----------------------------
//
// - Race condition bisa terjadi saat membaca maupun menulis data.
// - Gunakan mutex hanya jika ada akses bersamaan (concurrent access).
// - Gunakan defer Unlock() segera setelah Lock() untuk mencegah lupa membuka kunci.
// - Saat semua goroutine sudah selesai (setelah wg.Wait()),
//   biasanya pembacaan data tidak perlu mutex lagi.
//
// -----------------------------
// Kesimpulan:
// -----------------------------
//
// - Race condition = konflik akses data antara goroutine.
// - Solusinya pakai sync.Mutex untuk mengatur akses eksklusif.
// - Gunakan cara yang sesuai:
//     → embed mutex dalam struct jika struct-nya digunakan bersama
//     → mutex eksternal jika kontrol akses dilakukan dari luar
//
// - Selalu jalankan program dengan flag -race saat development,
//   untuk memastikan program bebas dari race condition.
//

package main

import (
	"fmt"
	"runtime"
	"sync"
)

// Struct counter dengan mutex di-embed langsung
type counter struct {
	sync.Mutex
	val int
}

func (c *counter) Add() {
	c.Lock()
	defer c.Unlock() // Gunakan defer agar Unlock tetap terpanggil walaupun terjadi error
	c.val++
}

func (c *counter) Value() int {
	return c.val
}

func main() {
	// Atur jumlah thread maksimal (core) yang digunakan
	runtime.GOMAXPROCS(2)

	var meter counter
	var wg sync.WaitGroup

	// Jalankan 1000 goroutine
	for i := 0; i < 1000; i++ {
		wg.Add(1)

		go func() {
			// Setiap goroutine menambah nilai sebanyak 1000x
			for j := 0; j < 1000; j++ {
				meter.Add()
			}
			wg.Done()
		}()
	}

	// Tunggu semua goroutine selesai
	wg.Wait()

	// Ambil dan tampilkan nilai akhir
	fmt.Println("Final value:", meter.Value())
}
