package configs

import (
	"chillroom/models"
	"log"
)

func SyncDB() {
	err := DB.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal("Error")
	}
}
