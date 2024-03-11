package service

import (
	"backend/pkg/constant"
	e "backend/pkg/error"

	"fmt"

	"github.com/gofiber/fiber/v2"
)

type Response struct {
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
	Token   string `json:"token,omitempty"`
}

type LoginBody struct {
	UserID   string `json:"user_id"`
	Pwd      string `json:"password"`
	DeviceID string `json:"device_id"`
}

func LoginHdl(service IService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var body LoginBody
		if err := c.BodyParser(&body); err != nil {
			return e.BadRequest(err.Error())
		}

		fmt.Printf("====Login Hdl %+v\n", body)
		userAgent := c.Get("User-Agent")
		ip := c.Get("Remote-Address")
		token, err := service.Login(c.Context(), body.UserID, ip, userAgent, body.DeviceID, body.Pwd)
		if err != nil {
			return err
		}

		return c.JSON(Response{"success", nil, token})
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

func GetAllUserHdl(service IService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		actorID := c.Locals(constant.LOCAL_USER_ID).(string)
		actorRole := c.Locals(constant.LOCAL_USER_ROLE).(string)
		status := c.Query("status")
		if status != "-1" && status != "0" && status != "1" {
			return e.BadRequest("Status must -1 or 0 or 1")
		}

		rs, err := service.GetUsers(c.Context(), actorID, actorRole, status)
		if err != nil {
			return err
		}

		return c.JSON(Response{"success", rs, ""})
	}
}

type AddUserBody struct {
	UserID string `json:"user_id"`
	Role   string `json:"role"`
}

func AddUserHdl(service IService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var body AddUserBody
		if err := c.BodyParser(&body); err != nil {
			return e.BadRequest(err.Error())
		}

		actorID := c.Locals(constant.LOCAL_USER_ID).(string)
		actorRole := c.Locals(constant.LOCAL_USER_ROLE).(string)
		err := service.AddUser(c.Context(), actorID, actorRole, body.UserID, body.Role)
		if err != nil {
			return err
		}

		return c.JSON(Response{"success", nil, ""})
	}
}

type UpdatePwdBody struct {
	OldPwd string `json:"old_password"`
	NewPwd string `json:"new_password"`
}

func UpdatePwdHdl(service IService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var body UpdatePwdBody
		if err := c.BodyParser(&body); err != nil {
			return e.BadRequest(err.Error())
		}

		actorID := c.Locals(constant.LOCAL_USER_ID).(string)
		err := service.UpdatePwd(c.Context(), actorID, body.OldPwd, body.NewPwd)
		if err != nil {
			return err
		}

		return c.JSON(Response{"success", nil, ""})
	}
}

func GetHistoryChangePasswordHdl(service IService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		actorID := c.Locals(constant.LOCAL_USER_ID).(string)
		actorRole := c.Locals(constant.LOCAL_USER_ROLE).(string)
		rs, err := service.GetHistoryChangePassword(c.Context(), actorID, actorRole, actorID)
		if err != nil {
			return err
		}

		return c.JSON(Response{"success", rs, ""})
	}
}

func GetHistoryLoginHdl(service IService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		actorID := c.Locals(constant.LOCAL_USER_ID).(string)
		actorRole := c.Locals(constant.LOCAL_USER_ROLE).(string)
		rs, err := service.GetHistoryLogin(c.Context(), actorID, actorRole, actorID)
		if err != nil {
			return err
		}

		return c.JSON(Response{"success", rs, ""})
	}
}

func AddEventHdl(service IService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var body Event
		if err := c.BodyParser(&body); err != nil {
			return e.BadRequest(err.Error())
		}

		err := service.AddEvent(c.Context(), body)
		if err != nil {
			return err
		}

		return nil
	}
}

type GetEventQuery struct {
	Parameter   string `query:"parameter"`
	ParameterID string `query:"parameter_id"`
}

func GetEventHdl(service IService) fiber.Handler {
	return func(c *fiber.Ctx) error {

		var query GetEventQuery
		if err := c.QueryParser(&query); err != nil {
			return e.BadRequest(err.Error())
		}

		actorID := c.Locals(constant.LOCAL_USER_ID).(string)
		actorRole := c.Locals(constant.LOCAL_USER_ROLE).(string)

		events, err := service.GetEvent(c.Context(), actorID, actorRole, query.ParameterID, query.Parameter)
		if err != nil {
			return err
		}

		return c.JSON(Response{"success", events, ""})
	}
}

func SearchEventHdl(service IService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var query GetEventQuery
		if err := c.QueryParser(&query); err != nil {
			return e.BadRequest(err.Error())
		}

		actorID := c.Locals(constant.LOCAL_USER_ID).(string)
		actorRole := c.Locals(constant.LOCAL_USER_ROLE).(string)

		events, err := service.SearchEvent(c.Context(), actorID, actorRole, query.ParameterID, query.Parameter)
		if err != nil {
			return err
		}

		return c.JSON(Response{"success", events, ""})
	}
}

type UserParams struct {
	UserID string `params:"user_id"`
}

func ResetPwdHdl(service IService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var params UserParams
		if err := c.ParamsParser(&params); err != nil {
			return e.BadRequest(err.Error())
		}
		actorID := c.Locals(constant.LOCAL_USER_ID).(string)
		actorRole := c.Locals(constant.LOCAL_USER_ROLE).(string)
		err := service.ResetPassword(c.Context(), actorID, actorRole, params.UserID)
		if err != nil {
			return err
		}

		return c.JSON(Response{"success", nil, ""})
	}
}

func DeleteUserHdl(service IService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var params UserParams
		if err := c.ParamsParser(&params); err != nil {
			return e.BadRequest(err.Error())
		}
		actorID := c.Locals(constant.LOCAL_USER_ID).(string)
		actorRole := c.Locals(constant.LOCAL_USER_ROLE).(string)
		err := service.DeleteUser(c.Context(), actorID, actorRole, params.UserID)
		if err != nil {
			return err
		}
		return c.JSON(Response{"success", nil, ""})
	}
}
