package usecase

import (
	"context"
	"fmt"
	"gokes/app/auth/delivery/http/request"
	"gokes/app/models"
	"gokes/pkg/utils"
	"gokes/platform/cache"
	"gokes/platform/database"

	fiber "github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
)

func SignInConfirmation(request request.SignInConfirmation) (*models.Tokens, error) {

	// connection redis and and key
	connRedis, err := cache.RedisConnection()
	if err != nil {
		log.WithFields(utils.LogFormat(models.LogLayerUsecase, models.LogServiceAuth, err.Error())).Error(models.LogErrorTypeConnectionRedis)

		err := fiber.ErrUnprocessableEntity
		err.Message = "UnprocessableEntity"

		return nil, err
	}

	// get data
	keySignIn := fmt.Sprintf("%v-%v", models.AuthUserOTPSignIn, request.Username)

	dataFromRedis, err := connRedis.Get(context.Background(), keySignIn).Result()
	if err != nil {
		log.WithFields(utils.LogFormat(models.LogLayerUsecase, models.LogServiceAuth, err.Error())).Error("failed to get data from redis, key : ", keySignIn)

		err := fiber.ErrUnprocessableEntity
		err.Message = "UnprocessableEntity"

		return nil, err
	}

	if dataFromRedis != request.Code {
		log.WithFields(utils.LogFormat(models.LogLayerUsecase, models.LogServiceAuth, err.Error())).Error("code and request didn't match : ", keySignIn)

		err := fiber.ErrUnprocessableEntity
		err.Message = "UnprocessableEntity"

		return nil, err
	}

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

	// Generate a new pair of access and refresh tokens.
	token, err := utils.GenerateNewTokens(userM.ID, userM.Username)
	if err != nil {
		log.WithFields(utils.LogFormat(models.LogLayerUsecase, models.LogServiceAuth, err.Error())).Error("error generate token")

		err := fiber.ErrUnprocessableEntity
		err.Message = "Please check your email or password didn't match"

		return nil, err

	}

	return token, nil
}
