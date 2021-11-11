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

	app.Use(middleware.GetCORSMiddleware())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"Hello": "World!"})
	})

	app.Get("/login", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"token": middleware.GetToken("clump")})
	})

	app.Use(middleware.GetJWTMiddleware())

	// Currently using Authorization header for token
	// What should be used will be decided with according to front-end
	app.Get("/restricted", func(c *fiber.Ctx) error {
		name := middleware.GetClaim(c.Locals("user"))
		return c.SendString("Welcome " + name)
	})

	app.Listen(":8080")
}
