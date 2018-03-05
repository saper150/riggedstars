package user

import (
	"encoding/json"
	"net/http"
	"riggedstars/app/db"
	"riggedstars/app/deck"
	"riggedstars/app/models"

	"github.com/gorilla/mux"
)

type createUserForm struct {
	Name string
}

func createUser(w http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	var userForm createUserForm
	err := decoder.Decode(&userForm)
	if err != nil {
		panic(err)
	}
	db := db.Db()
	user := models.User{Name: userForm.Name}
	db.Create(&user)
	js, _ := json.Marshal(user)
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func getUsers(w http.ResponseWriter, req *http.Request) {

	db := db.Db()

	var users []models.User
	db.Find(&users)

	w.Header().Set("Content-Type", "application/json")
	js, _ := json.Marshal(users)
	w.Write(js)
}

func deckk(w http.ResponseWriter, req *http.Request) {

	js, _ := json.Marshal(deck.ShufeledDeck())
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func RegisterRoutes(router *mux.Router) {
	router.HandleFunc("", createUser).Methods("POST")
	router.HandleFunc("", getUsers).Methods("GET")
}
