package migrations

import (
	"github.com/imatheus-lucas/golang_api/models"
	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) {
	db.AutoMigrate(models.Book{})
}
