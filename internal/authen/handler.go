package authen

import "github.com/gofiber/fiber/v2"

type Handler interface {
	Login(c *fiber.Ctx) error
	HandleGetUserFromToken(c *fiber.Ctx) error
}
