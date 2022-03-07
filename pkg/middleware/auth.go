package middleware

// package middleware

import (
	"os"
	"shirt-shop/interfaces"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v2"
)

// Authentication protect routes
func Authentication() func(*fiber.Ctx) error {
	return jwtware.New(jwtware.Config{
		// SigningKey: []byte("shirt-shop-secret"),
		SigningKey:   []byte(os.Getenv("APP_JWT_SECRET")),
		ErrorHandler: jwtError,
	})
}

func jwtError(c *fiber.Ctx, err error) error {
	if err.Error() == "Missing or malformed JWT" {
		var payload interfaces.ResponsePayload
		payload = interfaces.ResponsePayload{
			Code:    fiber.StatusBadRequest,
			Status:  false,
			Message: "Missing or malformed JWT",
			Payload: nil,
		}
		return c.Status(fiber.StatusBadRequest).JSON(payload)
	} else {
		var payload interfaces.ResponsePayload
		payload = interfaces.ResponsePayload{
			Code:    fiber.StatusUnauthorized,
			Status:  false,
			Message: "Invalid or expired JWT",
			Payload: nil,
		}
		return c.Status(fiber.StatusUnauthorized).JSON(payload)

	}
}
