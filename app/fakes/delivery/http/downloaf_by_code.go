package http

import (
	"fmt"

	fiber "github.com/gofiber/fiber/v2"
)

func Download(c *fiber.Ctx) error {

	code := c.Params("code")

	return c.SendFile(fmt.Sprintf("data/%v.pdf", code), true)

}
