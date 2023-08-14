package service

import "github.com/gofiber/fiber/v2"

func ErrorResp(e error, statusCode int) *fiber.Map {
	return &fiber.Map{
		"error":       e.Error(),
		"status_code": statusCode,
	}
}

type Error struct {
	Err  error
	Code FError
}

func (r Error) Error() string {
	return r.Err.Error()
}

type FError int

const (
	EndorsedError FError = iota
	SubmittedError
	CreateProposalError
)
