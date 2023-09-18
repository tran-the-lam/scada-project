package service

import (
	"backend/pkg/constant"
	e "backend/pkg/error"

	"fmt"

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
			return e.BadRequest(err.Error())
		}

		fmt.Printf("====PutStateHdl %+v\n", body)

		token, err := service.Login(c.Context(), body.UserID, body.Pwd)
		if err != nil {
			return err
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
			return e.BadRequest(err.Error())
		}

		actorID := c.Locals(constant.LOCAL_USER_ID).(string)
		actorRole := c.Locals(constant.LOCAL_USER_ROLE).(string)
		rs, err := service.GetState(c.Context(), actorID, actorRole, query.Key)
		if err != nil {
			return err
		}

		return c.JSON(Response{"success", rs, ""})
	}
}

type AddUserBody struct {
	UserID string `json:"user_id"`
	Role   string `json:"role"`
	Pwd    string `json:"password"`
}

func AddUserHdl(service IService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var body AddUserBody
		if err := c.BodyParser(&body); err != nil {
			return e.BadRequest(err.Error())
		}

		actorID := c.Locals(constant.LOCAL_USER_ID).(string)
		actorRole := c.Locals(constant.LOCAL_USER_ROLE).(string)
		err := service.AddUser(c.Context(), actorID, actorRole, body.UserID, body.Pwd, body.Role)
		if err != nil {
			return err
		}

		return c.JSON(Response{"success", "", ""})
	}
}
