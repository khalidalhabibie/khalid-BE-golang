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
	user.Post(":code", middleware.JWTProtectedUser(), fakesDelivery.FindByCode)

}
