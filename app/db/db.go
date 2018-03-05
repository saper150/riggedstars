package db

import (
	"fmt"
	"log"
	"riggedstars/app/models"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

const (
	host     = "postgres"
	port     = 5432
	user     = "postgres"
	password = "pass"
	dbname   = "postgres"
)

var connection *gorm.DB

func Db() *gorm.DB {

	if connection != nil {
		return connection
	}

	connectionString := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	connection, err := gorm.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}
	migrate(connection)
	return connection
}

func migrate(db *gorm.DB) {

	db.AutoMigrate(&models.User{})

}
