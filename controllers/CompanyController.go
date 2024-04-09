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

type ResponseCompanies struct {
	StatusCode int                                 `json:"statusCode"`
	Data       services.PaginationCompaniesService `json:"data"`
	Message    string                              `json:"message"`
}
type ResponseCompany struct {
	StatusCode int                  `json:"statusCode"`
	Data       *models.CompanyModel `json:"data"`
	Message    string               `json:"message"`
}

type RequestCreateCompany struct {
	Title     string `json:"title" xml:"title" form:"title"`
	ImagePath string `json:"image_path" xml:"image_path" form:"image_path"`
	Location  string `json:"location" xml:"location" form:"location"`
	Industry  string `json:"industry" xml:"industry" form:"industry"`
	UserID    int    `json:"user_id" xml:"user_id" form:"user_id"`
}

func GetCompanyController(c *fiber.Ctx) error {
	companyIdParam := c.Params("id")
	fmt.Printf("companyIdParam: %s\n", companyIdParam)
	if len(companyIdParam) > 0 {
		companyId, err := strconv.Atoi(companyIdParam)
		fmt.Printf("companyId: %v\n", companyId)
		if err == nil {
			company := services.GetCompanyService(companyId)
			if company != nil {
				return c.Status(http.StatusOK).JSON(responses.Response{
					Status:  http.StatusOK,
					Message: "Company found",
					Data:    company,
				})
			} else {
				return c.Status(http.StatusNotFound).JSON(responses.Response{
					Status:  http.StatusNotFound,
					Message: "Company not found",
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
		Message: "Company not found",
		Data:    nil,
	})
}

func SetCreateCompanyController(c *fiber.Ctx) error {
	params := new(RequestCreateCompany)
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
	fmt.Printf("params.ImagePath: %s\n", params.ImagePath)
	if len(params.ImagePath) <= 0 {
		return c.Status(http.StatusBadRequest).JSON(responses.Response{
			Status:  http.StatusBadRequest,
			Message: "Error [difficulty] field is required",
			Data:    nil,
		})
	}
	fmt.Printf("params.Location: %s\n", params.Location)
	if len(params.Location) <= 0 {
		return c.Status(http.StatusBadRequest).JSON(responses.Response{
			Status:  http.StatusBadRequest,
			Message: "Error [location] field is required",
			Data:    nil,
		})
	}
	fmt.Printf("params.Industry: %s\n", params.Industry)
	if len(params.Industry) <= 0 {
		return c.Status(http.StatusBadRequest).JSON(responses.Response{
			Status:  http.StatusBadRequest,
			Message: "Error [industry] field is required",
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
	company := services.PostCompanyService(params.Title, params.ImagePath, params.Location, params.Industry, uint(params.UserID))
	if company != nil {
		return c.Status(http.StatusCreated).JSON(responses.Response{
			Status:  http.StatusCreated,
			Message: "company created",
			Data:    company,
		})
	}
	return c.Status(http.StatusBadRequest).JSON(responses.Response{
		Status:  http.StatusBadRequest,
		Message: "company not created",
		Data:    nil,
	})
}

func SetUpdateCompanyController(c *fiber.Ctx) error {
	companyIdParam := c.Params("id")
	fmt.Printf("companyIdParam: %s\n", companyIdParam)
	if len(companyIdParam) > 0 {
		companyId, _ := strconv.Atoi(companyIdParam)
		params := new(RequestCreateCompany)
		err := c.BodyParser(params)
		if err != nil {
			return c.Status(http.StatusBadRequest).JSON(responses.Response{
				Status:  http.StatusBadRequest,
				Message: err.Error(),
				Data:    nil,
			})
		}
		fmt.Printf("params.Title: %s\n", params.Title)
		fmt.Printf("params.ImagePath: %s\n", params.ImagePath)
		fmt.Printf("params.Location: %s\n", params.Location)
		fmt.Printf("params.Industry: %s\n", params.Industry)
		fmt.Printf("params.UserID: %s\n", params.UserID)
		company := services.PutCompanyService(companyId, params.Title, params.ImagePath, params.Location, params.Industry, uint(params.UserID))
		if company != nil {
			return c.Status(http.StatusOK).JSON(responses.Response{
				Status:  http.StatusOK,
				Message: "company updated",
				Data:    company,
			})
		}
		return c.Status(http.StatusBadRequest).JSON(responses.Response{
			Status:  http.StatusBadRequest,
			Message: "company not updated",
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

func GetCompaniesController(c *fiber.Ctx) error {
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
	companys := services.GetCompaniesService(page, perPage, dateStart, dateEnd)
	fmt.Printf("companys: %v\n", companys)
	if companys == nil {
		return c.Status(http.StatusOK).JSON(ResponseCompanies{
			StatusCode: http.StatusOK,
			Message:    "Companies not found",
			Data: services.PaginationCompaniesService{
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
		Message: "Companies found",
		Data:    companys,
	})
}
