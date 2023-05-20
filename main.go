package main

import (
	"net/http"

	controllers "github.com/ceejay1000/go_mongo_intro/controllers"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
)

func main() {

	router := httprouter.New()

	uc := controllers.NewUserController(getSession())

	router.GET("/user/:id", uc.GetUser)
	router.POST("/user", uc.CreateUser)
	router.DELETE("/user", uc.DeleteUser)

	http.ListenAndServe(":9000", router)
}

func getSession() *mgo.Session {
	mongoUrl := "mongodb://localhost:27017"
	session, err := mgo.Dial(mongoUrl)

	if err != nil {
		panic(err)
	}
	return session
}
