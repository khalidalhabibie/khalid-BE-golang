package http

import (
	"gokes/app/fakes/delivery/http/request"
	fakesUsecase "gokes/app/fakes/usecase"
	"gokes/app/models"
	"gokes/pkg/utils"

	fiber "github.com/gofiber/fiber/v2"
)

func Register(c *fiber.Ctx) error {
	log := utils.NewLog()

	log.WithFields(utils.LogFormat(models.LogLayerDelivery, models.LogServiceFakes, "start")).Info("register")

	userDataM, err := utils.ExtractTokenMetadata(c)

	if err != nil {
		log.WithFields(utils.LogFormat(models.LogLayerDelivery, models.LogServiceFakes, err.Error())).Error("failed to extract data user")

		// Return status 401 and error message.
		return utils.ReturnFormat(c, fiber.StatusUnprocessableEntity, true, err.Error(), nil)
	}

	// Create a register struct.
	request := &request.Register{}

	// Checking received data from JSON body.
	if err := c.BodyParser(request); err != nil {
		log.WithFields(utils.LogFormat(models.LogLayerDelivery, models.LogServiceFakes, err.Error())).Error("failed to parse request register")

		// Return status 400 and error message.
		return utils.ReturnFormat(c, fiber.StatusBadRequest, true, err.Error(), nil)
	}

	// Create a new validator for a model.
	validate := utils.NewValidator()

	// Validate sign up fields.
	if err := validate.Struct(request); err != nil {
		log.WithFields(utils.LogFormat(models.LogLayerDelivery, models.LogServiceFakes, err.Error())).Error("error validate request body")

		// Return, if some fields are not valid.
		return utils.ReturnFormat(c, fiber.StatusBadRequest, true, err.Error(), nil)
	}

	fakesM, err := fakesUsecase.Register(userDataM.UserID, *request)
	if err != nil {
		return utils.ReturnFormat(c, fiber.StatusUnprocessableEntity, true, err.Error(), nil)
	}

	dataM, err := utils.MarshalUsers(fakesM, models.AuthRoleNameUser)
	if err != nil {
		log.WithFields(utils.LogFormat(models.LogLayerDelivery, models.LogServiceFakes, err.Error())).Error("error marshal to user")

		// Return, if some fields are not valid.
		return utils.ReturnFormat(c, fiber.StatusBadRequest, true, err.Error(), nil)
	}

	return utils.ReturnFormat(c, fiber.StatusOK, false, nil, dataM)

}
