package models

import (
	"github.com/danieljvx/talent-pitch-api/config"
	"time"
)

// ProgramParticipantsModel struct to describe book object.
type ProgramParticipantsModel struct {
	ID          uint            `db:"id" json:"id" validate:"required,uuid"`
	ProgramID   uint            `db:"user_id" json:"user_id" validate:"required,uuid"`
	Program     *ProgramModel   `json:"program" gorm:"foreignKey:UserID;references:ID;"`
	ChallengeID uint            `db:"challenge_id" json:"challenge_id" validate:"required,uuid"`
	Challenge   *ChallengeModel `json:"challenge" gorm:"foreignKey:UserID;references:ID;"`
	CompanyID   uint            `db:"company_id" json:"company_id" validate:"required,uuid"`
	Company     *CompanyModel   `json:"company" gorm:"foreignKey:UserID;references:ID;"`
	CreatedAt   time.Time       `db:"created_at" json:"created_at"`
	UpdatedAt   time.Time       `db:"updated_at" json:"updated_at"`
}

func (ProgramParticipantsModel) TableName() string {
	return "program_participants"
}

func GetProgramParticipants(id int) *ProgramModel {
	var program ProgramModel

	err := config.DB.Find(&program, "id = ?", id).Error
	if err == nil {
		return &program
	}
	return nil
}

func GetProgramsParticipants() *[]ProgramModel {
	var programs []ProgramModel

	err := config.DB.Find(&programs, "email is not null").Error
	if err == nil {
		return &programs
	}
	return nil
}
