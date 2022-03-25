package http

import (
	"fmt"
	fakesUsecase "gokes/app/fakes/usecase"
	"gokes/pkg/utils"

	fiber "github.com/gofiber/fiber/v2"
)

func Download(c *fiber.Ctx) error {

	code := c.Params("code")

	_, err := fakesUsecase.FindByCode(code)
	if err != nil {
		utils.ReturnFormat(c, fiber.StatusNotFound, true, "data not found", nil)
	}

	return c.SendFile(fmt.Sprintf("data/%v.pdf", code), true)

}
