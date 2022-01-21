package controllers

import (
	"fmt"

	"github.com/naruepanart/gorm/database"
	"github.com/naruepanart/gorm/model"
)

func CreateUsers(name string) {
	user := model.User{}
	user.Name = name

	tx := database.ConDB.Create(&user)
	if tx.Error != nil {
		fmt.Println(tx.Error)
	}
}
