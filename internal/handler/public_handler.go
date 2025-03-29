package handler

import "github.com/gofiber/fiber/v2"

func PublicHandler(c *fiber.Ctx) error {
	resp := fiber.Map{
		"message": "Hello World",
	}

	return c.JSON(resp)
}
