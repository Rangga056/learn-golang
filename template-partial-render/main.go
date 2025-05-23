package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type M map[string]interface{}

// Tipe M merupakan alias dari map[string]interface{}, disiapkan untuk mempersingkat penulisan tipe map tersebut
func main() {
	http.HandleFunc("/index", func(w http.ResponseWriter, r *http.Request) {
		data := M{"name": "Batman"}
		tmpl := template.Must(template.ParseFiles(
			"views/index.html",
			"views/_header.html",
			"views/_message.html",
		))

		err := tmpl.ExecuteTemplate(w, "index", data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		data := M{"name": "Batman"}
		tmpl := template.Must(template.ParseFiles(
			"views/about.html",
			"views/_header.html",
			"views/_message.html",
		))
		err := tmpl.ExecuteTemplate(w, "about", data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	fmt.Println("server started at localhost:9000")
	http.ListenAndServe(":9000", nil)
}
