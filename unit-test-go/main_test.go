package main

import (
	"testing"

	"github.com/stretchr/testify/assert" // Menggunakan 3rd party library untuk testing
)

// Inisialisasi object dan expected value untuk perhitungan
var (
	kubus                Kubus   = Kubus{4} // Membuat object Kubus dengan sisi 4
	expectedVolume       float64 = 64       // Volume seharusnya (4^3)
	expectedArea         float64 = 96       // Luas permukaan seharusnya (4^2 * 6)
	expectedCirumference float64 = 48       // Keliling seharusnya (4 * 12)
)

/* ============================
   UNIT TEST BAWAAN PACKAGE testing
   ============================ */

// Fungsi unit test harus diawali dengan "Test" dan menerima parameter *testing.T

func TestCalculateVolume(t *testing.T) {
	t.Logf("Volume : %.2f", kubus.Volume())
	// Logf berfungsi mencetak log informasi selama testing. Mirip fmt.Printf

	if kubus.Volume() != expectedVolume {
		t.Errorf("Salah! Seharusnya %.2f", expectedVolume)
		// Errorf menampilkan error dan menandakan bahwa test gagal
	}
}

func TestCalculateArea(t *testing.T) {
	t.Logf("Area : %.2f", kubus.Luas())

	if kubus.Luas() != expectedArea {
		t.Errorf("Salah! Seharusnya %.2f", expectedArea)
	}
}

func TestCalculateCircumference(t *testing.T) {
	t.Logf("Circumference : %.2f", kubus.Keliling())

	if kubus.Keliling() != expectedCirumference {
		t.Errorf("Salah! Seharusnya %.2f", expectedCirumference)
	}
}

/* ============================
   BENCHMARK TESTING
   ============================ */

// Fungsi benchmark diawali dengan "Benchmark" dan menerima parameter *testing.B
// Benchmark digunakan untuk menguji performa / kecepatan eksekusi fungsi

func BenchmarkCalculateArea(b *testing.B) {
	for i := 0; i < b.N; i++ {
		kubus.Luas()
	}
}

// Cara menjalankan: go test -bench=.
// Contoh output: 30000000  51.1 ns/op → 30 juta kali, 51 nanosecond per operasi

/* ============================
   UNIT TEST DENGAN testify
   ============================ */

// Testify adalah library eksternal untuk mempermudah assertion
// assert.Equal(t, actual, expected, message) → jika actual != expected, test gagal

func TestHitungLuas(t *testing.T) {
	assert.Equal(t, kubus.Luas(), expectedArea, "Perhitungan luas salah")
}

func TestHitungVolume(t *testing.T) {
	assert.Equal(t, kubus.Volume(), expectedVolume, "Perhitungan volume salah")
}

func TestHitungKeliling(t *testing.T) {
	assert.Equal(t, kubus.Keliling(), expectedCirumference, "Perhitungan keliling salah")
}
