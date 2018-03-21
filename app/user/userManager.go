package user

import (
	"encoding/json"
	"net/http"
	"riggedstars/app/db"
	"riggedstars/app/models"
	"strings"

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
	Data   models.User
	Token  string
	Status int
}

type customClaims struct {
	Name string
	ID   uint
	jwt.StandardClaims
}

var riggedKey = []byte("rigged")

func checkPassword(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
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
	var userInDb models.User
	db.Where(&models.User{Name: userForm.Name}).First(&userInDb)
	if userInDb.ID != 0 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Name already exist"))
		return
	}
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

func login(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	decoder := json.NewDecoder(req.Body)
	var userForm createUserForm
	err := decoder.Decode(&userForm)
	if err != nil {
		http.Error(w, "Error in decoding request body", http.StatusBadRequest)
		return
	}

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
		http.Error(w, "", 401)
		w.WriteHeader(http.StatusUnauthorized)
		w.Write(jsErr)
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &customClaims{
		Name: user.Name,
		ID:   user.ID,
	})
	tokenString, _ := token.SignedString(riggedKey)
	jsAuth, _ := json.Marshal(loginAuthForm{Data: user, Token: "Bearer " + tokenString, Status: http.StatusOK})
	w.Write(jsAuth)

}

func tokenAuthWithClaimsExample(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	authToken := req.Header.Get("Authorization")

	if len(authToken) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Authorization field not found"))
		return
	}

	claims, err := getBearerTokenClaims(authToken)
	if err == nil {
		w.Header().Set("Content-Type", "application/json")
		js, _ := json.Marshal(customClaims{Name: claims.Name, ID: claims.ID})
		w.Write(js)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("error"))
	}
}

func getBearerTokenClaims(bearerTokenString string) (*customClaims, error) {
	tokenSplit := strings.Split(bearerTokenString, " ")
	var tokenString string
	if len(tokenSplit) == 2 {
		tokenString = tokenSplit[1]
	}
	token, err := jwt.ParseWithClaims(tokenString, &customClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(riggedKey), nil
	})
	if err == nil {
		claims := token.Claims.(*customClaims)
		return claims, nil
	} else {
		return nil, err
	}

}
