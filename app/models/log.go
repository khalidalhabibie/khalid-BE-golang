package models

const (

	// layer
	LogLayerDelivery   = "delivery"
	LogLayerUsecase    = "usecase"
	LogLayerRepository = "repository"

	// service
	LogServiceUser = "user"
	LogServiceAuth = "auth"
	LogServiceFakes = "fakes"

	// error type
	LogErrorTypeConnectionDatabase = "database connection"
	LogErrorTypeConnectionRedis    = "redis connection"
)
