package model

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

const dsn string = "root:@tcp(127.0.0.1:3306)/gourly?charset=utf8mb4&parseTime=True&loc=Local"

func Setup() error {
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	err = DB.AutoMigrate(&URL{})
	if err != nil {
		return err
	}

	return err
}
