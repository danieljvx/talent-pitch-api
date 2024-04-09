package services

import (
	"github.com/danieljvx/talent-pitch-api/config"
	"github.com/danieljvx/talent-pitch-api/models"
	"log"
	"math"
	"time"
)

type PaginationUsersService struct {
	Page     int                `json:"page"`
	PerPage  int                `json:"perPage"`
	PrevPage int                `json:"prevPage"`
	NextPage int                `json:"nextPage"`
	LastPage int                `json:"lastPage"`
	Total    int64              `json:"total"`
	Data     []models.UserModel `json:"data"`
}

type PostUserServiceResponse struct {
	ID int `json:"id"`
}

func GetUserService(userId int) *models.UserModel {
	user := models.GetUser(userId)
	return user
}

func PostUserService(name string, email string, image string) *models.UserModel {
	user := models.UserModel{Name: name, Email: email, ImagePath: image}

	err := config.DB.Create(&user).Error

	if err != nil {
		log.Println(err)
		return nil
	}

	return &user
}

func PutUserService(id int, name string, email string, image string) *models.UserModel {
	user := models.UserModel{UpdatedAt: time.Now()}
	if name != "" {
		user.Name = name
	}
	if email != "" {
		user.Email = email
	}
	if image != "" {
		user.ImagePath = image
	}

	err := config.DB.Model(&models.UserModel{}).Where("id = ?", id).UpdateColumns(user).Error

	if err != nil {
		log.Println(err)
		return nil
	}

	return GetUserService(id)
}

func GetUsersService(page int, perPage int, dateStart time.Time, dateEnd time.Time) *PaginationUsersService {
	if page == 0 {
		page = 1
	}
	var paginationUsersService = PaginationUsersService{
		Page:     page,
		PerPage:  perPage,
		PrevPage: page,
		NextPage: page,
	}
	var usersCount []models.UserModel
	Query := config.DB.Where("id is not null")
	if !dateStart.IsZero() && dateEnd.IsZero() {
		Query = Query.Where("created_at = ?", dateStart)
		log.Printf("created_at = %s\n", dateStart)
	} else if !dateStart.IsZero() && !dateEnd.IsZero() {
		log.Printf("created_at >= %s AND created_at <= %s\n", dateStart, dateEnd)
		Query = Query.Where("created_at >= ? AND created_at <= ?", dateStart, dateEnd)
	}
	errUsersCount := Query.Find(&usersCount).Count(&paginationUsersService.Total).Error
	if errUsersCount == nil {
		paginationUsersService.LastPage = int(math.Ceil(float64(paginationUsersService.Total) / float64(perPage)))
		if paginationUsersService.LastPage == 0 {
			paginationUsersService.LastPage = 1
		}
	}
	var offset = paginationUsersService.PerPage * (paginationUsersService.Page - 1)
	if paginationUsersService.Page > 1 {
		paginationUsersService.PrevPage = paginationUsersService.Page - 1
	}
	if paginationUsersService.Page < paginationUsersService.LastPage {
		paginationUsersService.NextPage = paginationUsersService.Page + 1
	}
	var users []models.UserModel
	errUsers := Query.Offset(offset).Limit(paginationUsersService.PerPage).Order("id desc").Find(&users).Error
	for _, u := range users {
		u.ProgramParticipants = models.GetUserProgramParticipants(int(u.ID))
		paginationUsersService.Data = append(paginationUsersService.Data, u)
	}
	if errUsers == nil {
		return &paginationUsersService
	}
	return nil
}
