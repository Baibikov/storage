package http

import (
	"github.com/gofiber/fiber/v2"

	"storage/internal/app/delivery/http/response"
	"storage/pkg/fibererr"
)

func (h Handler) v1PostFileHandler(c *fiber.Ctx) error {
	file, err := c.FormFile("file")
	if err != nil {
		return fibererr.CauseBadRequest(c, err)
	}

	uid, err := h.service.SaveFile(c.Context(), c.Params("uid"), file)
	if err != nil {
		return fibererr.CauseBadRequest(c, err)
	}

	return fibererr.CauseBadRequest(c, c.JSON(response.V1FilePost{
		UID: uid,
	}))
}
