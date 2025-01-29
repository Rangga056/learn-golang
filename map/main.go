package main

import "fmt"

func main() {
	// Map adalah tipe data asosiatif yang berbentuk key-value pair
	// Penggunaan Map
	var chicken map[string]int // map[<tipe-data key>]<tipe-data value>
	chicken = map[string]int{} // NOTE: karena default value map = nil maka penting untuk menambah {}

	chicken["januari"] = 50
	chicken["februari"] = 40

	fmt.Println("Januari", chicken["januari"]) // Januari 50
	fmt.Println("mei", chicken["mei"])         // mei 0

	// Inisialisasi nilai map
	var data map[string]int
	data["one"] = 1
	// akan muncul error

	data = map[string]int{}
	data["one"] = 1
	// tidak ada error

	// Nilai variable map bisa didefinisikan di awal
	// cara horizontal
	chicken1 := map[string]int{"januari": 50, "februari": 40}

	// cara vertical
	chicken2 := map[string]int{
		"januari":  50,
		"februari": 40,
	}

	// Map juga bisa diinisialisasikan menggunakan make() dan new()
	chicken1 := map[string]int
	chicken2 := make(map[string]int)
	chicken3 := *new(map[string]int)

	// Iterasi item map menggunakan for - range
	chicken := map[string]int{
		"januari":  50,
		"februari": 40,
		"maret":    34,
		"april":    67,
	}

	for key, val := range chicken {
		fmt.Println(key, "  \t:", val)
	}

	// Menghapus item Map
	// menggunakan fungsi delete() menghapus item dengan key tertentu
	chicken := map[string]int{"januari": 50, "februari": 40}

	fmt.Println(len(chicken)) // 2
	fmt.Println(chicken)

	delete(chicken, "januari")

	fmt.Println(len(chicken)) // 1
	fmt.Println(chicken)

	// Deteksi keberadaan item denga key tertentu
	chicken := map[string]int{"januari": 50, "februari": 40}
	value, isExist := chicken["mei"] // cek apakah ada key mei dalam map

	if isExist {
		fmt.Println(value)
	} else {
		fmt.Println("item is not exists")
	}

	// Kombinasi slice dan map
	chickens := []map[string]int{
		map[string]string{"name": "chicken blue", "gender": "male"},
		map[string]string{"name": "chicken red", "gender": "male"},
		map[string]string{"name": "chicken yellow", "gender": "female"},
	}
	for _, chicken := range chickens {
		fmt.Println(chicken["gender"], chicken["name"])
	}

	// alternatif penulisan
	chickens := []map[string]string{
		{"name": "chicken blue", "gender": "male"},
		{"name": "chicken red", "gender": "male"},
		{"name": "chicken yellow", "gender": "female"},
	}
	// dalam []map[string]string tiap elemen bisa memiliki key yang berbeda beda
	data := []map[string]string{
		{"name": "chicken blue", "gender": "male", "color": "brown"},
		{"address": "mangga street", "id": "k001"},
		{"community": "chicken lovers"},
	}
}
