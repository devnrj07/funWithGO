package main

import(
	"net/http"
	"github.com/julienschmidt/httprouter"
)

fun main() {
	router := httprouter.New()
	//Get a UserController instance
	uc := controller.NewUserController(getSession())
	router.GET("/user/:id", uc.GetUser)
	router.POST("/user", uc.CreateUser)
	router.DELETE("/user/:id", uc.DeleteUser)
	log.Fatalln(http.ListenAndServe(":8080",router))
}

func getSession() *mgo.Session {
	//Connect to local mongodb server
	s, err := mgo.Dial("mongodb://localhost")

	//Check if connection error, is mongo running ?
	if err != nil {
		panic(err)
	}	
		return s
}