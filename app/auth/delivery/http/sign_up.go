package http

import (
	"gokes/app/auth/delivery/http/request"
	"gokes/app/models"
	"gokes/pkg/utils"

	authUsecase "gokes/app/auth/usecase"

	fiber "github.com/gofiber/fiber/v2"
)

// sign up
func SignUp(c *fiber.Ctx) error {

	log := utils.NewLog()

	log.WithFields(utils.LogFormat(models.LogLayerDelivery, models.LogServiceAuth, "start")).Info("sign up delvery")

	// Create a new jamaah auth struct.
	request := &request.SignUp{}

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

	userM, err := authUsecase.SignUp(*request)
	if err != nil {
		return utils.ReturnFormat(c, fiber.StatusUnprocessableEntity, true, err.Error(), nil)
	}

	log.WithFields(utils.LogFormat(models.LogLayerDelivery, models.LogServiceAuth, "end")).Info("sign up delvery")

	// filter by data by auth
	data, err := utils.MarshalUsers(userM, models.AuthRoleNameUser)
	if err != nil {
		log.WithFields(utils.LogFormat(models.LogLayerDelivery, models.LogServiceAuth, err)).Error("Error Marsal Body user after sign up")

		return utils.ReturnFormat(c, fiber.StatusUnprocessableEntity, true, err, nil)

	}

	return utils.ReturnFormat(c, fiber.StatusOK, false, nil, data)

}
