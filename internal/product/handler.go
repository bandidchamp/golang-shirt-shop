package product

import (
	"github.com/gofiber/fiber/v2"
)

type Handler interface {
	GetProductById(c *fiber.Ctx) error
	GetProductAll(c *fiber.Ctx) error
	InsertProduct(c *fiber.Ctx) error
	Size(c *fiber.Ctx) error
	Gender(c *fiber.Ctx) error
	Catagory(c *fiber.Ctx) error
	GetProductFilter(c *fiber.Ctx) error
	UpdateProduct(c *fiber.Ctx) error
	PaddingProduct(c *fiber.Ctx) error
	AsyncFunc(c *fiber.Ctx) error
}
