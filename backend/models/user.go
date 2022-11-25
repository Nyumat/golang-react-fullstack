package models

// Model for a user in the database
type User struct {
	Name  string		 `json:"name,omitempty" bson:"name,omitempty" validate:"required"`
}