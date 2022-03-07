package http

import (
	"shirt-shop/internal/user"
	"shirt-shop/pkg/middleware"

	"github.com/gofiber/fiber/v2"
)

func MapUserRoute(r fiber.Router, h user.Handler) {
	r.Get("/", middleware.Authentication(), h.GetUserById)
	r.Get("/all", middleware.Authentication(), h.GetUserAll)
	r.Post("/", h.InsertUser)
	r.Get("/role", middleware.Authentication(), h.GetRole)
	// require auth
	// r.Get("/", middleware.Authentication(), h.GetProductOne)
	// r.Get("/all", middleware.Authentication(), h.GetProductAll)
}
