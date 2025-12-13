package router

import (
	"github.com/fazriegi/go-boilerplate/internal/controller"
	"github.com/fazriegi/go-boilerplate/internal/infrastructure/repository"
	"github.com/fazriegi/go-boilerplate/internal/middleware"
	"github.com/fazriegi/go-boilerplate/internal/pkg"
	"github.com/fazriegi/go-boilerplate/internal/usecase"
	"github.com/gofiber/fiber/v2"
)

func NewRoute(app *fiber.App, jwt *pkg.JWT) {
	userRepo := repository.NewUserRepository()
	authRepo := repository.NewAuthRepository()
	authUC := usecase.NewAuthUsecase(userRepo, authRepo, jwt)
	authController := controller.NewAuthController(authUC)

	v1 := app.Group("/api/v1")
	{
		v1.Post("/register", authController.Register)
		v1.Post("/login", authController.Login)
		v1.Get("/check-token", middleware.Authentication(jwt), authController.CheckToken)
	}
}
