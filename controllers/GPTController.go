package controllers

import (
	"fmt"
	"github.com/danieljvx/talent-pitch-api/config"
	"github.com/danieljvx/talent-pitch-api/models"
	"github.com/danieljvx/talent-pitch-api/responses"
	"github.com/danieljvx/talent-pitch-api/services"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"strings"
	"time"
)

// GetGPTDataMigrationController @title API
// @version 1.0
// @description This is an auto-generated API Docs.
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email villanueva.danielx@gmail.com
// @GetGPTDataMigrationControllerPath /
func GetGPTDataMigrationController(c *fiber.Ctx) error {
	GptUserData := config.Config("GPT_USER_DATA")
	GptProgramData := config.Config("GPT_PROGRAM_DATA")
	GptCompanyData := config.Config("GPT_COMPANY_DATA")
	respUsers := services.GetDataGPTService(GptUserData)

	respUsersSplit := strings.Split(respUsers, "\n")
	fmt.Printf("respUsersSplit %s\n", respUsersSplit)
	var user models.UserModel
	for ui, us := range respUsersSplit {
		fmt.Printf("%s\n", us)
		if ui%2 == 0 {
			user.Name = us
		} else {
			user.Email = us
			config.DB.Create(&user)
			user = models.UserModel{}
		}
	}

	respPrograms := services.GetDataGPTService(GptProgramData)
	fmt.Printf("respPrograms %s\n", respPrograms)
	respProgramsSplit := strings.Split(respPrograms, "\n")
	var program models.ProgramModel
	for pi, ps := range respProgramsSplit {
		fmt.Printf("%s\n", ps)
		if pi%2 == 0 {
			program.Title = ps
		} else {
			t := time.Now()
			ts := t.Format("2006-01-02")
			program.StartDate = ts
			program.EndDate = ts
			program.Description = ps
			config.DB.Create(&program)
			program = models.ProgramModel{}
		}
	}

	respCompanies := services.GetDataGPTService(GptCompanyData)
	fmt.Printf("respCompanies %s\n", respCompanies)
	respCompaniesSplit := strings.Split(respCompanies, "\n")
	var companies models.CompanyModel
	for ci, cs := range respCompaniesSplit {
		fmt.Printf("%s\n", cs)
		if ci%2 == 0 {
			companies.Name = cs
		} else {
			companies.Location = cs
			config.DB.Create(&companies)
			companies = models.CompanyModel{}
		}
	}
	/*
	 *	El Formato de respuesta de Chat GPT varia mucho entre respuesta.
	 *  Se deja un mapeo sencillo que puede guardar datos aleatorios desordenados.
	 *  Se puede idear una lógica refinada para aceptar un formato específico.
	 */
	return c.Status(http.StatusOK).JSON(responses.Response{
		Status:  http.StatusOK,
		Message: "GPT API services",
		Data:    nil,
	})
}
