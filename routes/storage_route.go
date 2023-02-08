package routes

import (
	"gin-mongo-api/controllers"

	"github.com/gin-gonic/gin"
)

func StorageRoute(router *gin.Engine) {
	router.POST("/storage", controllers.CreateStorage())
	router.GET("/storage/:storageId", controllers.GetStorage())
	router.PUT("/storage/:storageId", controllers.EditStorage())
	router.DELETE("/storage/:storageId", controllers.DeleteStorage())
	router.GET("/allstorage", controllers.GetAllStorage())
}
