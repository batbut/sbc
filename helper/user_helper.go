package helper

import (
	"context"
	"fmt"
	"gin-mongo-api/configs"
	"gin-mongo-api/models"
	"time"

	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

func SaveUser(user models.User) error {
	// hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	// set hashed password to user struct
	user.Password = string(hashedPassword)
	// insert user to mongodb
	collection := configs.GetCollection(configs.DB, "user")
	_, err = collection.InsertOne(context.TODO(), user)
	return err
}

func GetUserByUsername(username string) (*models.User, error) {
	var user models.User
	collection := configs.GetCollection(configs.DB, "user")
	err := collection.FindOne(context.TODO(), bson.M{"username": username}).Decode(&user)
	if err == mongo.ErrNoDocuments {
		return nil, nil
	}
	return &user, err
}

func GetUserByUserID(userID string) (*models.User, error) {
	var user models.User
	collection := configs.GetCollection(configs.DB, "user")
	id, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return nil, err
	}
	err = collection.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&user)

	if err == mongo.ErrNoDocuments {
		return nil, nil
	}
	return &user, err
}

func ComparePassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func CreateJWT(userID, role string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = userID
	claims["role"] = role
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
	tokenString, err := token.SignedString([]byte(configs.JwtSecret))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func GenerateToken(userID string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = userID
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
	tokenString, err := token.SignedString([]byte(configs.JwtSecret))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func SaveToken(userID, token string) error {
	collection := configs.GetCollection(configs.DB, "user")
	_, err := collection.UpdateOne(context.TODO(), bson.M{"username": userID}, bson.M{"$set": bson.M{"token": token}})
	if err != nil {
		return err
	}
	return nil
}

func GetTokenByUserID(userID string) (string, error) {
	collection := configs.GetCollection(configs.DB, "user")
	var result struct {
		Token string
	}
	err := collection.FindOne(context.TODO(), bson.M{"username": userID}).Decode(&result)
	if err != nil {
		return "", err
	}
	if result.Token == "" {
		return "", fmt.Errorf("token not found")
	}
	return result.Token, nil
}

func DeleteToken(userID string) error {
	collection := configs.GetCollection(configs.DB, "user")
	_, err := collection.UpdateOne(context.TODO(), bson.M{"_id": userID}, bson.M{"$unset": bson.M{"token": ""}})
	if err != nil {
		return err
	}
	return nil
}
