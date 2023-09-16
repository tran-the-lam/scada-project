package service

import "github.com/gofiber/fiber/v2"

func Route(app fiber.Router, service IService) {
	api := app.Group("/users")
	api.Post("/login", LoginHdl(service))
	api.Get("/state", GetStateHdl(service))

	api.Post("/state", PutStateHdl(service))
	// api.Put("/password", UpdatePwdHdl(service))

}
