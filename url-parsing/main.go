package main

import (
	"fmt"
	"net/url"
)

func main() {
	// Data string berisi informasi URL bisa dikonversi ke tipe data url.URL . Dengan menggunakan tipe data url.URL akan ada banyak informasi yang bisa kita
	// dapatkan dengan mudah, diantaranya seperti jenis protokol yang digunakan, path yang diakses, query, dll.

	urlString := "http://kalipare.com:80/hello?name=muhammad rangga&age=21"
	// Fungsi url.Parse() digunakan untuk parsing string ke bentuk URL. Mengembalikan 2 data, variable objek bertipe url.URL dan error
	// lewat variable tersebut pengaksesan informasi url akan menjadi lebih mudah seperti nama host, protocol dan path
	// selain itu, query yang ada pada url akan otomatis di parsing juga, menjadi bentuk map[string]string, dengan key nama elemen query dan value berisi value query
	u, e := url.Parse(urlString)
	if e != nil {
		fmt.Println(e.Error())
		return
	}

	fmt.Printf("url: %s\n", urlString)

	fmt.Printf("protocol: %s\n", u.Scheme) // http
	fmt.Printf("host: %s\n", u.Host)       // kalipare:80
	fmt.Printf("path: %s\n", u.Path)       // /hello

	name := u.Query()["name"][0]
	age := u.Query()["age"][0]

	fmt.Printf("name: %s, age: %s\n", name, age)
}
