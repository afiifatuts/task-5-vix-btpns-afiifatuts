package initializer

import "github.com/afiifatuts/go-authentication/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.User{}, models.Photo{})
}
