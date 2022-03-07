package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

// app.Use(logger.New(logger.Config{
// 	Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
// }))

func MiddleWareLogger() func(*fiber.Ctx) error {
	lg := logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	})
	return lg
}
