package handlers

import (
	"github.com/gofiber/fiber/v2"
	"redis-example/pkg/services"
)

type Handler struct {
	Service *services.Service
}

func NewHandler(services *services.Service) *Handler {
	return &Handler{
		Service: services,
	}
}

func (h *Handler) InitRoutes(app *fiber.App) {

	app.Post("/:score/:key/add", h.add)
	app.Get("/:key/all", h.getAll)
	app.Get("/:key", h.getByKey)
	app.Delete("/:set/:key/remove", h.removeElemFromSet)

}
