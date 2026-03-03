package seed

import (
	"log"

	"github.com/farhandwiyan/project-himaose/config"
	"github.com/farhandwiyan/project-himaose/models"
	"github.com/farhandwiyan/project-himaose/utils"
	"github.com/google/uuid"
)

func SeedAdmin() {
	password, _ := utils.HashPassword("admin123")
	
	admin := models.User{
		Username: "Super Admin",
		Password: password,
		Name: "Admin123",
		Role: "admin",
		PublicID: uuid.New(),
	}

	if err := config.DB.FirstOrCreate(&admin, models.User{Username:admin.Username, Name: admin.Name}).Error; err != nil {
		log.Println("Failed to seed admin", err)
	} else {
		log.Println("Admin user seeded")
	}
	
}