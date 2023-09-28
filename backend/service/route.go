package service

import (
	"backend/middleware"

	"github.com/gofiber/fiber/v2"
)

func Route(app fiber.Router, service IService) {
	loginApi := app.Group("/login")
	loginApi.Post("", LoginHdl(service))

	userApi := app.Group("/users", middleware.Auth)
	userApi.Get("/state", GetStateHdl(service))
	userApi.Post("", AddUserHdl(service))
	userApi.Put("/password", UpdatePwdHdl(service))
	userApi.Get("/history/change-password", GetHistoryChangePasswordHdl(service))
	userApi.Get("/history/login", GetHistoryLoginHdl(service))

	eventApi := app.Group("/events")
	eventApi.Post("", AddEventHdl(service))
	eventApi.Get("", GetEventHdl(service))
}
