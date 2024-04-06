package main

import (
	"log"

	"github.com/danieljvx/talent-pitch-api/config"
	"github.com/danieljvx/talent-pitch-api/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	config.ConnectDB()
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*", // Se aceptan desde cualquier origen a efecto de la prueba
		AllowMethods:     "GET, POST, PUT, DELETE",
		AllowCredentials: false,
		AllowHeaders:     "Origin, Content-Type, Accept",
	}))

	router.SetupRoutes(app)

	log.Fatal(app.Listen(":" + config.Config("APP_PORT")))
}
