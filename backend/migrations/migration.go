package migrations

import (
	"gofiber-api/models"
	"log"

	"gorm.io/gorm"
)

var ModelsList []interface{}

func init() {
	// Registering all models in the list
	ModelsList = []interface{}{
		&models.User{},
		&models.Business{},
	}
}

func Migrate(db *gorm.DB) {

	for _, model := range ModelsList {
		if err := db.AutoMigrate(model); err != nil {
			log.Fatalln("Migration Error:", err)
		}
	}
	log.Println("Migration Completed 😀")
}
