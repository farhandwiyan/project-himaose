package controllers

import (
	"math"
	"strconv"

	"github.com/farhandwiyan/project-himaose/models"
	"github.com/farhandwiyan/project-himaose/services"
	"github.com/farhandwiyan/project-himaose/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

type LombaController struct {
	service services.LombaService
}

func NewLombaController(s services.LombaService) *LombaController {
	return &LombaController{service: s}
}

func (c *LombaController) CreateLomba(ctx *fiber.Ctx) error {
	lomba := new(models.Lomba)
	user := ctx.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	if err := ctx.BodyParser(lomba); err != nil {
		return utils.BadRequest(ctx, "Failed to read request", err.Error())
	}

	userID, ok := claims["user_id"].(float64)

	if ok {
		lomba.CreatedBy = int64(userID)
	} else {
		return utils.BadRequest(ctx, "Failed to read request", "user_id not found in token")
	}

	if err := c.service.Create(lomba); err != nil {
		return utils.BadRequest(ctx, "Failed to saved data", err.Error())
	}

	return utils.Success(ctx, "Lomba was successfully created", lomba)
}

func (c *LombaController) UpdateLomba(ctx *fiber.Ctx) error {
	publicID := ctx.Params("id")
	lomba := new (models.Lomba)

	if err := ctx.BodyParser(lomba); err != nil {
		return utils.BadRequest(ctx, "Failed Parsing Data", err.Error())
	}

	if _, err := uuid.Parse(publicID); err != nil {
		return utils.BadRequest(ctx, "Invalid ID", err.Error())
	}

	existingLomba, err := c.service.GetByPublicID(publicID)
	if err != nil {
		return utils.NotFound(ctx, "Lomba not found", err.Error())
	}

	lomba.PublicID = existingLomba.PublicID
	lomba.InternalID = existingLomba.InternalID
	lomba.CreatedBy = existingLomba.CreatedBy

	if err := c.service.Update(lomba); err != nil {
		return utils.BadRequest(ctx, "Failed Update Lomba", err.Error())
	}

	return utils.Success(ctx, "Lomba Success Updated", lomba)
}

func (c *LombaController) GetMyLombaPaginate(ctx *fiber.Ctx) error {
	user := ctx.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userID := claims["user_id"].(float64)

	page, _ := strconv.Atoi(ctx.Query("page", "1"))
	limit, _ := strconv.Atoi(ctx.Query("limit", "10"))
	offset := (page - 1) * limit

	filter := ctx.Query("filter", "")
	sort := ctx.Query("sort", "")

	lomba, total, err := c.service.GetAllByUserPaginate(int64(userID),filter,sort,limit,offset)
	if err != nil {
		return utils.InternalServerError(ctx, "Failed get lomba", err.Error())
	}

	meta := utils.PaginationMeta{
		Page: page,
		Limit: limit,
		Total: int(total),
		TotalPage: int(math.Ceil(float64(total) / float64(limit))),
		Filter: filter,
		Sort: sort,
	}

	return utils.SuccessPagination(ctx, "Success get data lomba", lomba, meta)
}

func (c *LombaController) DeleteLomba(ctx *fiber.Ctx) error {	
	pubID := ctx.Params("id")
	user := ctx.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	_, ok := claims["user_id"].(float64)

	if !ok {
		return utils.BadRequest(ctx, "Failed to read request", "user_id not found in token")
	} 

	if err := c.service.DeleteLombaByID(pubID); err != nil {
		return utils.InternalServerError(ctx, "Failed delete data", err.Error())
	}

	return utils.Success(ctx, "Succes deleted data", pubID)
}