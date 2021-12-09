package api

import (
	"github.com/clumpapp/clump-be/model"

	"github.com/gofiber/fiber/v2"
)

func (obj *API) getmessages(c *fiber.Ctx) error {
	var groupDTO model.GroupDTO
	if err := c.BodyParser(&groupDTO); err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	messagesDTO := obj.service.GetMessages(groupDTO)
	return c.JSON(messagesDTO)
}

func (obj *API) getusers(c *fiber.Ctx) error {
	var groupDTO model.GroupDTO
	if err := c.BodyParser(&groupDTO); err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	usersDTO := obj.service.GetUsers(groupDTO)
	return c.JSON(usersDTO)
}

func (obj *API) updategroup(c *fiber.Ctx) error {
	var groupDTO model.GroupDTO
	if err := c.BodyParser(&groupDTO); err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	obj.service.UpdateGroup(groupDTO)
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
