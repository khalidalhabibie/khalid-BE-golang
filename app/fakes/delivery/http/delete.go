package http

import (
	fakesUsecase "gokes/app/fakes/usecase"
	"gokes/app/models"
	"gokes/pkg/utils"

	fiber "github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
)

func Delete(c *fiber.Ctx) error {
	log.WithFields(utils.LogFormat(models.LogLayerDelivery, models.LogServiceFakes, "start")).Info("delete fakes")

	code := c.Params("code")

	fakesM, err := fakesUsecase.Delete(code)
	if err != nil {
		log.WithFields(utils.LogFormat(models.LogLayerDelivery, models.LogServiceFakes, err.Error())).Error("error update fakes")

		return utils.ReturnFormat(c, fiber.ErrUnprocessableEntity.Code, true, err.Error(), nil)
	}

	dataM, err := utils.MarshalUsers(fakesM, models.AuthRoleNameUser)
	if err != nil {
		log.WithFields(utils.LogFormat(models.LogLayerDelivery, models.LogServiceAuth, err)).Error("Error Marsal Body get data fakes by code for consume")

		return utils.ReturnFormat(c, fiber.StatusUnprocessableEntity, true, err, nil)
	}

	log.WithFields(utils.LogFormat(models.LogLayerDelivery, models.LogServiceFakes, "end")).Info("delete fakes")

	return utils.ReturnFormat(c, fiber.StatusOK, false, nil, dataM)

}
