package api

import (
	"github.com/gofiber/fiber/v2"
)

func (s Server) healthCheck(c *fiber.Ctx) error {
	return c.SendString("OK")
}
