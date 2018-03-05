package main

import (
	"net/http"
	"riggedstars/app/game"
	"riggedstars/app/user"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	user.RegisterRoutes(r.PathPrefix("/user").Subrouter())
	game.RegisterRoutes(r.PathPrefix("/game").Subrouter())
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./app/static"))))
	http.ListenAndServe(":3001", r)
}
