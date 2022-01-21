package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/naruepanart/gorm/database"
	"github.com/naruepanart/gorm/model"
)

func GetAllPermission(c *fiber.Ctx) error {
	permission := []model.Permission{}

	database.ConDB.Find(&permission)
	return c.JSON(permission)
}

func CreatePermission(c *fiber.Ctx) error {
	permission := model.Permission{}

	if err := c.BodyParser(&permission); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	database.ConDB.Create(&permission)
	return c.JSON(permission)
}
