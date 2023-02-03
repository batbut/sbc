package controllers

import (
	"context"
	"gin-mongo-api/configs"
	"gin-mongo-api/models"
	"gin-mongo-api/responses"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var jeniskoleksiCollection *mongo.Collection = configs.GetCollection(configs.DB, "Jenis_Koleksi")
var validate_jeniskoleksi = validator.New()

func CreateJenisKoleksi() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var jeniskoleksi models.JenisKoleksi
		defer cancel()

		//validate the request body
		if err := c.BindJSON(&jeniskoleksi); err != nil {
			c.JSON(http.StatusBadRequest, responses.JenisKoleksiResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		//use the validator library to validate required fields
		if validationErr := validate_jeniskoleksi.Struct(&jeniskoleksi); validationErr != nil {
			c.JSON(http.StatusBadRequest, responses.JenisKoleksiResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
			return
		}

		newJenisKoleksi := models.JenisKoleksi{
			Id:          primitive.NewObjectID(),
			NamaKoleksi: jeniskoleksi.NamaKoleksi,
			JenisKoleksiFosil: struct {
				JenisKoleksi     string "bson:\"Jenis_Koleksi\" json:\"jenis_koleksi\" validate:\"required\""
				SubJenisKoleksi  string "bson:\"Sub_Jenis_Koleksi\" json:\"sub_jenis_koleksi\" validate:\"required\""
				KodeJenisKoleksi string "bson:\"Kode_Jenis_Koleksi\" json:\"kode_jenis_koleksi\" validate:\"required\""
			}{
				JenisKoleksi:     jeniskoleksi.JenisKoleksiFosil.JenisKoleksi,
				SubJenisKoleksi:  jeniskoleksi.JenisKoleksiFosil.SubJenisKoleksi,
				KodeJenisKoleksi: jeniskoleksi.JenisKoleksiFosil.KodeJenisKoleksi,
			},
			JenisKoleksiBatuan: struct {
				JenisKoleksi     string "bson:\"Jenis_Koleksi\" json:\"jenis_koleksi\" validate:\"required\""
				SubJenisKoleksi  string "bson:\"Sub_Jenis_Koleksi\" json:\"sub_jenis_koleksi\" validate:\"required\""
				KodeJenisKoleksi string "bson:\"Kode_Jenis_Koleksi\" json:\"kode_jenis_koleksi\" validate:\"required\""
			}{
				JenisKoleksi:     jeniskoleksi.JenisKoleksiBatuan.JenisKoleksi,
				SubJenisKoleksi:  jeniskoleksi.JenisKoleksiBatuan.SubJenisKoleksi,
				KodeJenisKoleksi: jeniskoleksi.JenisKoleksiBatuan.KodeJenisKoleksi,
			},
		}

		result, err := jeniskoleksiCollection.InsertOne(ctx, newJenisKoleksi)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.JenisKoleksiResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusCreated, responses.JenisKoleksiResponse{Status: http.StatusCreated, Message: "success", Data: map[string]interface{}{"data": result}})
	}
}

func GetJenisKoleksi() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		jeniskoleksiId := c.Param("jeniskoleksiId")
		var jeniskoleksi models.JenisKoleksi
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(jeniskoleksiId)

		err := jeniskoleksiCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&jeniskoleksi)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.JenisKoleksiResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusOK, responses.JenisKoleksiResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": jeniskoleksi}})
	}
}

func EditJenisKoleksi() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		jeniskoleksiId := c.Param("jeniskoleksiId")
		var jeniskoleksi models.JenisKoleksi
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(jeniskoleksiId)

		//validate the request body
		if err := c.BindJSON(&jeniskoleksi); err != nil {
			c.JSON(http.StatusBadRequest, responses.JenisKoleksiResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		//use the validator library to validate required fields
		if validationErr := validate_jeniskoleksi.Struct(&jeniskoleksi); validationErr != nil {
			c.JSON(http.StatusBadRequest, responses.JenisKoleksiResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
			return
		}

		update := bson.M{
			"Nama_Koleksi": jeniskoleksi.NamaKoleksi,
			"Jenis_Koleksi_Fosil": bson.M{
				"Jenis_Koleksi":      jeniskoleksi.JenisKoleksiFosil.JenisKoleksi,
				"Sub_Jenis_Koleksi":  jeniskoleksi.JenisKoleksiFosil.SubJenisKoleksi,
				"Kode_Jenis_Koleksi": jeniskoleksi.JenisKoleksiFosil.KodeJenisKoleksi,
			},
			"Jenis_Koleksi_Batuan": bson.M{
				"Jenis_Koleksi":      jeniskoleksi.JenisKoleksiBatuan.JenisKoleksi,
				"Sub_Jenis_Koleksi":  jeniskoleksi.JenisKoleksiBatuan.SubJenisKoleksi,
				"Kode_Jenis_Koleksi": jeniskoleksi.JenisKoleksiBatuan.KodeJenisKoleksi,
			},
		}
		result, err := jeniskoleksiCollection.UpdateOne(ctx, bson.M{"_id": objId}, bson.M{"$set": update})

		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.JenisKoleksiResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		//get updated invertebrata details
		var updatedJenisKoleksi models.JenisKoleksi
		if result.MatchedCount == 1 {
			err := jeniskoleksiCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&updatedJenisKoleksi)
			if err != nil {
				c.JSON(http.StatusInternalServerError, responses.JenisKoleksiResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
				return
			}
		}

		c.JSON(http.StatusOK, responses.JenisKoleksiResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": updatedJenisKoleksi}})
	}
}

func DeleteJenisKoleksi() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		jeniskoleksiId := c.Param("jeniskoleksiId")
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(jeniskoleksiId)

		result, err := jeniskoleksiCollection.DeleteOne(ctx, bson.M{"_id": objId})

		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.JenisKoleksiResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		if result.DeletedCount < 1 {
			c.JSON(http.StatusNotFound,
				responses.JenisKoleksiResponse{Status: http.StatusNotFound, Message: "error", Data: map[string]interface{}{"data": "Jenis Koleksi with specified ID not found!"}},
			)
			return
		}

		c.JSON(http.StatusOK,
			responses.JenisKoleksiResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": "Jenis Koleksi successfully deleted!"}},
		)
	}
}

func GetAllJenisKoleksi() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var alljeniskoleksi []models.JenisKoleksi
		defer cancel()

		results, err := jeniskoleksiCollection.Find(ctx, bson.M{})

		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.JenisKoleksiResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		//reading from the db in an optimal way
		defer results.Close(ctx)
		for results.Next(ctx) {
			var singleJenisKoleksi models.JenisKoleksi
			if err = results.Decode(&singleJenisKoleksi); err != nil {
				c.JSON(http.StatusInternalServerError, responses.JenisKoleksiResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			}

			alljeniskoleksi = append(alljeniskoleksi, singleJenisKoleksi)
		}

		c.JSON(http.StatusOK,
			responses.JenisKoleksiResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": alljeniskoleksi}},
		)
	}
}
