package api

import (
	"clump/middleware"
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
		return c.JSON(fiber.Map{"Hello": "World!"})
	})

	app.Get("/login", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"token": middleware.GetToken("clump")})
	})

	app.Use(middleware.GetJWTMiddleware())

	app.Listen(":8080")
}
