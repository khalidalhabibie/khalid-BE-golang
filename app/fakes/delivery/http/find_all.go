package http

// func FindAllFakes(c *fiber.Ctx) error {

// 	log.WithFields(utils.LogFormat(models.LogLayerDelivery, models.LogServiceFakes, "start")).Info("index fakes")

// 	fakesM, fakesPagination, err := fakesUsecase.FindAll(request.PaginationConfig(c.Request().URI().QueryArgs().))
// 	if err != nil {
// 		log.WithFields(utils.LogFormat(models.LogLayerDelivery, models.LogServiceFakes, err.Error())).Error("failed to extract data user")

// 		// Return status 401 and error message.
// 		return utils.ReturnFormat(c, fiber.StatusUnprocessableEntity, true, err.Error(), nil)

// 	}

// 	log.WithFields(utils.LogFormat(models.LogLayerDelivery, models.LogServiceFakes, "start")).Info("index fakes")

// }
