package main

import (
	"math/rand"
	"net/http"
	"riggedstars/app/game"
	"riggedstars/app/user"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	game.RunTests()
	rand.Seed(time.Now().UTC().UnixNano())
	r := mux.NewRouter()
	user.RegisterRoutes(r.PathPrefix("/user").Subrouter())
	game.RegisterRoutes(r.PathPrefix("/game").Subrouter())
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	headersAllowed := handlers.AllowedHeaders([]string{"content-type"})
	originsAllowed := handlers.AllowedOrigins([]string{"*"})
	methodsAllowed := handlers.AllowedMethods([]string{"GET", "DELETE", "POST", "PUT"})
	http.ListenAndServe(":3001", handlers.CORS(headersAllowed, originsAllowed, methodsAllowed)(r))
}
