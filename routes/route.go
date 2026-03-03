package routes

import (
	"log"

	"github.com/farhandwiyan/project-himaose/config"
	"github.com/farhandwiyan/project-himaose/controllers"
	"github.com/farhandwiyan/project-himaose/utils"
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/joho/godotenv"
)

func Setup(app *fiber.App, uc *controllers.UserController, 
	pc *controllers.ProgramKerjaController, lc *controllers.LombaController,
	bc *controllers.BeasiswaController) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	app.Post("/v1/auth/register", uc.Register)
	app.Post("/v1/auth/login", uc.Login)
	
	// JWT protected routes
	api := app.Group("/api/v1", jwtware.New(jwtware.Config{
		SigningKey: []byte(config.AppConfig.JWTSecret),
		ContextKey: "user",
		ErrorHandler: func (c *fiber.Ctx, err error) error {
			return utils.Unauthorized(c, "Error unauthorized", err.Error())
		},
	}))

	prokerGroup := api.Group("/proker")
	prokerGroup.Post("/", pc.CreateProker)
	prokerGroup.Put("/:id", pc.UpdateProker)
	prokerGroup.Get("/my", pc.GetMyProkerPaginate)
	prokerGroup.Delete("/:id", pc.DeleteProgramKerja)

	lombaGroup := api.Group("/lomba")
	lombaGroup.Post("/", lc.CreateLomba)
	lombaGroup.Put("/:id", lc.UpdateLomba)
	lombaGroup.Get("/my", lc.GetMyLombaPaginate)
	lombaGroup.Delete("/:id", lc.DeleteLomba)

	beasiswaGroup := api.Group("/beasiswa")
	beasiswaGroup.Post("/", bc.CreateBeasiswa)
	beasiswaGroup.Put("/:id", bc.UpdateBeasiswa)
	beasiswaGroup.Get("/my", bc.GetMyBeasiswaPaginate)
	beasiswaGroup.Delete("/:id", bc.DeleteBeasiswa)
}