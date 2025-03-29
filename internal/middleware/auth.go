package middleware

import (
	"strings"

	"github.com/gofiber/fiber/v2"
)

const StaticToken = "mySecretTok3n" // diganti dengan hasil encoding user ID

func AuthMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Missing Authorization header",
			})
		}

		splitToken := strings.Split(authHeader, "Bearer")
		if len(splitToken) != 2 {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Invalid Authorization header format",
			})
		}

		token := strings.TrimSpace(splitToken[1])
		if token != StaticToken {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Invalid token",
			})
		}

		// Simulasi decode userID dari token
		c.Locals("userID", "123")

		return c.Next()
	}
}
