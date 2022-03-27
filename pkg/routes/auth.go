package routes

import (
	authDelivery "gokes/app/auth/delivery/http"

	fiber "github.com/gofiber/fiber/v2"
)

// per service

// CompanyRoutes func for describe group of company routes.
func AuthRoutes(a *fiber.App) {
	route := a.Group("/auth")

	// public
	public := route.Group("/public")

	public.Post("/sign/up", authDelivery.SignUp)
	public.Post("/sign/up/confirmation", authDelivery.SignUpConfirmation)
	public.Post("/sign/in", authDelivery.SignIn)
	public.Post("/sign/in/confirmation", authDelivery.SignInConfirmaton)

}
