package main

import (
	"gin-mongo-api/configs"
	"gin-mongo-api/routes" //add this
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "https://resplendent-dragon-4ca5a6.netlify.app")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "https://sbc-sebatcabut.herokuapp.com" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func main() {
	router := gin.Default()

	//run database
	configs.ConnectDB()

	//routes
	routes.UserRoute(router)              //add this
	routes.InvertebrataRoute(router)      //add this
	routes.VertebrataRoute(router)        //add this
	routes.FosilRoute(router)             //add this
	routes.BatuanRoute(router)            //add this
	routes.SumberDayaGeologiRoute(router) //add this
	routes.LokasiTemuanRoute(router)      //add this
	routes.KoordinatRoute(router)         //add this

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"https://resplendent-dragon-4ca5a6.netlify.app"},
		AllowMethods:     []string{"PUT", "PATCH", "POST", "OPTIONS", "GET", "DELETE"},
		AllowHeaders:     []string{"Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization", "Accept", "Origin", "Cache-Control", "X-Requested-With"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "https://sbc-sebatcabut.herokuapp.com"
		},
		MaxAge: 12 * time.Hour,
	}))

	router.Run(":" + SetPort())
}

func SetPort() string {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "80"
	}
	return port
}
