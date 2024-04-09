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
	fmt.Println(port)
	if err != nil {
		fmt.Println("Faild to parse database port")
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", Config("DB_USER"), Config("DB_PASS"), Config("DB_HOST"), port, Config("DB_NAME"))
	fmt.Println(dsn)
	DB, err = gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,   // data source name
		DefaultStringSize:         256,   // default size for string fields
		DisableDatetimePrecision:  true,  // disable datetime precision, which not supported before MySQL 5.6
		DontSupportRenameIndex:    true,  // drop & create when rename index, rename index not supported before MySQL 5.7, MariaDB
		DontSupportRenameColumn:   true,  // `change` when rename column, rename column not supported before MySQL 8, MariaDB
		SkipInitializeWithVersion: false, // auto configure based on currently MySQL version
	}), &gorm.Config{})

	if err != nil {
		fmt.Println("Fail connection to database")
	}

	fmt.Println("Connection open to database")
	// DB.AutoMigrate(&models.Task{})
	// fmt.Println("Database migrated")
}
