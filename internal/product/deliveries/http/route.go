package http

import (
	"shirt-shop/internal/product"
	"shirt-shop/pkg/middleware"

	"github.com/gofiber/fiber/v2"
)

func MapProductRoute(r fiber.Router, h product.Handler) {
	r.Get("/", middleware.Authentication(), h.GetProductOne)
	r.Get("/all", middleware.Authentication(), h.GetProductAll)
	r.Post("/", middleware.Authentication(), h.InsertProduct)
	r.Get("/size", middleware.Authentication(), h.Size)
	r.Get("/gender", middleware.Authentication(), h.Gender)
	r.Get("/catagory", middleware.Authentication(), h.Catagory)
}
