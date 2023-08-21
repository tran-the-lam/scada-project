package service

import (
	"testing"

	"github.com/gofiber/fiber/v2"
)

func TestRoute(t *testing.T) {
	var testcase = []struct {
		name    string
		service IService
	}{
		{
			name:    "Test route",
			service: svc,
		},
	}

	for _, tc := range testcase {
		t.Run(tc.name, func(t *testing.T) {
			app := fiber.New()
			Route(app, tc.service)
		})
	}
}
