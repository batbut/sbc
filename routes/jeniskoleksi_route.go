package routes

import (
	"gin-mongo-api/controllers"

	"github.com/gin-gonic/gin"
)

func JenisKoleksiRoute(router *gin.Engine) {
	router.POST("/jeniskoleksi", controllers.CreateJenisKoleksi())
	router.GET("/jeniskoleksi/:jeniskoleksiId", controllers.GetJenisKoleksi())
	router.PUT("/jeniskoleksi/:jeniskoleksiId", controllers.EditJenisKoleksi())
	router.DELETE("/jeniskoleksi/:jeniskoleksiId", controllers.DeleteJenisKoleksi())
	router.GET("/alljeniskoleksi", controllers.GetAllJenisKoleksi())
}
