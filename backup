package main

import (
	"fmt"

	"github.com/naruepanart/gorm/database"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type User struct {
	gorm.Model
	Name    string
	Product []Product
}

type Product struct {
	gorm.Model
	Title  string
	UserId int
	User   User `gorm:"foreignKey:UserId"`
}

func main() {
	database.ConnectDatabaseSQL()
	// database.ConDB.Migrator().DropColumn(&User{}, "product_id")
	// database.ConDB.Migrator().DropTable(&Product{}, &User{})

	database.ConDB.AutoMigrate(&Product{}, &User{})

	// Users
	//CreateUsers("A01")

	// Product
	//CreateProduct("P03", 1)
	//GetAllProduct()
	//GetProduct(1)
	GetProductByTitle("P02")
	//UpdateProduct(2, "n2")
	//DeleteProduct(1)
	//DeleteProductByTitle("n2")
}

func CreateUsers(name string) {
	user := User{}
	user.Name = name

	tx := database.ConDB.Create(&user)
	if tx.Error != nil {
		fmt.Println(tx.Error)
	}
}

func CreateProduct(title string, userid int) {
	product := Product{}
	product.Title = title
	product.UserId = userid

	tx := database.ConDB.Create(&product)
	if tx.Error != nil {
		fmt.Println(tx.Error)
	}
}

func GetAllProduct() {
	product := []Product{}

	tx := database.ConDB.Preload(clause.Associations).Find(&product)
	if tx.Error != nil {
		fmt.Println(tx.Error)
	}

	for _, t := range product {
		fmt.Printf("%v|%v|%v\n", t.ID, t.UserId, t.User)
	}

	/* fmt.Println(product) */
}

func GetProduct(id int) {
	product := Product{}

	tx := database.ConDB.First(&product, id)
	if tx.Error != nil {
		fmt.Println(tx.Error)
	}

	fmt.Println(product)
}

func GetProductByTitle(title string) {
	product := Product{}
	tx := database.ConDB.Where("title = ?", &title).First(&product)
	if tx.Error != nil {
		fmt.Println(tx.Error)
	}

	fmt.Println(product)

}

func UpdateProduct(id int, title string) {
	product := Product{}
	product.ID = uint(id)
	product.Title = title

	database.ConDB.Model(&product).Updates(&product)

	fmt.Println(product)
}

func DeleteProduct(id int) {
	product := Product{}
	product.ID = uint(id)

	database.ConDB.Delete(&product, id)

	fmt.Println(product)
}

func DeleteProductByTitle(title string) {
	product := Product{}
	product.Title = title

	database.ConDB.Where("title = ?", title).Delete(&product)
	fmt.Println(product)
}
