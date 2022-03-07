package http

import (
	"shirt-shop/internal/authen"
	"shirt-shop/pkg/middleware"

	"github.com/gofiber/fiber/v2"
)

func MapAuthRoute(r fiber.Router, h authen.Handler) {
	r.Post("/", h.Login)
	r.Get("/", middleware.Authentication(), h.HandleGetUserFromToken)
}
