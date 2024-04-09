package routes

import (
	"github.com/danieljvx/talent-pitch-api/controllers"
	"github.com/gofiber/fiber/v2"
)

func ChallengeRoute(router fiber.Router) {
	base := router.Group("/challenge")
	base.Get("/:id", controllers.GetChallengeController)
	base.Post("/", controllers.SetCreateChallengeController)
	base.Put("/:id", controllers.SetUpdateChallengeController)
	bases := router.Group("/challenges")
	bases.Get("/", controllers.GetChallengesController)
}
