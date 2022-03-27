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
	"gokes/platform/email"
	"time"

	fiber "github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func SignUp(request request.SignUp) (*models.User, error) {

	if request.Password != request.Repassword {
		log.WithFields(utils.LogFormat(models.LogLayerUsecase, models.LogServiceAuth, "Password and confirm password does not match")).Error("Password and confirm password does not match")

		err := fiber.ErrUnprocessableEntity
		err.Message = "Password and confirm password does not match"

		return nil, err

	}

	db, err := database.OpenDBConnection()
	if err != nil {
		log.WithFields(utils.LogFormat(models.LogLayerUsecase, models.LogServiceAuth, err.Error())).Error(models.LogErrorTypeConnectionDatabase)

		err := fiber.ErrUnprocessableEntity
		err.Message = "UnprocessableEntity"

		return nil, err
	}

	// check username is exist?
	userCheckM, _ := db.UserRepository.FindByUsername(request.Username)
	if userCheckM != nil {

		log.WithFields(utils.LogFormat(models.LogLayerUsecase, models.LogServiceAuth, fmt.Sprintf("username is exist %v", request.Username))).Error("username is exist")

		err := fiber.ErrUnprocessableEntity
		err.Message = "Username is exist"

		return nil, err
	}

	// check email is exist?
	userCheckM, _ = db.UserRepository.FindByEmail(request.Email)
	if userCheckM != nil {
		log.WithFields(utils.LogFormat(models.LogLayerUsecase, models.LogServiceAuth, fmt.Sprintf("email is exist %v", request.Username))).Error("email is exist")

		err := fiber.ErrUnprocessableEntity
		err.Message = "Email is exist"

		return nil, err
	}

	// get id didn't exist
	newID := uuid.New()
	for {

		_, err := db.UserRepository.FindByID(newID)
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				break
			} else {
				log.WithFields(utils.LogFormat(models.LogLayerUsecase, models.LogServiceAuth, err.Error())).Error("error to find id user")

				err := fiber.ErrUnprocessableEntity
				err.Message = "UnprocessableEntity"

				return nil, err
			}
		}
		newID = uuid.New()
	}

	userM := &models.User{
		ID:       newID,
		Username: request.Username,
		Email:    request.Email,
		Password: utils.GeneratePassword(request.Password),
	}

	connRedis, err := cache.RedisConnection()
	if err != nil {
		log.WithFields(utils.LogFormat(models.LogLayerUsecase, models.LogServiceAuth, err)).Error(models.LogErrorTypeConnectionRedis)

		err := fiber.ErrUnprocessableEntity
		err.Message = "UnprocessableEntity"

		return nil, err
	}

	// Save key(email-otp) = user data to Redis.
	otpLength := 4
	otpNumber := utils.RandomNumber(otpLength)
	keySignUp := fmt.Sprintf("%v-%v-%v", models.AuthUserOTPSignUp, userM.Email, otpNumber)

	dataMarshal, err := json.Marshal(userM)
	if err != nil {
		log.WithFields(utils.LogFormat(models.LogLayerUsecase, models.LogServiceAuth, err.Error())).Error("failed to marshal user to redis")

		err := fiber.ErrUnprocessableEntity
		err.Message = "UnprocessableEntity"

		return nil, err
	}
	err = connRedis.Set(context.Background(), keySignUp, dataMarshal, time.Minute*models.AuthOTPTimeDurationMinutes).Err()
	if err != nil {
		log.WithFields(utils.LogFormat(models.LogLayerDelivery, models.LogServiceAuth, err.Error())).Error("failed to insert redis")

		err := fiber.ErrUnprocessableEntity
		err.Message = "UnprocessableEntity"

		return nil, err
	}

	go email.SendEmailDestination(request.Email, "Sign Up OTP", fmt.Sprintf("Nomer OTP untuk sign up kamu %v : %v \nTolong konfirmasi sebelum %v menit setelah sing up",
		userM.Username,
		otpNumber,
		models.AuthOTPTimeDurationMinutes))

	return userM, nil

}
