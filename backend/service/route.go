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
	userApi.Get("", GetAllUserHdl(service))
	userApi.Put("/password", UpdatePwdHdl(service))
	userApi.Get("/history/change-password", GetHistoryChangePasswordHdl(service))
	userApi.Get("/history/login", GetHistoryLoginHdl(service))
	userApi.Post("/:user_id/reset-password", ResetPwdHdl(service))
	userApi.Delete("/:user_id/delete", DeleteUserHdl(service))

	eventApi := app.Group("/events", middleware.Auth)
	eventApi.Post("", AddEventHdl(service))
	eventApi.Get("", GetEventHdl(service))
	eventApi.Get("/search", SearchEventHdl(service))
}
