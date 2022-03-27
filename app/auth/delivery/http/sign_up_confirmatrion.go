package http

import (
	"gokes/app/auth/delivery/http/request"
	"gokes/app/models"
	"gokes/pkg/utils"

	authUsecase "gokes/app/auth/usecase"

	fiber "github.com/gofiber/fiber/v2"

	log "github.com/sirupsen/logrus"
)

// sign up
func SignUpConfirmation(c *fiber.Ctx) error {
	log.WithFields(utils.LogFormat(models.LogLayerDelivery, models.LogServiceAuth, "start")).Info("sign up confirmation delivery")

	// Create a new jamaah auth struct.
	request := &request.SignUpConfirmation{}

	// Checking received data from JSON body.
	if err := c.BodyParser(request); err != nil {
		log.WithFields(utils.LogFormat(models.LogLayerDelivery, models.LogServiceAuth, err.Error())).Error("failed to parse request")

		// Return status 400 and error message.
		return utils.ReturnFormat(c, fiber.StatusBadRequest, true, err.Error(), nil)
	}

	// Create a new validator for a Author model.
	validate := utils.NewValidator()

	// Validate sign up fields.
	if err := validate.Struct(request); err != nil {
		log.WithFields(utils.LogFormat(models.LogLayerDelivery, models.LogServiceAuth, err.Error())).Error("Error Validate Sign Up Body")

		// Return, if some fields are not valid.
		return utils.ReturnFormat(c, fiber.StatusBadRequest, true, err.Error(), nil)
	}

	err := authUsecase.SignUpConfirmation(*request)
	if err != nil {
		return utils.ReturnFormat(c, fiber.StatusUnprocessableEntity, true, err.Error(), nil)
	}

	log.WithFields(utils.LogFormat(models.LogLayerDelivery, models.LogServiceAuth, "end")).Info("sign up confirmation delvery")

	return utils.ReturnFormat(c, fiber.StatusOK, false, nil, "OK")

}
