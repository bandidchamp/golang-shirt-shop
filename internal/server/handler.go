package server

import (
	"github.com/gofiber/fiber/v2"

	authenHttp "shirt-shop/internal/authen/deliveries/http"
	authenRepository "shirt-shop/internal/authen/repositories"
	authenUseCase "shirt-shop/internal/authen/usecases"

	userHttp "shirt-shop/internal/user/deliveries/http"
	userRepository "shirt-shop/internal/user/repositories"
	userUseCase "shirt-shop/internal/user/usecases"

	productHttp "shirt-shop/internal/product/deliveries/http"
	productRepository "shirt-shop/internal/product/repositories"
	productUseCase "shirt-shop/internal/product/usecases"
	"shirt-shop/pkg/middleware"
)

func (s *Server) MapHandlers(a *fiber.App) error {

	// setup middleware here.
	a.Use(middleware.MiddleWareLogger())

	// >> setup clean stuct <<

	//* Version 1 group.
	api := a.Group("/api")
	v1 := api.Group("/v1")

	//* authen
	authenGroup := v1.Group("/login")
	authenRepo := authenRepository.NewAuthRepository(s.db)
	authenUC := authenUseCase.NewAuthUseCase(authenRepo)
	authenHandler := authenHttp.NewAuthenHandler(authenUC)
	authenHttp.MapAuthRoute(authenGroup, authenHandler)

	//* User group.
	userGroup := v1.Group("/users")
	userRepo := userRepository.NewUserRepository(s.db)
	userUC := userUseCase.NewUserUseCase(userRepo)
	userHandler := userHttp.NewUserHandler(userUC)
	userHttp.MapUserRoute(userGroup, userHandler)

	//* Product group.
	productGroup := v1.Group("/products")
	productRepo := productRepository.NewProductRepository(s.db, s.cache)
	productUC := productUseCase.NewProductUseCase(productRepo)
	productHandler := productHttp.NewProductHandler(productUC)
	productHttp.MapProductRoute(productGroup, productHandler)

	return nil
}
