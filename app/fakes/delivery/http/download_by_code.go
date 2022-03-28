package http

import (
	"fmt"
	fakesUsecase "gokes/app/fakes/usecase"
	"gokes/pkg/utils"

	fiber "github.com/gofiber/fiber/v2"
)

func DownloadByCode(c *fiber.Ctx) error {

	code := c.Params("code")

	fakesM, _ := fakesUsecase.FindByCode(code)
	if fakesM != nil {
		utils.ReturnFormat(c, fiber.StatusNotFound, true, "data not found", nil)
	}

	return c.SendFile(fmt.Sprintf("data/%v.pdf", code), true)

}
