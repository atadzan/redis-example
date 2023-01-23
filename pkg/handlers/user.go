package handlers

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"math"
	"redis-example/models"
	"strconv"
)

const elemPerPage = 2

func (h *Handler) add(c *fiber.Ctx) error {
	key := c.Params("key")
	score, err := strconv.Atoi(c.Params("score"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	var user models.User
	if err = c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	ctx := context.Background()

	result, rErr := h.Service.User.Add(ctx, float64(score), key, user)
	if rErr != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": rErr.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"Added": result})
}

func (h *Handler) getAll(c *fiber.Ctx) error {
	key := c.Params("key")
	page, err := strconv.Atoi(c.Query("page", "1"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid page param"})
	}
	ctx := context.Background()
	total, err := h.Service.User.CountSetElems(ctx, key)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	pageCount := int(math.Ceil(float64(total) / float64(elemPerPage)))
	if page < 1 || page > pageCount {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"Result": "no elems"})
	}
	var offset int64
	offset = int64((page - 1) * elemPerPage)
	result, resErr := h.Service.User.GetAll(ctx, key, offset, int64(elemPerPage))
	if resErr != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": resErr.Error()})
	}
	if pageCount > page {
		page++
	} else {
		page = 0
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"Result": result, "Next page": page})
}

func (h *Handler) getByKey(c *fiber.Ctx) error {
	key := c.Params("key")
	ctx := context.Background()
	result, err := h.Service.User.GetByKey(ctx, key)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"Result": result})
}

func (h *Handler) removeElemFromSet(c *fiber.Ctx) error {
	set := c.Params("set")
	key := c.Params("key")
	ctx := context.Background()
	result, err := h.Service.User.RemoveElemFromSet(ctx, set, key)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"Result": result})
}
