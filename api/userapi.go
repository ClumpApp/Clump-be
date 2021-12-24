package api

import (
	"github.com/clumpapp/clump-be/middleware"
	"github.com/clumpapp/clump-be/model"

	"github.com/gofiber/fiber/v2"
)

func (obj *API) login(c *fiber.Ctx) error {
	var loginDTO model.LoginDTO
	if err := c.BodyParser(&loginDTO); err != nil {
		return c.SendStatus(fiber.StatusUnprocessableEntity)
	}
	uid, gid, matches := obj.service.Login(loginDTO)
	if matches {
		token := middleware.CreateToken(uid, gid)
		return c.SendString(token)
	}
	return c.SendStatus(fiber.StatusUnauthorized)
}

func (obj *API) signup(c *fiber.Ctx) error {
	var signupDTO model.SignUpDTO
	if err := c.BodyParser(&signupDTO); err != nil {
		return c.SendStatus(fiber.StatusUnprocessableEntity)
	}
	uid, gid, available := obj.service.SignUp(signupDTO)
	if available {
		token := middleware.CreateToken(uid, gid)
		return c.SendString(token)
	}
	return c.SendStatus(fiber.StatusUnauthorized)
}

func (obj *API) getGroupUsers(c *fiber.Ctx) error {
	id := obj.getGroupIDFromToken(c)
	usersDTO := obj.service.GetGroupUsers(id)
	return c.JSON(usersDTO)
}

func (obj *API) getUser(c *fiber.Ctx) error {
	id := obj.getUserIDFromToken(c)
	usersDTO := obj.service.GetUser(id)
	return c.JSON(usersDTO)
}

func (obj *API) putUser(c *fiber.Ctx) error {
	var userDTO model.UserDTO
	if err := c.BodyParser(&userDTO); err != nil {
		return c.SendStatus(fiber.StatusUnprocessableEntity)
	}
	id := obj.getIDFromParam(c)
	obj.service.UpdateUser(id, userDTO)
	return c.SendStatus(fiber.StatusOK)
}

func (obj *API) deleteUser(c *fiber.Ctx) error {
	id := obj.getIDFromParam(c)
	obj.service.DeleteUser(id)
	return c.SendStatus(fiber.StatusNoContent)
}
