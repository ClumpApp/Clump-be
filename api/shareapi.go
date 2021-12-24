package api

import (
	"github.com/clumpapp/clump-be/model"

	"github.com/gofiber/fiber/v2"
)

func (obj *API) getgroupmessages(c *fiber.Ctx) error {
	id := obj.getGroupIDFromToken(c)
	messagesDTO := obj.service.GetGroupMessages(id)
	return c.JSON(messagesDTO)
}

func (obj *API) postmessage(c *fiber.Ctx) error {
	var messageInDTO model.MessageInDTO
	if err := c.BodyParser(&messageInDTO); err != nil {
		return c.SendStatus(fiber.StatusUnprocessableEntity)
	}
	gid := obj.getGroupIDFromToken(c)
	uid := obj.getUserIDFromToken(c)
	obj.service.CreateMessage(gid, uid, messageInDTO)
	return c.SendStatus(fiber.StatusCreated)
}

func (obj *API) deletemessage(c *fiber.Ctx) error {
	id := obj.getIDFromParam(c)
	obj.service.DeleteMessage(id)
	return c.SendStatus(fiber.StatusNoContent)
}
