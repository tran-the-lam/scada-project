package service

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type PutStateBody struct {
	Key string `json:"key"`
	Val any    `json:"val"`
}

func PutStateHdl(service IService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var body PutStateBody
		if err := c.BodyParser(&body); err != nil {
			return c.Status(http.StatusBadRequest).JSON(ErrorResp(err, http.StatusBadRequest))
		}

		err := service.PutState(c.Context(), body.Key, body.Val)
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(ErrorResp(err, http.StatusInternalServerError))
		}

		return c.JSON(nil)
	}
}

type GetStateParams struct {
	Key string `params:"key"`
}

func GetStateHdl(service IService) fiber.Handler {
	return func(c *fiber.Ctx) error {

		var params GetStateParams
		if err := c.ParamsParser(&params); err != nil {
			return c.Status(http.StatusBadRequest).JSON(ErrorResp(err, http.StatusBadRequest))
		}

		// err := service.GetState(c.Context(), headers.AppId, headers.UserId)
		// if err != nil {
		// 	return c.Status(http.StatusInternalServerError).JSON(InternalErrorResponse())
		// }

		return c.JSON(nil)
	}
}
