package http

import (
	"github.com/gofiber/fiber/v2"
	"time"
)

func NewConfiguredFiberApp() *fiber.App {
	app := fiber.New(fiber.Config{
		AppName:               "SysPeak",
		DisableStartupMessage: true,
		EnablePrintRoutes:     true,
		GETOnly:               true,
		Network:               "tcp4",
		ReadTimeout:           5 * time.Second,
	})
	return app
}
