package api

import (
	"github.com/clumpapp/clump-be/model"

	"github.com/gofiber/fiber/v2"
)

func (obj *API) textshare(c *fiber.Ctx) error {
	var messageDTO model.MessageDTO
	if err := c.BodyParser(&messageDTO); err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	obj.service.TextShare(messageDTO)
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

func (obj *API) deletemessage(c *fiber.Ctx) error {
	var messageDTO model.MessageDTO
	if err := c.BodyParser(&messageDTO); err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	obj.service.DeleteMessage(messageDTO)
	return c.SendStatus(fiber.StatusOK)
}
