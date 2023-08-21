package service

import (
	"bytes"
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
)

func TestPutStateHdl(t *testing.T) {
	var testcase = []struct {
		name    string
		service IService
	}{
		{
			name:    "Test PutStateHdl",
			service: svc,
		},
	}

	for _, tc := range testcase {
		t.Run(tc.name, func(t *testing.T) {
			app := fiber.New()
			Route(app, tc.service)

			payload := PutStateBody{
				ChainCodeID: "basic",
				ChannelID:   "mychannel",
				Function:    "CreateKey",
				Args:        []string{"k", "13"},
			}

			jsonPayload, err := json.Marshal(payload)
			if err != nil {
				panic(err)
			}

			req := httptest.NewRequest("POST", "/scadas/state", bytes.NewBuffer(jsonPayload))
			req.Header.Set("Content-Type", "application/json")
			resp, err := app.Test(req)
			if err != nil {
				t.Fatal(err)
			}
			if resp.StatusCode != 200 {
				t.Fatalf("Expected status code 200, got %v", resp.StatusCode)
			}
		})
	}
}

func TestGetStateHdl(t *testing.T) {
	var testcase = []struct {
		name    string
		service IService
	}{
		{
			name:    "Test GetStateHdl",
			service: svc,
		},
	}

	for _, tc := range testcase {
		t.Run(tc.name, func(t *testing.T) {
			app := fiber.New()
			Route(app, tc.service)

			req := httptest.NewRequest("GET", "/scadas/state?chain_code_id=basic&channel_id=mychannel&function=QueryKey&args=k", nil)
			req.Header.Set("Content-Type", "application/json")
			resp, err := app.Test(req)
			if err != nil {
				t.Fatal(err)
			}
			if resp.StatusCode != 200 {
				t.Fatalf("Expected status code 200, got %v", resp.StatusCode)
			}
		})
	}
}
