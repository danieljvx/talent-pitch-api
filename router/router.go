package router

import (
	"github.com/danieljvx/talent-pitch-api/controllers"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/", controllers.Home)
}
