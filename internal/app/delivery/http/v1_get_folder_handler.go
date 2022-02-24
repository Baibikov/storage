package http

import (
	"time"

	"github.com/gofiber/fiber/v2"

	"storage/internal/app/delivery/http/response"
	"storage/pkg/fibererr"
)

func (h Handler) v1GetFolderHandler(c *fiber.Ctx) error {
	return fibererr.CauseBadRequest(c, c.JSON(response.V1FolderGet{
		Name: "Messages",
		CreatedDate: time.Now(),
	}))
}