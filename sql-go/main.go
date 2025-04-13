package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql" // Import driver MySQL, gunakan _ karena hanya untuk inisialisasi (tidak digunakan langsung)
)

// Struktur data student untuk mapping hasil query database
type student struct {
	id    string
	name  string
	age   int
	grade int
}

// Fungsi untuk menghubungkan ke database MySQL
func connect() (*sql.DB, error) {
	// Format DSN: username:password@tcp(host:port)/dbname
	db, err := sql.Open("mysql", "rangga:rangga@tcp(127.0.0.1:3306)/db_belajar_golang")
	if err != nil {
		return nil, err
	}
	return db, nil
}

// Fungsi untuk menampilkan semua siswa berdasarkan usia (menggunakan Query dan looping hasil)
func sqlQuery() {
	db, err := connect()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	age := 27 // usia yang ingin dicari
	rows, err := db.Query("select id, name, grade from tb_student where age = ?", age)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer rows.Close()

	var result []student

	// Loop melalui setiap baris hasil query
	for rows.Next() {
		each := student{}
		// Scan isi row ke dalam struct
		err := rows.Scan(&each.id, &each.name, &each.grade)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		result = append(result, each)
	}

	// Cek apakah terjadi error saat iterasi rows
	if err = rows.Err(); err != nil {
		fmt.Println(err.Error())
		return
	}

	// Cetak nama setiap siswa
	for _, each := range result {
		fmt.Println(each.name)
	}
}

// Fungsi untuk mengambil 1 baris data berdasarkan id (menggunakan QueryRow)
func sqlQueryRow() {
	db, err := connect()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	result := student{}
	id := "E001" // id siswa yang ingin dicari
	// Gunakan QueryRow karena hanya satu hasil
	err = db.QueryRow("select name, grade from tb_student where id = ?", id).Scan(&result.name, &result.grade)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// Tampilkan hasil
	fmt.Printf("Name: %s\ngrade: %d\n", result.name, result.grade)
}

// Fungsi untuk melakukan query berulang dengan prepare statement
func sqlPrepare() {
	db, err := connect()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	// Siapkan statement agar lebih efisien bila digunakan berulang
	stmt, err := db.Prepare("select name, grade from tb_student where id = ?")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// Query 1
	result1 := student{}
	stmt.QueryRow("E001").Scan(&result1.name, &result1.grade)
	fmt.Printf("name: %s\ngrade: %d\n", result1.name, result1.grade)

	// Query 2
	result2 := student{}
	stmt.QueryRow("W001").Scan(&result2.name, &result2.grade)
	fmt.Printf("name: %s\ngrade: %d\n", result2.name, result2.grade)

	// Query 3
	result3 := student{}
	stmt.QueryRow("B001").Scan(&result3.name, &result3.grade)
	fmt.Printf("name: %s\ngrade: %d\n", result3.name, result3.grade)
}

// Fungsi untuk melakukan operasi INSERT, UPDATE, DELETE (menggunakan Exec)
func sqlExec() {
	db, err := connect()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	// INSERT data baru
	_, err = db.Exec("insert into tb_student values (?,?,?,?)", "G001", "Galahud", 29, 2)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Insert Success!")

	// UPDATE usia berdasarkan id
	_, err = db.Exec("update tb_student set age = ? where id = ?", 28, "G001")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Update Success!")

	// DELETE data berdasarkan id
	_, err = db.Exec("delete from tb_student where id = ?", "G001")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Delete Success!")
}

// Fungsi main untuk mengeksekusi semua fungsi
func main() {
	sqlQuery()    // Menampilkan semua siswa dengan usia tertentu
	sqlQueryRow() // Menampilkan satu siswa berdasarkan ID
	sqlPrepare()  // Menggunakan prepare statement untuk beberapa query
	sqlExec()     // Menjalankan insert, update, dan delete
}
