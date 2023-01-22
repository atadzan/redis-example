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

	app.Post("/:id/add", h.add)
	app.Get("/all", h.getAll)
	//app.Get("/all", h.getAll)
	//app.Delete("/:key", h.removeSet)
	//app.Put("/:key", h.updateSet)

}
