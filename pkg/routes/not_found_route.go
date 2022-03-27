package routes

import (
	"gokes/pkg/utils"

	fiber "github.com/gofiber/fiber/v2"
)

// NotFoundRoute func for describe 404 Error route.
func NotFoundRoute(a *fiber.App) {
	// Register new special route.
	a.Use(
		// Anonymous function.
		func(c *fiber.Ctx) error {
			// Return HTTP 404 status and JSON response.
			return utils.ReturnFormat(c, fiber.StatusNotFound, true, "sorry, endpoint is not found", nil)
		},
	)
}
