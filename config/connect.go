package config

import (
	"fmt"
	"strconv"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDB() {
	var err error

	p := Config("DB_PORT")
	port, err := strconv.ParseUint(p, 10, 32)

	if err != nil {
		panic("Faild to parse database port")
	}

	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		Config("DB_HOST"), port, Config("DB_USER"), Config("DB_PASSWORD"),
		Config("DB_NAME"))

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Fail connection to database")
	}

	fmt.Println("Connection open to database")
	// DB.AutoMigrate(&models.Task{})
	// fmt.Println("Database migrated")
}
