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

// CreateUser creates a new user
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
		Id: primitive.NewObjectID(),
		Name: user.Name,
	}


	// Create a new MongoDB document
	result, err := userCollection.InsertOne(context, userToAppend)
	if err != nil {
		log.Fatal(err)
	}

	// Return the new user if successful
	return c.Status(http.StatusCreated).JSON(fiber.Map{
		"status":  http.StatusCreated,
		"msg": "User created successfully",
		"data":   result,
	})
	
}

// Get Single User by ID
func GetUser(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	// Get the ID from the URL
	id := c.Params("Id")
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

	// Find the user in the database
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


