package user

import (
	"encoding/json"
	"net/http"
	"riggedstars/app/db"
	"riggedstars/app/deck"
	"riggedstars/app/models"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
)

type createUserForm struct {
	Name     string
	Password string
}

type responseForm struct {
	Status int
	Data   string
}

type loginAuthForm struct {
	Data  models.User
	Token string
}

type authForm struct {
	Name string
	jwt.StandardClaims
}

var riggedKey = []byte("rigged")

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

func checkPassword(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}

func login(w http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	var userForm createUserForm
	err := decoder.Decode(&userForm)
	if err != nil {
		http.Error(w, "Error in decoding request body", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")

	jsErr, _ := json.Marshal(responseForm{Data: "Wrong login or password", Status: http.StatusUnauthorized})
	db := db.Db()
	var user models.User
	db.Where(&models.User{Name: userForm.Name}).First(&user)
	if user.ID == 0 {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write(jsErr)
		return
	}

	match := checkPassword(user.Password, userForm.Password)

	if !match {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write(jsErr)
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &authForm{
		Name: user.Name,
	})

	tokenString, _ := token.SignedString(riggedKey)
	jsAuth, _ := json.Marshal(loginAuthForm{Data: user, Token: tokenString})
	w.Write(jsAuth)
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
	router.HandleFunc("/login/authTest", authTest).Methods("GET")
}

func authTest(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	authToken := req.Header.Get("Authorization")

	if len(authToken) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Authorization field not found"))
	}

	_, authErr := jwt.Parse(authToken, func(token *jwt.Token) (interface{}, error) {
		return []byte(riggedKey), nil
	})
	if authErr == nil {
		w.Write([]byte("Authorization token works"))
	} else {
		w.Write([]byte("Bad authorization token"))
	}

}
