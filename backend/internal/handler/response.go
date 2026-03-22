package handler

import "github.com/gofiber/fiber/v2"

func Success(c *fiber.Ctx, status int, data interface{}, message string) error {
	return c.Status(status).JSON(fiber.Map{
		"data":    data,
		"message": message,
	})
}

func Error(c *fiber.Ctx, status int, code string, message string) error {
	return c.Status(status).JSON(fiber.Map{
		"error": message,
		"code":  code,
	})
}
