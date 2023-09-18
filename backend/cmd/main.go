package main

import (
	"backend/config"
	e "backend/pkg/error"
	"backend/service"
	"fmt"

	fiber "github.com/gofiber/fiber/v2"
)

func main() {
	cfg := config.InitConfig()
	app := fiber.New(fiber.Config{
		AppName:      "Scada Project",
		ErrorHandler: e.HandleError,
	})

	svc := service.NewService(cfg)
	service.Route(app, svc)
	app.Listen(fmt.Sprintf(":%s", cfg.PORT))
}
