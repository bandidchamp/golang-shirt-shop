package server

import (
	"fmt"
	"shirt-shop/configs"

	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"gorm.io/gorm"
)

type Server struct {
	fiber *fiber.App
	db    *gorm.DB
	cache *redis.Client
	cfg   *configs.Config
}

func Setup(db *gorm.DB, cache *redis.Client, config *configs.Config) *Server {
	return &Server{
		fiber: fiber.New(),
		db:    db,
		cache: cache,
		cfg:   config,
	}
}

func (s *Server) StartServer() error {

	// map handle here.
	s.MapHandlers(s.fiber)

	s.fiber.Use(
		cors.New(cors.Config{
			AllowOrigins:     "example.com, dev.info, bandid.info, localhost:3000",
			AllowMethods:     "GET,POST,PUT,PATCH,DELETE",
			AllowCredentials: true,
		}),
	)

	err := s.fiber.Listen(fmt.Sprintf(`%s:%s`, s.cfg.Fiber.Host, s.cfg.Fiber.Port))
	if err != nil {
		return err
	}
	return nil
}
