package handlers

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"redis-example/models"
	"strconv"
)

const elemPerPage = 5

func (h *Handler) add(c *fiber.Ctx) error {
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

func (h *Handler) getAll(c *fiber.Ctx) error {
	page, err := strconv.Atoi(c.Query("page", "1"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid page param"})
	}
	var offset int64
	offset = int64((page - 1) * elemPerPage)
	ctx := context.Background()
	result, resErr := h.Service.User.Get(ctx, offset)
	if resErr != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": resErr.Error()})
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
