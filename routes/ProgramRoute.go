package routes

import (
	"github.com/danieljvx/talent-pitch-api/controllers"
	"github.com/gofiber/fiber/v2"
)

func ProgramRoute(router fiber.Router) {
	base := router.Group("/program")
	base.Get("/:id", controllers.GetProgramController)
	base.Post("/", controllers.SetCreateProgramController)
	base.Put("/:id", controllers.SetUpdateProgramController)
	base.Post("/:id/participant", controllers.SetCreateProgramParticipantController)
	base.Put("/:id/participant", controllers.SetUpdateProgramParticipantController)
	bases := router.Group("/programs")
	bases.Get("/", controllers.GetProgramsController)
}
