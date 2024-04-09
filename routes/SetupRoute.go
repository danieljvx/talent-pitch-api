package routes

import "github.com/gofiber/fiber/v2"

func SetupRoute(app *fiber.App) {
	BaseRoute(app)
	UserRoute(app)
	ProgramRoute(app)
	CompanyRoute(app)
	ChallengeRoute(app)
	SwaggerRoute(app)
}
