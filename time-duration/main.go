package main

import (
	"fmt"
	"time"
)

func main() {
	// time.Duration() ini merepresentasikan durasi seperti 1 menit, 2 jam, dll
	// data dengan tipe ini bisa dihasilkan dari operasi pencarian delta atau selisih 2 buah objek time.Time
	start := time.Now()

	time.Sleep(4 * time.Second)

	duration := time.Since(start) // durasi antara waktu dari argumen vs statement time.Since() akan dihitung

	fmt.Println("time elapsed in seconds:", duration.Seconds()) // 4 Second
	fmt.Println("time elapsed in minutes:", duration.Minutes())
	fmt.Println("time elapsed in house:", duration.Hours())

	// Kalkulasi durasi antara 2 objek waktu

	t1 := time.Now()
	time.Sleep(3 * time.Second)
	t2 := time.Now()

	duration2 := t2.Sub(t1) // .Sub() digunakan untuk mencari selisih waktu

	fmt.Println("time elapsed in seconds:", duration2.Seconds())
	fmt.Println("time elapsed in minutes:", duration2.Minutes())
	fmt.Println("time elapsed in hours:", duration2.Hours())

	// Konversi angka dari time.Duration
	// gunakan konstanta time.Duration untuk menciptakan variable/ objek bertipe durasi
	var n time.Duration = 5
	duration3 := n * time.Second
	// atau bisa melalui casting
	n2 := 5
	duration4 := time.Duration(n2) * time.Second

	fmt.Println(duration4, duration3)
}
