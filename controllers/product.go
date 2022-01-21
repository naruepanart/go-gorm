package controllers

import (
	"fmt"

	"github.com/naruepanart/gorm/database"
	"github.com/naruepanart/gorm/model"
	"gorm.io/gorm/clause"
)


func CreateProduct(title string, userid int) {
	product := model.Product{}
	product.Title = title
	product.UserId = userid

	tx := database.ConDB.Create(&product)
	if tx.Error != nil {
		fmt.Println(tx.Error)
	}
}

func GetAllProduct() {
	product := []model.Product{}

	tx := database.ConDB.Preload(clause.Associations).Find(&product)
	if tx.Error != nil {
		fmt.Println(tx.Error)
	}

	for _, t := range product {
		fmt.Printf("%v|%v|%v\n", t.ID, t.UserId, t.User)
	}

}

func GetProduct(id int) {
	product := model.Product{}

	tx := database.ConDB.First(&product, id)
	if tx.Error != nil {
		fmt.Println(tx.Error)
	}

	fmt.Println(product)
}

func GetProductByTitle(title string) {
	product := model.Product{}
	tx := database.ConDB.Where("title = ?", &title).First(&product)
	if tx.Error != nil {
		fmt.Println(tx.Error)
	}

	fmt.Println(product)

}

func UpdateProduct(id int, title string) {
	product := model.Product{}
	product.ID = uint(id)
	product.Title = title

	database.ConDB.Model(&product).Updates(&product)

	fmt.Println(product)
}

func DeleteProduct(id int) {
	product := model.Product{}
	product.ID = uint(id)

	database.ConDB.Delete(&product, id)

	fmt.Println(product)
}

func DeleteProductByTitle(title string) {
	product := model.Product{}
	product.Title = title

	database.ConDB.Where("title = ?", title).Delete(&product)
	fmt.Println(product)
}
