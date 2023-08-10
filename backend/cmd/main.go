package main

import (
	"backend/config"
	"backend/service"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func main() {
	cfg := config.InitConfig()
	app := fiber.New()

	svc := service.NewService(cfg)
	service.Route(app, svc)

	app.Listen(fmt.Sprintf(":%s", cfg.PORT))
}
