package database

import (
	"fmt"
	"time"
)

var dbase *gorm.DB

func Init() *gorm.DB{
	db, err := gorm.Open("sqlite")

	if err != nil {
		log.Fatal(err)
	}
	db.AutoMigrate(&models.Task, &models.Word, &models.User, &models.UserWord)
	return db
}

func GetDB() *gorm.DB {
	if dbase == nil {
		dbase = Init()
		var sleep = time.Duration(1)
		for dbase == nil {
			sleep = sleep * 2
			fmt.Print("database is unavailable. Wait %d sec \n", sleep)
			time.Sleep(sleep * time.Second)
		}
	}

	return dbase
}