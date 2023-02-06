package routes

import (
	"gin-mongo-api/controllers"
	// "gin-mongo-api/middleware"

	"github.com/gin-gonic/gin"
)

func UserRoute(router *gin.Engine) {
	// Register route tidak perlu autentikasi
	router.POST("/register", controllers.Register)

	// Login dan logout memerlukan autentikasi
	// router.Group("/").Use(middleware.JwtMiddleware())
	router.POST("/login", controllers.Login)
	router.POST("/logout", controllers.Logout)
	router.GET("/users", controllers.GetAllUsers())
}
