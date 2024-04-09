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

type ResponseChallenges struct {
	StatusCode int                                  `json:"statusCode"`
	Data       services.PaginationChallengesService `json:"data"`
	Message    string                               `json:"message"`
}
type ResponseChallenge struct {
	StatusCode int                    `json:"statusCode"`
	Data       *models.ChallengeModel `json:"data"`
	Message    string                 `json:"message"`
}

type RequestCreateChallenge struct {
	Title       string `json:"title" xml:"title" form:"title"`
	Description string `json:"description" xml:"description" form:"description"`
	Difficulty  int    `json:"difficulty" xml:"difficulty" form:"difficulty"`
	UserID      int    `json:"user_id" xml:"user_id" form:"user_id"`
}

func GetChallengeController(c *fiber.Ctx) error {
	challengeIdParam := c.Params("id")
	fmt.Printf("challengeIdParam: %s\n", challengeIdParam)
	if len(challengeIdParam) > 0 {
		challengeId, err := strconv.Atoi(challengeIdParam)
		fmt.Printf("challengeId: %v\n", challengeId)
		if err == nil {
			challenge := services.GetChallengeService(challengeId)
			if challenge != nil {
				return c.Status(http.StatusOK).JSON(responses.Response{
					Status:  http.StatusOK,
					Message: "Challenge found",
					Data:    challenge,
				})
			} else {
				return c.Status(http.StatusNotFound).JSON(responses.Response{
					Status:  http.StatusNotFound,
					Message: "Challenge not found",
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
		Message: "Challenge not found",
		Data:    nil,
	})
}

func SetCreateChallengeController(c *fiber.Ctx) error {
	params := new(RequestCreateChallenge)
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
	fmt.Printf("params.Difficulty: %d\n", params.Difficulty)
	if params.Difficulty <= 0 {
		return c.Status(http.StatusBadRequest).JSON(responses.Response{
			Status:  http.StatusBadRequest,
			Message: "Error [difficulty] field is required",
			Data:    nil,
		})
	}
	fmt.Printf("params.UserID: %d\n", params.UserID)
	if params.UserID <= 0 {
		return c.Status(http.StatusBadRequest).JSON(responses.Response{
			Status:  http.StatusBadRequest,
			Message: "Error [user_id] field is required",
			Data:    nil,
		})
	}
	challenge := services.PostChallengeService(params.Title, params.Description, params.Difficulty, uint(params.UserID))
	if challenge != nil {
		return c.Status(http.StatusCreated).JSON(responses.Response{
			Status:  http.StatusCreated,
			Message: "challenge created",
			Data:    challenge,
		})
	}
	return c.Status(http.StatusBadRequest).JSON(responses.Response{
		Status:  http.StatusBadRequest,
		Message: "challenge not created",
		Data:    nil,
	})
}

func SetUpdateChallengeController(c *fiber.Ctx) error {
	challengeIdParam := c.Params("id")
	fmt.Printf("challengeIdParam: %s\n", challengeIdParam)
	if len(challengeIdParam) > 0 {
		challengeId, _ := strconv.Atoi(challengeIdParam)
		params := new(RequestCreateChallenge)
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
		fmt.Printf("params.Difficulty: %d\n", params.Difficulty)
		fmt.Printf("params.UserID: %d\n", params.UserID)
		challenge := services.PutChallengeService(challengeId, params.Title, params.Description, params.Difficulty, uint(params.UserID))
		if challenge != nil {
			return c.Status(http.StatusOK).JSON(responses.Response{
				Status:  http.StatusOK,
				Message: "challenge updated",
				Data:    challenge,
			})
		}
		return c.Status(http.StatusBadRequest).JSON(responses.Response{
			Status:  http.StatusBadRequest,
			Message: "challenge not updated",
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

func GetChallengesController(c *fiber.Ctx) error {
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
	challenges := services.GetChallengesService(page, perPage, dateStart, dateEnd)
	fmt.Printf("challenges: %v\n", challenges)
	if challenges == nil {
		return c.Status(http.StatusOK).JSON(ResponseChallenges{
			StatusCode: http.StatusOK,
			Message:    "Challenges not found",
			Data: services.PaginationChallengesService{
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
		Message: "Challenges found",
		Data:    challenges,
	})
}
