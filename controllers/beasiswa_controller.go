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

type BeasiswaController struct {
	service services.BeasiswaService
}

func NewBeasiswaController(s services.BeasiswaService) *BeasiswaController {
	return &BeasiswaController{service: s}
}

func (c *BeasiswaController) CreateBeasiswa(ctx *fiber.Ctx) error {
	beasiswa := new(models.Beasiswa)
	user := ctx.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	if err := ctx.BodyParser(beasiswa); err != nil {
		return utils.BadRequest(ctx, "Failed to read request", err.Error())
	}

	userID, ok := claims["user_id"].(float64)

	if ok {
		beasiswa.CreatedBy = int64(userID)
	} else {
		return utils.BadRequest(ctx, "Failed to read request", "user_id not found in token")
	}

	if err := c.service.Create(beasiswa); err != nil {
		return utils.BadRequest(ctx, "Failed to saved data", err.Error())
	}

	return utils.Success(ctx, "Beasiswa was successfully created", beasiswa)
}

func (c *BeasiswaController) UpdateBeasiswa(ctx *fiber.Ctx) error {
	publicID := ctx.Params("id")
	beasiswa := new (models.Beasiswa)

	if err := ctx.BodyParser(beasiswa); err != nil {
		return utils.BadRequest(ctx, "Failed Parsing Data", err.Error())
	}

	if _, err := uuid.Parse(publicID); err != nil {
		return utils.BadRequest(ctx, "Invalid ID", err.Error())
	}

	existingBeasiswa, err := c.service.GetByPublicID(publicID)
	if err != nil {
		return utils.NotFound(ctx, "Beasiswa not found", err.Error())
	}

	beasiswa.PublicID = existingBeasiswa.PublicID
	beasiswa.InternalID = existingBeasiswa.InternalID
	beasiswa.CreatedBy = existingBeasiswa.CreatedBy

	if err := c.service.Update(beasiswa); err != nil {
		return utils.BadRequest(ctx, "Failed Update Beasiswa", err.Error())
	}

	return utils.Success(ctx, "Beasiswa Success Updated", beasiswa)
}

func (c *BeasiswaController) GetMyBeasiswaPaginate(ctx *fiber.Ctx) error {
	user := ctx.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userID := claims["user_id"].(float64)

	page, _ := strconv.Atoi(ctx.Query("page", "1"))
	limit, _ := strconv.Atoi(ctx.Query("limit", "10"))
	offset := (page - 1) * limit

	filter := ctx.Query("filter", "")
	sort := ctx.Query("sort", "")

	beasiswa, total, err := c.service.GetAllByUserPaginate(int64(userID),filter,sort,limit,offset)
	if err != nil {
		return utils.InternalServerError(ctx, "Failed get beasiswa", err.Error())
	}

	meta := utils.PaginationMeta{
		Page: page,
		Limit: limit,
		Total: int(total),
		TotalPage: int(math.Ceil(float64(total) / float64(limit))),
		Filter: filter,
		Sort: sort,
	}

	return utils.SuccessPagination(ctx, "Success get data beasiswa", beasiswa, meta)
}

func (c *BeasiswaController) DeleteBeasiswa(ctx *fiber.Ctx) error {	
	pubID := ctx.Params("id")
	user := ctx.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	_, ok := claims["user_id"].(float64)

	if !ok {
		return utils.BadRequest(ctx, "Failed to read request", "user_id not found in token")
	} 

	if err := c.service.DeleteBeasiswaByID(pubID); err != nil {
		return utils.InternalServerError(ctx, "Failed delete data", err.Error())
	}

	return utils.Success(ctx, "Succes deleted data", pubID)
}