package models

import (
	"github.com/danieljvx/talent-pitch-api/config"
	"time"
)

// ProgramModel struct to describe book object.
type ProgramModel struct {
	ID          uint       `db:"id" json:"id" validate:"required,uuid"`
	Title       string     `db:"title" json:"title" validate:"required,lte=255"`
	Description string     `db:"description" json:"description" validate:"required,lte=255"`
	StartDate   string     `db:"start_date" json:"start_date"`
	EndDate     string     `db:"end_date" json:"end_date"`
	UserID      uint       `db:"user_id" json:"user_id" validate:"required,uuid"`
	User        *UserModel `json:"user" gorm:"foreignKey:UserID;references:ID;"`
	CreatedAt   time.Time  `db:"created_at" json:"created_at"`
	UpdatedAt   time.Time  `db:"updated_at" json:"updated_at"`
}

func (ProgramModel) TableName() string {
	return "programs"
}

func GetProgram(id int) *ProgramModel {
	var program ProgramModel

	err := config.DB.Find(&program, "id = ?", id).Error
	if err == nil {
		return &program
	}
	return nil
}

func GetPrograms() *[]ProgramModel {
	var programs []ProgramModel

	err := config.DB.Find(&programs, "email is not null").Error
	if err == nil {
		return &programs
	}
	return nil
}
