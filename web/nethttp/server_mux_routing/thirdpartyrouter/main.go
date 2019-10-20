package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {

	router := httprouter.New()
	router.GET("/", index)
	router.GET("/about", about)
	router.GET("/contact", contact)
	router.GET("/like", like)
	router.POST("/like", liked)
	router.GET("/user/:name", user)
	router.GET("/blog/:category/:article", blogRead)
	router.POST("/blog/:category/:article", blogWrite)

	http.ListenAndServe(":8080", router)

}

//IN schdmit router we need a trid parameter "httprouter.Params"
func index(res http.ResponseWriter, _ *http.Request, _ httprouter.Params) {

	err := tpl.ExecuteTemplate(res, "index.gohtml", nil)
	handleError(res, err)
}

func about(res http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	err := tpl.ExecuteTemplate(res, "about.gohtml", nil)
	handleError(res, err)
}

func contact(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	err := tpl.ExecuteTemplate(res, "contact.gohtml", nil)
	handleError(res, err)
}

func like(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	err := tpl.ExecuteTemplate(res, "like.gohtml", nil)
	handleError(res, err)
}

func liked(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	err := tpl.ExecuteTemplate(res, "liked.gohtml", nil)
	handleError(res, err)
}

func user(res http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	fmt.Fprintf(res, "USER, %s!\n", ps.ByName("name"))
}

func blogRead(res http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	fmt.Fprintf(res, "READ CATEGORY, %s!\n", ps.ByName("category"))
	fmt.Fprintf(res, "READ ARTICLE, %s!\n", ps.ByName("article"))
}

func blogWrite(res http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	fmt.Fprintf(res, "WRITE CATEGORY, %s!\n", ps.ByName("category"))
	fmt.Fprintf(res, "WRITE ARTICLE, %s!\n", ps.ByName("article"))
}

func handleError(res http.ResponseWriter, err error) {

	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		log.Fatalln(err)
	}
}
