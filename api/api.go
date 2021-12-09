package api

import (
	"github.com/clumpapp/clump-be/middleware"
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
	app.Use(middleware.GetLimiterMiddleware())

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

	//starting here
	api.Get("/group/messages", obj.getmessages)
	api.Get("/group/users", obj.getusers)

	api.Post("/signup", obj.signup)

	api.Patch("/group", obj.updategroup)
	api.Patch("group/user", obj.updateuser)
	api.Patch("group/user/message", obj.updatemessage)

	api.Delete("/group", obj.deletegroup)
	api.Delete("/group/user", obj.deleteuser)
	api.Delete("group/user/message", obj.deletemessage)

	//ending here

	app.Use(obj.NotFound)

	app.Listen(":8080")
}

func (obj *API) NotFound(c *fiber.Ctx) error {
	return c.Status(fiber.StatusNotFound).SendString("Hey, there are no friends to make here.")
}
