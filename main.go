package main

import (
	"fmt"

	"github.com/naruepanart/gorm/database"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name string
}

func main() {
	database.ConnectDatabaseSQL()
	database.ConDB.AutoMigrate(&Product{})

	CreateProduct("a2")
	//GetAllProduct()
	//GetProduct(1)
	//GetProductByName("a2")
	//UpdateProduct(2, "n2")
	//DeleteProduct(1)
	//DeleteProductByName("n2")
}

func CreateProduct(name string) {
	product := Product{}
	product.Name = name

	tx := database.ConDB.Create(&product)
	if tx.Error != nil {
		fmt.Println(tx.Error)
	}
}

func GetAllProduct() {
	product := []Product{}

	tx := database.ConDB.Find(&product)
	if tx.Error != nil {
		fmt.Println(tx.Error)
	}

	fmt.Println(product)
}

func GetProduct(id int) {
	product := Product{}

	tx := database.ConDB.First(&product, id)
	if tx.Error != nil {
		fmt.Println(tx.Error)
	}

	fmt.Println(product)
}

func GetProductByName(name string) {
	product := Product{}
	tx := database.ConDB.First(&product, "name = ?", &name)
	if tx.Error != nil {
		fmt.Println(tx.Error)
	}
	fmt.Println(product)
}

func UpdateProduct(id int, name string) {
	product := Product{}
	product.ID = uint(id)
	product.Name = name

	database.ConDB.Model(&product).Updates(&product)

	fmt.Println(product)
}

func DeleteProduct(id int) {
	product := Product{}
	product.ID = uint(id)

	database.ConDB.Delete(&product, id)

	fmt.Println(product)
}

func DeleteProductByName(name string) {
	product := Product{}
	product.Name = name

	database.ConDB.Where("name = ?", name).Delete(&product)
	fmt.Println(product)
}
