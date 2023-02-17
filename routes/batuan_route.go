package routes

import (
	"gin-mongo-api/controllers"

	"github.com/gin-gonic/gin"
)

func BatuanRoute(router *gin.Engine) {
	router.POST("/batuan", controllers.CreateBatuan())
	router.GET("/batuan/:batuanId", controllers.GetBatuan())
	router.PUT("/batuan/:batuanId", controllers.EditBatuan())
	router.DELETE("/batuan/:batuanId", controllers.DeleteBatuan())
	router.GET("/batuans", controllers.GetAllBatuans())
	router.GET("/batuans/export", controllers.ExportBatuanToExcel())
}
