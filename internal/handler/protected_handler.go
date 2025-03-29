package handler

import "github.com/gofiber/fiber/v2"

func ProtectedHandler(c *fiber.Ctx) error {
	userID := c.Locals("userID")
	return c.JSON(fiber.Map{
		"message": "This is a protected endpoint.",
		"userID":  userID,
	})
}
