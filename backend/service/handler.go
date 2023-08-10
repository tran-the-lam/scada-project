package service

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type Response struct {
	Message string `json:"message"`
	Data    string `json:"data,omitempty"`
}

type PutStateBody struct {
	ChainCodeID string   `json:"chain_code_id"`
	ChannelID   string   `json:"channel_id"`
	Function    string   `json:"function"`
	Args        []string `json:"args"`
}

func PutStateHdl(service IService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var body PutStateBody
		if err := c.BodyParser(&body); err != nil {
			return c.Status(http.StatusBadRequest).JSON(ErrorResp(err, http.StatusBadRequest))
		}

		fmt.Printf("====PutStateHdl %+v\n", body)

		err := service.PutState(c.Context(), body.ChainCodeID, body.ChannelID, body.Function, body.Args)
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(ErrorResp(err, http.StatusInternalServerError))
		}

		return c.JSON(Response{"success", ""})
	}
}

type GetStateQuery struct {
	ChannelID   string   `query:"channel_id"`
	ChainCodeID string   `query:"chain_code_id"`
	Function    string   `query:"function"`
	Args        []string `query:"args"`
}

func GetStateHdl(service IService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var query GetStateQuery
		if err := c.QueryParser(&query); err != nil {
			return c.Status(http.StatusBadRequest).JSON(ErrorResp(err, http.StatusBadRequest))
		}

		fmt.Printf("====GetStateHdl %+v\n", query)
		rs, err := service.GetState(c.Context(), query.ChainCodeID, query.ChannelID, query.Function, query.Args)
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(ErrorResp(err, http.StatusInternalServerError))
		}

		return c.JSON(Response{"success", rs})
	}
}
