package api

import (
	"github.com/clumpapp/clump-be/middleware"
	"github.com/clumpapp/clump-be/model"
	"github.com/clumpapp/clump-be/service"

	"github.com/gofiber/fiber/v2"
)

const prefix = "/api/v1"

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

	app.Post("/login", obj.login)

	app.Use(middleware.GetJWTMiddleware())

	api := app.Group(prefix)

	// Currently using Authorization header for token
	// What should be used will be decided with according to front-end
	api.Get("/restricted", func(c *fiber.Ctx) error {
		name := middleware.GetClaim(c.Locals("user"))
		return c.SendString("Welcome " + name)
	})

	api.Post("/chat", obj.textshare)

	app.Use(obj.NotFound)

	app.Listen(":8080")
}

func (obj *API) NotFound(c *fiber.Ctx) error {
	return c.Status(fiber.StatusNotFound).SendString("Hey, there are no friends to make here.")
}

func (obj *API) login(c *fiber.Ctx) error {
	var loginDTO model.LoginDTO
	if err := c.BodyParser(&loginDTO); err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	if obj.service.Login(loginDTO) {
		return c.JSON(fiber.Map{"token": middleware.GetToken(loginDTO.UserName)})
	}
	return c.SendStatus(fiber.StatusUnauthorized)
}

func (obj *API) textshare(c *fiber.Ctx) error {
	var messageDTO model.MessageDTO
	if err := c.BodyParser(&messageDTO); err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	obj.service.TextShare(messageDTO)
	return c.SendStatus(fiber.StatusOK)
}
