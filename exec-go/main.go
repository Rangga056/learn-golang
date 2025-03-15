package main

import (
	"fmt"
	"os/exec"
	"runtime"
)

func main() {
	// Exec digunakan untuk eksekusi perintah command melalui kode program.
	// command yang eksekusi adalah semua command yang bisa dieksekusi di command line sesuai sistem operasinya

	// Penggunaan exec.Command() untuk menjalankan command yang dituliskan pada argument dan .Output() untuk mendapatkan outputnya
	// NOTE: Nilai dari Output() berbentuk []byte pastikan untuk cast ke string terlebih dahulu untuk dapat membaca isi output
	output1, _ := exec.Command("ls").Output()
	fmt.Printf("-> ls\n%s\n", string(output1))

	output2, _ := exec.Command("pwd").Output()
	fmt.Printf("-> pwd\n%s\n", string(output2))

	output3, _ := exec.Command("git", "config", "user.name").Output()
	fmt.Printf("-> git config user.name\n%s\n", string(output3))

	// Ada kalanya saat execute command terdapat kendala seperti command not found yang diakibatkan perbadaan command di tiap OS
	// untuk mengatasi itu tambahkan bash -c untuk sistem operasi Linux, Mac OS, Unix dan cmd /C untuk OS Windows
	if runtime.GOOS == "windows" { // runtime.GOOS untuk mendapatkan informasi sistem operasi yang digunakan
		output4, err := exec.Command("cmd", "/C", "git config user.name").Output()
	} else {
		output4, err := exec.Command("bash", "-c", "git config user.name").Output()
	}
}
