package main

import (
	"io"
	"log"
	"net/http"
	"text/template"
)

func main() {

	http.HandleFunc("/", index)
	http.HandleFunc("/dog/", dog)
	http.Handle("/me/", http.HandlerFunc(me)) // using HandlerFunc type //It takes a Func of signature similar to ServeHTTP
	//Hence, a handler
	http.ListenAndServe(":8080", nil)
}

func index(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "index func called")
}

func dog(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "dog func called")
}

func me(res http.ResponseWriter, req *http.Request) {

	temp, err := template.ParseFiles("index.gohtml")
	if err != nil {
		log.Fatalln("Error in parsing the template", err.Error())
	}

	err = temp.ExecuteTemplate(res, "index.gohtml", "nrj")
	if err != nil {
		log.Fatalln("error executing template", err)
	}
}
