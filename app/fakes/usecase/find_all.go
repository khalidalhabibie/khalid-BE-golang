package usecase

import (
	"gokes/app/models"
	"gokes/pkg/utils"
	"gokes/platform/database"

	fiber "github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
)

func FindAll(request utils.PaginationConfig) ([]models.Fakes, utils.PaginationMeta, error) {
	meta := utils.PaginationMeta{
		Offset: request.Offset(),
		Limit:  request.Limit(),
		Total:  0,
	}

	db, err := database.OpenDBConnection()
	if err != nil {

		log.WithFields(utils.LogFormat(models.LogLayerUsecase, models.LogServiceFakes, err.Error())).Error(models.LogErrorTypeConnectionDatabase)

		err := fiber.ErrUnprocessableEntity
		return nil, meta, err
	}

	// campaigns, err := u.campaignRepo.FindAll(paginationConfig)
	// if err != nil {
	// 	return nil, meta, err
	// }

	// total, err := u.campaignRepo.Count(paginationConfig)
	// if err != nil {
	// 	return nil, meta, err
	// }

	// meta.Total = total

	fakesM, err := db.FakesRepository.FindAll(request)
	if err != nil {
		log.WithFields(utils.LogFormat(models.LogLayerUsecase, models.LogServiceFakes, err.Error())).Error("error find fakes all fakes by pagination ")

		err := fiber.ErrUnprocessableEntity

		return nil, meta, err
	}

	total, err := db.FakesRepository.Count(request)
	if err != nil {
		log.WithFields(utils.LogFormat(models.LogLayerUsecase, models.LogServiceFakes, err.Error())).Error("error count find fakes all fakes by pagination ")

		err := fiber.ErrUnprocessableEntity

		return nil, meta, err
	}

	meta.Total = total

	return fakesM, meta, nil
}
