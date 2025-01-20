package server

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type Status string

const (
	StatusRunning Status = "running"
)

type StatusResponse struct {
	Status Status   `json:"status"`
	Errors []string `json:"errors,omitempty"`
}

func HandleGetStatus(ctx *fiber.Ctx) error {
	ctx.JSON(StatusResponse{Status: StatusRunning})
	return ctx.SendStatus(http.StatusOK)
}
