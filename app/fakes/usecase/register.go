package usecase

import (
	"gokes/app/fakes/delivery/http/request"
	"gokes/app/models"
	"gokes/pkg/helper"
	"gokes/pkg/utils"
	"gokes/platform/database"

	fiber "github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func Register(userID uuid.UUID, request request.Register) (*models.Fakes, error) {

	// log := utils.NewLog()

	// check valid fakes type
	if !helper.ValidateFakes(request.Type) {

		log.WithFields(utils.LogFormat(models.LogLayerUsecase, models.LogServiceFakes, "fakes type invalid")).Error("type fakes validation")

		err := fiber.ErrUnprocessableEntity
		err.Message = "fakes type invalid"
		return nil, err
	}

	db, err := database.OpenDBConnection()
	if err != nil {
		log.WithFields(utils.LogFormat(models.LogLayerUsecase, models.LogServiceFakes, err.Error())).Error(models.LogErrorTypeConnectionDatabase)

		err := fiber.ErrUnprocessableEntity
		return nil, err
	}

	// get id
	newID := uuid.New()
	for {
		_, err := db.FakesRepository.FindByID(newID)
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				break
			} else {
				log.WithFields(utils.LogFormat(models.LogLayerUsecase, models.LogServiceFakes, err.Error())).Error("error to find id fakes")

				err := fiber.ErrUnprocessableEntity
				err.Message = "UnprocessableEntity"

				return nil, err
			}
		}
		newID = uuid.New()
	}

	// get fakes code
	newCode := utils.GenerateFakesCode()

	for {
		_, err = db.FakesRepository.FindByCode(newCode)
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				break
			} else {
				log.WithFields(utils.LogFormat(models.LogLayerUsecase, models.LogServiceFakes, err.Error())).Error("error to find code fakes")

				err := fiber.ErrUnprocessableEntity
				err.Message = "UnprocessableEntity"

				return nil, err
			}
		}
		newCode = utils.GenerateFakesCode()
	}

	fakesM := &models.Fakes{
		Code:       newCode,
		ID:         newID,
		Name:       request.Name,
		Type:       request.Type,
		NakesCount: request.NakesCount,

		CreatedBy: userID,
	}

	isError := utils.ConvertDataDataFakesToPDF(*fakesM)
	if isError {
		log.WithFields(utils.LogFormat(models.LogLayerUsecase, models.LogServiceFakes, err.Error())).Error("convert pdf")

		err := fiber.ErrUnprocessableEntity
		err.Message = "Unprocessable entity"

		return nil, err
	}

	err = db.FakesRepository.Insert(fakesM, nil)
	if err != nil {
		log.WithFields(utils.LogFormat(models.LogLayerUsecase, models.LogServiceFakes, err.Error())).Error("error to insert fakes data")

		err := fiber.ErrUnprocessableEntity
		err.Code = fiber.ErrUnprocessableEntity.Code 
		err.Message = "Failed insert data"

		return nil, err
	}

	return fakesM, nil

}
