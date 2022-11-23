package handlers

import (
	"github.com/gofiber/fiber/v2"
)

func NotFound(c *fiber.Ctx, err error) error {
	return c.Status(fiber.StatusNotFound).Render("layouts/not_found", nil)
}
