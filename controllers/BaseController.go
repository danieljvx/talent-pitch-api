package controllers

import (
	"github.com/danieljvx/talent-pitch-api/responses"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

// BaseController @title API
// @version 1.0
// @description This is an auto-generated API Docs.
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email villanueva.danielx@gmail.com
// @BasePath /
func BaseController(c *fiber.Ctx) error {
	return c.Status(http.StatusOK).JSON(responses.Response{
		Status:  http.StatusOK,
		Message: "TalentPitch Api Test by Daniel Villanueva - @danieljvx",
		Data:    nil,
	})
}
