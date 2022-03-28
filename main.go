package main

import (
	"os"

	"gokes/pkg/configs"
	"gokes/pkg/middleware"
	"gokes/pkg/routes"
	"gokes/pkg/utils"

	// "github.com/gofiber/fiber/v2"
	fiber "github.com/gofiber/fiber/v2"

	// _ "gokes/docs" // load API Docs files (Swagger)

	_ "github.com/joho/godotenv/autoload" // load .env file automatically
)

func main() {

	// Define Fiber config.
	config := configs.FiberConfig()

	// Define a new Fiber app with config.
	app := fiber.New(config)

	// Middlewares.
	middleware.FiberMiddleware(app) // Register Fiber's middleware for app.

	// Routes.
	routes.AuthRoutes(app)
	routes.FakesRoutes(app)
	routes.NotFoundRoute(app) // Register route for 404 Error.

	// Start server (with or without graceful shutdown).
	if os.Getenv("STAGE_STATUS") == "dev" {
		utils.StartServer(app)
	} else {
		utils.StartServerWithGracefulShutdown(app)
	}
}
