package app

import (
	"github.com/danieljvx/talent-pitch-api/config"
	"github.com/danieljvx/talent-pitch-api/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func App() *fiber.App {
	config.ConnectDB()
	config.ConnectGPT()
	app := fiber.New()
	app.Use(cors.New())
	routes.SetupRoute(app)
	return app
}
