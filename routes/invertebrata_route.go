package routes

import (
	"gin-mongo-api/controllers"
	// "gin-mongo-api/middleware"

	"github.com/gin-gonic/gin"
)

func InvertebrataRoute(router *gin.Engine) {
	// router.Use(middleware.JwtMiddleware())
	router.POST("/invertebrata", controllers.CreateInvertebrata())
	router.GET("/invertebrata/:invertebrataId", controllers.GetInvertebrata())
	router.PUT("/invertebrata/:invertebrataId", controllers.EditInvertebrata())
	router.DELETE("/invertebrata/:invertebrataId", controllers.DeleteInvertebrata())
	router.GET("/invertebratas", controllers.GetAllInvertebratas())
}
