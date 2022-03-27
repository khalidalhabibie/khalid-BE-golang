package routes

import (
	fakesDelivery "gokes/app/fakes/delivery/http"
	"gokes/pkg/middleware"

	fiber "github.com/gofiber/fiber/v2"
)

// per service

// fakesRoutes func for describe group of fakes routes.
func FakesRoutes(a *fiber.App) {
	route := a.Group("/fakes")

	// user
	user := route.Group("/user")
	user.Post("", middleware.JWTProtectedUser(), fakesDelivery.Register)
	user.Patch("/:code", middleware.JWTProtectedUser(), fakesDelivery.Update)
	user.Delete("/:code", middleware.JWTProtectedUser(), fakesDelivery.Delete)
	user.Get("/:code", middleware.JWTProtectedUser(), fakesDelivery.FindByCode)
	user.Get("/download/:code", middleware.JWTProtectedUser(), fakesDelivery.DownloadByCode)
}
