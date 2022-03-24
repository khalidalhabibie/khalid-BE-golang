package utils

import "github.com/gofiber/fiber/v2"

func ReturnFormat(c *fiber.Ctx, httpCode int, isError bool, err, data interface{}) error {
	return c.Status(httpCode).JSON(fiber.Map{
		"http_code": httpCode,
		"is_error":  isError,
		"message":   err,
		"data":      data,
	})
}
