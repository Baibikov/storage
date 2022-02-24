package fibererr

import (
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v2"
)

type errMessage struct {
	Message string `json:"message"`
	Status int     `json:"status"`
}

func CauseBadRequest(c *fiber.Ctx, err error) error {
	return Err(c, http.StatusBadRequest, err)
}

func CauseNotFound(c *fiber.Ctx, err error) error {
	return Err(c, http.StatusNotFound, err)
}

func Err(c *fiber.Ctx, status int, err error) error {
	if err == nil {
		return nil
	}
	c.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)
	c.Status(status)

	message := err.Error()
	message = strings.ToUpper(string(message[0])) + message[1:]

	return c.JSON(errMessage{
		Message: message,
		Status: status,
	})
}
