package handler

import (
	"my-crud-management/middleware"

	"github.com/gofiber/fiber/v2"
)

type Router struct {
	*fiber.App
}

type RouterParams struct {
}

func NewRouter(p RouterParams) (*Router, error) {
	app := fiber.New(fiber.Config{
		BodyLimit: 5000 * 1024 * 1024,
	})

	app.Use(middleware.Recover())

	return &Router{
		App: app,
	}, nil
}

func (r *Router) Serve(listenAddr string) error {
	return r.Listen(listenAddr)
}
