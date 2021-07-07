package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", foo)
	http.HandleFunc("/dog/", dog)
	http.HandleFunc("/dog.jpeg", dogImage)
	log.Fatalln(http.ListenAndServe(":8080", nil))
}

func foo(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "foo just ran")
}

func dog(res http.ResponseWriter, req *http.Request) {
	tpl, err := template.ParseFiles("dog.gohtml")
	if err != nil {
		log.Fatalln(err)
	}

	tpl.ExecuteTemplate(res, "dog.gohtml", "kaalu")
}

func dogImage(res http.ResponseWriter, req *http.Request) {
	http.ServeFile(res, req, "dog.jpeg")
}
