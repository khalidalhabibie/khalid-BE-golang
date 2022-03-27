package usecase

import (
	"gokes/app/auth/delivery/http/request"
	"gokes/app/models"
	"gokes/pkg/utils"
	"gokes/platform/database"

	fiber "github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
)

func SignIn(request request.SignIn) (*models.Tokens, error) { // Create database connection.
	// log := utils.NewLog()

	db, err := database.OpenDBConnection()
	if err != nil {

		log.WithFields(utils.LogFormat(models.LogLayerUsecase, models.LogServiceAuth, err.Error())).Error(models.LogErrorTypeConnectionDatabase)

		err := fiber.ErrUnprocessableEntity
		return nil, err
	}

	// find user by email
	userM, err := db.UserRepository.FindByUsername(request.Username)
	if err != nil {
		log.WithFields(utils.LogFormat(models.LogLayerUsecase, models.LogServiceAuth, err)).Error("user not found")

		err := fiber.ErrNotFound
		err.Message = "User Not Found"

		return nil, err

	}

	if !utils.ComparePasswords(userM.Password, request.Password) {
		log.WithFields(utils.LogFormat(models.LogLayerUsecase, models.LogServiceAuth, "password not same")).Error("password not same")

		err := fiber.ErrUnprocessableEntity
		err.Message = "Please check your email or password didn't match"

		return nil, err
	}

	// Generate a new pair of access and refresh tokens.
	token, err := utils.GenerateNewTokens(userM.ID, userM.Email)
	if err != nil {
		log.WithFields(utils.LogFormat(models.LogLayerUsecase, models.LogServiceAuth, err.Error())).Error("error generate token")

		err := fiber.ErrUnprocessableEntity
		err.Message = "Please check your email or password didn't match"

		return nil, err

	}

	return token, nil
}
