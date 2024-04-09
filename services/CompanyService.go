package services

import (
	"github.com/danieljvx/talent-pitch-api/config"
	"github.com/danieljvx/talent-pitch-api/models"
	"log"
	"math"
	"time"
)

type PaginationCompaniesService struct {
	Page     int                   `json:"page"`
	PerPage  int                   `json:"perPage"`
	PrevPage int                   `json:"prevPage"`
	NextPage int                   `json:"nextPage"`
	LastPage int                   `json:"lastPage"`
	Total    int64                 `json:"total"`
	Data     []models.CompanyModel `json:"data"`
}

type PostCompanyServiceResponse struct {
	ID int `json:"id"`
}

func GetCompanyService(companyId int) *models.CompanyModel {
	company := models.GetCompany(companyId)
	return company
}

func PostCompanyService(name string, image string, location string, industry string, userId uint) *models.CompanyModel {
	company := models.CompanyModel{Name: name, ImagePath: image, Location: location, Industry: industry, UserID: userId}
	err := config.DB.Create(&company).Error

	if err != nil {
		log.Println(err)
		return nil
	}

	return &company
}

func PutCompanyService(id int, name string, image string, location string, industry string, userId uint) *models.CompanyModel {
	company := models.CompanyModel{UpdatedAt: time.Now()}
	if name != "" {
		company.Name = name
	}
	if image != "" {
		company.ImagePath = image
	}
	if location != "" {
		company.Location = location
	}
	if industry != "" {
		company.Industry = industry
	}
	if userId <= 0 {
		company.UserID = userId
	}

	err := config.DB.Model(&models.CompanyModel{}).Where("id = ?", id).UpdateColumns(company).Error

	if err != nil {
		log.Println(err)
		return nil
	}

	return GetCompanyService(id)
}

func GetCompaniesService(page int, perPage int, dateStart time.Time, dateEnd time.Time) *PaginationCompaniesService {
	if page == 0 {
		page = 1
	}
	var paginationCompaniesService = PaginationCompaniesService{
		Page:     page,
		PerPage:  perPage,
		PrevPage: page,
		NextPage: page,
	}
	var companysCount []models.CompanyModel
	Query := config.DB.Where("id is not null")
	if !dateStart.IsZero() && dateEnd.IsZero() {
		Query = Query.Where("created_at = ?", dateStart)
		log.Printf("created_at = %s\n", dateStart)
	} else if !dateStart.IsZero() && !dateEnd.IsZero() {
		log.Printf("created_at >= %s AND created_at <= %s\n", dateStart, dateEnd)
		Query = Query.Where("created_at >= ? AND created_at <= ?", dateStart, dateEnd)
	}
	errCompaniesCount := Query.Find(&companysCount).Count(&paginationCompaniesService.Total).Error
	if errCompaniesCount == nil {
		paginationCompaniesService.LastPage = int(math.Ceil(float64(paginationCompaniesService.Total) / float64(perPage)))
		if paginationCompaniesService.LastPage == 0 {
			paginationCompaniesService.LastPage = 1
		}
	}
	var offset = paginationCompaniesService.PerPage * (paginationCompaniesService.Page - 1)
	if paginationCompaniesService.Page > 1 {
		paginationCompaniesService.PrevPage = paginationCompaniesService.Page - 1
	}
	if paginationCompaniesService.Page < paginationCompaniesService.LastPage {
		paginationCompaniesService.NextPage = paginationCompaniesService.Page + 1
	}
	errCompanies := Query.Offset(offset).Limit(paginationCompaniesService.PerPage).Order("id desc").Find(&paginationCompaniesService.Data).Error
	if errCompanies == nil {
		return &paginationCompaniesService
	}
	return nil
}
