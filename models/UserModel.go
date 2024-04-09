package models

import (
	"fmt"
	"github.com/danieljvx/talent-pitch-api/config"
	"time"
)

// UserModel struct to describe book object.
type UserModel struct {
	ID                  uint                       `db:"id" json:"id" gorm:"primaryKey;unique"`
	Name                string                     `db:"name" json:"name"`
	Email               string                     `db:"email" json:"email" gorm:"type:varchar(100);unique_index"`
	ImagePath           string                     `db:"image_path" json:"image_path"`
	ProgramParticipants []ProgramParticipantsModel `json:"program_participants" gorm:"foreignKey:ID;references:ProgramParticipants"`
	CreatedAt           time.Time                  `db:"created_at" json:"created_at"`
	UpdatedAt           time.Time                  `db:"updated_at" json:"updated_at"`
}

func (UserModel) TableName() string {
	return "users"
}

func GetUser(id int) *UserModel {
	fmt.Println("GetUser")
	var user UserModel
	err := config.DB.Find(&user, "id = ?", id).Error
	if err == nil {
		user.ProgramParticipants = GetUserProgramParticipants(id)
		return &user
	}
	return nil
}

func GetUserProgramParticipants(userId int) []ProgramParticipantsModel {
	var programParticipants []ProgramParticipantsModel
	var programParticipantsTmp []ProgramParticipantsModel
	config.DB.Find(&programParticipantsTmp, "user_id = ?", userId)
	fmt.Printf("programParticipantsTmp %v\n", programParticipantsTmp)
	for _, pp := range programParticipantsTmp {
		var program ProgramModel
		errProgram := config.DB.Find(&program, "id = ?", pp.ProgramID).Error
		fmt.Printf("pp.ProgramID %d\n", pp.ProgramID)
		fmt.Printf("errProgram %s\n", errProgram)
		if errProgram == nil {
			pp.Program = &program
		}
		var challenge ChallengeModel
		errChallenge := config.DB.Find(&challenge, "id = ?", pp.ChallengeID).Error
		fmt.Printf("pp.ChallengeID %d\n", pp.ChallengeID)
		fmt.Printf("errChallenge %s\n", errChallenge)
		if errChallenge == nil {
			pp.Challenge = &challenge
		}
		var company CompanyModel
		errCompany := config.DB.Find(&company, "id = ?", pp.CompanyID).Error
		fmt.Printf("pp.CompanyID %d\n", pp.CompanyID)
		fmt.Printf("errCompany %s\n", errCompany)
		if errCompany == nil {
			pp.Company = &company
		}
		programParticipants = append(programParticipants, pp)
	}
	fmt.Printf("programParticipants %v\n", programParticipants)
	return programParticipants
}
