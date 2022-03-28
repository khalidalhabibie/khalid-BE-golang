package http

import (
	"gokes/app/models"
	"gokes/pkg/utils"

	fakesUsecase "gokes/app/fakes/usecase"

	fiber "github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
)

func FindByCode(c *fiber.Ctx) error {

	log.WithFields(utils.LogFormat(models.LogLayerDelivery, models.LogServiceFakes, "start")).Info("get by code")

	code := c.Params("code")
	fakesM, err := fakesUsecase.FindByCode(code)
	if err != nil {
		return utils.ReturnFormat(c, fiber.StatusNotFound, true, "data not found", nil)
	}

	dataM, err := utils.MarshalUsers(fakesM, models.AuthRoleNameUser)
	if err != nil {
		log.WithFields(utils.LogFormat(models.LogLayerDelivery, models.LogServiceAuth, err)).Error("Error Marsal Body get data fakes by code for consume")

		return utils.ReturnFormat(c, fiber.StatusUnprocessableEntity, true, err, nil)

	}

	log.WithFields(utils.LogFormat(models.LogLayerDelivery, models.LogServiceFakes, "end")).Info("get by code")

	return utils.ReturnFormat(c, fiber.StatusOK, false, nil, dataM)


}
