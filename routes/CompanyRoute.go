package routes

import (
	"github.com/danieljvx/talent-pitch-api/controllers"
	"github.com/gofiber/fiber/v2"
)

func CompanyRoute(router fiber.Router) {
	base := router.Group("/company")
	base.Get("/:id", controllers.GetCompanyController)
	base.Post("/", controllers.SetCreateCompanyController)
	base.Put("/:id", controllers.SetUpdateCompanyController)
	bases := router.Group("/companies")
	bases.Get("/", controllers.GetCompaniesController)
}
