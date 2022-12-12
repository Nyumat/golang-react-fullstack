package controllers

import (
	"backend/configs"
	"backend/models"
	"context"
	"log"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var validate = validator.New()
var userCollection *mongo.Collection = configs.GetMongoCollection(configs.DB, "users")

// POST CreateUser creates a new user
func CreateUser(c *fiber.Ctx) error {

	context, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	// New user
	var user models.User

	// Parse the body into the user model
	if err := c.BodyParser(&user); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status": http.StatusBadRequest,
			"msg":  "Invalid user data",
			"data":   &fiber.Map{
				"error": err.Error(),
			},
		})
	}

	// Validate the user model
	if err := validate.Struct(&user); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status": http.StatusBadRequest,
			"msg": "Invalid user data",
			"data":   &fiber.Map{
				"error": err.Error(),
			},
		})
	}

	userToAppend := models.User{
		Name: user.Name,
	}


	// Insert new MongoDB document to the users collection
	result, err := userCollection.InsertOne(context, userToAppend)
	if err != nil {
		log.Fatal(err)
	}

	// Return the new user if successfully created
	return c.Status(http.StatusCreated).JSON(fiber.Map{
		"status":  http.StatusCreated,
		"msg": "User created successfully",
		"data":   result,
	})
	
}

// GET Single User by ID
func GetUser(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	// Get the user's object _id from the request params
	id := c.Params("id")
	var user models.User
	defer cancel()

	// Convert the ID to an ObjectID
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status": http.StatusBadRequest,
			"msg":  "Invalid user ID",
			"data":   &fiber.Map{
				"error": err.Error(),
			},
		})
	}

	// Find the user in the database based on the '_id'/ objectID
	err = userCollection.FindOne(ctx, bson.M{"_id": objectID}).Decode(&user)
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			"status": http.StatusNotFound,
			"msg":  "User not found",
			"data":   &fiber.Map{
				"error": err.Error(),
			},
		})
	}

	// Return the user if successful
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":  http.StatusOK,
		"msg": "User found successfully",
		"data":   user,
	})
}

// GET Single User by Name
func GetUserByName(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	// Get the user's object _id from the request params
	var user models.User
	defer cancel()
	
	name := c.Params("name")
	// Find the user in the database based on the name field
	err := userCollection.FindOne(ctx, bson.M{"name": name}).Decode(&user)
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			"status": http.StatusNotFound,
			"msg":  "User not found",
			"data":   &fiber.Map{
				"error": err.Error(),
			},
		})
	}


	// Return the user if successful
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":  http.StatusOK,
		"msg": "User found successfully",
		"data":   user,
	})
}

// GET All Users
func GetUsers(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Find all users in the database
	cursor, err := userCollection.Find(ctx, bson.M{})
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status": http.StatusInternalServerError,
			"msg":  "Error getting users",
			"data":   &fiber.Map{
				"error": err.Error(),
			},
		})
	}

	// Create a slice of users
	var users []models.User

	// Iterate through the cursor and append the users to the slice
	for cursor.Next(ctx) {
		var user models.User
		cursor.Decode(&user)
		users = append(users, user)
	}

	// Return the users if successful
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":  http.StatusOK,
		"msg": "Users found successfully",
		"data":   users,
	})
}

// PUT Update User by ID
func UpdateUser(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Get the user's object _id from the request params
	id := c.Params("id")

	// Convert the ID to an ObjectID
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status": http.StatusBadRequest,
			"msg":  "Invalid user ID",
			"data":   &fiber.Map{
				"error": err.Error(),
			},
		})
	}

	// New user
	var user models.User

	// Parse the body into the user model
	if err := c.BodyParser(&user); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status": http.StatusBadRequest,
			"msg":  "Invalid user data",
			"data":   &fiber.Map{
				"error": err.Error(),
			},
		})
	}

	// Validate the user model
	if err := validate.Struct(&user); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status": http.StatusBadRequest,
			"msg": "Invalid user data",
			"data":   &fiber.Map{
				"error": err.Error(),
			},
		})
	}

	// Update the user in the database
	result, err := userCollection.UpdateOne(ctx, bson.M{"_id": objectID}, bson.M{"$set": user})
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status": http.StatusInternalServerError,
			"msg":  "Error updating user",
			"data":   &fiber.Map{
				"error": err.Error(),
			},
		})
	}

	// Signal that the user was updated
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":  http.StatusOK,
		"msg": "User updated successfully",
		"data":   result,
	})
}

// DELETE User by ID
func DeleteUser(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Get the user's object _id from the request params
	id := c.Params("id")

	// Convert the ID to an ObjectID
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status": http.StatusBadRequest,
			"msg":  "Invalid user ID",
			"data":   &fiber.Map{
				"error": err.Error(),
			},
		})
	}

	// Delete the user in the database
	result, err := userCollection.DeleteOne(ctx, bson.M{"_id": objectID})
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status": http.StatusInternalServerError,
			"msg":  "Error deleting user",
			"data":   &fiber.Map{
				"error": err.Error(),
			},
		})
	}

	// Signal that the user was deleted
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":  http.StatusOK,
		"msg": "User deleted successfully",
		"data":   result,
	})
}

