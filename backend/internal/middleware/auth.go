package middleware

import (
	"strings"

	"github.com/gofiber/fiber/v2"

	"github.com/dogannx/SDU_RandevuSistemi/backend/pkg/jwt"
)

func AuthRequired(jwtSecret string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Token gerekli",
				"code":  "UNAUTHORIZED",
			})
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Geçersiz token formatı",
				"code":  "UNAUTHORIZED",
			})
		}

		claims, err := jwt.ValidateToken(parts[1], jwtSecret)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Token geçersiz veya süresi dolmuş",
				"code":  "UNAUTHORIZED",
			})
		}

		c.Locals("userID", claims.UserID)
		c.Locals("email", claims.Email)

		return c.Next()
	}
}
