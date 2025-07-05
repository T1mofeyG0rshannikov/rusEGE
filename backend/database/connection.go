package database

import (
	"rusEGE/database/models"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var dbase *gorm.DB

func Init() *gorm.DB {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Ошибка при загрузке файла .env: %v", err)
	}

	databaseURL := os.Getenv("DATABASEURL")

	db, err := gorm.Open(sqlite.Open(databaseURL), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&models.Task{}, &models.Word{}, &models.User{}, &models.UserWord{}, &models.IndexSeo{})
	return db
}

func GetDB() *gorm.DB {
	if dbase == nil {
		dbase = Init()
		var sleep = time.Duration(1)
		for dbase == nil {
			sleep = sleep * 2
			fmt.Println("database is unavailable. Wait", sleep, "sec")
			time.Sleep(sleep * time.Second)
		}
	}

	return dbase
}
