package routes

import (
	"github.com/danieljvx/talent-pitch-api/controllers"
	"github.com/gofiber/fiber/v2"
)

func UserRoute(router fiber.Router) {
	base := router.Group("/user")
	base.Get("/:id", controllers.GetUserController)
	base.Post("/", controllers.SetCreateUserController)
	base.Put("/:id", controllers.SetUpdateUserController)
	bases := router.Group("/users")
	bases.Get("/", controllers.GetUsersController)
}
