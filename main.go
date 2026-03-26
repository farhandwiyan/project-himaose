package main

import (
	"log"
	"os"

	"github.com/farhandwiyan/project-himaose/config"
	"github.com/farhandwiyan/project-himaose/controllers"
	"github.com/farhandwiyan/project-himaose/database/seed"
	"github.com/farhandwiyan/project-himaose/repositories"
	"github.com/farhandwiyan/project-himaose/routes"
	"github.com/farhandwiyan/project-himaose/services"
	"github.com/gofiber/fiber/v2"
)

func main() {
	config.LoadEnv()
	config.ConnectDB()

	seed.SeedAdmin()
	app := fiber.New()

	// user
	userRepo := repositories.NewUserRepository()
	userService := services.NewUserService(userRepo)
	userController := controllers.NewUserController(userService)

	// program kerja
	prokerRepo := repositories.NewProgramKerjaRepository()
	prokerService := services.NewProgramKerjaService(prokerRepo, userRepo)
	prokerController := controllers.NewProgramKerjaController(prokerService)

	// lomba 
	lombaRepo := repositories.NewLombaRepository()
	lombaService := services.NewLombaService(lombaRepo, userRepo)
	lombaController := controllers.NewLombaController(lombaService)

	// beasiswa
	beasiswaRepo := repositories.NewBeasiswaRepository()
	beasiswaService := services.NewBeasiswaService(beasiswaRepo, userRepo)
	beasiswaController := controllers.NewBeasiswaController(beasiswaService, )
	
	routes.Setup(app, userController, prokerController, lombaController, beasiswaController)

	port := os.Getenv("PORT")

	if port == "" {
		port = config.AppConfig.AppPort 
	}

	log.Println("Server is running on port :", port)
	log.Fatal(app.Listen(":"+port))
}