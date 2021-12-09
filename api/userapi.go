package api

import (
	"github.com/clumpapp/clump-be/middleware"
	"github.com/clumpapp/clump-be/model"

	"github.com/gofiber/fiber/v2"
)

func (obj *API) login(c *fiber.Ctx) error {
	var loginDTO model.LoginDTO
	if err := c.BodyParser(&loginDTO); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	if obj.service.Login(loginDTO) {
		return c.JSON(fiber.Map{"token": middleware.GetToken(loginDTO.UserName)})
	}
	return c.SendStatus(fiber.StatusUnauthorized)
}

func (obj *API) signup(c *fiber.Ctx) error {
	var userDTO model.UserDTO
	if err := c.BodyParser(&userDTO); err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	obj.service.SignUp(userDTO)
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

func (obj *API) deleteuser(c *fiber.Ctx) error {
	var userDTO model.UserDTO
	if err := c.BodyParser(&userDTO); err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	obj.service.DeleteUser(userDTO)
	return c.SendStatus(fiber.StatusOK)
}
