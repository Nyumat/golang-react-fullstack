package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// User is the model for a user
type User struct {
	Id    primitive.ObjectID `json:"id,omitempty" bson:"id,omitempty"`
	Name  string		 `json:"name,omitempty" bson:"name,omitempty" validate:"required"`
}