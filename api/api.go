package api

import (
	"github.com/clumpapp/clump-be/middleware"
	"github.com/clumpapp/clump-be/service"

	"github.com/gofiber/fiber/v2"
)

const prefix = "/api/v1"
const id = "id"

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
	app.Post("/signup", obj.signup)

	app.Use(middleware.GetJWTMiddleware())

	api := app.Group(prefix)

	api.Get("/messages", obj.getgroupmessages)
	api.Get("/users", obj.getgroupusers)

	api.Post("/messages", obj.postmessage)

	api.Put("/groups/:"+id, obj.putgroup)
	api.Put("/users/:"+id, obj.putuser)
	api.Put("/messages/:"+id, obj.putmessage)

	api.Delete("/groups/:"+id, obj.deletegroup)
	api.Delete("/users/:"+id, obj.deleteuser)
	api.Delete("/messages/:"+id, obj.deletemessage)

	app.Use(obj.notFound)

	app.Listen(":8080")
}

func (obj *API) notFound(c *fiber.Ctx) error {
	return c.Status(fiber.StatusNotFound).SendString("Hey, there are no friends to make here.")
}

func (obj *API) getIDFromParam(c *fiber.Ctx) string {
	return c.Params(id)
}

func (obj *API) getGroupIDFromToken(c *fiber.Ctx) float64 {
	return middleware.GetGroupID(c.Locals(middleware.ContextKey))
}

func (obj *API) getUserIDFromToken(c *fiber.Ctx) float64 {
	return middleware.GetUserID(c.Locals(middleware.ContextKey))
}
