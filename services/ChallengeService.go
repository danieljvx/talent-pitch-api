package services

import (
	"github.com/danieljvx/talent-pitch-api/config"
	"github.com/danieljvx/talent-pitch-api/models"
	"log"
	"math"
	"time"
)

type PaginationChallengesService struct {
	Page     int                     `json:"page"`
	PerPage  int                     `json:"perPage"`
	PrevPage int                     `json:"prevPage"`
	NextPage int                     `json:"nextPage"`
	LastPage int                     `json:"lastPage"`
	Total    int64                   `json:"total"`
	Data     []models.ChallengeModel `json:"data"`
}

type PostChallengeServiceResponse struct {
	ID int `json:"id"`
}

func GetChallengeService(programId int) *models.ChallengeModel {
	program := models.GetChallenge(programId)
	return program
}

func PostChallengeService(title string, description string, difficulty int, userId uint) *models.ChallengeModel {
	program := models.ChallengeModel{Title: title, Description: description, Difficulty: difficulty, UserID: userId}
	err := config.DB.Create(&program).Error

	if err != nil {
		log.Println(err)
		return nil
	}

	return &program
}

func PutChallengeService(id int, title string, description string, difficulty int, userId uint) *models.ChallengeModel {
	program := models.ChallengeModel{UpdatedAt: time.Now()}
	if title != "" {
		program.Title = title
	}
	if description != "" {
		program.Description = description
	}
	if difficulty <= 0 {
		program.UserID = userId
	}
	if userId <= 0 {
		program.UserID = userId
	}

	err := config.DB.Model(&models.ChallengeModel{}).Where("id = ?", id).UpdateColumns(program).Error

	if err != nil {
		log.Println(err)
		return nil
	}

	return GetChallengeService(id)
}

func GetChallengesService(page int, perPage int, dateStart time.Time, dateEnd time.Time) *PaginationChallengesService {
	if page == 0 {
		page = 1
	}
	var paginationChallengesService = PaginationChallengesService{
		Page:     page,
		PerPage:  perPage,
		PrevPage: page,
		NextPage: page,
	}
	var programsCount []models.ChallengeModel
	Query := config.DB.Where("id is not null")
	if !dateStart.IsZero() && dateEnd.IsZero() {
		Query = Query.Where("created_at = ?", dateStart)
		log.Printf("created_at = %s\n", dateStart)
	} else if !dateStart.IsZero() && !dateEnd.IsZero() {
		log.Printf("created_at >= %s AND created_at <= %s\n", dateStart, dateEnd)
		Query = Query.Where("created_at >= ? AND created_at <= ?", dateStart, dateEnd)
	}
	errChallengesCount := Query.Find(&programsCount).Count(&paginationChallengesService.Total).Error
	if errChallengesCount == nil {
		paginationChallengesService.LastPage = int(math.Ceil(float64(paginationChallengesService.Total) / float64(perPage)))
		if paginationChallengesService.LastPage == 0 {
			paginationChallengesService.LastPage = 1
		}
	}
	var offset = paginationChallengesService.PerPage * (paginationChallengesService.Page - 1)
	if paginationChallengesService.Page > 1 {
		paginationChallengesService.PrevPage = paginationChallengesService.Page - 1
	}
	if paginationChallengesService.Page < paginationChallengesService.LastPage {
		paginationChallengesService.NextPage = paginationChallengesService.Page + 1
	}
	errChallenges := Query.Offset(offset).Limit(paginationChallengesService.PerPage).Order("id desc").Find(&paginationChallengesService.Data).Error
	if errChallenges == nil {
		return &paginationChallengesService
	}
	return nil
}
