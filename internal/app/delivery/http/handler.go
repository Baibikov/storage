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

	// folder handlers
	v1.Get("/folders/:uid", handler.v1GetFolderHandler)
	v1.Delete("/folders/:uid", handler.v1DeleteFolderHandler)
	v1.Post("/folders", handler.v1PostFolderHandler)
	v1.Get("/folders/:uid/:level", handler.v1GetFolderDirectoryHandler)

	// file handlers
	v1.Post("/folders/file/:uid", handler.v1PostFileHandler)
	return handler
}

func (h *Handler) Listen(addr string) error {
	return h.app.Listen(addr)
}
