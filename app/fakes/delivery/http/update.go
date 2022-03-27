package http

import (
	"gokes/app/fakes/delivery/http/request"
	fakesUsecase "gokes/app/fakes/usecase"
	"gokes/app/models"
	"gokes/pkg/utils"

	fiber "github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
)

func Update(c *fiber.Ctx) error {
	log.WithFields(utils.LogFormat(models.LogLayerDelivery, models.LogServiceFakes, "start")).Info("update fakes")

	code := c.Params("code")

	// Create a update struct.
	request := &request.Update{}

	// Checking received data from JSON body.
	if err := c.BodyParser(request); err != nil {
		log.WithFields(utils.LogFormat(models.LogLayerDelivery, models.LogServiceFakes, err.Error())).Error("failed to parse request update")

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

	fakesM, err := fakesUsecase.Update(code, *request)
	if err != nil {
		log.WithFields(utils.LogFormat(models.LogLayerDelivery, models.LogServiceFakes, err.Error())).Error("error update fakes")

		utils.ReturnFormat(c, fiber.StatusNotFound, true, err.Error(), nil)
	}

	dataM, err := utils.MarshalUsers(fakesM, models.AuthRoleNameUser)
	if err != nil {
		log.WithFields(utils.LogFormat(models.LogLayerDelivery, models.LogServiceAuth, err)).Error("Error Marsal Body get data fakes by code for consume")

		return utils.ReturnFormat(c, fiber.StatusUnprocessableEntity, true, err, nil)
	}

	log.WithFields(utils.LogFormat(models.LogLayerDelivery, models.LogServiceFakes, "end")).Info("update fakes")

	return utils.ReturnFormat(c, fiber.StatusOK, false, nil, dataM)

}
