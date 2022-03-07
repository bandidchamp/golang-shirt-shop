package product

import (
	"github.com/gofiber/fiber/v2"
)

type Handler interface {
	GetProductOne(c *fiber.Ctx) error
	GetProductAll(c *fiber.Ctx) error
	InsertProduct(c *fiber.Ctx) error
	Size(c *fiber.Ctx) error
	Gender(c *fiber.Ctx) error
	Catagory(c *fiber.Ctx) error
}
