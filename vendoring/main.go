package main

import (
	"fmt"

	gubrak "github.com/novalagung/gubrak/v2"
)

// Vendoring di go memberikan kita kapabilitas untuk mengunduh semua dependency atau 3rd party, untuk disimpan di lokal dalam folder project bernama subfolder /vendor
// dengan adanya folder tersebut go tidak perlu look up 3rd party ke cache folder atau GOPATH karena sudah ada di dalam vendor sehingga tidak perlu menggunakan
// go mod download atau go mod tidy
func main() {
	// NOTE: agar go lookup ke folder vendor saat build adalah dengan menambahkan flag -mod=vendor  sewaktu build atau run project
	// go run -mod=vendor main.go
	// go build -mod=vendor -o executable
	// vendoring bermanfaat pada sisi kompatibilitas dan kestabilan 3rd party dan tidak perlu repot mendownload dependency karena semua sudah ada di lokal
	// konsekuensinya adalah size project menjadi cukup besar jadi digunakan sesuai dengan kebutuhan saja
	fmt.Println(gubrak.RandomInt(10, 20))
}
