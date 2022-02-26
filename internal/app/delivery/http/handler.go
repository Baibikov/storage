package http

import (
	"github.com/gofiber/fiber/v2"

	"storage/internal/app/service"
)

type Handler struct {
	app     *fiber.App
	service *service.UseCase
}

func New(service *service.UseCase) *Handler {
	handler := &Handler{
		app:     fiber.New(),
		service: service,
	}

	v1 := handler.app.Group("/api/v1")
	v1.Get("/folder/:uid", handler.v1GetFolderHandler)
	v1.Delete("/folder/:uid", handler.v1DeleteFolderHandler)
	v1.Post("/folder", handler.v1PostFolderHandler)
	v1.Get("/folder/directory/:uid/:level", handler.v1GetFolderDirectoryHandler)
	return handler
}

func (h *Handler) Listen(addr string) error {
	return h.app.Listen(addr)
}
