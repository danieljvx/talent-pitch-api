package models

import (
	"github.com/danieljvx/talent-pitch-api/config"
	"time"
)

// ProgramParticipantsModel struct to describe book object.
type ProgramParticipantsModel struct {
	ID          uint            `db:"id" json:"id" validate:"required,uuid"`
	ProgramID   uint            `db:"program_id" json:"program_id" validate:"required,uuid"`
	Program     *ProgramModel   `json:"program"`
	ChallengeID uint            `db:"challenge_id" json:"challenge_id" validate:"required,uuid"`
	Challenge   *ChallengeModel `json:"challenge"`
	CompanyID   uint            `db:"company_id" json:"company_id" validate:"required,uuid"`
	Company     *CompanyModel   `json:"company"`
	UserID      uint            `db:"user_id" json:"user_id" validate:"required,uuid"`
	User        *UserModel      `json:"user"`
	CreatedAt   time.Time       `db:"created_at" json:"created_at"`
	UpdatedAt   time.Time       `db:"updated_at" json:"updated_at"`
}

func (ProgramParticipantsModel) TableName() string {
	return "program_participants"
}

func GetProgramParticipant(id int) *ProgramParticipantsModel {
	var programParticipant ProgramParticipantsModel

	err := config.DB.Find(&programParticipant, "id = ?", id).Error
	if err == nil {
		return &programParticipant
	}
	return nil
}
