package service

import (
	"backend/middleware"

	"github.com/gofiber/fiber/v2"
)

func Route(app fiber.Router, service IService) {
	api2 := app.Group("/login")
	api2.Post("", LoginHdl(service))

	api := app.Group("/users", middleware.Auth)
	api.Get("/state", GetStateHdl(service))
	api.Post("", AddUserHdl(service))

	// api.Put("/password", UpdatePwdHdl(service))
}
