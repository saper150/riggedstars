package db

import (
	"riggedstars/app/models"

	"github.com/jinzhu/gorm"
)

func ChangeStack(user models.User, ammount int) {
	db := Db()
	var userInDb models.User
	db.Where(&models.User{Name: user.Name}).First(&userInDb)
	if userInDb.ID == 0 {
		return
	}
	db.Model(&userInDb).Update("Stack", gorm.Expr("Stack + ?", ammount))
}
