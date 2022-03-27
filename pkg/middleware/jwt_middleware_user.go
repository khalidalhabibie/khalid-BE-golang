package middleware

import (
	"gokes/pkg/utils"
	"os"

	fiber "github.com/gofiber/fiber/v2"

	jwtMiddleware "github.com/gofiber/jwt/v2"
)

// JWTProtected func for specify routes group with JWT authentication.
func JWTProtectedUser() func(*fiber.Ctx) error {
	// Create config for JWT authentication middleware.
	config := jwtMiddleware.Config{
		SigningKey:    []byte(os.Getenv("JWT_SECRET_KEY_USER")),
		ContextKey:    "jwt",
		ErrorHandler:  jwtUserError,
		SigningMethod: "HS512",
	}

	return jwtMiddleware.New(config)
}

func jwtUserError(c *fiber.Ctx, err error) error {

	// extract jwt and 


	// Return status 401 and failed authentication error.
	if err.Error() == "Missing or malformed JWT" {
		return utils.ReturnFormat(c, fiber.StatusBadRequest, true, err.Error(), nil)
	}

	// Return status 401 and failed authentication error.
	return utils.ReturnFormat(c, fiber.StatusUnauthorized, true, err.Error(), nil)
}
