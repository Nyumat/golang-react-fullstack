package responses

import (
	"github.com/gofiber/fiber/v2"
)

// UserResponse is the resuable struct which will be used to describe a user api response.
type UserResponse struct {
	Status int   `json:"status"`
	Msg    string `json:"msg"`
	Data   *fiber.Map `json:"data"`
}