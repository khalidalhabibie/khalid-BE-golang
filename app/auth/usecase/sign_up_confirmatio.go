package usecase

import (
	"context"
	"encoding/json"
	"fmt"
	"gokes/app/auth/delivery/http/request"
	"gokes/app/models"
	"gokes/pkg/utils"
	"gokes/platform/cache"
	"gokes/platform/database"

	fiber "github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
)

func SignUpConfirmation(request request.SignUpConfirmation) error {

	// connection redis and and key
	connRedis, err := cache.RedisConnection()
	if err != nil {
		log.WithFields(utils.LogFormat(models.LogLayerUsecase, models.LogServiceAuth, err.Error())).Error(models.LogErrorTypeConnectionRedis)

		err := fiber.ErrUnprocessableEntity
		err.Message = "UnprocessableEntity"

		return err
	}

	keySignUp := fmt.Sprintf("%v-%v-%v", models.AuthUserOTPSignUp, request.Email, request.Code)

	dataFromRedis, err := connRedis.Get(context.Background(), keySignUp).Result()
	if err != nil {
		log.WithFields(utils.LogFormat(models.LogLayerUsecase, models.LogServiceAuth, err.Error())).Error("failed to get data from redis, key : ", keySignUp)

		err := fiber.ErrUnprocessableEntity
		err.Message = "UnprocessableEntity"

		return err
	}

	userM := &models.User{}

	dataFromRedisString := fmt.Sprint(dataFromRedis)

	err = json.Unmarshal([]byte(dataFromRedisString), userM)
	if err != nil {
		log.WithFields(utils.LogFormat(models.LogLayerUsecase, models.LogServiceAuth, err.Error())).Error("failed to unmarshal value from redis , data : ", dataFromRedisString)

		err := fiber.ErrUnprocessableEntity
		err.Message = "UnprocessableEntity"

		return err
	}

	// Create database connection.
	db, err := database.OpenDBConnection()
	if err != nil {

		log.WithFields(utils.LogFormat(models.LogLayerUsecase, models.LogServiceAuth, err.Error())).Error(models.LogErrorTypeConnectionDatabase)

		err := fiber.ErrUnprocessableEntity
		return err

	}

	// Create a new user with validated data
	if err := db.UserRepository.Insert(userM, nil); err != nil {
		log.WithFields(utils.LogFormat(models.LogLayerUsecase, models.LogServiceAuth, err)).Error("Error insert user")
		err := fiber.ErrUnprocessableEntity
		return err

	}

	return nil
}
