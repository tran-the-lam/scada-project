package main

import (
	"backend/config"
	e "backend/pkg/error"
	"backend/service"
	"fmt"

	fiber "github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	cfg := config.InitConfig()
	app := fiber.New(fiber.Config{
		AppName:      "Scada Project",
		ErrorHandler: e.HandleError,
	})

	app.Use(cors.New(cors.Config{
		AllowHeaders:     "*",
		AllowOrigins:     "*",
		AllowCredentials: true,
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
	}))

	svc := service.NewService(cfg)
	service.Route(app, svc)
	app.Listen(fmt.Sprintf(":%s", cfg.PORT))
}
