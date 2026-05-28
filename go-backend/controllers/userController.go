package controllers

import (
	"benchmark/database"
	"benchmark/models"
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func CreateUser(c *gin.Context) {

	var user models.User

	err := c.BindJSON(&user)

	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	collection := database.DB.Collection("users")

	result, err := collection.InsertOne(
		context.TODO(),
		user,
	)

	if err != nil {
		c.JSON(500, gin.H{
			"error": "Failed to create user",
		})
		return
	}

	c.JSON(201, gin.H{
		"message": "User created successfully",
		"id":      result.InsertedID,
	})
}

func GetUsers(c *gin.Context) {

	collection := database.DB.Collection("users")

	cursor, err := collection.Find(
		context.TODO(),
		bson.M{},
	)

	if err != nil {
		c.JSON(500, gin.H{
			"error": "Failed to fetch users",
		})
		return
	}

	var users []models.User

	err = cursor.All(context.TODO(), &users)

	if err != nil {
		c.JSON(500, gin.H{
			"error": "Failed to decode users",
		})
		return
	}

	c.JSON(200, users)
}

func GetUserByID(c *gin.Context) {

	id := c.Param("id")

	objectId, err :=
		primitive.ObjectIDFromHex(id)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Invalid user id",
		})
		return
	}

	collection := database.DB.Collection("users")

	var user models.User

	err = collection.FindOne(
		context.TODO(),
		bson.M{
			"_id": objectId,
		},
	).Decode(&user)

	if err != nil {
		c.JSON(404, gin.H{
			"error": "User not found",
		})
		return
	}

	c.JSON(200, user)
}