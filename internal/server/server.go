package server

import (
	"fmt"
	"shirt-shop/configs"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Server struct {
	fiber *fiber.App
	db    *gorm.DB
	cfg   *configs.Config
}

func Setup(db *gorm.DB, config *configs.Config) *Server {
	return &Server{
		fiber: fiber.New(),
		db:    db,
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

// func (s *Server) Start() error {
// 	if err := s.MapHandlers(s.fiber); err != nil {
// 		return fmt.Errorf("failed to map handlers with error: %w", err)
// 	}

// 	if s.cfg.StateStatus == "dev" {
// 		if err := StartServer(s.fiber, s.cfg); err != nil {
// 			return fmt.Errorf("failed to start server with error: %w", err)
// 		}
// 	} else {
// 		if err := StartServerWithGracefulShutdown(s.fiber, s.cfg); err != nil {
// 			return fmt.Errorf("failed to start server with graceful shutdown with error: %w", err)
// 		}
// 	}

// 	return nil
// }

// func StartServerWithGracefulShutdown(a *fiber.App, c *configs.Config) error {
// 	//* Create channel for idle connections.
// 	idleConnsClosed := make(chan struct{})

// 	go func() {
// 		sigint := make(chan os.Signal, 1)
// 		signal.Notify(sigint, os.Interrupt) //* Catch OS signals.
// 		<-sigint

// 		//* Received an interrupt signal, shutdown.
// 		if err := a.Shutdown(); err != nil {
// 			return
// 		}

// 		close(idleConnsClosed)
// 	}()
// 	fiberConnURL, _ := utils.ConnectionURLBuilder("fiber", c)

// 	if err := a.Listen(fiberConnURL); err != nil {
// 		return fmt.Errorf("oops... server is not running! reason: %v", err)
// 	}

// 	<-idleConnsClosed
// 	return nil
// }

// func StartServer(a *fiber.App, c *configs.Config) error {
// 	fiberConnURL, _ := utils.ConnectionURLBuilder("fiber", c)

// 	if err := a.Listen(fiberConnURL); err != nil {
// 		return fmt.Errorf("oops... server is not running! reason: %v", err)
// 	}

// 	return nil
// }
