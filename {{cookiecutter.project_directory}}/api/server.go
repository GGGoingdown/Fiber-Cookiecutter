package api

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"

	"github.com/GGGoingdown/Fiber-Cookiecutter/config"
	"github.com/GGGoingdown/Fiber-Cookiecutter/docs"
	"github.com/GGGoingdown/Fiber-Cookiecutter/pkg"
	"github.com/GGGoingdown/Fiber-Cookiecutter/utils"
)

type Server struct {
	config *config.Config
	api    *fiber.App
	logger *utils.LogHelper
}

func NewServer(cfg *config.Config, logger *utils.LogHelper) *Server {
	server := &Server{
		config: cfg,
		logger: logger,
	}
	server.setupRouter()
	return server
}

func (s *Server) setupRouter() {
	prefork := false
	if s.config.Mode == "production" {
		prefork = true
	}

	app := fiber.New(fiber.Config{
		Prefork: prefork,
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			// Status code defaults to 500
			code := fiber.StatusInternalServerError

			// Retrieve the custom status code if it's a *fiber.Error
			var e *fiber.Error
			if errors.As(err, &e) {
				code = e.Code
			}

			c.Status(code)
			if code >= 500 {
				s.logger.Error(err.Error())
			}
			return c.SendString(err.Error())
		},
	})

	recoverMiddleware := newRecoverMiddleware()
	loggerMiddleware := newLoggerMiddleware()
	app.Use(recoverMiddleware, loggerMiddleware)

	if s.config.SentryDsn != "" {
		sentryOptions := pkg.NewSentryOptions(
			s.config.SentryDsn,
			s.config.SentryTraceSampleRate,
			s.config.Mode,
			docs.SwaggerInfo.Title,
			docs.SwaggerInfo.Version,
		)
		pkg.InitSentry(sentryOptions)
		sentryMiddleware := newSentryMiddleware()
		app.Use(sentryMiddleware)
	}

	// Health check
	app.Get("/health", s.healthCheck)

	// Swagger
	app.Get("/swagger/*", swagger.HandlerDefault) // default

	// Do not remove this line
	s.api = app
}

func (s *Server) Run() error {
	return s.api.Listen(s.config.ServerAddress)
}

func (s *Server) Shutdown() error {
	return s.api.Shutdown()
}
