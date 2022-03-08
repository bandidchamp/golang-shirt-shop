package http

import (
	"shirt-shop/internal/product"
	"shirt-shop/pkg/middleware"

	"github.com/gofiber/fiber/v2"
)

func MapProductRoute(r fiber.Router, h product.Handler) {
	r.Get("/", middleware.Authentication(), h.GetProductById)
	r.Get("/all", middleware.Authentication(), h.GetProductAll)
	r.Get("/findby", middleware.Authentication(), h.GetProductFilter)

	r.Put("/", middleware.Authentication(), h.InsertProduct)
	r.Patch("/", middleware.Authentication(), h.UpdateProduct)
	r.Delete("/padding", middleware.Authentication(), h.PaddingProduct)

	r.Get("/size", middleware.Authentication(), h.Size)
	r.Get("/gender", middleware.Authentication(), h.Gender)
	r.Get("/catagory", middleware.Authentication(), h.Catagory)
}
