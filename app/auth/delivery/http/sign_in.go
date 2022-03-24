package http

import (
	"gokes/app/auth/delivery/http/request"
	authUsecase "gokes/app/auth/usecase"
	"gokes/app/models"
	"gokes/pkg/utils"

	"github.com/gofiber/fiber/v2"
)

func SignIn(c *fiber.Ctx) error {

	log := utils.NewLog()

	log.WithFields(utils.LogFormat(models.LogLayerDelivery, models.LogServiceAuth, "start")).Info("sign in")

	// Create a sign in struct.
	request := &request.SignIn{}

	// Checking received data from JSON body.
	if err := c.BodyParser(request); err != nil {
		log.WithFields(utils.LogFormat(models.LogLayerDelivery, models.LogServiceAuth, err.Error())).Error("failed to parse request")

		// Return status 400 and error message.
		return utils.ReturnFormat(c, fiber.StatusBadRequest, true, err.Error(), nil)
	}

	// Create a new validator for a jamaah model.
	validate := utils.NewValidator()

	// Validate sign up fields.
	if err := validate.Struct(request); err != nil {
		log.WithFields(utils.LogFormat(models.LogLayerDelivery, models.LogServiceAuth, err.Error())).Error("error validate request body")

		// Return, if some fields are not valid.
		return utils.ReturnFormat(c, fiber.StatusBadRequest, true, err.Error(), nil)
	}

	tokenM, err := authUsecase.SignIn(*request)
	if err != nil {
		return utils.ReturnFormat(c, fiber.StatusUnprocessableEntity, true, err.Error(), nil)
	}

	// Return status 200 OK.
	return utils.ReturnFormat(c, fiber.StatusOK, false, nil,
		fiber.Map{
			"access": tokenM.Access,
			// "refresh": tokens.Refresh,
		})

}
