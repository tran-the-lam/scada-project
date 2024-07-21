package error

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type ApiError struct {
	HttpCode int    `json:"http_code"`
	Code     int    `json:"code"`
	Message  string `json:"message"`
}

func (e ApiError) Error() string {
	return e.Message
}

func InternalServerError() ApiError {
	return ApiError{
		HttpCode: 500,
		Code:     10001,
		Message:  "Internal server error",
	}
}

func Unauthorized() ApiError {
	return ApiError{
		HttpCode: 401,
		Code:     10002,
		Message:  "Unauthorized",
	}
}

func BadRequest(msg string) ApiError {
	return ApiError{
		HttpCode: 400,
		Code:     10003,
		Message:  msg,
	}
}

func NotFound() ApiError {
	return ApiError{
		HttpCode: 404,
		Code:     10004,
		Message:  "Not found",
	}
}

func LoginFailed() ApiError {
	return ApiError{
		HttpCode: 401,
		Code:     10005,
		Message:  "User or password is incorrect",
	}
}

func Forbidden() ApiError {
	return ApiError{
		HttpCode: 403,
		Code:     10006,
		Message:  "Forbidden",
	}
}

func TxErr(msg string) ApiError {
	return ApiError{
		HttpCode: 500,
		Code:     10007,
		Message:  msg,
	}
}

func HandleError(ctx *fiber.Ctx, err error) error {
	switch err.(type) {
	case ApiError:
		apiError := err.(ApiError)
		return ctx.Status(apiError.HttpCode).JSON(apiError)
	default:
		fmt.Printf("Internal Error: %s", err.Error())
		return ctx.Status(500).JSON(InternalServerError())
	}
}

type FError int

const (
	EndorsedError FError = iota
	SubmittedError
	CreateProposalError
)
