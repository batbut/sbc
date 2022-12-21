package main

import (
	"gin-mongo-api/configs"
	"gin-mongo-api/routes" //add this
	"os"

	"github.com/gin-gonic/gin"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "https://resplendent-dragon-4ca5a6.netlify.app/")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "http://sbc-sebatcabut.herokuapp.com/" {
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

	router.Run(":" + SetPort())
}

func SetPort() string {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "80"
	}
	return port
}
