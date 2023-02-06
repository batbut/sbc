package controllers

import (
	"context"
	"fmt"
	"gin-mongo-api/configs"
	"gin-mongo-api/helper"
	"gin-mongo-api/models"
	"gin-mongo-api/responses"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *gin.Context) {
	// Ambil input dari client
	var input models.User
	if err := c.ShouldBindJSON(&input); err != nil {
		responses.Error(http.StatusBadRequest, "Invalid input", c)
		c.Abort()
		return
	}
	// Validasi input
	if input.Name == "" || input.Username == "" || input.Password == "" || input.Role == "" {
		responses.Error(http.StatusBadRequest, "Name, username, password and role are required", c)
		c.Abort()
		return
	}

	// Cek apakah username sudah digunakan atau belum
	existingUser, err := helper.GetUserByUsername(input.Username)
	if err != nil {
		responses.Error(http.StatusInternalServerError, "Failed to check existing username", c)
		c.Abort()
		return
	}
	if existingUser != nil {
		responses.Error(http.StatusBadRequest, "Username already exists", c)
		c.Abort()
		return
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		responses.Error(http.StatusInternalServerError, "Failed to hash password", c)
		c.Abort()
		return
	}

	// Simpan user baru ke dalam database
	user := models.User{
		Name:     input.Name,
		Username: input.Username,
		Password: string(hashedPassword),
		Role:     input.Role,
	}
	collection := configs.GetCollection(configs.DB, "user")
	_, err = collection.InsertOne(context.TODO(), user)
	if err != nil {
		responses.Error(http.StatusInternalServerError, "Failed to insert new user", c)
		c.Abort()
		return
	}

	// Hapus password dari response
	user.Password = ""

	// Tampilkan respon "Register Success"
	responses.Success(http.StatusCreated, "Register Success", c)
}

func Login(c *gin.Context) {
	// Ambil input dari client
	var input models.User
	if err := c.ShouldBindJSON(&input); err != nil {
		responses.Error(http.StatusBadRequest, "Invalid input", c)
		c.Abort()
		return
	}
	// Validasi input
	if input.Username == "" || input.Password == "" {
		responses.Error(http.StatusBadRequest, "Username and Password are required", c)
		c.Abort()
		return
	}

	// Cek apakah username ada di database
	user, err := helper.GetUserByUsername(input.Username)
	if err != nil {
		responses.Error(http.StatusInternalServerError, "Failed to check existing username", c)
		c.Abort()
		return
	}
	if user == nil {
		responses.Error(http.StatusBadRequest, "Username not found", c)
		c.Abort()
		return
	}

	// Cek apakah password cocok dengan password yang ada di database
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
	if err != nil {
		responses.Error(http.StatusBadRequest, "Invalid password", c)
		c.Abort()
		return
	}

	// Setelah user berhasil login, keluarkan token baru
	token, err := helper.GenerateToken(user.Id.Hex())
	if err != nil {
		responses.Error(http.StatusInternalServerError, "Failed to generate token", c)
		c.Abort()
		return
	}

	// Simpan token baru ke dalam database
	err = helper.SaveToken(user.Username, token)
	if err != nil {
		responses.Error(http.StatusInternalServerError, "Failed to save token", c)
		c.Abort()
		return
	}

	// Kirim token ke client sebagai respon
	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

var userCollection *mongo.Collection = configs.GetCollection(configs.DB, "user")

func GetAllUsers() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var users []models.User
		defer cancel()

		results, err := userCollection.Find(ctx, bson.M{})

		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		//reading from the db in an optimal way
		defer results.Close(ctx)
		for results.Next(ctx) {
			var singleUser models.User
			if err = results.Decode(&singleUser); err != nil {
				c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			}

			users = append(users, singleUser)
		}

		c.JSON(http.StatusOK,
			responses.UserResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": users}},
		)
	}
}

func Logout(c *gin.Context) {
	var err error

	// Ambil token dari header
	tokenString := c.GetHeader("Authorization")
	if tokenString == "" {
		responses.Error(http.StatusBadRequest, "Token is required", c)
		c.Abort()
		return
	}

	// validasi token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method")
		}
		return []byte(configs.JwtSecret), nil
	})
	if err != nil || !token.Valid {
		responses.Error(http.StatusBadRequest, "Invalid token", c)
		c.Abort()
		return
	}

	// ambil userID dari token
	claims := token.Claims.(jwt.MapClaims)
	userID, ok := claims["user_id"].(string)
	if !ok {
		responses.Error(http.StatusBadRequest, "Invalid Token", c)
		c.Abort()
		return
	}
	// ambil user dari database berdasarkan userID
	user, err := helper.GetUserByUserID(userID)
	if err != nil {
		responses.Error(http.StatusInternalServerError, "Failed to get user", c)
		c.Abort()
		return
	}
	if user == nil {
		responses.Error(http.StatusBadRequest, "Invalid Token", c)
		c.Abort()
		return
	}

	// hapus token dari database
	err = helper.DeleteToken(userID)
	if err != nil {
		responses.Error(http.StatusInternalServerError, "Failed to delete token", c)
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully logged out",
	})
}
