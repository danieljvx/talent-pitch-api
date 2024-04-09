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

type ResponseUsers struct {
	StatusCode int                             `json:"statusCode"`
	Data       services.PaginationUsersService `json:"data"`
	Message    string                          `json:"message"`
}
type ResponseUser struct {
	StatusCode int               `json:"statusCode"`
	Data       *models.UserModel `json:"data"`
	Message    string            `json:"message"`
}

type RequestCreateUser struct {
	Name  string `json:"name" xml:"name" form:"name"`
	Email string `json:"email" xml:"email" form:"email"`
	Image string `json:"image" xml:"image" form:"image"`
}

// swagger:operation GET /user/{id} User GetUserController
// Get User by id
//
// ---
// responses:
//
//  404: CommonError
//  200: CommonSuccess
func GetUserController(c *fiber.Ctx) error {
	userIdParam := c.Params("id")
	fmt.Printf("userIdParam: %s\n", userIdParam)
	if len(userIdParam) > 0 {
		userId, err := strconv.Atoi(userIdParam)
		fmt.Printf("userId: %v\n", userId)
		if err == nil {
			user := services.GetUserService(userId)
			if user != nil {
				return c.Status(http.StatusOK).JSON(responses.Response{
					Status:  http.StatusOK,
					Message: "User found",
					Data:    user,
				})
			} else {
				return c.Status(http.StatusNotFound).JSON(responses.Response{
					Status:  http.StatusNotFound,
					Message: "User not found",
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
		Message: "User not found",
		Data:    nil,
	})
}

// swagger:operation POST /user User SetCreateUserController
// Create User
//
// ---
// responses:
//
//  400: CommonError
//  201: CommonSuccess
func SetCreateUserController(c *fiber.Ctx) error {
	params := new(RequestCreateUser)
	err := c.BodyParser(params)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.Response{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
	}
	fmt.Printf("params.Name: %s\n", params.Name)
	if len(params.Name) <= 0 {
		return c.Status(http.StatusBadRequest).JSON(responses.Response{
			Status:  http.StatusBadRequest,
			Message: "Error [name] field is required",
			Data:    nil,
		})
	}
	fmt.Printf("params.Email: %s\n", params.Email)
	if len(params.Email) <= 0 {
		return c.Status(http.StatusBadRequest).JSON(responses.Response{
			Status:  http.StatusBadRequest,
			Message: "Error [email] field is required",
			Data:    nil,
		})
	}
	user := services.PostUserService(params.Name, params.Email, params.Image)
	if user != nil {
		return c.Status(http.StatusCreated).JSON(responses.Response{
			Status:  http.StatusCreated,
			Message: "user created",
			Data:    user,
		})
	}
	return c.Status(http.StatusBadRequest).JSON(responses.Response{
		Status:  http.StatusBadRequest,
		Message: "user not created",
		Data:    nil,
	})
}

// swagger:operation PUT /user/{id} Product SetUpdateUserController
// Update User
//
// ---
// responses:
//
//  400: CommonError
//  200: CommonSuccess
func SetUpdateUserController(c *fiber.Ctx) error {
	userIdParam := c.Params("id")
	fmt.Printf("userIdParam: %s\n", userIdParam)
	if len(userIdParam) > 0 {
		userId, _ := strconv.Atoi(userIdParam)
		params := new(RequestCreateUser)
		err := c.BodyParser(params)
		if err != nil {
			return c.Status(http.StatusBadRequest).JSON(responses.Response{
				Status:  http.StatusBadRequest,
				Message: err.Error(),
				Data:    nil,
			})
		}
		fmt.Printf("params.Name: %s\n", params.Name)
		fmt.Printf("params.Email: %s\n", params.Email)
		fmt.Printf("params.Image: %s\n", params.Image)
		user := services.PutUserService(userId, params.Name, params.Email, params.Image)
		if user != nil {
			return c.Status(http.StatusOK).JSON(responses.Response{
				Status:  http.StatusOK,
				Message: "user updated",
				Data:    user,
			})
		}
		return c.Status(http.StatusBadRequest).JSON(responses.Response{
			Status:  http.StatusBadRequest,
			Message: "user not updated",
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

// swagger:operation GET /users [User] GetUsersController
// Get Users list
//
// ---
// responses:
//
//  400: CommonError
//  200: CommonSuccess
func GetUsersController(c *fiber.Ctx) error {
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
	users := services.GetUsersService(page, perPage, dateStart, dateEnd)
	fmt.Printf("users: %v\n", users)
	if users == nil {
		return c.Status(http.StatusOK).JSON(ResponseUsers{
			StatusCode: http.StatusOK,
			Message:    "Users not found",
			Data: services.PaginationUsersService{
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
		Message: "Users found",
		Data:    users,
	})
}
