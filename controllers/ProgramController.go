package controllers

import (
	"fmt"
	"github.com/danieljvx/talent-pitch-api/models"
	"github.com/danieljvx/talent-pitch-api/responses"
	"github.com/danieljvx/talent-pitch-api/services"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"strconv"
	"time"
)

type ResponsePrograms struct {
	StatusCode int                                `json:"statusCode"`
	Data       services.PaginationProgramsService `json:"data"`
	Message    string                             `json:"message"`
}
type ResponseProgram struct {
	StatusCode int                  `json:"statusCode"`
	Data       *models.ProgramModel `json:"data"`
	Message    string               `json:"message"`
}

type RequestCreateProgram struct {
	Title       string `json:"title" xml:"title" form:"title"`
	Description string `json:"description" xml:"description" form:"description"`
	StartDate   string `json:"start_date" xml:"start_date" form:"start_date"`
	EndDate     string `json:"end_date" xml:"end_date" form:"end_date"`
	UserID      int    `json:"user_id" xml:"user_id" form:"user_id"`
}

func GetProgramController(c *fiber.Ctx) error {
	programIdParam := c.Params("id")
	fmt.Printf("programIdParam: %s\n", programIdParam)
	if len(programIdParam) > 0 {
		programId, err := strconv.Atoi(programIdParam)
		fmt.Printf("programId: %v\n", programId)
		if err == nil {
			program := services.GetProgramService(programId)
			if program != nil {
				return c.Status(http.StatusOK).JSON(responses.Response{
					Status:  http.StatusOK,
					Message: "Program found",
					Data:    program,
				})
			} else {
				return c.Status(http.StatusNotFound).JSON(responses.Response{
					Status:  http.StatusNotFound,
					Message: "Program not found",
					Data:    nil,
				})
			}
		} else {
			return c.Status(http.StatusBadRequest).JSON(responses.Response{
				Status:  http.StatusBadRequest,
				Message: "Error id field required",
				Data:    nil,
			})
		}
	}
	return c.Status(http.StatusNotFound).JSON(responses.Response{
		Status:  http.StatusNotFound,
		Message: "Program not found",
		Data:    nil,
	})
}

func SetCreateProgramController(c *fiber.Ctx) error {
	params := new(RequestCreateProgram)
	err := c.BodyParser(params)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.Response{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
	}
	fmt.Printf("params.Title: %s\n", params.Title)
	if len(params.Title) <= 0 {
		return c.Status(http.StatusBadRequest).JSON(responses.Response{
			Status:  http.StatusBadRequest,
			Message: "Error [title] field is required",
			Data:    nil,
		})
	}
	fmt.Printf("params.Description: %s\n", params.Description)
	if len(params.Description) <= 0 {
		return c.Status(http.StatusBadRequest).JSON(responses.Response{
			Status:  http.StatusBadRequest,
			Message: "Error [email] field is required",
			Data:    nil,
		})
	}
	fmt.Printf("params.StartDate: %s\n", params.StartDate)
	if len(params.StartDate) <= 0 {
		return c.Status(http.StatusBadRequest).JSON(responses.Response{
			Status:  http.StatusBadRequest,
			Message: "Error [start_date] field is required",
			Data:    nil,
		})
	}
	fmt.Printf("params.EndDate: %s\n", params.EndDate)
	if len(params.EndDate) <= 0 {
		return c.Status(http.StatusBadRequest).JSON(responses.Response{
			Status:  http.StatusBadRequest,
			Message: "Error [end_date] field is required",
			Data:    nil,
		})
	}
	fmt.Printf("params.UserID: %s\n", params.UserID)
	if params.UserID <= 0 {
		return c.Status(http.StatusBadRequest).JSON(responses.Response{
			Status:  http.StatusBadRequest,
			Message: "Error [user_id] field is required",
			Data:    nil,
		})
	}
	program := services.PostProgramService(params.Title, params.Description, params.StartDate, params.EndDate, uint(params.UserID))
	if program != nil {
		return c.Status(http.StatusCreated).JSON(responses.Response{
			Status:  http.StatusCreated,
			Message: "program created",
			Data:    program,
		})
	}
	return c.Status(http.StatusBadRequest).JSON(responses.Response{
		Status:  http.StatusBadRequest,
		Message: "program not created",
		Data:    nil,
	})
}

func SetUpdateProgramController(c *fiber.Ctx) error {
	programIdParam := c.Params("id")
	fmt.Printf("programIdParam: %s\n", programIdParam)
	if len(programIdParam) > 0 {
		programId, _ := strconv.Atoi(programIdParam)
		params := new(RequestCreateProgram)
		err := c.BodyParser(params)
		if err != nil {
			return c.Status(http.StatusBadRequest).JSON(responses.Response{
				Status:  http.StatusBadRequest,
				Message: err.Error(),
				Data:    nil,
			})
		}
		fmt.Printf("params.Title: %s\n", params.Title)
		fmt.Printf("params.Description: %s\n", params.Description)
		fmt.Printf("params.StartDate: %s\n", params.StartDate)
		fmt.Printf("params.EndDate: %s\n", params.EndDate)
		fmt.Printf("params.UserID: %s\n", params.UserID)
		program := services.PutProgramService(programId, params.Title, params.Description, params.StartDate, params.EndDate, uint(params.UserID))
		if program != nil {
			return c.Status(http.StatusCreated).JSON(responses.Response{
				Status:  http.StatusCreated,
				Message: "program updated",
				Data:    program,
			})
		}
		return c.Status(http.StatusBadRequest).JSON(responses.Response{
			Status:  http.StatusBadRequest,
			Message: "program not updated",
			Data:    nil,
		})
	} else {
		return c.Status(http.StatusBadRequest).JSON(responses.Response{
			Status:  http.StatusBadRequest,
			Message: "Error id field required",
			Data:    nil,
		})
	}
}

func GetProgramsController(c *fiber.Ctx) error {
	page := 0
	pageParam := c.Query("page")
	if len(pageParam) > 0 {
		page, _ = strconv.Atoi(pageParam)
	}
	perPage := 10
	perPageParam := c.Query("perPage")
	if len(perPageParam) > 0 {
		perPage, _ = strconv.Atoi(perPageParam)
	}
	var dateStart time.Time
	dateStartParam := c.Query("dateStart")
	if len(dateStartParam) > 0 {
		dateStartInt, _ := strconv.Atoi(dateStartParam)
		dateStart = time.Unix(int64(dateStartInt)/int64(time.Microsecond), 0)
	}
	var dateEnd time.Time
	dateEndParam := c.Query("dateEnd")
	if len(dateEndParam) > 0 {
		dateEndInt, _ := strconv.Atoi(dateEndParam)
		dateEnd = time.Unix(int64(dateEndInt)/int64(time.Microsecond), 0)
	}
	programs := services.GetProgramsService(page, perPage, dateStart, dateEnd)
	fmt.Printf("programs: %v\n", programs)
	if programs == nil {
		return c.Status(http.StatusOK).JSON(ResponsePrograms{
			StatusCode: http.StatusOK,
			Message:    "Programs not found",
			Data: services.PaginationProgramsService{
				Page:     page,
				PerPage:  perPage,
				NextPage: 0,
				PrevPage: 0,
				LastPage: 0,
				Total:    0,
				Data:     nil,
			},
		})
	}

	return c.Status(http.StatusOK).JSON(responses.Response{
		Status:  http.StatusOK,
		Message: "Programs found",
		Data:    programs,
	})
}
