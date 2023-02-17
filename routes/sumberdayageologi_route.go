package routes

import (
	"gin-mongo-api/controllers"

	"github.com/gin-gonic/gin"
)

func SumberDayaGeologiRoute(router *gin.Engine) {
	router.POST("/sumberdayageologi", controllers.CreateSumberDayaGeologi())
	router.GET("/sumberdayageologi/:sumberdayageologiId", controllers.GetSumberDayaGeologi())
	router.PUT("/sumberdayageologi/:sumberdayageologiId", controllers.EditSumberDayaGeologi())
	router.DELETE("/sumberdayageologi/:sumberdayageologiId", controllers.DeleteSumberDayaGeologi())
	router.GET("/sumberdayageologis", controllers.GetAllSumberDayaGeologis())
	router.GET("/sumberdayageologis/export", controllers.ExportSumberDayaGeologiToExcel())
}
