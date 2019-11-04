//Package main helps to keep a count of number of users that visit website or that particular endpoints
package main

import (
	"io"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/mywebsite", set)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func set(res http.ResponseWriter, req *http.Request) {
	//read the cookies first if exists
	cookie, err := req.Cookie("xyz")

	//if not create one
	if err == http.ErrNoCookie {
		cookie = &http.Cookie{
			Name:  "xyz",
			Value: "0",
		}
	}

	//Get the value of cookie, convert to int
	counter, _ := strconv.Atoi(cookie.Value)
	//Increment it and set to /update cookie.Value
	counter++
	//But convert to string first
	cookie.Value = strconv.Itoa(counter)

	http.SetCookie(res, cookie)
	io.WriteString(res, "Number of page visites : "+cookie.Value)
}
