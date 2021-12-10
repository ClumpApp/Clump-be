package api

import (
	"strconv"

	"github.com/clumpapp/clump-be/middleware"
	"github.com/clumpapp/clump-be/model"

	"github.com/gofiber/fiber/v2"
)

func (obj *API) login(c *fiber.Ctx) error {
	var loginDTO model.LoginDTO
	if err := c.BodyParser(&loginDTO); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	matches, user := obj.service.Login(loginDTO)
	if matches {
		uid := strconv.FormatUint(uint64(user.ID), 10)
		gid := strconv.FormatUint(uint64(user.GroupID), 10)
		token := middleware.GetToken(uid, gid)
		c.Cookie(&fiber.Cookie{Name: middleware.Token, Value: token})
		c.SendStatus(fiber.StatusOK)
	}
	return c.SendStatus(fiber.StatusUnauthorized)
}

func (obj *API) signup(c *fiber.Ctx) error {
	var userDTO model.UserDTO
	if err := c.BodyParser(&userDTO); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	obj.service.SignUp(userDTO)
	return c.SendStatus(fiber.StatusOK)
}

func (obj *API) getgroupusers(c *fiber.Ctx) error {
	id := obj.getGroupIDFromToken(c)
	usersDTO := obj.service.GetGroupUsers(id)
	return c.JSON(usersDTO)
}

func (obj *API) putuser(c *fiber.Ctx) error {
	var userDTO model.UserDTO
	if err := c.BodyParser(&userDTO); err != nil {
		return c.SendStatus(fiber.StatusUnprocessableEntity)
	}
	id := obj.getIDFromParam(c)
	obj.service.UpdateUser(id, userDTO)
	return c.SendStatus(fiber.StatusOK)
}

func (obj *API) deleteuser(c *fiber.Ctx) error {
	id := obj.getIDFromParam(c)
	obj.service.DeleteUser(id)
	return c.SendStatus(fiber.StatusNoContent)
}
