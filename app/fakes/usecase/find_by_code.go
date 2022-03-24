package fakes

import (
	"fmt"
	"gokes/app/models"
	"gokes/pkg/utils"
	"gokes/platform/database"

	fiber "github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
)

func FindByCode(code string) (*models.Fakes, error) {

	db, err := database.OpenDBConnection()
	if err != nil {

		log.WithFields(utils.LogFormat(models.LogLayerUsecase, models.LogServiceFakes, err.Error())).Error(models.LogErrorTypeConnectionDatabase)

		err := fiber.ErrUnprocessableEntity
		return nil, err
	}

	fakesM, err := db.FakesRepository.FindByCode(code)
	if err != nil {

		log.WithFields(utils.LogFormat(models.LogLayerUsecase, models.LogServiceFakes, err.Error())).Error(fmt.Sprintf("error find fakes by code %v ", code))

		err := fiber.ErrUnprocessableEntity
		return nil, err
	}

	return fakesM, nil

}
