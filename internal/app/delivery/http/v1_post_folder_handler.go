package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"

	"storage/internal/app/delivery/http/response"
	"storage/internal/app/types"
	"storage/pkg/fibererr"
)

type FolderRequest struct {
	Name   string `json:"name"`
	Parent int    `json:"parent"`
}

func (h Handler) v1PostFolderHandler(c *fiber.Ctx) error {
	fr := FolderRequest{}

	err := c.BodyParser(&fr)
	if err != nil {
		return fibererr.CauseBadRequest(
			c,
			errors.Wrap(err, "parse body information"),
		)
	}

	uid, err := h.service.CreateFolder(c.Context(), types.Folder{
		Name:  fr.Name,
		Level: fr.Parent,
	})
	if err != nil {
		return fibererr.CauseBadRequest(c, err)
	}

	return fibererr.CauseBadRequest(c, c.JSON(response.V1FolderPost{
		UID: uid,
	}))
}
