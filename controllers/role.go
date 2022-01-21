package controllers

import (
	"strconv"

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

/* func GetRole(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")

	role := model.Role{}
	role.ID = uint(id)

	database.ConDB.Preload(clause.Associations).First(&role)
	return c.JSON(role)
} */

func GetRole(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	role := model.Role{
		Id: uint(id),
	}

	database.ConDB.Preload("Permissions").Find(&role)

	return c.JSON(role)
}

func CreateRole(c *fiber.Ctx) error {
	var roleDto fiber.Map

	if err := c.BodyParser(&roleDto); err != nil {
		return err
	}

	list := roleDto["permissions"].([]interface{})

	permissions := make([]model.Permission, len(list))

	for i, permissionId := range list {
		id, _ := strconv.Atoi(permissionId.(string))

		permissions[i] = model.Permission{
			Id: uint(id),
		}
	}

	role := model.Role{
		Name:        roleDto["name"].(string),
		Permissions: permissions,
	}

	database.ConDB.Create(&role)

	return c.JSON(role)
}

func UpdateRole(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	var roleDto map[string]interface{}

	if err := c.BodyParser(&roleDto); err != nil {
		return err
	}

	list := roleDto["permissions"].([]interface{})

	permissions := make([]model.Permission, len(list))

	for i, permissionId := range list {
		id, _ := permissionId.(float64)

		permissions[i] = model.Permission{
			Id: uint(id),
		}
	}

	var result interface{}

	database.ConDB.Table("role_permissions").Where("role_id", id).Delete(&result)

	role := model.Role{
		Id:          uint(id),
		Name:        roleDto["name"].(string),
		Permissions: permissions,
	}

	database.ConDB.Model(&role).Updates(role)

	return c.JSON(role)
}

func DeleteRole(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")
	role := model.Role{}

	database.ConDB.Delete(&role, id)
	return c.JSON("ok")
}
