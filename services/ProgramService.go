package services

import (
	"github.com/danieljvx/talent-pitch-api/config"
	"github.com/danieljvx/talent-pitch-api/models"
	"log"
	"math"
	"time"
)

type PaginationProgramsService struct {
	Page     int                   `json:"page"`
	PerPage  int                   `json:"perPage"`
	PrevPage int                   `json:"prevPage"`
	NextPage int                   `json:"nextPage"`
	LastPage int                   `json:"lastPage"`
	Total    int64                 `json:"total"`
	Data     []models.ProgramModel `json:"data"`
}

type PostProgramServiceResponse struct {
	ID int `json:"id"`
}

func GetProgramService(programId int) *models.ProgramModel {
	program := models.GetProgram(programId)
	return program
}

func PostProgramService(title string, description string, startDate string, endDate string, userId uint) *models.ProgramModel {
	program := models.ProgramModel{Title: title, Description: description, StartDate: startDate, EndDate: endDate, UserID: userId}
	err := config.DB.Create(&program).Error

	if err != nil {
		log.Println(err)
		return nil
	}

	return &program
}

func PutProgramService(id int, title string, description string, startDate string, endDate string, userId uint) *models.ProgramModel {
	program := models.ProgramModel{UpdatedAt: time.Now()}
	if title != "" {
		program.Title = title
	}
	if description != "" {
		program.Description = description
	}
	if startDate != "" {
		program.StartDate = startDate
	}
	if endDate != "" {
		program.EndDate = endDate
	}
	if userId <= 0 {
		program.UserID = userId
	}

	err := config.DB.Model(&models.ProgramModel{}).Where("id = ?", id).UpdateColumns(program).Error

	if err != nil {
		log.Println(err)
		return nil
	}

	return GetProgramService(id)
}

func GetProgramsService(page int, perPage int, dateStart time.Time, dateEnd time.Time) *PaginationProgramsService {
	if page == 0 {
		page = 1
	}
	var paginationProgramsService = PaginationProgramsService{
		Page:     page,
		PerPage:  perPage,
		PrevPage: page,
		NextPage: page,
	}
	var programsCount []models.ProgramModel
	Query := config.DB.Where("id is not null")
	if !dateStart.IsZero() && dateEnd.IsZero() {
		Query = Query.Where("created_at = ?", dateStart)
		log.Printf("created_at = %s\n", dateStart)
	} else if !dateStart.IsZero() && !dateEnd.IsZero() {
		log.Printf("created_at >= %s AND created_at <= %s\n", dateStart, dateEnd)
		Query = Query.Where("created_at >= ? AND created_at <= ?", dateStart, dateEnd)
	}
	errProgramsCount := Query.Find(&programsCount).Count(&paginationProgramsService.Total).Error
	if errProgramsCount == nil {
		paginationProgramsService.LastPage = int(math.Ceil(float64(paginationProgramsService.Total) / float64(perPage)))
		if paginationProgramsService.LastPage == 0 {
			paginationProgramsService.LastPage = 1
		}
	}
	var offset = paginationProgramsService.PerPage * (paginationProgramsService.Page - 1)
	if paginationProgramsService.Page > 1 {
		paginationProgramsService.PrevPage = paginationProgramsService.Page - 1
	}
	if paginationProgramsService.Page < paginationProgramsService.LastPage {
		paginationProgramsService.NextPage = paginationProgramsService.Page + 1
	}
	errPrograms := Query.Offset(offset).Limit(paginationProgramsService.PerPage).Order("id desc").Find(&paginationProgramsService.Data).Error
	if errPrograms == nil {
		return &paginationProgramsService
	}
	return nil
}
