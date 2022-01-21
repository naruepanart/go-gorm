package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/naruepanart/gorm/database"
	"github.com/naruepanart/gorm/model"
	"github.com/naruepanart/gorm/routes"
)

func main() {
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	database.ConnectDatabaseSQL()
	// database.ConDB.Migrator().DropColumn(&User{}, "product_id")
	//database.ConDB.Migrator().DropTable(&model.Role{}, &model.Permission{})

	database.ConDB.AutoMigrate(&model.Role{}, &model.Permission{})

	// Users
	//CreateUsers("A01")

	// Product
	//CreateProduct("P05", 1)
	//GetAllProduct()
	//GetProduct(1)
	//GetProductByTitle("P02")
	//UpdateProduct(2, "n2")
	//DeleteProduct(1)
	//DeleteProductByTitle("n2")

	routes.SetupRoutes(app)

	app.Listen(":3000")

}
