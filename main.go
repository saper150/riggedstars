package main

import (
	"net/http"
	"riggedstars/app/game"
	"riggedstars/app/user"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	user.RegisterRoutes(r.PathPrefix("/user").Subrouter())
	game.RegisterRoutes(r.PathPrefix("/game").Subrouter())
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./app/static"))))

	headersAllowed := handlers.AllowedHeaders([]string{"content-type"})
	originsAllowed := handlers.AllowedOrigins([]string{"*"})
	methodsAllowed := handlers.AllowedMethods([]string{"GET", "DELETE", "POST", "PUT"})
	http.ListenAndServe(":3001", handlers.CORS(headersAllowed, originsAllowed, methodsAllowed)(r))
}
