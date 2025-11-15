package middleware

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/fazriegi/go-boilerplate/internal/entity"
	"github.com/fazriegi/go-boilerplate/internal/pkg"
	"github.com/gofiber/fiber/v2"
)

func Authentication(jwt *pkg.JWT) func(ctx *fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		var response = pkg.Response{}
		header := ctx.Get("Authorization")
		isHasBearer := strings.HasPrefix(header, "Bearer")

		if !isHasBearer {
			response = response.Create(http.StatusUnauthorized, "sign in to proceed", nil)

			return ctx.Status(response.Code).JSON(response)
		}

		tokenString := strings.Split(header, " ")[1]

		verifiedToken, err := jwt.VerifyJWTTOken(tokenString)
		if err != nil {
			response = response.Create(http.StatusUnauthorized, err.Error(), nil)

			return ctx.Status(response.Code).JSON(response)
		}

		jsonData, _ := json.Marshal(verifiedToken)

		var user entity.User
		json.Unmarshal(jsonData, &user)

		ctx.Locals("user", user)

		return ctx.Next()
	}
}
