package handlers

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"redis-example/models"
)

func (h *Handler) addSet(c *fiber.Ctx) error {
	id := c.Params("id")
	var user models.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	ctx := context.Background()

	result, err := h.Service.User.Add(ctx, id, user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"Result": result})
}

func (h *Handler) getByKey(c *fiber.Ctx) error {
	id := c.Params("/:id")
	ctx := context.Background()
	result, err := h.Service.User.GetById(ctx, id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"Result": result})
}

//
//func (h *Handler) getAll(c *fiber.Ctx) error {
//
//}
//
//func (h *Handler) removeSet(c *fiber.Ctx) error {
//
//}
//
//func (h *Handler) updateSet(c *fiber.Ctx) error {
//
//}
