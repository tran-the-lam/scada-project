package service

import "github.com/gofiber/fiber/v2"

func Route(app fiber.Router, service IService) {
	api := app.Group("/scadas")
	api.Post("/state", PutStateHdl(service))
	api.Get("/state/:key", GetStateHdl(service))
}
