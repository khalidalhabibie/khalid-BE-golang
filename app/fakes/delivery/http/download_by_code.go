package http

import (
	"fmt"
	fakesUsecase "gokes/app/fakes/usecase"
	"gokes/pkg/utils"
	"log"

	fiber "github.com/gofiber/fiber/v2"
)

func DownloadByCode(c *fiber.Ctx) error {

	code := c.Params("code")

	_, err := fakesUsecase.FindByCode(code)
	if err != nil {
		return utils.ReturnFormat(c, fiber.StatusNotFound, true, "data not found", nil)
	}

	log.Println(err)

	return c.SendFile(fmt.Sprintf("data/%v.pdf", code), true)

}
