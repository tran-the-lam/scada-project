package service

import "github.com/gofiber/fiber/v2"

func ErrorResp(e error, statusCode int) *fiber.Map {
	return &fiber.Map{
		"error":       e.Error(),
		"status_code": statusCode,
	}
}
