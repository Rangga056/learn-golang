package main

import (
	"fmt"
	"runtime"
)

func getAverage(numbers []int, ch chan float64) {
	sum := 0
	for _, e := range numbers {
		sum += e
	}

	ch <- float64(sum) / float64(len(numbers))
}

func getMax(numbers []int, ch chan int) {
	max := numbers[0]
	for _, e := range numbers {
		if max < e {
			max = e
		}
	}
	ch <- max
}

func main() {
	// Ada kalanya kita butuh lebih dari satu channel untuk komunikasi antar goroutine
	// selain menggunakan keyword  <- ada juga cara lain menggunakan select
	// select digunakan untuk mempermudah dalam kontrol penerimaan data via 1 atau lebih channel
	// Cara penggunaan select mirip dengan switch

	// Penerapan keyword select
	runtime.GOMAXPROCS(2)

	numbers := []int{3, 4, 3, 5, 6, 3, 2, 2, 6, 3, 4, 6, 3}
	fmt.Println("Numbers:", numbers)

	ch1 := make(chan float64)
	go getAverage(numbers, ch1)

	ch2 := make(chan int)
	go getMax(numbers, ch2)

	// NOTE: karena ada 2 channel maka perlu siapkan 2x perulangan
	for i := 0; i < 2; i++ {
		// case berjalan setiap ada penerimaan atau transfer data pada channel tersebut
		select {
		case avg := <-ch1:
			fmt.Printf("Avg \t: %.2f \n", avg)
		case max := <-ch2:
			fmt.Printf("Max \t: %d \n", max)
		}
	}
}
