package database

import (
	"log"

	"sw-api-go/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConectDB() {
	url := "host=127.0.0.1 user=postgres password=root dbname=sw port=5432"
	DB, err = gorm.Open(postgres.Open(url))
	if err != nil {
		log.Panic("Error to conect database")
	}
	DB.AutoMigrate(&model.Character{})
}
