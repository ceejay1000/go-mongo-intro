package controllers

import (
	"fmt"
	"net/http"

	"encoding/json"

	models "github.com/ceejay1000/go_mongo_intro/models"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type UserController struct {
	session *mgo.Session
}

func NewUserController(session *mgo.Session) *UserController {
	return &UserController{session: session}
}

func (uc *UserController) GetUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	id := p.ByName("id")

	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(http.StatusNotFound)
	}

	oid := bson.ObjectIdHex(id)

	u := models.User{}

	if err := uc.session.DB("mongo-golang").C("users").FindId(oid).One(&u); err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	uj, err := json.Marshal(u)

	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s\n", uj)
}

func (uc *UserController) CreateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	u := models.User{}

	err := json.NewDecoder(r.Body).Decode(&u)

	u.Id = bson.NewObjectId()

	if err != nil {
		fmt.Println("Unable to parse json data")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err = uc.session.DB("mongo-golang").C("users").Insert(u); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	uj, err := json.Marshal(u)

	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "%s\n", uj)
}

func (uc *UserController) DeleteUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	id := p.ByName("id")

	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	oid := bson.ObjectIdHex(id)

	if err := uc.session.DB("mongo-golang").C("users").Remove(oid); err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Deleted user ", oid, "\n")

}
