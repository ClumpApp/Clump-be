package api

import (
	"clump/service"

	"github.com/gofiber/fiber/v2"
)

type API struct {
	service *service.Service
}

func New(service *service.Service) *API {
	return &API{service}
}

func (obj *API) Run() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World! This is Clump.")
	})

	app.Listen(":8080")
}
