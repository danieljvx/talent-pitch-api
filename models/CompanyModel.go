package models

import (
	"github.com/danieljvx/talent-pitch-api/config"
	"time"
)

// CompanyModel struct to describe book object.
type CompanyModel struct {
	ID        uint       `db:"id" json:"id" validate:"required,uuid"`
	Name      string     `db:"name" json:"name" validate:"required,lte=255"`
	ImagePath string     `db:"image_path" json:"image_path"`
	Location  string     `db:"location" json:"location" validate:"required,lte=255"`
	Industry  string     `db:"industry" json:"industry" validate:"required,lte=255"`
	UserID    uint       `db:"user_id" json:"user_id" validate:"required,uuid"`
	User      *UserModel `json:"user" gorm:"foreignKey:UserID;references:ID;"`
	CreatedAt time.Time  `db:"created_at" json:"created_at"`
	UpdatedAt time.Time  `db:"updated_at" json:"updated_at"`
}

func (CompanyModel) TableName() string {
	return "companies"
}

func GetCompany(id int) *CompanyModel {
	var company CompanyModel

	err := config.DB.Find(&company, "id = ?", id).Error
	if err == nil {
		return &company
	}
	return nil
}
