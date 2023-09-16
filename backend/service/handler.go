package service

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type Response struct {
	Message string `json:"message"`
	Data    string `json:"data,omitempty"`
	Token   string `json:"token,omitempty"`
}

type LoginBody struct {
	UserID string `json:"user_id"`
	Pwd    string `json:"password"`
}

func LoginHdl(service IService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var body LoginBody
		if err := c.BodyParser(&body); err != nil {
			return c.Status(http.StatusBadRequest).JSON(ErrorResp(err, http.StatusBadRequest))
		}

		fmt.Printf("====PutStateHdl %+v\n", body)

		token, err := service.Login(c.Context(), body.UserID, body.Pwd)
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(ErrorResp(err, http.StatusInternalServerError))
		}

		return c.JSON(Response{"success", "", token})
	}
}

type GetStateQuery struct {
	Key string `query:"key"`
}

func GetStateHdl(service IService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var query GetStateQuery
		if err := c.QueryParser(&query); err != nil {
			return c.Status(http.StatusBadRequest).JSON(ErrorResp(err, http.StatusBadRequest))
		}

		rs, err := service.GetState(c.Context(), query.Key)
		fmt.Printf("====GetStateHdl %+v\n == %+v", query, rs)

		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(ErrorResp(err, http.StatusInternalServerError))
		}

		return c.JSON(Response{"success", rs, ""})
	}
}

type PutStateBody struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func PutStateHdl(service IService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var body PutStateBody
		if err := c.BodyParser(&body); err != nil {
			return c.Status(http.StatusBadRequest).JSON(ErrorResp(err, http.StatusBadRequest))
		}

		err := service.PutState(c.Context(), body.Key, body.Value)
		fmt.Printf("====GetStateHdl %+v\n == ", err)

		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(ErrorResp(err, http.StatusInternalServerError))
		}

		return c.JSON(Response{"success", "", ""})
	}
}
