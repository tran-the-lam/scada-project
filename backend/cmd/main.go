package main

import (
	"backend/config"
	"backend/service"
	"fmt"

	fiber "github.com/gofiber/fiber/v2"
)

func main() {
	cfg := config.InitConfig()
	app := fiber.New()

	// app.Use(jwtware.New(jwtware.Config{
	// 	SigningKey: jwtware.SigningKey{Key: []byte("secret")},
	// }))

	svc := service.NewService(cfg)
	service.Route(app, svc)
	app.Listen(fmt.Sprintf(":%s", cfg.PORT))
}
