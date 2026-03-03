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

type ProgramKerjaController struct {
	service services.ProgramKerjaService
}

func NewProgramKerjaController(s services.ProgramKerjaService) *ProgramKerjaController {
	return &ProgramKerjaController{service: s}
}

func (c *ProgramKerjaController) CreateProker(ctx *fiber.Ctx) error {
	proker := new(models.ProgramKerja)
	user := ctx.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	if err := ctx.BodyParser(proker); err != nil {
		return utils.BadRequest(ctx, "Failed to read request", err.Error())
	}

	userID, ok := claims["user_id"].(float64)
	
	if ok {
		proker.CreatedBy = int64(userID)
	} else {
		return utils.BadRequest(ctx, "Failed to read request", "user_id not found in token")
	}

	if err := c.service.Create(proker); err != nil {
		return utils.BadRequest(ctx, "Failed to saved data", err.Error())
	}

	return utils.Success(ctx, "Program Kerja was successfully created", proker)
}

func (c *ProgramKerjaController) UpdateProker(ctx *fiber.Ctx) error {
	publicID := ctx.Params("id")
	proker := new (models.ProgramKerja)

	if err := ctx.BodyParser(proker); err != nil {
		return utils.BadRequest(ctx, "Failed Parsing Data", err.Error())
	}

	if _, err := uuid.Parse(publicID); err != nil {
		return utils.BadRequest(ctx, "Invalid ID", err.Error())
	}

	existingProker, err := c.service.GetByPublicID(publicID)
	if err != nil {
		return utils.NotFound(ctx, "Proker not found", err.Error())
	}

	proker.InternalID = existingProker.InternalID
	proker.PublicID = existingProker.PublicID
	proker.CreatedBy = existingProker.CreatedBy
	
	if err := c.service.Update(proker); err != nil {
		return utils.BadRequest(ctx, "Failed Update Proker", err.Error())
	}

	return utils.Success(ctx, "Proker Succes Updated", proker)
}

func (c *ProgramKerjaController) GetMyProkerPaginate(ctx *fiber.Ctx) error {
	user := ctx.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userID := claims["user_id"].(float64)

	page, _ := strconv.Atoi(ctx.Query("page", "1"))
	limit, _ := strconv.Atoi(ctx.Query("limit", "10"))
	offset := (page - 1) * limit

	filter := ctx.Query("filter", "")
	sort := ctx.Query("sort", "")

	proker, total, err := c.service.GetAllByUserPaginate(int64(userID),filter,sort,limit,offset)
	if err != nil {
		return utils.InternalServerError(ctx, "Failed get proker", err.Error())
	}

	meta := utils.PaginationMeta{
		Page: page,
		Limit: limit,
		Total: int(total),
		TotalPage: int(math.Ceil(float64(total) / float64(limit))),
		Filter: filter,
		Sort: sort,
	}

	return utils.SuccessPagination(ctx, "Success get data proker", proker, meta)
}

func (c *ProgramKerjaController) DeleteProgramKerja(ctx *fiber.Ctx) error {	
	pubID := ctx.Params("id")
	user := ctx.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	_, ok := claims["user_id"].(float64)

	if !ok {
		return utils.BadRequest(ctx, "Failed to read request", "user_id not found in token")
	} 

	if err := c.service.DeleteProgramKerjaByID(pubID); err != nil {
		return utils.InternalServerError(ctx, "Failed delete data", err.Error())
	}

	return utils.Success(ctx, "Succes deleted data", pubID)
}