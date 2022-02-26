package http

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"

	"storage/internal/app/delivery/http/response"
	"storage/pkg/fibererr"
)

func (h Handler) v1GetFolderDirectoryHandler(c *fiber.Ctx) error {
	level, err := strconv.Atoi(c.Params("level"))
	if err != nil {
		return fibererr.CauseBadRequest(c, errors.Wrapf(
			err,
			"converting symbol to int %s",
			c.Params("level"),
		))
	}

	folders, err := h.service.GetFolderDirectory(c.Context(), c.Params("uid"), level)
	if err != nil {
		return fibererr.CauseBadRequest(c, err)
	}

	res := make(response.V1FolderDirectoryGet, 0, len(folders))

	for _, f := range folders {
		res = append(res, response.V1FolderGet{
			UID:         f.UID,
			Name:        f.Name,
			Level:       level,
			CreatedDate: f.CreatedAt,
		})
	}

	return fibererr.CauseBadRequest(c, c.JSON(res))
}
