package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/", index)
	http.HandleFunc("/set", set)
	http.HandleFunc("/read", read)
	http.HandleFunc("/expire", expire)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":9090", nil)
}

func index(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(res, `<h1><a href="/set">Set a cookie</a></h1>`)
}
func set(res http.ResponseWriter, req *http.Request) {
	http.SetCookie(res, &http.Cookie{
		Name:   "lolcookies",
		Value:  "0",
		MaxAge: 2,
	})
	fmt.Fprintln(res, `<h1><a href="/read">Read cookies</a></h1>`)
}
func read(res http.ResponseWriter, req *http.Request) {

	cookie, err := req.Cookie("lolcookies")
	if err != nil {
		http.Redirect(res, req, "/set", http.StatusSeeOther)
		return
	}
	fmt.Fprintf(res, `<h1>Cookie : <br> %v</h1><h2><a href ="/expire">Delete cookie </a></h2>`, cookie)

}
func expire(res http.ResponseWriter, req *http.Request) {

	cookie, err := req.Cookie("lolcookies")
	if err != nil {
		http.Redirect(res, req, "/set", http.StatusSeeOther)
		return // return is imp bcoz we want to exit this func
	}
	cookie.MaxAge = -1 // delete cookie Expire is deprecated
	http.SetCookie(res, cookie)
	http.Redirect(res, req, "/", http.StatusSeeOther)
}
