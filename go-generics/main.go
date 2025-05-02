package main

import "fmt"

// func Sum(numbers []int) int {
// 	var total int
// 	for _, e := range numbers {
// 		total += e
// 	}
// 	return total
// }

// Modifikasi func Sum untuk dapat menampung tipe data slice lainnya diluar []int
// func FuncName[dataType <ComparableType>](params)
func Sum[V int | float32 | float64](numbers []V) V {
	// v dideklarasikan dengan notasi [V int], dimana tipe data v artinya kompatibel atau comparable dengan tipe int
	var total V
	for _, e := range numbers {
		total += e
	}
	return total
}

func SumNumbers1(m map[string]int64) int64 {
	var s int64
	for _, v := range m {
		s += v
	}
	return s
}

// Gunakan keyword comparable untuk kompatibel dengan semua tipe data
func SumNumbers2[K comparable, V int64 | float64](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}

// Generic type Constraint
type Number interface {
	int64 | float64
}

func SumNumbers3[K comparable, V Number](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}

// Struct Generic
type UserModel[T int | float64] struct {
	Name   string
	Scores []T
}

func (m *UserModel[int]) SetScoresA(scores []int) {
	m.Scores = scores
}

func (m *UserModel[float64]) SetScoresB(scores []float64) {
	m.Scores = scores
}

func main() {
	// Generic programming adalah salah satu metode dalam penulisan kode program, di mana tipe data dalam kode didefinisikan menggunakan
	// suatu tipe yang tipe pastinya dituliskan saat kode tersebut di-call atau dieksekusi.
	total1 := Sum([]int{1, 2, 3, 4, 5})
	fmt.Println("total:", total1)

	ints := map[string]int64{"first": 34, "second": 12}
	floats := map[string]float64{"first": 35.98, "second": 26.99}

	fmt.Printf("Generic Sums with Constraint: %v and %v\n",
		SumNumbers2(ints),
		SumNumbers2(floats))

	var m1 UserModel[int]
	m1.Name = "Noval"
	m1.Scores = []int{1, 2, 3}
	fmt.Println("scores:", m1.Scores)

	var m2 UserModel[float64]
	m2.Name = "Noval"
	m2.SetScoresB([]float64{10, 11})
	fmt.Println("scores:", m2.Scores)
}
