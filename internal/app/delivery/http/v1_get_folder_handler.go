package http

import (
	"github.com/gofiber/fiber/v2"

	"storage/internal/app/delivery/http/response"
	"storage/pkg/fibererr"
)

func (h Handler) v1GetFolderHandler(c *fiber.Ctx) error {
	folder, err := h.service.GetFolder(c.Context(), c.Params("uid"))
	if err != nil {
		return fibererr.CauseBadRequest(c, err)
	}

	return fibererr.CauseBadRequest(c, c.JSON(response.V1FolderGet{
		UID:         folder.UID,
		Name:        folder.Name,
		Level:       folder.Level,
		CreatedDate: folder.CreatedAt,
	}))
}
