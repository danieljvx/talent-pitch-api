package models

import (
	"github.com/danieljvx/talent-pitch-api/config"
	"time"
)

// UserModel struct to describe book object.
type UserModel struct {
	ID        uint      `db:"id" json:"id" gorm:"primaryKey;unique"`
	Name      string    `db:"name" json:"name"`
	Email     string    `db:"email" json:"email" gorm:"type:varchar(100);unique_index"`
	ImagePath string    `db:"image_path" json:"image_path"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}

func (UserModel) TableName() string {
	return "users"
}

func GetUser(id int) *UserModel {
	var user UserModel

	err := config.DB.Find(&user, "id = ?", id).Error
	if err == nil {
		return &user
	}
	return nil
}
