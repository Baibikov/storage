package http

import (
	"github.com/gofiber/fiber/v2"

	"storage/internal/app/delivery/http/response"
	"storage/pkg/fibererr"
)

func (h Handler) v1DeleteFolderHandler(c *fiber.Ctx) error {
	err := h.service.DeleteFolder(c.Context(), c.Params("uid"))
	if err != nil {
		return fibererr.CauseBadRequest(c, err)
	}

	return fibererr.CauseBadRequest(c, c.JSON(response.V1FolderDelete{
		Message: "ok",
	}))
}
