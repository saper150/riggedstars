package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

type User struct {
	gorm.Model
	Name     string
	Password string
}

const (
	host     = "postgres"
	port     = 5432
	user     = "postgres"
	password = "pass"
	dbname   = "postgres"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {

	db := db()
	db.Create(&User{Name: "super name", Password: "pass"})

	var users []User
	db.Find(&users)

	str := "<html><ul>"
	for _, user := range users {
		str += "<li>" + user.Name + " " + user.Password + "</li>"
	}
	str += "</ul></html>"

	fmt.Fprintf(w, str)
}

func main() {
	db := db()

	db.AutoMigrate(&User{})

	r := mux.NewRouter()
	r.HandleFunc("/", homeHandler)
	http.ListenAndServe(":3001", r)
}

func db() *gorm.DB {
	connectionString := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := gorm.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}
	return db
}
