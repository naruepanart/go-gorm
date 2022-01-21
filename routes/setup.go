package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/naruepanart/gorm/controllers"
)

func SetupRoutes(router fiber.Router) {

	api := router.Group("api")
	api.Get("roles", controllers.GetAllRole)
	api.Get("role/:id", controllers.GetRole)
	api.Post("role", controllers.CreateRole)
	api.Patch("role/:id", controllers.UpdateRole)
	api.Delete("role/:id", controllers.DeleteRole)

	api.Get("permissions", controllers.GetAllPermission)
	api.Post("permission", controllers.CreatePermission)
}
