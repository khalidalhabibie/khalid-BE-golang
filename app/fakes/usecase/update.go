package usecase

import (
	"fmt"
	"gokes/app/fakes/delivery/http/request"
	"gokes/app/models"
	"gokes/pkg/helper"
	"gokes/pkg/utils"
	"gokes/platform/database"

	fiber "github.com/gofiber/fiber/v2"
	"github.com/jinzhu/copier"
	log "github.com/sirupsen/logrus"
)

func Update(code string, request request.Update) (*models.Fakes, error) {

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

	// copier
	err = copier.Copy(fakesM, &request)
	if err != nil {
		log.WithFields(utils.LogFormat(models.LogLayerUsecase, models.LogServiceFakes, err.Error())).Error(fmt.Sprintf("error find fakes by code %v ", code))

		err := fiber.ErrUnprocessableEntity
		err.Code = fiber.ErrUnprocessableEntity.Code
		err.Message = "data not found"

		return nil, err
	}

	// validate
	if !helper.ValidateFakes(fakesM.Type) {

		log.WithFields(utils.LogFormat(models.LogLayerUsecase, models.LogServiceFakes, "fakes type invalid")).Error("type fakes validation")

		err := fiber.ErrUnprocessableEntity
		err.Message = "Fakes type invalid"
		return nil, err
	}

	// update
	err = db.FakesRepository.Insert(fakesM, nil)
	if err != nil {
		log.WithFields(utils.LogFormat(models.LogLayerUsecase, models.LogServiceFakes, err.Error())).Error("error to update fakes data")

		err := fiber.ErrUnprocessableEntity
		err.Message = "Failed update data"

		return nil, err
	}

	return fakesM, nil

}
