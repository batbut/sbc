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

var storageCollection *mongo.Collection = configs.GetCollection(configs.DB, "Lokasi_Penyimpanan")
var validate_storage = validator.New()

func CreateStorage() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var storage models.Storage
		defer cancel()

		//validate the request body
		if err := c.BindJSON(&storage); err != nil {
			c.JSON(http.StatusBadRequest, responses.StorageResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		//use the validator library to validate required fields
		if validationErr := validate_storage.Struct(&storage); validationErr != nil {
			c.JSON(http.StatusBadRequest, responses.StorageResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
			return
		}

		newStorage := models.Storage{
			Id: primitive.NewObjectID(),
			Storage: struct {
				RuangStorage string "bson:\"Ruang_Storage\" json:\"ruang_storage\" validate:\"required\""
				Lantai       string "bson:\"Lantai\" json:\"lantai\" validate:\"required\""
				Lajur        string "bson:\"Lajur\" json:\"lajur\" validate:\"required\""
				Lemari       string "bson:\"Lemari\" json:\"lemari\" validate:\"required\""
				Laci         string "bson:\"Laci\" json:\"laci\" validate:\"required\""
				Slot         string "bson:\"Slot\" json:\"slot\" validate:\"required\""
			}{
				RuangStorage: storage.Storage.RuangStorage,
				Lantai:       storage.Storage.RuangStorage,
				Lajur:        storage.Storage.Lajur,
				Lemari:       storage.Storage.Lemari,
				Laci:         storage.Storage.Laci,
				Slot:         storage.Storage.Slot,
			},
			NonStorage: struct {
				NamaNonStorage string "bson:\"Nama_Non_Storage\" json:\"nama_non_storage\" validate:\"required\""
			}{
				NamaNonStorage: storage.NonStorage.NamaNonStorage,
			},
		}

		result, err := storageCollection.InsertOne(ctx, newStorage)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.StorageResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusCreated, responses.StorageResponse{Status: http.StatusCreated, Message: "success", Data: map[string]interface{}{"data": result}})
	}
}

func GetStorage() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		storageId := c.Param("storageId")
		var storage models.Storage
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(storageId)

		err := storageCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&storage)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.StorageResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusOK, responses.StorageResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": storage}})
	}
}

func EditStorage() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		storageId := c.Param("storageId")
		var storage models.Storage
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(storageId)

		//validate the request body
		if err := c.BindJSON(&storage); err != nil {
			c.JSON(http.StatusBadRequest, responses.StorageResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		//use the validator library to validate required fields
		if validationErr := validate_storage.Struct(&storage); validationErr != nil {
			c.JSON(http.StatusBadRequest, responses.StorageResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
			return
		}

		update := bson.M{
			"Storage": bson.M{
				"Ruang_Storage": storage.Storage.RuangStorage,
				"Lantai":        storage.Storage.Lantai,
				"Lajur":         storage.Storage.Lajur,
				"Lemari":        storage.Storage.Lemari,
				"Laci":          storage.Storage.Laci,
				"Slot":          storage.Storage.Slot,
			},
			"Non_Storage": bson.M{
				"Nama_Non_Storage": storage.NonStorage.NamaNonStorage,
			},
		}
		result, err := storageCollection.UpdateOne(ctx, bson.M{"_id": objId}, bson.M{"$set": update})

		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.StorageResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		//get updated invertebrata details
		var updatedStorage models.Storage
		if result.MatchedCount == 1 {
			err := storageCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&updatedStorage)
			if err != nil {
				c.JSON(http.StatusInternalServerError, responses.StorageResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
				return
			}
		}

		c.JSON(http.StatusOK, responses.StorageResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": updatedStorage}})
	}
}

func DeleteStorage() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		storageId := c.Param("storageId")
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(storageId)

		result, err := storageCollection.DeleteOne(ctx, bson.M{"_id": objId})

		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.StorageResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		if result.DeletedCount < 1 {
			c.JSON(http.StatusNotFound,
				responses.StorageResponse{Status: http.StatusNotFound, Message: "error", Data: map[string]interface{}{"data": "Storage with specified ID not found!"}},
			)
			return
		}

		c.JSON(http.StatusOK,
			responses.StorageResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": "Storage successfully deleted!"}},
		)
	}
}

func GetAllStorage() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var allstorage []models.Storage
		defer cancel()

		results, err := storageCollection.Find(ctx, bson.M{})

		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.StorageResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		//reading from the db in an optimal way
		defer results.Close(ctx)
		for results.Next(ctx) {
			var singleStorage models.Storage
			if err = results.Decode(&singleStorage); err != nil {
				c.JSON(http.StatusInternalServerError, responses.StorageResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			}

			allstorage = append(allstorage, singleStorage)
		}

		c.JSON(http.StatusOK,
			responses.StorageResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": allstorage}},
		)
	}
}
