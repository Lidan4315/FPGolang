package routes

import (
	"github.com/Caknoooo/go-gin-clean-starter/controller"
	"github.com/gin-gonic/gin"
)

func MobilRoutes(router *gin.Engine, mobilController *controller.MobilController) {
	mobilGroup := router.Group("/api/mobils")
	{
		mobilGroup.POST("/", mobilController.CreateMobil)
		mobilGroup.GET("/all", mobilController.GetAllMobil)
		mobilGroup.GET("/:id", mobilController.GetMobilById)
		mobilGroup.PUT("/", mobilController.UpdateMobil)
		mobilGroup.DELETE("/:id", mobilController.DeleteMobil)
	}
}
