package models

import (
	"github.com/danieljvx/talent-pitch-api/config"
	"time"
)

// ChallengeModel struct to describe book object.
type ChallengeModel struct {
	ID          uint       `db:"id" json:"id" validate:"required,uuid"`
	Title       string     `db:"title" json:"title" validate:"required,lte=255"`
	Description string     `db:"description" json:"description" validate:"required,lte=255"`
	Difficulty  int        `db:"difficulty" json:"difficulty" validate:"required,len=1"`
	UserID      uint       `db:"user_id" json:"user_id" validate:"required,uuid"`
	User        *UserModel `json:"user" gorm:"foreignKey:UserID;references:ID;"`
	CreatedAt   time.Time  `db:"created_at" json:"created_at"`
	UpdatedAt   time.Time  `db:"updated_at" json:"updated_at"`
}

func (ChallengeModel) TableName() string {
	return "challenges"
}

func GetChallenge(id int) *ChallengeModel {
	var challenge ChallengeModel

	err := config.DB.Find(&challenge, "id = ?", id).Error
	if err == nil {
		return &challenge
	}
	return nil
}
