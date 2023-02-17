package routes

import (
	"gin-mongo-api/controllers"

	"github.com/gin-gonic/gin"
)

func FosilRoute(router *gin.Engine) {
	router.POST("/fosil", controllers.CreateFosil())
	router.GET("/fosil/:fosilId", controllers.GetFosil())
	router.PUT("/fosil/:fosilId", controllers.EditFosil())
	router.DELETE("/fosil/:fosilId", controllers.DeleteFosil())
	router.GET("/fosils", controllers.GetAllFosils())
	router.GET("/fosils/export", controllers.ExportBatuanToExcel())
}
