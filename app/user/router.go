package user

import (
	"encoding/json"
	"net/http"
	"riggedstars/app/db"
	"riggedstars/app/deck"
	"riggedstars/app/models"

	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
)

type createUserForm struct {
	Name     string
	Password string
}

func createUser(w http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	var userForm createUserForm
	err := decoder.Decode(&userForm)
	if err != nil {
		http.Error(w, "Error in decoding request body", http.StatusBadRequest)
		return
	}
	db := db.Db()
	bytesHash, err := bcrypt.GenerateFromPassword([]byte(userForm.Password), 12)
	if err != nil {
		http.Error(w, "Error while generating a hash", http.StatusBadRequest)
		return
	}
	user := models.User{Name: userForm.Name, Password: string(bytesHash)}
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

func deleteUser(w http.ResponseWriter, req *http.Request) {
	db := db.Db()

	vars := mux.Vars(req)
	id := vars["id"]
	var user models.User
	db.First(&user, id)
	w.Header().Set("Content-Type", "application/json")

	if user.ID != 0 {
		db.Delete(&user)
		js, _ := json.Marshal(user)
		w.Write(js)
	} else {
		http.Error(w, "No user with id:"+id, http.StatusBadRequest)
	}
}

func updateUser(w http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	var userForm createUserForm
	err := decoder.Decode(&userForm)
	if err != nil {
		http.Error(w, "Error in decoding request body", http.StatusBadRequest)
		return
	}
	db := db.Db()
	vars := mux.Vars(req)
	id := vars["id"]
	var user models.User
	db.First(&user, id)
	if user.ID == 0 {
		http.Error(w, "No user with id:"+id, http.StatusBadRequest)
		return
	}
	db.Model(&user).Update(userForm)
	js, _ := json.Marshal(user)
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
	router.HandleFunc("/id/{id:[0-9]+}", deleteUser).Methods("DELETE")
	router.HandleFunc("/id/{id:[0-9]+}", updateUser).Methods("PUT")
}
