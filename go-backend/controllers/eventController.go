package controllers

import (
	"benchmark/database"
	"benchmark/models"
	"context"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func CreateEvent(c *gin.Context) {

	var event models.Event

	err := c.BindJSON(&event)

	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	event.Timestamp = time.Now()

	collection := database.DB.Collection("events")

	result, err := collection.InsertOne(
		context.TODO(),
		event,
	)

	if err != nil {
		c.JSON(500, gin.H{
			"error": "Failed to create event",
		})
		return
	}

	c.JSON(201, gin.H{
		"message": "Event created successfully",
		"id":      result.InsertedID,
	})
}

func GetEventsByUser(c *gin.Context) {

	userId := c.Param("id")

	collection := database.DB.Collection("events")

	cursor, err := collection.Find(
		context.TODO(),
		bson.M{
			"userId": userId,
		},
	)

	if err != nil {
		c.JSON(500, gin.H{
			"error": "Failed to fetch events",
		})
		return
	}

	var events []models.Event

	err = cursor.All(context.TODO(), &events)

	if err != nil {
		c.JSON(500, gin.H{
			"error": "Failed to decode events",
		})
		return
	}

	c.JSON(200, events)
}

func GetRecentEvents(c *gin.Context) {

	collection := database.DB.Collection("events")

	findOptions := options.Find()

	findOptions.SetSort(
		bson.D{{Key: "timestamp", Value: -1}},
	)

	findOptions.SetLimit(100)

	cursor, err := collection.Find(
		context.TODO(),
		bson.M{},
		findOptions,
	)

	if err != nil {
		c.JSON(500, gin.H{
			"error": "Failed to fetch recent events",
		})
		return
	}

	var events []models.Event

	err = cursor.All(context.TODO(), &events)

	if err != nil {
		c.JSON(500, gin.H{
			"error": "Failed to decode recent events",
		})
		return
	}

	c.JSON(200, events)
}