package http

import (
	"github.com/gofiber/fiber/v2"

	"storage/internal/app/service"
)

type Handler struct {
	app *fiber.App
	service *service.UseCase
}

func New(service *service.UseCase) *Handler {
	handler := &Handler{
		app: fiber.New(),
		service: service,
	}

	v1 := handler.app.Group("/api/v1")
	v1.Get("/folder", handler.v1GetFolderHandler)
	v1.Post("/folder", handler.v1PostFolderHandler)

	return handler
}

func (h *Handler) Listen(addr string) error {
	return h.app.Listen(addr)
}