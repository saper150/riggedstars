package user

import (
	"encoding/json"
	"net/http"
	"riggedstars/app/deck"

	"github.com/gorilla/mux"
)

func loginOptions(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Access-Control-Allow-Origin, content-type")
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
	router.HandleFunc("/login", login).Methods("POST")
	router.HandleFunc("/login", loginOptions).Methods("OPTIONS")
	router.HandleFunc("/login/authTest", tokenAuthWithClaimsExample).Methods("GET")
}
