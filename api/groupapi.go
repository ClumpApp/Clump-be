package api

import (
	"github.com/clumpapp/clump-be/model"

	"github.com/gofiber/fiber/v2"
)

func (obj *API) putgroup(c *fiber.Ctx) error {
	var groupDTO model.GroupDTO
	if err := c.BodyParser(&groupDTO); err != nil {
		return c.SendStatus(fiber.StatusUnprocessableEntity)
	}
	id := obj.getIDFromParam(c)
	obj.service.UpdateGroup(id, groupDTO)
	return c.SendStatus(fiber.StatusOK)
}

func (obj *API) deletegroup(c *fiber.Ctx) error {
	id := obj.getIDFromParam(c)
	obj.service.DeleteGroup(id)
	return c.SendStatus(fiber.StatusNoContent)
}
