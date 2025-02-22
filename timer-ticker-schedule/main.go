package main

import (
	"fmt"
	"os"
	"time"
)

// Kombinasi Timer dan Goroutine
func timer(timeout int, ch chan<- bool) {
	time.AfterFunc(time.Duration(timeout)*time.Second, func() {
		ch <- true
	})
}

func watcher(timeout int, ch <-chan bool) {
	<-ch
	fmt.Println("\ntime out! no answer more than", timeout, "seconds")
	os.Exit(0)
}

func main() {
	// time.Sleep yang menghentikan program sejenak yang bersifat blocking
	fmt.Println("start")
	time.Sleep(time.Second * 3)
	fmt.Println("after 3 seconds")

	// selain untuk blocking time.Sleep juga bisa digunakan sebagai scheduler sederhana
	// for true { // akan print Hello setiap 1 detik
	// 	fmt.Println("Hello")
	// 	time.Sleep(1 * time.Second)
	// }

	// Fungsi time.NewTimer() yang mengembalikan objek bertipe *time.Timer yang memiliki property c bertipe channel
	// cara kerja fungsi ini adalah setelah jeda waktu yang ditentukan data akan dikirim lewat channel c
	// NOTE: Penggunaan fungsi ini harur diikuti statement untuk penerimaan data dari channel c

	// timer := time.NewTimer(4 * time.Second)
	// fmt.Println("start")
	// <-timer.C
	// fmt.Println("finish")

	// Fungsi time.AfterFunc(), memiliki 2 parameter yaitu durasi timer dan callback func yang dieksekusi setelah
	// waktu timer sudah habis
	ch := make(chan bool)

	time.AfterFunc(4*time.Second, func() {
		fmt.Println("expired")
		ch <- true
	})

	// jika tidak ada transfer data lewat channel maka eksekusi time.AfterFunc() asynchronous hingga baris
	// penerimaan data channel dilakukan

	fmt.Println("start")
	<-ch // setelah data terkirim ke channel (sesuai durasi yang ditentukan)
	fmt.Println("finish")

	// Fungsi time.After(), mirip dengan time.Sleep() namun time.After mengembalikan data channel sehingga perlu
	// <- dalam Penggunaanya
	<-time.After(3 * time.Second)
	fmt.Println("expired")

	// Scheduler menggunakan Ticker dengan fungi time.NewTicker()
	// memiliki argumen durasi waktu dan dapat mengakses property seperti .C yang merupakan channel
	// done := make(chan bool)
	// ticker := time.NewTicker(time.Second)
	//
	// go func() {
	// 	time.Sleep(10 * time.Second)
	// 	done <- true
	// }()
	//
	// for {
	// 	select {
	// 	case <-done:
	// 		ticker.Stop() // stopping the ticker (in this case is 10 seconds)
	// 		// return
	// 	case t := <-ticker.C:
	// 		fmt.Println("Hello !!", t) // print untuk setiap durasi yang sudah ditetapkan (in this case every 1 second)
	// 	}
	// }

	// Kombinasi Timer dan Goroutine
	timeout := 5
	ch2 := make(chan bool)

	go timer(timeout, ch2)
	go watcher(timeout, ch2)

	var input string
	fmt.Println("what is 725/25 ? :")
	fmt.Scan(&input)

	if input == "29" {
		fmt.Println("the answer is right")
	} else {
		fmt.Println("the answer is wrong")
	}
}
