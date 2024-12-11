package controller

import (
	"net/http"

	"github.com/Caknoooo/go-gin-clean-starter/dto"
	"github.com/Caknoooo/go-gin-clean-starter/entity"
	"github.com/Caknoooo/go-gin-clean-starter/service"
	"github.com/gin-gonic/gin"
)

type MobilController struct {
	mobilService service.MobilService
}

func NewMobilController(mobilService service.MobilService) *MobilController {
	return &MobilController{mobilService: mobilService}
}

func (c *MobilController) CreateMobil(ctx *gin.Context) {
	var mobil entity.Mobil
	if err := ctx.ShouldBindJSON(&mobil); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newMobil, err := c.mobilService.CreateMobil(ctx, mobil)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, newMobil)
}

func (c *MobilController) GetAllMobil(ctx *gin.Context) {
	var req dto.PaginationRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	mobils, err := c.mobilService.GetAllMobil(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, mobils)
}

func (c *MobilController) GetMobilById(ctx *gin.Context) {
	id := ctx.Param("id")

	mobil, err := c.mobilService.GetMobilById(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, mobil)
}

func (c *MobilController) UpdateMobil(ctx *gin.Context) {
	var mobil entity.Mobil
	if err := ctx.ShouldBindJSON(&mobil); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedMobil, err := c.mobilService.UpdateMobil(ctx, mobil)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, updatedMobil)
}

func (c *MobilController) DeleteMobil(ctx *gin.Context) {
	id := ctx.Param("id")

	if err := c.mobilService.DeleteMobil(ctx, id); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Mobil deleted successfully"})
}
