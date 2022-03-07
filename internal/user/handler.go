package user

import "github.com/gofiber/fiber/v2"

type Handler interface {
	GetUserById(c *fiber.Ctx) error
	GetUserAll(c *fiber.Ctx) error
	InsertUser(c *fiber.Ctx) error
	UpdateUser(c *fiber.Ctx) error
	DeleteUser(c *fiber.Ctx) error
	GetRole(c *fiber.Ctx) error
}
