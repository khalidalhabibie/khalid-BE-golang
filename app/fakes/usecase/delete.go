package usecase

import (
	"fmt"
	"gokes/app/models"
	"gokes/pkg/utils"
	"gokes/platform/database"

	fiber "github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
)

func Delete(code string) (*models.Fakes, error) {
	// find fakes by code
	db, err := database.OpenDBConnection()
	if err != nil {
		log.WithFields(utils.LogFormat(models.LogLayerUsecase, models.LogServiceFakes, err.Error())).Error(models.LogErrorTypeConnectionDatabase)

		err := fiber.ErrUnprocessableEntity
		return nil, err
	}

	fakesM, err := db.FakesRepository.FindByCode(code)
	if err != nil {
		log.WithFields(utils.LogFormat(models.LogLayerUsecase, models.LogServiceFakes, err.Error())).Error(fmt.Sprintf("error find fakes by code %v ", code))

		err := fiber.ErrNotFound
		err.Code = fiber.ErrNotFound.Code
		err.Message = "data not found"

		return nil, err
	}

	err = db.FakesRepository.Delete(fakesM, nil)
	if err != nil {
		log.WithFields(utils.LogFormat(models.LogLayerUsecase, models.LogServiceFakes, err.Error())).Error(fmt.Sprintf("error deleted fakes by code %v ", code))

		err := fiber.ErrUnprocessableEntity
		err.Code = fiber.ErrUnprocessableEntity.Code

		return nil, err
	}

	return nil, err

}
