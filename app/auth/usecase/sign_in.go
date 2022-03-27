package usecase

import (
	"context"
	"fmt"
	"gokes/app/auth/delivery/http/request"
	"gokes/app/models"
	"gokes/pkg/utils"
	"gokes/platform/cache"
	"gokes/platform/database"
	"gokes/platform/email"
	"time"

	fiber "github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
)

func SignIn(request request.SignIn) error {

	db, err := database.OpenDBConnection()
	if err != nil {

		log.WithFields(utils.LogFormat(models.LogLayerUsecase, models.LogServiceAuth, err.Error())).Error(models.LogErrorTypeConnectionDatabase)

		err := fiber.ErrUnprocessableEntity
		return err
	}

	// find user by email
	userM, err := db.UserRepository.FindByUsername(request.Username)
	if err != nil {
		log.WithFields(utils.LogFormat(models.LogLayerUsecase, models.LogServiceAuth, err)).Error("user not found")

		err := fiber.ErrNotFound
		err.Message = "User Not Found"

		return err

	}

	if !utils.ComparePasswords(userM.Password, request.Password) {
		log.WithFields(utils.LogFormat(models.LogLayerUsecase, models.LogServiceAuth, "password not same")).Error("password not same")

		err := fiber.ErrUnprocessableEntity
		err.Message = "Please check your email or password didn't match"

		return err
	}

	connRedis, err := cache.RedisConnection()
	if err != nil {
		log.WithFields(utils.LogFormat(models.LogLayerUsecase, models.LogServiceAuth, err)).Error(models.LogErrorTypeConnectionRedis)

		err := fiber.ErrUnprocessableEntity
		err.Message = "UnprocessableEntity"

		return err
	}

	// Save key(email) = user data to Redis.
	otpLength := 4
	otpNumber := utils.RandomNumber(otpLength)
	keySignIn := fmt.Sprintf("%v-%v", models.AuthUserOTPSignIn, userM.Username)

	err = connRedis.Set(context.Background(), keySignIn, otpNumber, time.Minute*models.AuthOTPTimeDurationMinutes).Err()
	if err != nil {
		log.WithFields(utils.LogFormat(models.LogLayerDelivery, models.LogServiceAuth, err.Error())).Error("failed to insert redis")

		err := fiber.ErrUnprocessableEntity
		err.Message = "UnprocessableEntity"

		return err
	}

	go email.SendEmailDestination(userM.Email, "Sign In OTP", fmt.Sprintf("Nomer OTP untuk sign in kamu %v : %v \nTolong konfirmasi sebelum %v menit setelah sing in",
		userM.Username,
		otpNumber,
		models.AuthOTPTimeDurationMinutes))

	return nil

}
