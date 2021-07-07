package main

import (
	"io"
	"net/http"
)

func main() {

	http.HandleFunc("/", index)
	http.HandleFunc("/dog/", dog)
	http.HandleFunc("/me/", me)

	http.ListenAndServe(":8080", nil)
}

func index(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "index func called")
}

func dog(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "dog func called")
}

func me(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "hello, its nrj")
}
