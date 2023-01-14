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

var invertebrataCollection *mongo.Collection = configs.GetCollection(configs.DB, "invertebrata")
var validate_invertebrata = validator.New()

func CreateInvertebrata() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var invertebrata models.Invertebrata
		defer cancel()

		//validate the request body
		if err := c.BindJSON(&invertebrata); err != nil {
			c.JSON(http.StatusBadRequest, responses.InvertebrataResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		//use the validator library to validate required fields
		if validationErr := validate_invertebrata.Struct(&invertebrata); validationErr != nil {
			c.JSON(http.StatusBadRequest, responses.InvertebrataResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
			return
		}

		newInvertebrata := models.Invertebrata{
			Id:               primitive.NewObjectID(),
			Nama:             invertebrata.Nama,
			Lokasi_Ditemukan: invertebrata.Lokasi_Ditemukan,
			Waktu_Ditemukan:  invertebrata.Waktu_Ditemukan,
		}

		result, err := invertebrataCollection.InsertOne(ctx, newInvertebrata)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.InvertebrataResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusCreated, responses.InvertebrataResponse{Status: http.StatusCreated, Message: "success", Data: map[string]interface{}{"data": result}})
	}
}

func GetInvertebrata() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		invertebrataId := c.Param("invertebrataId")
		var invertebrata models.Invertebrata
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(invertebrataId)

		err := invertebrataCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&invertebrata)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.InvertebrataResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusOK, responses.InvertebrataResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": invertebrata}})
	}
}

func EditInvertebrata() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		invertebrataId := c.Param("invertebrataId")
		var invertebrata models.Invertebrata
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(invertebrataId)

		//validate the request body
		if err := c.BindJSON(&invertebrata); err != nil {
			c.JSON(http.StatusBadRequest, responses.InvertebrataResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		//use the validator library to validate required fields
		if validationErr := validate_invertebrata.Struct(&invertebrata); validationErr != nil {
			c.JSON(http.StatusBadRequest, responses.InvertebrataResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
			return
		}

		update := bson.M{"nama": invertebrata.Nama, "lokasi_ditemukan": invertebrata.Lokasi_Ditemukan, "waktu_ditemukan": invertebrata.Waktu_Ditemukan}
		result, err := invertebrataCollection.UpdateOne(ctx, bson.M{"_id": objId}, bson.M{"$set": update})

		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.InvertebrataResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		//get updated invertebrata details
		var updatedInvertebrata models.Invertebrata
		if result.MatchedCount == 1 {
			err := invertebrataCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&updatedInvertebrata)
			if err != nil {
				c.JSON(http.StatusInternalServerError, responses.InvertebrataResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
				return
			}
		}

		c.JSON(http.StatusOK, responses.InvertebrataResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": updatedInvertebrata}})
	}
}

func DeleteInvertebrata() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		invertebrataId := c.Param("invertebrataId")
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(invertebrataId)

		result, err := invertebrataCollection.DeleteOne(ctx, bson.M{"_id": objId})

		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.InvertebrataResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		if result.DeletedCount < 1 {
			c.JSON(http.StatusNotFound,
				responses.InvertebrataResponse{Status: http.StatusNotFound, Message: "error", Data: map[string]interface{}{"data": "Invertebrata with specified ID not found!"}},
			)
			return
		}

		c.JSON(http.StatusOK,
			responses.InvertebrataResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": "Invertebrata successfully deleted!"}},
		)
	}
}

func GetAllInvertebratas() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var invertebratas []models.Invertebrata
		defer cancel()

		results, err := invertebrataCollection.Find(ctx, bson.M{})

		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.InvertebrataResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		//reading from the db in an optimal way
		defer results.Close(ctx)
		for results.Next(ctx) {
			var singleInvertebrata models.Invertebrata
			if err = results.Decode(&singleInvertebrata); err != nil {
				c.JSON(http.StatusInternalServerError, responses.InvertebrataResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			}

			invertebratas = append(invertebratas, singleInvertebrata)
		}

		c.JSON(http.StatusOK,
			responses.InvertebrataResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": invertebratas}},
		)
	}
}
