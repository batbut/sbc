package main

import (
	"gin-mongo-api/configs"
	"gin-mongo-api/routes" //add this

	"os"

	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	configs.ConnectDB()
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "https://sbc-sebatcabut.herokuapp.com"
		},
		MaxAge: 12 * time.Hour,
	}))

	routes.InvertebrataRoute(router)      //add this
	routes.VertebrataRoute(router)        //add this
	routes.FosilRoute(router)             //add this
	routes.BatuanRoute(router)            //add this
	routes.SumberDayaGeologiRoute(router) //add this
	routes.BmnRoute(router)               //add this
	routes.LokasiTemuanRoute(router)      //add this
	routes.KoordinatRoute(router)         //add this
	routes.JenisKoleksiRoute(router)      //add this

	router.Run(":" + SetPort())
}

func SetPort() string {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "80"
	}
	return port
}
