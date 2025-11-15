package controller

import (
	"net/http"

	"github.com/fazriegi/go-boilerplate/internal/entity"
	"github.com/fazriegi/go-boilerplate/internal/infrastructure/logger"
	"github.com/fazriegi/go-boilerplate/internal/pkg"
	"github.com/fazriegi/go-boilerplate/internal/usecase"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type AuthController interface {
	Register(ctx *fiber.Ctx) error
	Login(ctx *fiber.Ctx) error
}

type authController struct {
	usecase usecase.AuthUsecase
	logger  *logrus.Logger
}

func NewAuthController(usecase usecase.AuthUsecase) AuthController {
	logger := logger.Get()
	return &authController{
		usecase,
		logger,
	}
}

func (c *authController) Register(ctx *fiber.Ctx) error {
	var (
		response pkg.Response
		reqBody  entity.RegisterRequest
	)

	if err := ctx.BodyParser(&reqBody); err != nil {
		c.logger.Errorf("error parsing request body: %s", err.Error())
		return ctx.Status(http.StatusBadRequest).JSON(response.Create(http.StatusBadRequest, pkg.ErrParseQueryParam.Error(), nil))
	}

	// validate reqBody struct
	validationErr := pkg.ValidateRequest(&reqBody)
	if len(validationErr) > 0 {
		errResponse := map[string]any{
			"errors": validationErr,
		}

		return ctx.Status(http.StatusUnprocessableEntity).JSON(response.Create(http.StatusUnprocessableEntity, pkg.ErrValidation.Error(), errResponse))
	}

	response = c.usecase.Register(&reqBody)

	return ctx.Status(response.Status.Code).JSON(response)
}

func (c *authController) Login(ctx *fiber.Ctx) error {
	var (
		response pkg.Response
		reqBody  entity.LoginRequest
	)

	if err := ctx.BodyParser(&reqBody); err != nil {
		c.logger.Errorf("error parsing request body: %s", err.Error())
		return ctx.Status(http.StatusBadRequest).JSON(response.Create(http.StatusBadRequest, pkg.ErrParseQueryParam.Error(), nil))
	}

	// validate reqBody struct
	validationErr := pkg.ValidateRequest(&reqBody)
	if len(validationErr) > 0 {
		errResponse := map[string]any{
			"errors": validationErr,
		}

		return ctx.Status(http.StatusUnprocessableEntity).JSON(response.Create(http.StatusUnprocessableEntity, pkg.ErrValidation.Error(), errResponse))
	}

	response = c.usecase.Login(&reqBody)

	return ctx.Status(response.Status.Code).JSON(response)
}
