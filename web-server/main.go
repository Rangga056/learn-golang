package main

import (
	"fmt"
	"net/http"
	"text/template"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "apa Kabar!")
}

func main() {
	// Go menyediakan package net/http, yang berisi berbagai macam fitur untuk keperluan pembuatan aplikasi berbasis web, termasuk di dalamnya web server, routing, templating, dll

	// http.HandleFunc digunakan untuk routing, memiliki 2 parameter yang harus diisi yaitu route dan callback yang dijalankan ketika route tersebut diakses
	// callback tersebut bertipe func(w http.ResponseWriter, r *http.Request)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := map[string]string{
			"Name":    "Muhammad Rangga",
			"Message": "Have a Good Day",
		}

		// template.ParseFiles() digunakan untuk parsing template, return 2 data yaitu instance templatenya dan error
		// dan method Execute() akan membuat hasil parsing template ditampilkan
		t, err := template.ParseFiles("template.html")
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		t.Execute(w, data)
	})

	http.HandleFunc("/index", index)

	fmt.Println("Starting web server at http://localhost:8000/")
	// http.ListenAndServe() digunakan untuk menghidupkan dan menjalankan server serta aplikasi,
	// NOTE: di GO server 1 web aplikasi adalah 1 buah server yang berbeda
	http.ListenAndServe(":8000", nil)

	// Penggunaan Template Web
	// Template engine memberikan kemudahan dalam mendasain tampilan view aplikasi website dan GO sendiri memiliki template engine sendiri
}
