package api

import (
	"github.com/clumpapp/clump-be/model"

	"github.com/gofiber/fiber/v2"
)

func (obj *API) createInterest(c *fiber.Ctx) error {
	var interestDTO model.InterestDTO
	if err := c.BodyParser(&interestDTO); err != nil {
		return c.SendStatus(fiber.StatusUnprocessableEntity)
	}
	obj.service.CreateInterest(interestDTO)
	return c.SendStatus(fiber.StatusCreated)
}

func (obj *API) getInterests(c *fiber.Ctx) error {
	interestsDTO := obj.service.GetInterests()
	return c.JSON(interestsDTO)
}

func (obj *API) addInterests(c *fiber.Ctx) error {
	var interestDTOs []model.InterestDTO
	if err := c.BodyParser(&interestDTOs); err != nil {
		return c.SendStatus(fiber.StatusUnprocessableEntity)
	}
	uid := obj.getUserIDFromToken(c)
	obj.service.AddInterests(interestDTOs, uint(uid))
	return c.SendStatus(fiber.StatusCreated)
}
