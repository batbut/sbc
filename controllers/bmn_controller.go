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

var bmnCollection *mongo.Collection = configs.GetCollection(configs.DB, "bmn")
var validate_bmn = validator.New()

func CreateBmn() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var bmn models.Bmn
		defer cancel()

		//validate the request body
		if err := c.BindJSON(&bmn); err != nil {
			c.JSON(http.StatusBadRequest, responses.BmnResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		//use the validator library to validate required fields
		if validationErr := validate_bmn.Struct(&bmn); validationErr != nil {
			c.JSON(http.StatusBadRequest, responses.BmnResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
			return
		}

		newBmn := models.Bmn{
			Id:             primitive.NewObjectID(),
			NoRegister:     bmn.NoRegister,
			KategoriBMN:    bmn.KategoriBMN,
			TipeBMN:        bmn.TipeBMN,
			NilaiPerolehan: bmn.NilaiPerolehan,
			NilaiBuku:      bmn.NilaiBuku,
		}

		result, err := bmnCollection.InsertOne(ctx, newBmn)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.BmnResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusCreated, responses.BmnResponse{Status: http.StatusCreated, Message: "success", Data: map[string]interface{}{"data": result}})
	}
}

func GetBmn() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		bmnId := c.Param("bmnId")
		var bmn models.Bmn
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(bmnId)

		err := bmnCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&bmn)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.BmnResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusOK, responses.BmnResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": bmn}})
	}
}

func EditBmn() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		bmnId := c.Param("bmnId")
		var bmn models.Bmn
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(bmnId)

		//validate the request body
		if err := c.BindJSON(&bmn); err != nil {
			c.JSON(http.StatusBadRequest, responses.BmnResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		//use the validator library to validate required fields
		if validationErr := validate_bmn.Struct(&bmn); validationErr != nil {
			c.JSON(http.StatusBadRequest, responses.BmnResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
			return
		}

		update := bson.M{"No Register": bmn.NoRegister, "Kategori BMN": bmn.KategoriBMN, "Tipe BMN": bmn.TipeBMN, "Nilai Perolehan": bmn.NilaiPerolehan, "Nilai Buku": bmn.NilaiBuku}
		result, err := bmnCollection.UpdateOne(ctx, bson.M{"_id": objId}, bson.M{"$set": update})

		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.BmnResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		//get updated invertebrata details
		var updatedBmn models.Bmn
		if result.MatchedCount == 1 {
			err := bmnCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&updatedBmn)
			if err != nil {
				c.JSON(http.StatusInternalServerError, responses.BmnResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
				return
			}
		}

		c.JSON(http.StatusOK, responses.BmnResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": updatedBmn}})
	}
}

func DeleteBmn() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		bmnId := c.Param("bmnId")
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(bmnId)

		result, err := bmnCollection.DeleteOne(ctx, bson.M{"_id": objId})

		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.BmnResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		if result.DeletedCount < 1 {
			c.JSON(http.StatusNotFound,
				responses.BmnResponse{Status: http.StatusNotFound, Message: "error", Data: map[string]interface{}{"data": "BMN with specified ID not found!"}},
			)
			return
		}

		c.JSON(http.StatusOK,
			responses.BmnResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": "BMN successfully deleted!"}},
		)
	}
}

func GetAllBmn() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var allbmn []models.Bmn
		defer cancel()

		results, err := bmnCollection.Find(ctx, bson.M{})

		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.BmnResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		//reading from the db in an optimal way
		defer results.Close(ctx)
		for results.Next(ctx) {
			var singleBmn models.Bmn
			if err = results.Decode(&singleBmn); err != nil {
				c.JSON(http.StatusInternalServerError, responses.BmnResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			}

			allbmn = append(allbmn, singleBmn)
		}

		c.JSON(http.StatusOK,
			responses.BmnResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": allbmn}},
		)
	}
}
