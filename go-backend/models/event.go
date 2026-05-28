package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Event struct {
	ID        primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	UserID    string             `json:"userId" bson:"userId"`
	Type      string             `json:"type" bson:"type"`
	Timestamp time.Time          `json:"timestamp" bson:"timestamp"`
}