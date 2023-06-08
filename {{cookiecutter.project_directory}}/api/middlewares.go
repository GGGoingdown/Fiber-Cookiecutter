package api

import (
	"github.com/gofiber/fiber/v2"

	"github.com/gofiber/fiber/v2/middleware/basicauth"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"

	sentryfiber "github.com/aldy505/sentry-fiber"
)

func newLoggerMiddleware() fiber.Handler {
	middleware := logger.New(logger.Config{
		TimeFormat: "2006-01-02T15:04:05",
		Format:     "${time} - [${ip}] - ${status} - ${method} ${path} ${resBody}\n",
	})

	return middleware
}

func newRecoverMiddleware() fiber.Handler {
	middleware := recover.New(
		recover.Config{
			EnableStackTrace: true,
		},
	)
	return middleware
}

func newBasicAuthMiddleware(users map[string]string) fiber.Handler {
	middleware := basicauth.New(basicauth.Config{
		Users: users,
		Unauthorized: func(c *fiber.Ctx) error {
			return c.Status(fiber.StatusUnauthorized).JSON(MessageResponse{Message: "Unauthorized"})
		},
	})
	return middleware
}

func newSentryMiddleware() fiber.Handler {
	middleware := sentryfiber.New(sentryfiber.Options{
		Repanic:         true,
		WaitForDelivery: false,
	})
	return middleware
}
