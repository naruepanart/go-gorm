package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/naruepanart/gorm/database"
	"github.com/naruepanart/gorm/model"
	"gorm.io/gorm/clause"
)

func GetAllRole(c *fiber.Ctx) error {
	role := []model.Role{}

	database.ConDB.Preload(clause.Associations).Find(&role)
	return c.JSON(role)
}

func GetRole(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")

	role := model.Role{}
	role.ID = uint(id)

	database.ConDB.First(&role)
	return c.JSON(role)
}

func CreateRole(c *fiber.Ctx) error {
	role := model.Role{}

	if err := c.BodyParser(&role); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	database.ConDB.Create(&role)
	return c.JSON(role)
}

func UpdateRole(c *fiber.Ctx) error { 
	id, _ := c.ParamsInt("id")

	role := model.Role{}
	role.ID = uint(id)
	if err := c.BodyParser(&role); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	database.ConDB.Model(&role).Updates(&role)
	return c.JSON(role)
}

func DeleteRole(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")
	role := model.Role{}

	database.ConDB.Delete(&role, id)
	return c.JSON("ok")
}
