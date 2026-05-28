package database

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Database

func Connect() {

	client, err := mongo.Connect(
		context.TODO(),
		options.Client().ApplyURI(
			"mongodb://localhost:27017",
		),
	)

	if err != nil {
		log.Fatal(err)
	}

	DB = client.Database("benchmark_db")

	log.Println("MongoDB Connected")
}