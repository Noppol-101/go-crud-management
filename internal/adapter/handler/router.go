package handler

import (
	"my-crud-management/middleware"

	"github.com/gofiber/fiber/v2"
)

type Router struct {
	*fiber.App
}

type RouterParams struct {
	CategoryHandler CategoryHandler
}

func NewRouter(p RouterParams) (*Router, error) {
	app := fiber.New(fiber.Config{
		BodyLimit: 5000 * 1024 * 1024,
	})

	app.Use(middleware.Recover())

	api := app.Group("/api")

	version1 := api.Group("/v1")

	categoryv1 := version1.Group("/category")
	categoryv1.Post("/", p.CategoryHandler.CreateCategory)

	return &Router{
		App: app,
	}, nil
}

func (r *Router) Serve(listenAddr string) error {
	return r.Listen(listenAddr)
}
