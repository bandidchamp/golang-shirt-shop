package server

import (
	"fmt"
	"shirt-shop/configs"

	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
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

	err := s.fiber.Listen(fmt.Sprintf(`%s:%s`, s.cfg.Fiber.Host, s.cfg.Fiber.Port))
	if err != nil {
		return err
	}
	return nil
}
