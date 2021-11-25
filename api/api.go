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

	//starting here
	api.Get("/group/messages", obj.getmessages)
	api.Get("/group/users", obj.getusers)

	api.Post("/signup", obj.signup)

	api.Patch("/group", obj.updategroup)
	api.Patch("group/user", obj.updateuser)
	api.Patch("group/user/message", obj.updateuser)

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

func (obj *API) getmessages(c *fiber.Ctx) error {
	var groupDTO model.GroupDTO
	var messageDTO model.MessageDTO
	if err := c.BodyParser(&groupDTO); err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	messageDTO = obj.service.GetMessages(groupDTO)
	return c.JSON(fiber.Map{
		"messagetype": middleware.GetToken(messageDTO.MessageType),
		"messagetext": middleware.GetToken(messageDTO.MessageText),
	})
}

func (obj *API) getusers(c *fiber.Ctx) error {
	var groupDTO model.GroupDTO
	var userDTO model.UserDTO
	if err := c.BodyParser(&groupDTO); err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	userDTO = obj.service.GetUsers(groupDTO)
	return c.JSON(fiber.Map{
		"username": middleware.GetToken(userDTO.UserName),
		"userpfp":  middleware.GetToken(userDTO.ProfilePicture),
	})

}

func (obj *API) updategroup(c *fiber.Ctx) error {
	var groupDTO model.GroupDTO
	if err := c.BodyParser(&groupDTO); err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	obj.service.UpdateGroup(groupDTO)
	return c.SendStatus(fiber.StatusOK)
}
func (obj *API) updateuser(c *fiber.Ctx) error {
	var userDTO model.UserDTO
	if err := c.BodyParser(&userDTO); err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	obj.service.UpdateUser(userDTO)
	return c.SendStatus(fiber.StatusOK)
}
func (obj *API) updatemessage(c *fiber.Ctx) error {
	var messageDTO model.MessageDTO
	if err := c.BodyParser(&messageDTO); err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	obj.service.UpdateMessage(messageDTO)
	return c.SendStatus(fiber.StatusOK)
}

func (obj *API) deletegroup(c *fiber.Ctx) error {
	var groupDTO model.GroupDTO
	if err := c.BodyParser(&groupDTO); err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	obj.service.DeleteGroup(groupDTO)
	return c.SendStatus(fiber.StatusOK)
}
func (obj *API) deleteuser(c *fiber.Ctx) error {
	var userDTO model.UserDTO
	if err := c.BodyParser(&userDTO); err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	obj.service.DeleteUser(userDTO)
	return c.SendStatus(fiber.StatusOK)
}
func (obj *API) deletemessage(c *fiber.Ctx) error {
	var messageDTO model.MessageDTO
	if err := c.BodyParser(&messageDTO); err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	obj.service.DeleteMessage(messageDTO)
	return c.SendStatus(fiber.StatusOK)
}

func (obj *API) signup(c *fiber.Ctx) error {
	var userDTO model.UserDTO
	if err := c.BodyParser(&userDTO); err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	obj.service.SignUp(userDTO)
	return c.SendStatus(fiber.StatusOK)
}
