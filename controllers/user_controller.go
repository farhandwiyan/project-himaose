package controllers

import (
	"github.com/farhandwiyan/project-himaose/models"
	"github.com/farhandwiyan/project-himaose/services"
	"github.com/farhandwiyan/project-himaose/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/copier"
)

type UserController struct {
	service services.UserService
}

func NewUserController(s services.UserService) *UserController {
	return &UserController{service: s}
}

func (c *UserController) Register(ctx *fiber.Ctx) error {
	user := new(models.User)

	if err := ctx.BodyParser(user); err != nil {
		return utils.BadRequest(ctx, "Failed Parsing Data", err.Error())
	}

	if err := c.service.Register(user); err != nil {
		return utils.BadRequest(ctx, "Registration Failed", err.Error())
	}

	var userResp models.UserResponse
	_ = copier.Copy(&userResp, &user)

	return utils.Success(ctx, "Register Success", userResp)
}

func (c *UserController) Login(ctx *fiber.Ctx) error {
	var body struct {
		Username string `json:"username"`
		Password string `json:"password"`
	} 

	if err := ctx.BodyParser(&body); err != nil {
		return utils.BadRequest(ctx, "Invalid Request", err.Error())
	}

	user, err := c.service.Login(body.Username, body.Password)
	if err != nil {
		return utils.Unauthorized(ctx, "Login Failed", err.Error())
	}

	token, _ := utils.GenerateToken(user.InternalID, user.Role, body.Username, user.PublicID)
	refreshToken, _ := utils.GenerateRefreshToken(user.InternalID)

	var userResp models.UserResponse
	_ = copier.Copy(&userResp, &user)

	return utils.Success(ctx, "Login Success", fiber.Map{
		"access_token": token,
		"refresh_token": refreshToken,
		"user": userResp,
	})
}