package db

import (
	"riggedstars/app/models"
)

func ChangeStack(user models.User, ammount int) {
	db := Db()
	var userInDb models.User
	db.Where(&models.User{Name: user.Name}).First(&userInDb)
	if userInDb.ID == 0 {
		return
	}
	db.Model(&userInDb).Update("Stack", userInDb.Stack+ammount)
}
