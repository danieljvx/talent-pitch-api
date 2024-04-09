package routes

import (
	"github.com/danieljvx/talent-pitch-api/controllers"
	"github.com/gofiber/fiber/v2"
)

func BaseRoute(router fiber.Router) {
	base := router.Group("/")
	base.Get("/", controllers.BaseController)
}
