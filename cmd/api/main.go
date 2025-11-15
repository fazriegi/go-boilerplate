package main

import (
	"fmt"
	"log"

	"github.com/fazriegi/go-boilerplate/internal/infrastructure/config"
	"github.com/fazriegi/go-boilerplate/internal/infrastructure/database"
	"github.com/fazriegi/go-boilerplate/internal/infrastructure/logger"
	"github.com/fazriegi/go-boilerplate/internal/infrastructure/router.go"
	"github.com/fazriegi/go-boilerplate/internal/middleware"
	"github.com/fazriegi/go-boilerplate/internal/pkg"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	viperConfig := config.NewViper()
	database.NewMysql(viperConfig)
	jwt := pkg.InitJWT(viperConfig)
	file := logger.New(viperConfig)
	defer file.Close()

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,PUT,DELETE,OPTIONS",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
	}))

	app.Use(middleware.LogMiddleware())
	port := viperConfig.GetInt("web.port")
	router.NewRoute(app, jwt)

	log.Fatal(app.Listen(fmt.Sprintf(":%d", port)))
}
