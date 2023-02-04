package routes

import (
	"gin-mongo-api/controllers"

	"github.com/gin-gonic/gin"
)

func BmnRoute(router *gin.Engine) {
	router.POST("/bmn", controllers.CreateBmn())
	router.GET("/bmn/:bmnId", controllers.GetBmn())
	router.PUT("/bmn/:bmnId", controllers.EditBmn())
	router.DELETE("/bmn/:bmnId", controllers.DeleteBmn())
	router.GET("/allbmn", controllers.GetAllBmn())
}
