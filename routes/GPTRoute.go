package routes

import (
	"github.com/danieljvx/talent-pitch-api/controllers"
	"github.com/gofiber/fiber/v2"
)

func GPTRoute(router fiber.Router) {
	base := router.Group("/gpt")
	base.Get("/migration", controllers.GetGPTDataMigrationController)
}
