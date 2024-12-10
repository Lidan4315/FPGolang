package routes

import (
	"github.com/Caknoooo/go-gin-clean-starter/controller"
	"github.com/gin-gonic/gin"
)

func MobilRoutes(router *gin.Engine, mobilController *controller.MobilController) {
	mobilGroup := router.Group("/api/rentycar")
	{
		mobilGroup.POST("/create-car", mobilController.CreateMobil)
		mobilGroup.GET("/all-car", mobilController.GetAllMobil)
		mobilGroup.GET("single-car/:id", mobilController.GetMobilById)
		mobilGroup.PUT("/update/", mobilController.UpdateMobil)
		mobilGroup.DELETE("/delete-car/:id", mobilController.DeleteMobil)
	}
}
