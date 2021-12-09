package api

import (
	"github.com/clumpapp/clump-be/model"

	"github.com/gofiber/fiber/v2"
)

func (obj *API) getgroupmessages(c *fiber.Ctx) error {
	id := obj.getID(c)
	messagesDTO := obj.service.GetGroupMessages(id)
	return c.JSON(messagesDTO)
}

func (obj *API) postmessage(c *fiber.Ctx) error {
	var messageDTO model.MessageDTO
	if err := c.BodyParser(&messageDTO); err != nil {
		return c.SendStatus(fiber.StatusUnprocessableEntity)
	}
	obj.service.CreateMessage(messageDTO)
	return c.SendStatus(fiber.StatusOK)
}

func (obj *API) putmessage(c *fiber.Ctx) error {
	var messageDTO model.MessageDTO
	if err := c.BodyParser(&messageDTO); err != nil {
		return c.SendStatus(fiber.StatusUnprocessableEntity)
	}
	id := obj.getID(c)
	obj.service.UpdateMessage(id, messageDTO)
	return c.SendStatus(fiber.StatusOK)
}

func (obj *API) deletemessage(c *fiber.Ctx) error {
	id := obj.getID(c)
	obj.service.DeleteMessage(id)
	return c.SendStatus(fiber.StatusOK)
}
