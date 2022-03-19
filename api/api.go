package api

import (
	"github.com/clumpapp/clump-be/middleware"
	"github.com/clumpapp/clump-be/service"

	"github.com/gofiber/fiber/v2"
)

const (
	accountPrefix   = "/account"
	adminPrefix     = "/admin"
	apiPrefix       = "/api/v1"
	websocketPrefix = "/ws"
	id              = "id"
)

type API struct {
	service *service.Service
}

func New(service *service.Service) *API {
	return &API{service}
}

func (obj *API) Run() {
	app := fiber.New()

	app.Static("/", "./public")

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.SendStatus(fiber.StatusOK)
	})

	account := app.Group(accountPrefix)

	account.Use(middleware.GetLimiterMiddleware())

	account.Post("/login", obj.login)
	account.Post("/signup", obj.signup)

	admin := app.Group(adminPrefix)

	admin.Post("/interests", obj.postInterest)

	api := app.Group(apiPrefix)

	api.Use(middleware.GetJWTMiddleware())

	api.Get("/interests", obj.getInterests)
	api.Get("/messages", obj.getGroupMessages)
	api.Get("/users", obj.getGroupUsers)
	api.Get("/users/me", obj.getUser)
	api.Get("/users/assign", obj.assignGroup)

	api.Post("/users/interests", obj.addInterests)
	api.Post("/messages", obj.postMessage)
	api.Post("/messages/image", obj.postImage)
	api.Post("/messages/video", obj.postVideo)
	api.Post("/messages/other", obj.postOther)

	api.Put("/users", obj.putUser)

	api.Delete("/users", obj.deleteUser)
	api.Delete("/messages/:"+id, obj.deletemessage)

	ws := api.Group(websocketPrefix)

	ws.Use(obj.setup)

	ws.Get("/messages", obj.websocket())

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
