package main

import (
	"time"

	"github.com/naruepanart/gorm/database"
	"gorm.io/gorm"
)

type User struct {
	ID        string `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Name      string
}

type Product struct {
	gorm.Model
	Title  string
	UserId string
	User   User `gorm:"foreignKey:UserId"`
}

func main() {
	database.ConnectDatabaseSQL()
	// database.ConDB.Migrator().DropColumn(&User{}, "product_id")
	database.ConDB.Migrator().DropTable(&Product{}, &User{})

	database.ConDB.AutoMigrate(&Product{}, &User{})

	// Users
	CreateUsers("Name-01", "123123")

	// Product
	CreateProduct("Product-01", "123123")
	//GetAllProduct()
	//GetProduct(1)
	//GetProductByTitle("P02")
	//UpdateProduct(2, "n2")
	//DeleteProduct(1)
	//DeleteProductByTitle("n2")
}

func CreateUsers(name string, id string) {
	user := User{}
	user.Name = name
	user.ID = id
	database.ConDB.Create(&user)
}

func CreateProduct(title string, id string) {
	product := Product{}
	product.Title = title
	product.UserId = id

	database.ConDB.Create(&product)
}

/* func GetAllProduct() {
	product := []Product{}

	tx := database.ConDB.Preload(clause.Associations).Find(&product)
	if tx.Error != nil {
		fmt.Println(tx.Error)
	}

	for _, t := range product {
		fmt.Printf("%v|%v|%v\n", t.ID, t.Userkey, t.User)
	}

} */

/*
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
*/
